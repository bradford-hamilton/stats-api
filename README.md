<div align="center">
  <img
    alt="MLB logo"
    src="./assets/mlb-logo.png"
    height="200px"
  />
</div>
<h1 align="center">Welcome to the MLB stats api 👋</h1>
<p align="center">
  <a href="https://golang.org/dl" target="_blank">
    <img alt="Using go version 1.20" src="https://img.shields.io/badge/go-1.20-9cf.svg" />
  </a>
  <a href="https://goreportcard.com/report/github.com/bradford-hamilton/stats-api" target="_blank">
    <img alt="Go Report Card" src="https://goreportcard.com/badge/github.com/bradford-hamilton/stats-api" />
  </a>
  <a href="https://godoc.org/github.com/bradford-hamilton/stats-api" target="_blank">
    <img alt="godoc" src="https://godoc.org/github.com/bradford-hamilton/stats-api/pkg?status.svg" />
  </a>
  <a href="#" target="_blank">
    <img alt="License: MIT" src="https://img.shields.io/badge/License-MIT-yellow.svg" />
  </a>
</p>

### Run Application
```
go run cmd/server/main.go
```

### Build and Push Dockerfile
```
docker build -t bradfordhamilton/stats-api .
docker push bradfordhamilton/stats-api
```

### Usage
Currently the only supported API call is a GET to `/api/v1/schedule` which accepts two query params `date` and `teamID` where date is in `YYYY-MM-DD` format and teamID is a `number`/`int`.

#### Examples:

Local:
```
http://127.0.0.1:4000/api/v1/schedule?date=2022-07-21&teamID=147
```

Live Application:
```
http://134.209.123.132:4000/api/v1/schedule?date=2022-07-21&teamID=147
```
