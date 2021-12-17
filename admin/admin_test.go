package admin

import (
	"fmt"
	"os"
	"testing"
)

func TestAddLog(t *testing.T) {
	Log("Added test entry")
	for i := 0; i < 10; i++ {
		Log("Entry:" + fmt.Sprintf("%d", i))
	}
	if _, err := os.Open("logs"); err != nil {
		t.Fatalf("Could not create logs file")
	}
}
