BUF ?= buf

VALIDATE_VERSION ?= v1.0.2
BUILD_DIR := .build/

.DEFAULT: gen-prot

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
