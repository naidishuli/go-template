prepare-doc:
	swag init

# generate the swagger.json and a doc.html from it
doc: prepare-doc
	npx redoc-cli bundle -o docs/doc.html docs/swagger.json --options.nativeScrollbars

# quickly open the doc.html file in the browser
show-doc:
	open docs/doc.html

# run lint
lint:
	golangci-lint run --timeout 5m