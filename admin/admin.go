package admin

import (
	"log"
	"os"
	"sync"
	"time"
)

var lock sync.Mutex

func Log(entry string) {
	lock.Lock()
	defer lock.Unlock()
	err := os.WriteFile("logs", []byte(time.Now().String()+":"+entry+"\n"), os.ModeAppend)
	checkError(err)
}

func init() {
	os.WriteFile("logs", []byte("Created At"+time.Now().String()), os.ModeAppend)
}

func checkError(err error) {
	if err != nil {
		log.Fatalf("Error: %s", err.Error())
		os.Exit(1)
	}
}
