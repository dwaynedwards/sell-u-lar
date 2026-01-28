.PHONY: dev-web
dev-web:
	@cd web; make dev

.PHONY: dev-products
dev-products:
	@cd products; make dev

.PHONY: gen
gen:
	@cd web; make gen

.PHONY: css-build
css-build:
	@cd web; make css-build

.PHONY: tidy
tidy:
	@go mod tidy

.PHONY: gen-products
gen-products:
	@protoc -I ./pkg/proto/products \
		--go_out ./pkg/proto/products --go_opt paths=source_relative \
		--go-grpc_out ./pkg/proto/products --go-grpc_opt paths=source_relative \
		./pkg/proto/products/products.proto
