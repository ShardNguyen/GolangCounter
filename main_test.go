package main

import (
	"fmt"
	"testing"

	"github.com/ShardNguyen/GolangCounter/counter"
)

func TestMain(t *testing.T) {
	var counter counter.Counter
	counter.ReadFile("test.txt")
	fmt.Println(counter.StringSum())
}
