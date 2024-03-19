# Moke Project Template
[![Go Report Card](https://goreportcard.com/badge/github.com/gstones/moke-layout)](https://goreportcard.com/report/github.com/gstones/moke-layout)
[![Go Reference](https://pkg.go.dev/badge/github.com/GStones/moke-layout.svg)](https://pkg.go.dev/github.com/GStones/moke-layout)
[![Release](https://img.shields.io/github/v/release/gstones/moke-layout.svg?style=flat-square)](https://github.com/GStones/moke-layout)

## Create a service
```shell
    # install gonew
    go install golang.org/x/tools/cmd/gonew@latest
    # create a new project
    gonew github.com/gstones/moke-layout your.domain/myprog
```

## How to run
* deploy infrastructure:
  ```shell
   # first need fix .env file
   docker compose -f ./deployment/docker-compose/infrastructure.yaml up -d
  ```

* run service:
   ```shell
     go run cmd/demo/service/main.go
   ```
* build client:
   ```shell
   go build -o client.exe ./cmd/client/main.go 
   ```
* run client:
    ```shell
     # help
     ./client.exe help
      # run grpc client
     ./client.exe grpc
      # run tcp client
      ./client.exe tcp
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

      