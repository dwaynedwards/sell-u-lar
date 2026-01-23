.PHONY: dev-web
dev-web:
	cd web; make dev

.PHONY: dev-products
dev-products:
	cd products; make dev

.PHONY: gen
gen:
	cd web; make gen

.PHONY: css-build
css-build:
	cd web; make css-build

