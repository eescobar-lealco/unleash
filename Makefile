run:
	go run cmd/main.go

test:
	go test ./... -cover -coverprofile=coverage.out

gen-mocks:
	mockery --dir=./internal --all

# tidies dependencies
tidy-deps:
	go mod tidy

# vendors dependencies
deps:
	go mod tidy
	go mod vendor

# formats project with go's style guidelines
fmt:
	gofmt -w -s ./src

# checks the project for possible errors
check:
	staticcheck ./...

build-docker:
	docker build -f zarf/docker/Dockerfile -t orders .

gen-protos:
	protoc --go_out=. --go_opt=paths=source_relative ./internal/pkg/pb/orders.proto
