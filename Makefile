.DEFAULT_GOAL := help
.PHONY: build install help

build: ## Build all sub projects
	cd ping && make clean build pack

install: ## Install and deploy
	cd infra && make validate && make plan && make apply

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
