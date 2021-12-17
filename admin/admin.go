package admin

import (
	"fmt"
	"log"
	"os"
	"sync"
	"time"
)

var lock sync.Mutex
var file *os.File
var err error

func startAdmin() {
	for {
		command := startCli()
		switch command {
		case "logs":
			readLogs()
		case "exit":
			os.Exit(0)
		}
	}
}
func startCli() string {
	var s string
	fmt.Printf("goserver->")
	fmt.Scan(&s)
	fmt.Println()
	return s
}
func readLogs() {
	var buf []byte
	n, err := file.Read(buf)
	fmt.Println("[*]Total log size: ", n)
	checkError(err)
	fmt.Print(string(buf[0:n]))
}
func Log(entry string) {
	lock.Lock()
	defer lock.Unlock()
	_, err := file.WriteString(time.Now().String() + entry + ":" + "\n")
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		log.Fatalf("Error: %s", err.Error())
		os.Exit(1)
	}
}
func init() {
	file, err = os.Create("logs")
	checkError(err)
	file.WriteString(time.Now().String() + "Created at:" + "\n")
	go startAdmin()
}
