# <Network Monitoring>

## Description

The aim of this project is to learn and appreciate the feature golang has to provide.

It is a simple network monitoring system that monitors the registered ip addresses by ranking them based on their ping response speed

## Usage

Simply just run: `go run cmd/main.go config.json` to start the http server, it will listen at local port `5050`

To register an ip address
- get/post message to `http://localhost:5050/register`
- sample request body: {"ipaddress":"www.example.com"}
- sample response body: ok

To get based all registered ip addresses ranked in speed order
- get/post message to `http://localhost:5050/rank/networkspeed`
- sample request body: `""`
- sample response body: `[{"Addr":"www.youtube.com","AverageRtt":11},{"Addr":"www.google.com","AverageRtt":20},{"Addr":"cn.indeed.com","AverageRtt":20}]`