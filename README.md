# Moke Project Template

[![Go Report Card](https://goreportcard.com/badge/github.com/gstones/moke-layout)](https://goreportcard.com/report/github.com/gstones/moke-layout)
[![Go Reference](https://pkg.go.dev/badge/github.com/GStones/moke-layout.svg)](https://pkg.go.dev/github.com/GStones/moke-layout)
[![Release](https://img.shields.io/github/v/release/gstones/moke-layout.svg?style=flat-square)](https://github.com/GStones/moke-layout)

## Introduction

Moke is a template project for [moke-kit](https://github.com/GStones/moke-kit) that helps you quickly build gRPC/TCP/HTTP microservices with IoC (Inversion of Control). The template provides:

* Interactive command-line client
* Load testing tools
* Protocol buffer management:
  * Proto file generation
  * Proto file publishing to buf Schema Registry
  * SDK generation for multiple languages

## Quick Start

### Deploy Infrastructure
```shell
# Optional: Create ./deployment/docker-compose/.env for custom configuration
docker compose -f ./deployment/docker-compose/infrastructure.yaml up -d
```

### Run Service
```shell
go run ./cmd/demo/service/main.go
```

## Testing

### Integration Testing
1. Build the interactive client:
```shell
go build -o client.exe ./cmd/demo/client/main.go
```

2. Run the client:
```shell
# View available commands
./client.exe help

# Run specific clients
./client.exe grpc  # gRPC client
./client.exe tcp   # TCP client
```

Note: For HTTP testing, use Postman to connect to `localhost:8081`

### Load Testing
1. Install [k6](https://grafana.com/docs/k6/latest/get-started/installation/)
2. Run the tests:
```shell
k6 run ./tests/demo/demo.js
```

## Protocol Buffer Management

### Prerequisites
* Install [buf](https://buf.build/docs/installation)

### Working with Proto Files

1. Generate proto files:
```shell
buf generate
```

2. Publish to buf Schema Registry:
```shell
# Login to BSR (requires account)
# See: https://buf.build/docs/tutorials/getting-started-with-bsr#prerequisites
buf registry login username

# Push proto files
buf push
```

### Generate SDKs
1. Visit https://buf.build/gstones/moke-layout/sdks
2. Select your target language
3. Follow the provided instructions to import the SDK into your project

      