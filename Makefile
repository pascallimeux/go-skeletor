OUTPUT ?= hello
BUILD=go build
LDFLAGS=--ldflags "-w -s"
UPXFLAGS= --best --lzma
PACKAGE=examples
DOCKER_IMAGE=myimage
.DEFAULT_GOAL := build

build: lint
	@cd cmd && $(BUILD) $(LDFLAGS) -o ../bin/$(OUTPUT) && upx $(UPXFLAGS) ../bin/$(OUTPUT)
	@printf "\033[32m\xE2\x9c\x93 $(OUTPUT) compiled in /bin \n\033[0m"

lint:
	@go list ./... | grep -v /vendor/ | xargs -L1 golint -set_exit_status

test:
	@go test ./... -race -covermode=atomic

vet:
	@go vet ./...

clean:
	@rm -Rf ./bin/*
	@printf "\033[32m\xE2\x9c\x93 binaries removed on /bin \n\033[0m"

docker-build:
	@docker build . -t $(DOCKER_IMAGE)

docker-run:
	@docker run $(DOCKER_IMAGE)