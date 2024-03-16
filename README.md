# Autotrader

[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT)
![Go](https://img.shields.io/github/go-mod/go-version/noirstar/autotrader?filename=backend%252Fgo.mod)
![Vue](https://img.shields.io/npm/v/vue)

> 업비트 시세 수신 및 오토트레이딩 (개발중지)

## Introduce

A Golang HTTP server / Vue.js web application for autotrading
 
**Build info and requirements**

Built using Go 1.16.3 and Vue CLI 3

Build tools used:
- Vue CLI 
- NPM
- Golang 1.16+
- MongoDB Cloud

**List of URL routes**

Backend

URL Path | Request Type |Purpose
:-----:|:-----: |:-----:
`/candles`|GET|Get candle data
`/coins`|GET|Get coin info
`/signup`|POST|Register User
`/login`|POST|Login User

Frontend

URL Path | Request Type |Purpose
:-----:|:-----: |:-----:
`/`|GET|Main frontend


## Installation

Windows:

```powershell
cd ./frontend
npm init

cd ../backend
go mod tidy
```

## Usage example

```powershell
cd ./frontend
npm run serve

cd ../backend
go run main.go
```

## Part of WebPage

![1](https://user-images.githubusercontent.com/44075494/124548553-dbc4b980-de68-11eb-83b0-25e8a69c0e48.png)
