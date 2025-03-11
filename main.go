package main

import (
	"fmt"

	"github.com/ShardNguyen/GolangCounter/counter"
)

func main() {
	var counter counter.Counter
	counter.ReadFile("test.txt")
	fmt.Println(counter.StringSum())
}
