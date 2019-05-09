# Parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
BINARY_NAME=dist/sensu-sic-handler

all: test build
build:
	$(GOBUILD) -a -o $(BINARY_NAME)
test:
	$(GOTEST)
clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
run:
	$(GOBUILD) -o $(BINARY_NAME)
	./$(BINARY_NAME)
deps:
	dep ensure