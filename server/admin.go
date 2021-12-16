package server

import (
	"fmt"
	"os"
)

func StartAdmin() {
	for {
		command := startCli()
		execCommand(command)
	}
}

func startCli() string {
	var command string
	fmt.Printf("goserver-> ")
	fmt.Scan(&command)
	fmt.Println()
	return command
}
func execCommand(command string) {
	switch command {
	case "logs":
		readLogs()
	case "exit":
		os.Exit(0)
	default:
		return
	}
}
