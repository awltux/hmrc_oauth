// Code generated by goa v3.0.2, DO NOT EDIT.
//
// mtdServer HTTP client CLI support package
//
// Command:
// $ goa gen github.com/awltux/hmrc_oauth/design

package cli

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	mtdc "github.com/awltux/hmrc_oauth/gen/http/mtd/client"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// UsageCommands returns the set of commands and sub-commands using the format
//
//    command (subcommand1|subcommand2|...)
//
func UsageCommands() string {
	return `mtd (register|retrieve|hmrc-callback)
`
}

// UsageExamples produces an example of a valid invocation of the CLI tool.
func UsageExamples() string {
	return os.Args[0] + ` mtd register --state "Unde dicta rerum facere."` + "\n" +
		""
}

// ParseEndpoint returns the endpoint and payload as specified on the command
// line.
func ParseEndpoint(
	scheme, host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restore bool,
) (goa.Endpoint, interface{}, error) {
	var (
		mtdFlags = flag.NewFlagSet("mtd", flag.ContinueOnError)

		mtdRegisterFlags     = flag.NewFlagSet("register", flag.ExitOnError)
		mtdRegisterStateFlag = mtdRegisterFlags.String("state", "REQUIRED", "Key submitted to oAuth call; normally AES1 digest")

		mtdRetrieveFlags     = flag.NewFlagSet("retrieve", flag.ExitOnError)
		mtdRetrieveStateFlag = mtdRetrieveFlags.String("state", "REQUIRED", "Key submitted to oAuth call; normally AES1 digest")

		mtdHmrcCallbackFlags                = flag.NewFlagSet("hmrc-callback", flag.ExitOnError)
		mtdHmrcCallbackCodeFlag             = mtdHmrcCallbackFlags.String("code", "", "")
		mtdHmrcCallbackStateFlag            = mtdHmrcCallbackFlags.String("state", "", "")
		mtdHmrcCallbackErrorFlag            = mtdHmrcCallbackFlags.String("error", "", "")
		mtdHmrcCallbackErrorDescriptionFlag = mtdHmrcCallbackFlags.String("error-description", "", "")
		mtdHmrcCallbackErrorCodeFlag        = mtdHmrcCallbackFlags.String("error-code", "", "")
	)
	mtdFlags.Usage = mtdUsage
	mtdRegisterFlags.Usage = mtdRegisterUsage
	mtdRetrieveFlags.Usage = mtdRetrieveUsage
	mtdHmrcCallbackFlags.Usage = mtdHmrcCallbackUsage

	if err := flag.CommandLine.Parse(os.Args[1:]); err != nil {
		return nil, nil, err
	}

	if flag.NArg() < 2 { // two non flag args are required: SERVICE and ENDPOINT (aka COMMAND)
		return nil, nil, fmt.Errorf("not enough arguments")
	}

	var (
		svcn string
		svcf *flag.FlagSet
	)
	{
		svcn = flag.Arg(0)
		switch svcn {
		case "mtd":
			svcf = mtdFlags
		default:
			return nil, nil, fmt.Errorf("unknown service %q", svcn)
		}
	}
	if err := svcf.Parse(flag.Args()[1:]); err != nil {
		return nil, nil, err
	}

	var (
		epn string
		epf *flag.FlagSet
	)
	{
		epn = svcf.Arg(0)
		switch svcn {
		case "mtd":
			switch epn {
			case "register":
				epf = mtdRegisterFlags

			case "retrieve":
				epf = mtdRetrieveFlags

			case "hmrc-callback":
				epf = mtdHmrcCallbackFlags

			}

		}
	}
	if epf == nil {
		return nil, nil, fmt.Errorf("unknown %q endpoint %q", svcn, epn)
	}

	// Parse endpoint flags if any
	if svcf.NArg() > 1 {
		if err := epf.Parse(svcf.Args()[1:]); err != nil {
			return nil, nil, err
		}
	}

	var (
		data     interface{}
		endpoint goa.Endpoint
		err      error
	)
	{
		switch svcn {
		case "mtd":
			c := mtdc.NewClient(scheme, host, doer, enc, dec, restore)
			switch epn {
			case "register":
				endpoint = c.Register()
				data, err = mtdc.BuildRegisterPayload(*mtdRegisterStateFlag)
			case "retrieve":
				endpoint = c.Retrieve()
				data, err = mtdc.BuildRetrievePayload(*mtdRetrieveStateFlag)
			case "hmrc-callback":
				endpoint = c.HmrcCallback()
				data, err = mtdc.BuildHmrcCallbackPayload(*mtdHmrcCallbackCodeFlag, *mtdHmrcCallbackStateFlag, *mtdHmrcCallbackErrorFlag, *mtdHmrcCallbackErrorDescriptionFlag, *mtdHmrcCallbackErrorCodeFlag)
			}
		}
	}
	if err != nil {
		return nil, nil, err
	}

	return endpoint, data, nil
}

// mtdUsage displays the usage of the mtd command and its subcommands.
func mtdUsage() {
	fmt.Fprintf(os.Stderr, `Service is the mtd service interface.
Usage:
    %s [globalflags] mtd COMMAND [flags]

COMMAND:
    register: Store key that will store oauth token
    retrieve: Store key that will store oauth token
    hmrc-callback: Authentication code response

Additional help:
    %s mtd COMMAND --help
`, os.Args[0], os.Args[0])
}
func mtdRegisterUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] mtd register -state STRING

Store key that will store oauth token
    -state STRING: Key submitted to oAuth call; normally AES1 digest

Example:
    `+os.Args[0]+` mtd register --state "Unde dicta rerum facere."
`, os.Args[0])
}

func mtdRetrieveUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] mtd retrieve -state STRING

Store key that will store oauth token
    -state STRING: Key submitted to oAuth call; normally AES1 digest

Example:
    `+os.Args[0]+` mtd retrieve --state "Minima quaerat aut similique voluptas non ea."
`, os.Args[0])
}

func mtdHmrcCallbackUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] mtd hmrc-callback -code STRING -state STRING -error STRING -error-description STRING -error-code STRING

Authentication code response
    -code STRING: 
    -state STRING: 
    -error STRING: 
    -error-description STRING: 
    -error-code STRING: 

Example:
    `+os.Args[0]+` mtd hmrc-callback --code "Aut saepe dolor est." --state "Et in culpa." --error "Labore iure." --error-description "Non qui." --error-code "Eius illum amet."
`, os.Args[0])
}