package main

import (
	"fmt"

	"github.com/ShardNguyen/GolangCounter/counter"
)

func main() {
	var counter counter.Counter
	counter.ReadFile("test.txt")
	counter.ReadFile("test2.txt")
	counter.Add(1)
	fmt.Println(counter.StringSum())
}
