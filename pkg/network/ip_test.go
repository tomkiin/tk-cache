package network

import (
	"testing"
)

func TestGetLocalIP(t *testing.T) {
	if _, err := GetLocalIP(); err != nil {
		t.Fatal(err)
	}
}
