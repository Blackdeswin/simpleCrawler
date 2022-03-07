#!make
.PHONY: build test

install-swagger:
ifeq ($(wildcard $(SWAGGER_BIN)),)
	@echo "Downloading go-swagger $(SWAGGER_TAG)"
	@curl -o $(LOCAL_BIN)/swagger -L'#' "https://github.com/go-swagger/go-swagger/releases/download/v$(SWAGGER_TAG)/swagger_darwin_amd64"
	@chmod +x $(LOCAL_BIN)/swagger
endif

## Dependencies
tidy: checkpath
	@go mod tidy

deps: tidy

# Генерация сервера
gen:
	@mkdir -p cmd
	@mkdir -p internal/generated
	@mkdir -p internal/app
	@mkdir -p internal/config
	@goswagger generate server \
                     -f ./swagger-api/swagger.yml \
                     -t ./internal/generated -C ./swagger-templates/default-server.yml \
                     --template-dir ./swagger-templates/templates \
                     --name simpleCrawler
