# Healthcheck My PC

## Features

- [X] Get processes by process name
- [X] Get information each process such as CPU, Memory, Network, Uptime
- [X] Run server inside docker
- [ ] Zip data before send
- [ ] Encrypt with RSA

## How to get only code coverage

```sh
go test ./... -cover
```

## How to get code coverage for create html

```sh
go test ./... -cover -coverprofile=cover.out
go tool cover -html=cover.out
```

## How to build client

```sh
go build -o client client/cmd/main.go
```

## How to build server

```sh
docker build -t healthcheck-my-pc:0.0.1 .
```

## How to run server

```sh
cat > config/settings.json
{
    "min": <interval_time(number)>,
    "url": "http://<ip(string)>:<port(number)>/healthcheck",
    "target": "<app_target(string)>"
}

docker run --name healthcheck-my-pc-server -p <port>:5000 healthcheck-my-pc:0.0.1
```
