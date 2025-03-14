package main

import (
	"fmt"
	"sync"

	"github.com/ShardNguyen/GolangCounter/pkg/consumer"
	"github.com/ShardNguyen/GolangCounter/pkg/producer"
)

func addToSum(sum *int, value int) {
	*sum += value
}

// This is taken from the counter's method for the purpose of testing
func TestProCon(pro producer.Producer, con consumer.Consumer) {
	sum := 0
	var wg sync.WaitGroup

	for i := 0; i < 1001; i++ {
		wg.Add(1)
		go pro.Produce(i)
	}

	var sumFunc func(*int, int)
	sumFunc = addToSum

	for i := 0; i < 1001; i++ {
		// Consumer will call a function so that function can execute
		con.Consume(sumFunc, &sum)
		wg.Done()
	}

	wg.Wait()
	fmt.Println(sum)
}
