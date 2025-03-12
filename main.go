package main

import (
	"fmt"
	"sync"

	"github.com/ShardNguyen/GolangCounter/counter"
)

func main() {
	var counter counter.Counter
	var wg sync.WaitGroup

	wg.Add(3)
	go func() {
		defer wg.Done()
		counter.ReadFile("test.txt")
	}()

	go func() {
		defer wg.Done()
		counter.ReadFile("test2.txt")
	}()

	go func() {
		defer wg.Done()
		counter.Add(1)
	}()

	wg.Wait()
	fmt.Println(counter.StringSum())
}
