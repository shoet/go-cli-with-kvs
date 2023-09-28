.PHONY: help up destroy
.DEFAULT_GOAL := help

DOCKER_TAG := latest
up: ## up on plumi
	pulumi up

destroy: ## destroy from plumi
	pulumi destroy -s shoet/upstash-iac/dev

help: ## Show options
	@grep -E '^[a-zA-Z_]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'
	
