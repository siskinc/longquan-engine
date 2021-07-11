build: init
	go build -tags=jsoniter .
run:
	go run main.go
.PHONY: swagger, packr2, init
swagger:
	swag init --parseDependency --output swagger_docs
packr2:
	packr2.exe build
init: swagger packr2