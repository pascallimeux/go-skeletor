OUTPUT ?= hello
BUILD=go build
LDFLAGS=--ldflags "-w -s"
UPXFLAGS= --best --lzma
PACKAGE=examples
.DEFAULT_GOAL := build

build: lint
	@cd cmd && $(BUILD) $(LDFLAGS) -o ../bin/$(OUTPUT) && upx $(UPXFLAGS) ../bin/$(OUTPUT)
	@printf "\033[32m\xE2\x9c\x93 $(OUTPUT) compiled in /bin \n\033[0m"

lint:
	@golint $(PACKAGE)

test:
	@go test ./$(PACKAGE) -race -covermode=atomic -coverprofile=coverage.out

clean:
	@rm -Rf ./bin/*
	@printf "\033[32m\xE2\x9c\x93 binaries removed on /bin \n\033[0m"