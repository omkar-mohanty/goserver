# Go Server
## _A genral purpose server built using golang_

[![GitHub license](https://img.shields.io/github/license/Naereen/StrapDown.js.svg)](https://github.com/Naereen/StrapDown.js/blob/master/LICENSE)[![made-with-Go](https://img.shields.io/badge/Made%20with-Go-1f425f.svg)](http://golang.org)[![Windows](https://svgshare.com/i/ZhY.svg)](https://svgshare.com/i/ZhY.svg) [![GitHub go.mod Go version of a Go module](https://img.shields.io/github/go-mod/go-version/gomods/athens.svg)](https://github.com/gomods/athens)

Go Server is a genral purpose server meant to be simple to use and begginner friendly.

Currently the project supports only tcp connections but support for more is to come!

## Features

- TCP Connections
- Good for anyone attempting to understand networking concepts

This project was built by a student attempting to understand basic networking concepts and the TCP IP networking model.

> Every contributer is welcome!



## Tech

Go Server uses a number of open source projects to work properly:


- [Cobra ](https://github.com/spf13/cobra) - A CLI library
- [Golang standard library](https://pkg.go.dev/std)-The standard golang library



## Installation

Go Server requires [golang](https://go.dev/) v1.17+ to run.

Install the dependencies and devDependencies and start the server.

```sh
cd goserver
go mod get ./...
go run main.go 1200
```


## License

MIT


