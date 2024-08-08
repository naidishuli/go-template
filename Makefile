install-dep:
	go install github.com/swaggo/swag/cmd/swag@latest
	go install go.uber.org/mock/mockgen@latest
	npm install @redocly/cli -g

prepare-doc:
	swag init

# generate the swagger.json and a doc.html from it
doc: prepare-doc
	redocly build-docs docs/swagger.yaml -o docs/doc.html

# quickly open the doc.html file in the browser
show-doc:
	open docs/doc.html

# run lint
lint:
	golangci-lint run

# run tests & show total coverage at the end
coverage:
	go test ./... -coverpkg ./... -coverprofile docs/coverage.out
	go tool cover -func docs/coverage.out

.PHONY: migration
migration:
	go run cmd/main.go migration create $(filter-out $@,$(MAKECMDGOALS))
%:      # Do nothing recipe to avoid `make: *** No rule to make target 'create'.  Stop.`
	@:

.PHONY: migrate
migrate:
	go run cmd/main.go migration $(filter-out $@,$(MAKECMDGOALS))
%:      # Do nothing recipe to avoid `make: *** No rule to make target 'create'.  Stop.`
	@: