BUF ?= buf

VALIDATE_VERSION ?= v1.0.2
BUILD_DIR := .build/

.DEFAULT: run-server
all: gen-proto

# Build, Format, etc., Tools, Dependency checkouts

buf:
ifeq (, $(shell which buf))
	go install github.com/bufbuild/buf/cmd/buf@v1.26.1
endif

protoc-gen-validate: build_dir
	if test ! -d $(BUILD_DIR)validate-$(VALIDATE_VERSION)/; then \
		git clone --branch $(VALIDATE_VERSION) https://github.com/bufbuild/protoc-gen-validate.git $(BUILD_DIR)validate-$(VALIDATE_VERSION); \
	else \
		git -C $(BUILD_DIR)validate-$(VALIDATE_VERSION)/ pull --all; \
		git -C $(BUILD_DIR)validate-$(VALIDATE_VERSION)/ checkout $(VALIDATE_VERSION); \
	fi

	cd $(BUILD_DIR) && ln -sfn validate-$(VALIDATE_VERSION)/ validate

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

.PHONY: clean
clean:
	@npx nuxi cleanup
	rm -rf ./.nuxt/dist/

.PHONY: watch
watch:
	yarn dev

.PHONY: gen-proto
gen-proto: protoc-gen-validate
	$(BUF) generate

.PHONY: build-container
build-container: build-yarn
	docker build \
		--force-rm=true\
		-t docker.io/koorinc/data-control-center:latest .

.PHONY: release
release:
	docker tag docker.io/koorinc/data-control-center:latest docker.io/koorinc/data-control-center:$(TAG)

.PHONY: build-go
build-go:
	CGO_ENABLED=0 go build -a -installsuffix cgo -o data-control-center .

.PHONY: build-yarn
build-yarn:
	rm -rf ./.nuxt/dist/
	yarn build
	yarn generate

.PHONY: run-cephapidummy
run-cephapidummy:
	npx ts-node-esm ./node/server.ts

.PHONY: run-server
run-server:
	mkdir -p ./.nuxt/dist/
	go run . server

.PHONY: fmt
fmt:
	$(MAKE) fmt-proto gen-proto
	$(MAKE) fmt-js

.PHONY: fmt-proto
fmt-proto: buf
	$(BUF) format --write ./api

.PHONY: fmt-js
fmt-js:
	yarn prettier --write ./src

.PHONY: gen-licenses
gen-licenses:
	yarn licenses generate-disclaimer > ./src/public/licenses/frontend.txt
	go-licenses report . --template internal/scripts/licenses-backend.txt.tpl > ./src/public/licenses/backend.txt

.PHONY: helm-docs
helm-docs: bin-$(HELM_DOCS) ## Use helm-docs to generate documentation from helm charts
	$(HELM_DOCS) -c charts/data-control-center \
		-o README.md \
		-t README.gotmpl.md \
		-t _templates.gotmpl

.PHONY: lint
lint:
	$(BUF) lint
	$(BUF) breaking --against '.git#branch=main'
