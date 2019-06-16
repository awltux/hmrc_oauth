# Server profile to start with: [dev], prod
host?=dev
state?=0123456789012345678901234567890123456789
token?=9876543210987654321098765432109876543210

repo_name=hmrc_oauth
proj_path=github.com/awltux
repo_path=$(proj_path)/$(repo_name)
vbox_share=/media/sf_gopath/src

# test variables
shortState?=012345678901234567890123456789012345678
longState?=012345678901234567890123456789012345678901234567890
printHttpCode=-s -o /dev/null -X POST -w '%{http_code}' 


list:
	@echo "List of make targets:"
	@grep '^[^#\.[:space:]].*:' Makefile | sed 's/^\(.*\):.*/\1/'


# When using vbox, shared folders are slow and cause problems with go builds
# This rsync automatically updates a local copy of the project to the users GOPATH
# Only relevant if working with vbox
rsync:
# Only rsync if in vbox
ifneq ($(wildcard $(vbox_share)),)
  # Copy the project excluding generated files
	rsync -avzh --progress --exclude gen --exclude cmd $(vbox_share)/$(repo_path)/ .
else
	@echo "[rsync] vbox share not found: $(vbox_share)"
endif

# Setup my local git credentials
git-config:
	git config user.name "awltux" && \
	git config user.email "chris.d.holman@googlemail.com" && \
	git config color.ui true && \
	git config credential.helper store;

define GOLANG_PROFILE
#!/bin/bash

export GOROOT=/usr/local/go
export GOPATH=~/go

export PATH=$$PATH:/usr/local/go/bin:~/go/bin
export GO111MODULE=on
endef

export GOLANG_PROFILE
go-init:
	#Install goLang
	# Preparing Your Environment For Go Development
	# https://golang.org/dl/
	curl -o /tmp/go.tgz https://dl.google.com/go/go1.12.5.linux-amd64.tar.gz
	sudo mkdir -p /usr/local
	sudo tar -C /usr/local -zxvf /tmp/go.tgz
	sudo bash -c "echo \"$${GOLANG_PROFILE}\" > /etc/profile.d/golang.sh"
	sudo chmod 755 /etc/profile.d/golang.sh


# Install the goa libraries and command
# WARNING: May need to be run a few times to get all the dependencies downloaded.
goa-init: 
	export GO111MODULE=on && \
	go get -u goa.design/goa/v3 && \
	go get -u goa.design/goa/v3/...

# Generate service API code from the design directory
goa-gen: rsync
	export GO111MODULE=on && \
	goa gen  $(repo_path)/design

# Generate example application and client code from the design directory
goa-example: rsync
	export GO111MODULE=on && \
	goa example $(repo_path)/design

# Build and start the example server
goa-build: rsync
	export GO111MODULE=on && \
	pushd cmd/mtd_server && \
	go build &&\
	popd;

# Build and start the example server
goa-server-start:
	@echo "Starting mtd_server"
	@./cmd/mtd_server/mtd_server -host=$(host) &

goa-server-stop:
	@echo "Stopping mtd_server" &&\
	kill $(shell ps aux | grep ./cmd/mtd_server/mtd_server | grep -v grep | awk '{ print $$2 }' )


test-all: goa-build goa-server-start
	# Server starting as background task; wait for it to initialise.
	@httpCode=0; while [[ 405 -ne $${httpCode} ]]; do \
	      httpCode=$(shell curl $(printHttpCode) \
	          -X GET http://localhost:8088/v1/mtd/);  \
				echo "Waiting for server to start"; \
				sleep 1; \
				done;
	@httpCode=$(shell curl $(printHttpCode) \
	      -X GET http://localhost:8088/v1/mtd/); \
				if [[ 405 -ne $${httpCode} ]]; then \
				echo "[FAILED] Didnt reject missing state: $${httpCode}"; else \
				echo "[SUCCESS] Rejected missing state: $${httpCode}"; fi;
	@httpCode=$(shell curl $(printHttpCode) \
	      -X GET http://localhost:8088/v1/mtd/$(state)); \
				if [[ 404 -ne $${httpCode} ]]; then \
				echo "[FAILED] Didnt reject unknown state: $${httpCode}"; else \
				echo "[SUCCESS] Rejected unknown state: $${httpCode}"; fi;
	@httpCode=$(shell curl $(printHttpCode) \
	      -X POST http://localhost:8088/v1/mtd/$(shortState)); \
				if [[ 412 -ne $${httpCode} ]]; then \
				echo "[FAILED] Didnt reject short state: $${httpCode}"; else \
				echo "[SUCCESS] Rejected short state: $${httpCode}"; fi;
	@httpCode=$(shell curl $(printHttpCode) \
	      -X POST http://localhost:8088/v1/mtd/$(longState)); \
				if [[ 412 -ne $${httpCode} ]]; then \
				echo "[FAILED] Didnt reject long state: $${httpCode}"; else \
				echo "[SUCCESS] Rejected long state: $${httpCode}"; fi;
	@httpCode=$(shell curl $(printHttpCode) \
	      -X POST http://localhost:8088/v1/mtd/$(state)); \
				if [[ 201 -ne $${httpCode} ]]; then \
				echo "[FAILED] Didnt accept valid state: $${httpCode}"; else \
				echo "[SUCCESS] Accepted valid state: $${httpCode}"; fi;
	@httpCode=$(shell curl $(printHttpCode) \
	      -X POST http://localhost:8088/v1/mtd/$(state)); \
				if [[ 409 -ne $${httpCode} ]]; then \
				echo "[FAILED] Didnt reject repeated state: $${httpCode}"; else \
				echo "[SUCCESS] Rejected existing state: $${httpCode}"; fi;
	@httpCode=$(shell curl $(printHttpCode) \
	      -X GET http://localhost:8088/v1/mtd/$(state)); \
				if [[ 206 -ne $${httpCode} ]]; then \
				echo "[FAILED] Found unexpected token for previously registered state: $${httpCode}"; else \
				echo "[SUCCESS] Didnt find token for registered state: $${httpCode}"; fi;
	@httpCode=$(shell curl $(printHttpCode) \
	      -X POST http://localhost:8088/v1/mtd); \
				if [[ 400 -ne $${httpCode} ]]; then \
				echo "[FAILED] Didnt reject missing state parameter: $${httpCode}"; else \
				echo "[SUCCESS] Rejected missing state parameter: $${httpCode}"; fi;
	@httpCode=$(shell curl $(printHttpCode) \
	      -X POST http://localhost:8088/v1/mtd?state=$(shortState)); \
				if [[ 412 -ne $${httpCode} ]]; then \
				echo "[FAILED] Didnt reject unkown state parameter: $${httpCode}"; else \
				echo "[SUCCESS] Rejected unknown state parameter: $${httpCode}"; fi;
	@httpCode=$(shell curl $(printHttpCode) \
	      -X POST http://localhost:8088/v1/mtd?state=$(state)); \
				if [[ 400 -ne $${httpCode} ]]; then \
				echo "[FAILED] Didnt reject missing token parameter: $${httpCode}"; else \
				echo "[SUCCESS] Rejected missing token parameter: $${httpCode}"; fi;
	@httpCode=$(shell curl $(printHttpCode) \
	      -X POST "http://localhost:8088/v1/mtd?state=$(state)&code=$(token)"); $(shell # Must quote multiple parameters) \
				if [[ 200 -ne $${httpCode} ]]; then \
				echo "[FAILED] Didnt accept valid token parameter: $${httpCode}"; else \
				echo "[SUCCESS] Accepted valid token parameter: $${httpCode}"; fi;
	@httpCode=$(shell curl $(printHttpCode) \
	      -X GET http://localhost:8088/v1/mtd/$(state)); \
				if [[ 200 -ne $${httpCode} ]]; then \
				echo "[FAILED] Didnt find valid token for previously registered state: $${httpCode}"; else \
				echo "[SUCCESS] Found valid token for registered state: $${httpCode}"; fi;
	@echo "Stopping mtd_server" &&\
	kill $(shell ps aux | grep ./cmd/mtd_server/mtd_server | grep -v grep | awk '{ print $$2 }' )


test-post-state:
	echo "COMMAND: make state=$(state) test-post-state"
	curl -i -X POST http://localhost:8088/v1/mtd/$(state)

test-get-token:
	echo "COMMAND: make state=$(state) test-get-token"
	curl -i -X GET http://localhost:8088/v1/mtd/$(state)

test-post-token:
	echo "COMMAND: make state=$(state) test-post-token"
	curl -i -X POST http://localhost:8088/v1/mtd?state=$(state);code=$(token)
