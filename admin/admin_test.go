package admin

import (
	"os"
	"testing"
)

func TestAddLog(t *testing.T) {
	Log("Added Test entry")
	if _, err := os.Open("logs"); err != nil {
		t.Fatalf("Could not create logs file")
	}
}
