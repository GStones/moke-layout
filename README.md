# Moke Project Template
[![Go Report Card](https://goreportcard.com/badge/github.com/gstones/moke-layout)](https://goreportcard.com/report/github.com/gstones/moke-layout)
[![Go Reference](https://pkg.go.dev/badge/github.com/GStones/moke-layout.svg)](https://pkg.go.dev/github.com/GStones/moke-layout)
[![Release](https://img.shields.io/github/v/release/gstones/moke-layout.svg?style=flat-square)](https://github.com/GStones/moke-layout)

## How to run
* deploy infrastructure:
  ```shell
   # add ./deployment/docker-compose/.env file to custom your environment
   docker compose -f ./deployment/docker-compose/infrastructure.yaml up -d
  ```

*  service:
   ```shell
     go run ./cmd/demo/service/main.go
   ```
* build client:
   ```shell
     go build -o client.exe ./cmd/demo/client/main.go 
   ```
* run client:
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
* run load test:
   * install [k6](https://grafana.com/docs/k6/latest/get-started/installation/)
   * run 
      ``` shell
          k6 run ./tests/demo/demo.js
       ```
* proto generate 
  * install [buf](https://buf.build/docs/installation)
  ```shell
    buf generate
  ```

      