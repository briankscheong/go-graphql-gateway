# go-graphql-gateway

**go-graphql-gateway** is a Go-based GraphQL gateway that serves as an intermediary layer to various backend services, including RESTful APIs and gRPC services. This project leverages Go, GraphQL (via gqlgen), and integrates Kubernetes, providing a scalable and flexible solution for microservice architectures.

## Table of Contents

- [go-graphql-gateway](#go-graphql-gateway)
  - [Table of Contents](#table-of-contents)
  - [Overview](#overview)
  - [Features](#features)
  - [Tech Stack](#tech-stack)
  - [Project Structure](#project-structure)
  - [Installation and Setup](#installation-and-setup)
    - [Prerequisites](#prerequisites)
    - [Clone the Repository](#clone-the-repository)
    - [Install Dependencies](#install-dependencies)
  - [Running the Server in a Cluster](#running-the-server-in-a-cluster)
    - [Setup a k3d cluster](#setup-a-k3d-cluster)
    - [Build and Push Container Image](#build-and-push-container-image)
    - [Deploy GraphQL Server to k3d Cluster](#deploy-graphql-server-to-k3d-cluster)
  - [Running the Server Locally](#running-the-server-locally)
    - [Build the Go Server](#build-the-go-server)
    - [Run the Server](#run-the-server)
    - [Docker Setup](#docker-setup)
  - [Development Workflow](#development-workflow)
    - [Generating GraphQL Code](#generating-graphql-code)
    - [Cleaning Up](#cleaning-up)
  - [Contributing](#contributing)
    - [Guidelines](#guidelines)

---

## Overview

The **go-graphql-gateway** is a powerful and flexible API gateway built using Go. It integrates multiple backend services, including GraphQL, REST, and gRPC APIs, providing a seamless way for clients to query data from various sources in a single, unified interface. 🚀

This project is designed with scalability and performance in mind, offering the ability to serve high-traffic applications efficiently. With a strong focus on Kubernetes-based deployments, the gateway ensures smooth integration with microservices architectures, making it ideal for cloud-native environments. 🌐

With this gateway, you get the flexibility to query data from different backend services, ensuring a more cohesive and easy-to-manage API layer for your applications.

---

## Features

- 🖥️ **GraphQL Gateway**: Acts as a gateway for GraphQL queries, combining multiple backends.
- 🚀 **Go-based**: Written in Go for performance, scalability, and easy integration.
- 🐳 **Docker-Ready**: Docker configuration for easy containerization.

---

## Tech Stack

- 🏗️ **Go**: The backend language for building the API gateway.
- 🔗 **GraphQL (gqlgen)**: The library used for generating GraphQL schemas and resolvers.
- 🐳 **Docker**: Containerization for easy deployment.
- 🛠️ **Makefile**: To automate build and run tasks.

---

## Project Structure

```md
.
├─ proto/
│   └─ service.proto         - gRPC service definitions (optional)
├── main.go
├── Dockerfile
├── Makefile
├── README.md
├── go.mod
├── go.sum
├── gqlgen.yml               - The gqlgen config file, knobs for controlling the generated code.
├── graph
│   ├── generated            - A package that only contains the generated runtime
│   │   └── generated.go
│   ├── model                - A package for all your graph models, generated or otherwise
│   │   └── models_gen.go
│   ├── resolver.go          - The root graph resolver type. This file wont get regenerated
│   ├── schema.graphqls      - Some schema. You can split the schema into as many graphql files as you like
│   └── schema.resolvers.go  - The resolver implementation for schema.graphql
└── server.go                - The entry point to your app. Customize it however you see fit
```

---

## Installation and Setup

### Prerequisites

Before running the project, ensure the following tools are installed:

- **Go (v1.20+):** The Go programming language to build and run the project.
- **Protobuf Compiler (protoc):** For generating protobuf files and gRPC code.
- **gqlgen:** For generating GraphQL server code.

You can install these tools with the following commands:

- [Install Go](https://golang.org/doc/install)
- [Install Protobuf Compiler](https://grpc.io/docs/protoc-installation/)
- [Install gqlgen](https://github.com/99designs/gqlgen)

### Clone the Repository

Clone the project repository to your local machine:

```bash
git clone https://github.com/briankscheong/go-graphql-gateway.git
cd go-graphql-gateway
```

### Install Dependencies

Run the following commands to install necessary dependencies:

```bash
go mod tidy
```

This command will download and install any missing dependencies.

```bash
make gql_generate
```

This command will generate the Go model schema using the GraphQL schema definition.

---

## Running the Server in a Cluster

### Setup a k3d cluster

To setup a simple k3d cluster with a managed registry, run the following command:

```bash
make create_k3d_cluster
```

### Build and Push Container Image

To build and push the container image to the managed registry, run:

```bash
make build_push_image
```

This command will setup the container in the registry and automatically update the `deployment.yaml` file with the new image tag.

### Deploy GraphQL Server to k3d Cluster

To deploy the server to the cluster, run:

```bash
make deploy_server
```

This will deploy all the Kubernetes configurations and RBAC permissions needed for the server to be up and running.

**You can now port forward the server to your localhost and try it out!**

## Running the Server Locally

[!NOTE]
Building and running the server outside of a Kubernetes cluster will result in limited functionality.
Features that require interaction with the cluster (such as querying pods, deployments, or services) will not work unless the server is deployed inside the cluster.

Kindly refer to the deployment steps above to ensure the server is properly running within your Kubernetes environment for full functionality.

### Build the Go Server

To build the Go server, run the following Make command:

```bash
make build
```

This will compile the Go server into an executable named server.

### Run the Server

To run the server locally, use the run target from the Makefile. This will build the server and then run it:

```bash
make run
```

The server will start, and you can access the GraphQL endpoint on [http://localhost:8080](http://localhost:8080).

### Docker Setup

If you'd like to run the server in a container, you can build and run it using Docker.

To build the Docker container:

```bash
make build_container
```

To run the Docker container:

```bash
make run_container
```

This will start the server inside a Docker container and expose port 8080 to the host machine.

## Development Workflow

### Generating GraphQL Code

To refresh or generate GraphQL code, run the following command:

```bash
make gql_generate
```

This will update the generated GraphQL schema and resolvers based on your defined GraphQL types.

### Cleaning Up

To clean up build artifacts (e.g., executables), run:

```bash
make clean
```

This will remove the generated server binary.

## Contributing

We welcome contributions to improve this project! If you'd like to contribute:

```md
1. Fork the repository.
2. Create a new branch (git checkout -b feature-name).
3. Make your changes.
4. Commit your changes (git commit -am 'Add new feature').
5. Push to the branch (git push origin feature-name).
6. Open a pull request.
7. Please ensure your code follows the Go idioms and includes tests where applicable.
```

### Guidelines

- Ensure your code is properly formatted using `go fmt`.
- Provide detailed explanations in your pull request.
- Follow the project’s code conventions.
- Please include tests for new features or bug fixes.
- If your changes affect the documentation, update the README.

We appreciate your contributions, and thank you for improving **go-graphql-gateway**!
