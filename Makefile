ARTIFACT_NAME := weatherman

build:
	@go build -o bin/${ARTIFACT_NAME}/${ARTIFACT_NAME} cmd/${ARTIFACT_NAME}/main.go 

run:
	@go run cmd/${ARTIFACT_NAME}/main.go 

go-test:
	@go test -v $(shell go list ./... | grep -v /test/)

go-test-with-cover:
	@go test -coverprofile cover.out -v $(shell go list ./... | grep -v /test/)
	@go tool cover -html=cover.out

generate-mocks:
	@mockery --all --with-expecter --keeptree
	


