prepare-doc:
	swag init

doc: prepare-doc
	npx redoc-cli bundle -o docs/doc.html docs/swagger.json --options.nativeScrollbars

show-doc:
	open docs/doc.html