package slices

import (
	"log"
	"testing"
)

func TestRemove(t *testing.T) {
	type TestInterface struct {
		A string
		B string
	}
	s := []TestInterface{
		TestInterface{
			A: "A0",
			B: "B0",
		},
		TestInterface{
			A: "A1",
			B: "B1",
		},
		TestInterface{
			A: "A2",
			B: "B2",
		},
	}
	log.Println(Remove(s, 1))
	log.Println(Remove([]string{
		"a", "b", "c", "d",
	}, 1))
	log.Println(Remove([]string{
		"a", "b", "c", "d",
	}, 0))
}
