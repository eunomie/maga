.DEFAULT_GOAL := help
TERRAFORM_VERSION = 0.7.9

ifneq ($(shell which terraform-$(TERRAFORM_VERSION) 2>/dev/null),)
  # We have terraform-x.y.z on the PATH
  TERRAFORM_BIN=terraform-$(TERRAFORM_VERSION)
else ifeq ($(shell which terraform >/dev/null 2>&1 && terraform version | head -n 1),Terraform v$(TERRAFORM_VERSION))
  # The terraform binary on the PATH has the correct version
  TERRAFORM_BIN=terraform
else
  # Use Docker to ensure correct terraform version
  TERRAFORM_BIN = docker run -it --rm -v `pwd`:/data sjourdan/terraform:$(TERRAFORM_VERSION)
endif

validate: terraform-fmt terraform-validate	## Validate syntax

plan:	terraform-validate terraform-get terraform-plan ## Plan changes

apply: terraform-validate terraform-get terraform-apply ## Apply Changes

output: terraform-output	## Output display

destroy: terraform-destroy	## Destroy

terraform-validate:
	$(TERRAFORM_BIN) validate

terraform-get:
	$(TERRAFORM_BIN) get

terraform-plan:
	$(TERRAFORM_BIN) plan

terraform-apply:
	$(TERRAFORM_BIN) apply

terraform-fmt:
	$(TERRAFORM_BIN) fmt -list

terraform-output:
	$(TERRAFORM_BIN) output

terraform-destroy:
	$(TERRAFORM_BIN) destroy

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
