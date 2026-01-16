.PHONY: fmt
fmt:
	@templ fmt .

.PHONY: gen
gen: fmt
	@templ generate -path="."

.PHONY: dev
dev:
	@air

.PHONY: css-build
css-build:
	@npx --yes tailwindcss -i ./static/css/styles.css -o ./static/css/tailwind.css

.PHONY: css-build-minify
css-build-minify:
	@npx --yes tailwindcss -i ./static/css/styles.css -o ./static/css/tailwind.css --minify

.PHONY: css-install
css-install:
	@npm install -D tailwindcss @tailwindcss/cli