AUTHOR?=bannzai
PROJECT?=itree
DOMAIN?=github.com
PACKAGE?=$(DOMAIN)/$(AUTHOR)/$(PROJECT)

dry-run: build
	./$(PROJECT) 

build:
	go build -o $(PROJECT) main.go

install: go111moduleoff dependency
	go install

go111moduleoff:
	export GO111MODULE=off

go111moduleon:
	export GO111MODULE=on

test:
	go test $(PACKAGE)...

test-verbose:
	go test $(PACKAGE)... -v

update-dependency:
	go mod tidy

dependency: go111moduleon
	go mod download


