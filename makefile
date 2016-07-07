version=0.1.0

.PHONY: all

NOVENDOR := $(shell glide novendor)

all:
	@echo "make <cmd>"
	@echo ""
	@echo "commands:"
	@echo "  build		 - build the source code"
	@echo "  lint		  - lint the source code"
	@echo "  test		  - test the source code"
	@echo "  fmt		   - format the source code"
	@echo "  clean		 - clean the dist build"
	@echo ""
	@echo "  deps		  - install tool dependencies"

clean:
	@rm -rf ./build

lint:
	@go vet $(NOVENDOR)
	@go list ./... | grep -v /vendor/ | xargs -L1 golint

test:
	@go test $(NOVENDOR)

fmt:
	@gofmt -l -w $(NOVENDOR)
	@./node_modules/.bin/jsfmt -w ./public/scripts ./public/*.js

build: clean lint
	@go build -o ./build/server.bin ./main.go

deps:
	@npm install
	@go get github.com/golang/lint/golint
	@go get github.com/Masterminds/glide
	@glide install
