host?=dev

repo_name=hmrc_oauth
proj_path=github.com/awltux
repo_path=$(proj_path)/$(repo_name)
vbox_share=/media/sf_gopath/src

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

# Install the goa libraries and command
# WARNING: May need to be run a few times to get all the dependencies downloaded.
goa-init: 
	export GO111MODULE=on && \
	go get -u goa.design/goa/v3 && \
	go get -u goa.design/goa/v3/...

# Generate service API code from the design directory
goa-gen: rsync
	export GO111MODULE=on && \
	goa gen $(repo_path)/design

# Generate example application and client code from the design directory
goa-example: rsync
	export GO111MODULE=on && \
	goa example $(repo_path)/design

# Buidl and start the example server
goa-server: rsync
	export GO111MODULE=on && \
	pushd cmd/mtd_server && \
	go build && \
	./mtd_server -host=$(host); \
	popd;

