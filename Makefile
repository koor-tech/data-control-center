BUF ?= buf

VALIDATE_VERSION ?= v1.0.2
BUILD_DIR := .build/

.DEFAULT: proto-gen

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

.PHONY: proto-gen
proto-gen: protoc-gen-validate
	$(BUF) generate
