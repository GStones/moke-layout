# Demo

## How to run

* deploy infrastructure:
    * redis:6379
    * nats:4222
    * mongo:27017

  ```shell
   # first need fix .env file
   docker compose -f ./deployment/docker-compose/infrastructure.yaml up -d
  ```

* service:
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

* proto generate 
  * install [buf](https://buf.build/docs/installation)
  ```shell
  buf generate
  ```

      