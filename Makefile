APP := f1-bot
BIN_DIR := bin
INTERMEDIATES := $(BIN_DIR)/intermediates
SRC := main.go
BUILD := go build
BUCKET := f1-fantasy-bucket
STACK := f1-fantasy-stack
BUCKET_EXISTS := aws s3api head-bucket --bucket
CREATE_BUCKET := aws s3 mb
CREATE_PACKAGE := aws cloudformation package
DEPLOY := aws cloudformation deploy

BIN = $(BIN_DIR)/$(APP)
PACKAGE := $(INTERMEDIATES)/package.yml

.PHONY: clean deploy $(BUCKET) configure

all: $(BIN)

clean:
	@rm -rf $(BIN_DIR)

configure:
	@scripts/configure.sh

deploy: $(BIN) $(PACKAGE)
	@$(DEPLOY) --template-file $(PACKAGE) --stack-name $(STACK) --capabilities CAPABILITY_NAMED_IAM
	@echo "🚀 Deployed!"

$(PACKAGE): $(BUCKET)
	@mkdir -p $(@D)
	@$(CREATE_PACKAGE) --template-file template.yml --s3-bucket $(BUCKET) --output-template-file $(PACKAGE)

$(BUCKET):
	@$(BUCKET_EXISTS) $(BUCKET) || $(CREATE_BUCKET) s3://$(BUCKET)

$(BIN): $(SRC)
	@mkdir -p $(@D)
	@$(BUILD) -o $@ $<
