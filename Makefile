BUF ?= buf

VALIDATE_VERSION ?= v1.0.2
BUILD_DIR := .build/

.DEFAULT: gen-proto

# ====================================================================================
# Makefile helper functions for helm-docs: https://github.com/norwoodj/helm-docs
#

HELM_DOCS_VERSION := v1.11.0
HELM_DOCS := helm-docs
HELM_DOCS_REPO := github.com/norwoodj/helm-docs/cmd/helm-docs

bin-$(HELM_DOCS): ## Installs helm-docs
	@GO111MODULE=on go install $(HELM_DOCS_REPO)@$(HELM_DOCS_VERSION)

build_dir:
	mkdir -p $(BUILD_DIR)

protoc-gen-validate: build_dir
	if test ! -d $(BUILD_DIR)validate-$(VALIDATE_VERSION)/; then \
		git clone --branch $(VALIDATE_VERSION) https://github.com/bufbuild/protoc-gen-validate.git $(BUILD_DIR)validate-$(VALIDATE_VERSION); \
	else \
		git -C $(BUILD_DIR)validate-$(VALIDATE_VERSION)/ pull --all; \
		git -C $(BUILD_DIR)validate-$(VALIDATE_VERSION)/ checkout $(VALIDATE_VERSION); \
	fi

	cd $(BUILD_DIR) && ln -sfn validate-$(VALIDATE_VERSION)/ validate

.PHONY: gen-proto
gen-proto: protoc-gen-validate
	$(BUF) generate

build-image:
	docker build --force-rm=true -t data-control-center-api -f ./Dockerfile . && \
	docker tag data-control-center-api:latest koor/data-control-center-api:$(TAG)

run-cephapidummy:
	npx ts-node-esm ./node/server.ts

.PHONY: gen-licenses
gen-licenses:
	yarn licenses generate-disclaimer > ./src/public/licenses/frontend.txt
	go-licenses report . --template internal/scripts/licenses-backend.txt.tpl > ./src/public/licenses/backend.txt

helm-docs: bin-$(HELM_DOCS) ## Use helm-docs to generate documentation from helm charts
	$(HELM_DOCS) -c charts/data-control-center \
		-o README.md \
		-t README.gotmpl.md \
		-t _templates.gotmpl
