package hmrc

import (
	"context"
	"fmt"
	"log"

	mtd "github.com/awltux/hmrc_oauth/gen/mtd"
	memdb "github.com/hashicorp/go-memdb"
)

// mtdsrvc service example implementation.
// The example methods log the requests and return zero values.
type mtdsrvc struct {
	logger *log.Logger
	db     *memdb.MemDB
}

// KeyStore users key and hmrc token
type KeyStore struct {
	State            string
	Token            string
	ClientAddress    string
	Error            string
	ErrorDescription string
	ErrorCode        string
}

// NewMtd returns the mtd service implementation.
func NewMtd(logger *log.Logger) mtd.Service {

	// Create a memdb schema to hold KeyStore objects
	schema := &memdb.DBSchema{
		Tables: map[string]*memdb.TableSchema{
			"keystore": &memdb.TableSchema{
				Name: "keystore",
				Indexes: map[string]*memdb.IndexSchema{
					"id": &memdb.IndexSchema{
						Name:    "id",
						Unique:  true,
						Indexer: &memdb.StringFieldIndex{Field: "State"},
					},
				},
			},
		},
	}

	// Create a database of schema objects
	db, err := memdb.NewMemDB(schema)
	if err != nil {
		panic(err)
	}

	// Return a service object containing a logger and the memdb
	return &mtdsrvc{
		logger,
		db,
	}
}

// Store key that will store oauth token
func (s *mtdsrvc) Register(ctx context.Context, p *mtd.StatePayload) (err error) {
	stateMinLength := 40
	stateMaxLength := 50
	// Is State the correct length
	stateLength := len(*(p.State))
	if (stateLength < stateMinLength) || (stateLength > stateMaxLength) {
		return mtd.MakeKeyLengthError(fmt.Errorf("State length must be between %d and %d but is actually %d", stateMinLength, stateMaxLength, stateLength))
	}

	txn := s.db.Txn(true)

	raw, err := txn.First("keystore", "id", *(p.State))
	if err != nil {
		txn.Abort()
		panic(err)
	}

	if raw != nil {
		// TODO: check for matching ClientAddress
		txn.Abort()
		return mtd.MakeKeyAlreadyExists(fmt.Errorf("State has already been registered"))
	}

	// Insert a new keystore
	kStore := &KeyStore{*(p.State), "", "", "", "", ""}
	if err := txn.Insert("keystore", kStore); err != nil {
		txn.Abort()
		panic(err)
	}

	// Commit the transaction
	txn.Commit()

	return
}

// Retrieve key with oauth token
func (s *mtdsrvc) Retrieve(ctx context.Context, p *mtd.StatePayload) (res string, err error) {
	if p.State == nil {
		return "", mtd.MakeInvalidRequest(fmt.Errorf("State parameter is missing"))
	}

	//	s.logger.Print(*(p.State))
	txn := s.db.Txn(false)
	defer txn.Abort()

	raw, err := txn.First("keystore", "id", *(p.State))
	if err != nil {
		panic(err)
	}

	if raw == nil {
		return "", mtd.MakeMatchingKeyNotFound(fmt.Errorf("State not found in database"))
	}

	if raw.(*KeyStore).Token == "" {
		return "", mtd.MakeKeyHasNoToken(fmt.Errorf("State found but HMRC has not yet populated the Token"))
	}

	return raw.(*KeyStore).Token, nil
}

// Authentication code response
func (s *mtdsrvc) HmrcCallback(ctx context.Context, p *mtd.CodePayload) (err error) {

	if p.State == nil {
		return mtd.MakeInvalidRequest(fmt.Errorf("HMRC didn't provide a state parameter"))
	}

	stateMinLength := 40
	stateMaxLength := 50
	// Is State the correct length
	stateLength := len(*(p.State))
	if (stateLength < stateMinLength) || (stateLength > stateMaxLength) {
		return mtd.MakeKeyLengthError(fmt.Errorf("State length must be between %d and %d but is actually %d", stateMinLength, stateMaxLength, stateLength))
	}

	txn := s.db.Txn(true)
	defer txn.Abort()

	raw, err := txn.First("keystore", "id", *(p.State))
	if err != nil {
		panic(err)
	}

	if raw == nil {
		return mtd.MakeMatchingKeyNotFound(fmt.Errorf("State not found in database"))
	}

	if p.Error != nil {
		// Copy error info from HMRC in DB
		raw.(*KeyStore).Error = *(p.Error)

		if p.ErrorDescription != nil {
			raw.(*KeyStore).ErrorDescription = *(p.ErrorDescription)
		}

		if p.ErrorCode != nil {
			raw.(*KeyStore).ErrorCode = *(p.ErrorCode)
		}

		err = txn.Insert("keystore", raw)
		if err != nil {
			panic(err)
		}

		txn.Commit()
		return
	}

	if p.Code == nil || *(p.Code) == "" {
		return mtd.MakeInvalidRequest(fmt.Errorf("HMRC call had neither code nor error"))
	}
	// Store code as token in DB
	// TODO: Check length of Code
	raw.(*KeyStore).Token = *(p.Code)
	err = txn.Insert("keystore", raw)
	if err != nil {
		panic(err)
	}

	// Commit the transaction
	txn.Commit()
	return
}
