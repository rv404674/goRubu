# To run commands just use "make" and make function name
# ex- <make run> use for run go run main.go
# Go parameters

execution: echo "** Executing Makefile**"
GOCMD=go

GOINSTALL=$(GOCMD) install
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GORUN=$(GOCMD) run
BINARY_NAME=main
BINARY_UNIX=$(BINARY_NAME)_unix

# if we do "make run main", it will do "go run main.go"
execute:
	- @echo "** Please wait. Connecting with Mongo, Memcached and Spinning up the Go Server **"
	- ~/go/bin/goRubu

setup:
	- export GOBIN=~/go/bin/
	# my pwd is "/Users/home"
	- $(GOCMD) mod init goRubu

install:
	- @echo "** Will build the package into a single binary **"
	- $(GOINSTALL)

# make all -> will first install and then run execute
all:
	- install execute

docker:
	- @echo "** Spinning up the GoRubu, Mongo and Memcached Docker Containers **"
	- @chmod 777 scripts/build_docker.sh
	- @scripts/build_docker.sh

#NOTE: @ before a command will stop showing that command.
test:
	- @echo "** Running Tests **"
	- $(GOTEST) ./tests -v

