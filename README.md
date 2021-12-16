# Go Server
## _A genral purpose server built using golang_


[![Build Status](https://travis-ci.org/joemccann/dillinger.svg?branch=master)](https://travis-ci.org/joemccann/dillinger)
Go Server is a genral purpose server meant to be simple to use and begginner friendly.

Currently the project supports only tcp connections but support for more is to come!

## Features

- Import a HTML file and watch it magically convert to Markdown
- Drag and drop images (requires your Dropbox account be linked)
- Import and save files from GitHub, Dropbox, Google Drive and One Drive
- Drag and drop markdown and HTML files into Dillinger
- Export documents as Markdown, HTML and PDF

This project was built by a student attempting to understand basic networking concepts and the TCP IP networking model.

> Every contributer is welcome!



## Tech

Dillinger uses a number of open source projects to work properly:


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


