#To run command gust use "make" and make function name
#ex- <make run> use for run go run main.go
#Go parameters

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

#make functions
deps:
	echo "Installing dependencies"
	# use for routing
	$(GOGET) github.com/gorilla/mux
	#use to load .env file
	$(GOGET) github.com/joho/godotenv
	# mongo driver
	$(GOGET) go.mongodb.org/mongo-driver
	# can write a cron in go program itself
	$(GOGET) github.com/jasonlvhit/gocron
	# memcached client
	$(GOGET) github.com/rainycape/memcache


# if we do "make run main", it will do "go run main.go"
execute:
	~/go/bin/goRubu

setup:
	export GOBIN=~/go/bin/
	# my pwd is "/Users/home"
	$(GOCMD) mod init goRubu

install:
	echo "Executing go install"
	$(GOINSTALL)

# make all -> will first install and then run execute
all:
	install execute
