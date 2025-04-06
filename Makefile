# Makefile: Simplify common developer tasks

# 1. Generate GraphQL code
#    This command runs gqlgen to refresh/update generated files.
gqlgen_setup:
	@dir_name=$$(basename $$PWD)
	@go mod init github.com/briankscheong/$$dir_name
	@printf '//go:build tools\npackage tools\nimport (_ "github.com/99designs/gqlgen"\n _ "github.com/99designs/gqlgen/graphql/introspection")' | gofmt > tools.go
	@go mod tidy
	@go run github.com/99designs/gqlgen init
	@go mod tidy

gql_generate:
	@go mod tidy
	@go generate ./...

# 2. Generate protobuf/gRPC Go code
#    This requires protoc to be installed (along with the protoc-gen-go and protoc-gen-go-grpc plugins).
protoc_install:
	brew install protobuf

protoc_generate: protoc_setup
	@echo "Generating protobuf code..."
	@protoc --go_out=. --go_opt=paths=source_relative \
		--go-grpc_out=. --go-grpc_opt=paths=source_relative \
		proto/service.proto

protoc_setup:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

# 3. Build the Go server
build:
	@echo "Building Go server..."
	@go build server.go

# 4. Run the application directly
run: build
	@echo "Running the Go application..."
	@./server

# 5. Clean built artifacts
clean:
	@echo "Cleaning up..."
	@rm -f server
	@go clean
	@go clean -modcache

# Optional: Build container
build_container:
	@docker build . --tag go-graphql-gateway:latest

run_container: build_container
	@docker run -it -p 8080:8080 go-graphql-gateway:latest