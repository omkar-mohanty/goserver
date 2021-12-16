package server

import (
	"fmt"
	"os"
	"time"
)

var logFile *os.File

func init() {
	logFile, err := os.Create("logs.txt")
	checkError(err)
	logFile.WriteString("Created at " + time.Now().String())
}
func addLogEntry(entry string) {
	b := []byte(entry + time.Now().String() + "\n")
	logFile.Write(b)
}

func readLogs() {
	buf := make([]byte, 1024)
	n, err := logFile.Read(buf)
	checkError(err)
	fmt.Printf("Read %d data\n", n)
	fmt.Print(string(buf[:]))
	fmt.Println()
}
