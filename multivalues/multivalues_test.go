package multivalues

import (
	"log"
	"testing"
)

func multiValues(a int, b string) (int, string) {
	return a, b
}

func TestFirst(t *testing.T) {
	var a = 1
	var b = "abcd"
	if First(multiValues(a, b)) != a {
		t.Fatal("first multivalue should be", a)
	}
}

func TestLast(t *testing.T) {
	var (
		a = 1
		b = "abcd"
	)
	if Last(multiValues(a, b)) != b {
		t.Fatal("last multivalues should be", b)
	}
}

func TestToSlices(t *testing.T) {
	var (
		a = 1
		b = "abcd"
	)
	if len(ToSlices(multiValues(a, b))) != 2 {
		t.Fatal("multivalues should have", 2, "elements")
	}
	log.Println(ToSlices(multiValues(a, b)))
}
