- In order to create/generate swagger docs from the go code we need to use\
  https://github.com/swaggo/swag

- There is a makefile in the project for quick scripts\

- We use https://github.com/golangci/golangci-lint for lint -> install it locally https://golangci-lint.run/usage/install/  \

- TODO more on golangci-lint linter rules



// installs

go install github.com/swaggo/swag/cmd/swag@latest

go install go.uber.org/mock/mockgen@latest

go install github.com/cosmtrek/air@latest

go install github.com/pressly/goose/v3/cmd/goose@latest

yarn add global @redocly/cli
