# Healthcheck My PC

## Features

- [X] Get Processes By Process Name
- [X] Get Information Each Process Such As CPU, Memory, Network, Uptime
- [X] Get PC Uptimes
- [ ] Run Inside Docker
- [ ] Zip Data Before Send
- [ ] Encrypt with RSA

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
