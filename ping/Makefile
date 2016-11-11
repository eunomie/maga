.DEFAULT_GOAL := help
.PHONY: build pack clean help

do_build:
	go build -buildmode=c-shared -ldflags="-w -s" -o handler.so
	chown `stat -c "%u:%g" .` handler.so

do_pack:
	zip handler.zip handler.so
	chown `stat -c "%u:%g" .` handler.zip

do_deps:
	go get ./...

deps: ## Get dependencies
	GOPATH=$(PWD) go get ./...

build: ## Compile Go to Lambda
	@docker run --rm \
		-v $(GOPATH):/go \
		-v $(PWD):/usr/src/current \
		-w /usr/src/current \
		eawsy/aws-lambda-go make do_build

pack: ## Pack to a lambda zip
	@docker run --rm \
		-v $(GOPATH):/go \
		-v $(PWD):/usr/src/current \
		-w /usr/src/current \
		eawsy/aws-lambda-go make do_pack

clean: ## Clean temporary files
	@rm -rf handler.zip handler.so

cleanpack: clean deps build pack ## Create a clean pack

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
