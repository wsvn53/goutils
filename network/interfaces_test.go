package network

import (
    "testing"
    "log"
)

func TestGetInterfaces(t *testing.T) {
    log.Println(GetInterfaces())
}