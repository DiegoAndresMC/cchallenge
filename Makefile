BUILDPATH=$(CURDIR)
API_NAME=searching-products-wallmart-api

build:
	@echo "Creando Binario ..."
	@go build -mod=vendor -ldflags '-s -w' -o $(BUILDPATH)/build/bin/${API_NAME} cmd/main.go
	@echo "Binario generado en build/bin/${API_NAME}"

test:
	@echo "Ejecutando tests..."
	@go test ./... --coverprofile coverfile_out --covermode=atomic
	@go tool cover -func coverfile_out

coverage:
	@echo "Coverfile..."
	@go test ./... --coverprofile coverfile_out --covermode=atomic
	@go tool cover -func coverfile_out
	@go tool cover -html=coverfile_out -o coverfile_out.html

docker:
	@docker build . -t customer-debt-assesment:latest -f Dockerfile

	@#docker build . -t customer-debt-assesment:latest -f iaas/Dockerfile

.PHONY: test build
