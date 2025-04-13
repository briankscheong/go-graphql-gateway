#
#
# - gqlgen targets
###############################
MODULE_NAME?=github.com/briankscheong
gqlgen_setup:
	@dir_name=$$(basename $$PWD)
	@go mod init $(MODULE_NAME)/$$dir_name
	@printf '//go:build tools\npackage tools\nimport (_ "github.com/99designs/gqlgen"\n _ "github.com/99designs/gqlgen/graphql/introspection")' | gofmt > tools.go
	@go mod tidy
	@go run github.com/99designs/gqlgen init
	@go mod tidy

gql_generate:
	@go mod tidy
	@go generate ./...

#
#
# - protoc targets
###############################
protoc_install:
	@brew install protobuf

protoc_generate: protoc_setup
	@echo "Generating protobuf code..."
	@protoc --go_out=. --go_opt=paths=source_relative \
		--go-grpc_out=. --go-grpc_opt=paths=source_relative \
		proto/service.proto

protoc_setup:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

#
#
# - go targets
###############################
build:
	@echo "Building Go server..."
	@go build server.go

run: build
	@echo "Running the Go application..."
	@./server

clean:
	@echo "Cleaning up..."
	@rm -f server
	@go clean
	@go clean -modcache

#
#
# - docker targets
###############################
build_container:
	@docker build . --tag go-graphql-gateway:latest

run_container: build_container
	@docker run -it -p 8080:8080 go-graphql-gateway:latest

#
#
# - brew targets
###############################
brew_k3d_setup:
	brew install k3d
	brew install kubectl
	brew install k9s

brew_k3d_cleanup:
	brew uninstall k3d
	brew uninstall kubectl
	brew uninstall k9s

#
#
# - k3d targets
###############################
k3d_setup:
	@echo "Installing k3d CLI..."
	curl -s https://raw.githubusercontent.com/k3d-io/k3d/main/install.sh | TAG=v5.8.3 bash

kubectl_setup:
	@echo "Installing kubectl..."
	curl -LO "https://dl.k8s.io/release/$(curl -L -s https://dl.k8s.io/release/stable.txt)/bin/linux/arm64/kubectl"
	sudo install -o root -g root -m 0755 kubectl /usr/local/bin/kubectl

k3d_setup: k3d_setup kubectl_setup

k3d_cleanup:
	sudo rm $(where kubectl)
	sudo rm $(where k3d)


# https://github.com/k3d-io/k3d/issues/1449#issuecomment-2154672702
K3D_FIX_DNS=0
CLUSTER_NAME?=cluster-one
REGISTRY_PORT?=5432
KUBECONFIG_PATH=$(HOME)/.kube/config
create_k3d_cluster:
	@echo "Creating k3d cluster..."
	K3D_FIX_DNS=0 k3d cluster create $(CLUSTER_NAME) --servers 3 --agents 3 --api-port 0.0.0.0:6550 --registry-create $(CLUSTER_NAME)-registry:0.0.0.0:$(REGISTRY_PORT) --wait --timeout 120s 

delete_k3d_cluster:
	@echo "Deleting k3d cluster..."
	k3d cluster delete $(CLUSTER_NAME)

add_kubeconfig:
	k3d kubeconfig merge $(CLUSTER_NAME) --kubeconfig-switch-context