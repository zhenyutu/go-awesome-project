package router

import (
	"fmt"
	"testing"
)

func TestInsert(test *testing.T) {
	t := &Tire{}
	t.insert("/hello/world")
	if res := t.search("/hello"); res == nil {
		test.Errorf("node not found")
	}
	fmt.Println()
}

func TestInsert2(test *testing.T) {
	t := &Tire{}
	t.insert("/hello/world")
	if res := t.search("/hello/world"); res == nil {
		test.Errorf("node not found")
	}
	fmt.Println()
}

func TestFuzzyMatching(test *testing.T) {
	t := &Tire{}
	t.insert("/hello/*/world")
	if res := t.search("/hello/world/world"); res == nil {
		test.Errorf("node not found")
	}
	fmt.Println()
}
