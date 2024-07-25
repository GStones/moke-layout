# Moke Project Template

[![Go Report Card](https://goreportcard.com/badge/github.com/gstones/moke-layout)](https://goreportcard.com/report/github.com/gstones/moke-layout)
[![Go Reference](https://pkg.go.dev/badge/github.com/GStones/moke-layout.svg)](https://pkg.go.dev/github.com/GStones/moke-layout)
[![Release](https://img.shields.io/github/v/release/gstones/moke-layout.svg?style=flat-square)](https://github.com/GStones/moke-layout)

## Introduction

Moke is a template project for [moke-kit](https://github.com/GStones/moke-kit) to help you quickly build a gRPC/TCP/HTTP
Microservice IOC project.
It provides the following tools:

* Interact command line client
* Load test tool
* Proto file management: proto file generation, proto file push to buf Schema Registry, and SDKS generation
  for different languages.

## How to run

* deploy infrastructure:
  ```shell
   # add ./deployment/docker-compose/.env file to custom your environment if you have
   docker compose -f ./deployment/docker-compose/infrastructure.yaml up -d
  ```

* run service:
  ```shell
    go run ./cmd/demo/service/main.go
  ```

## How to test

### Integration Test

* build your interactive client:
   ```shell
     go build -o client.exe ./cmd/demo/client/main.go 
   ```
* run your interactive client:
    ```shell
     # help
     ./client.exe help
      # run grpc client
     ./client.exe grpc
      # run tcp client
     ./client.exe tcp
     # run help to see more command options. and follow the tips to run.
    ```
  tips: http client use Postman to connect `localhost:8081`.

### Load Test

* install [k6](https://grafana.com/docs/k6/latest/get-started/installation/)
* run k6 load test
   ``` shell
       k6 run ./tests/demo/demo.js
    ```

## Proto file Manage

* install [buf](https://buf.build/docs/installation)

* manage proto file
  ```shell
   #  generate proto file
    buf generate
  ```
  ```shell
   # use buf Schema Registry to manage proto file
   # you need to sign up and login to buf Schema Registry,follow the steps below:
   # https://buf.build/docs/tutorials/getting-started-with-bsr#prerequisites
    buf registry login username 
   # push proto file to buf Schema Registry
    buf push
  ```
* generate SDKS for different languages
    * visit https://buf.build/gstones/moke-layout/sdks
    * choose the language you want to generate, and follow the cmd to import the SDKS to your project.
  

      