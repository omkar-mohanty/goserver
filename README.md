# Go Server
## _A genral purpose server built using golang_


[![Build Status](https://travis-ci.org/joemccann/dillinger.svg?branch=master)](https://travis-ci.org/joemccann/dillinger)
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


