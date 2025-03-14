package main

import (
	"fmt"
	"sync"

	"github.com/ShardNguyen/GolangCounter/pkg/consumer"
	"github.com/ShardNguyen/GolangCounter/pkg/producer"
)

// This is taken from the counter's method for the purpose of testing
func TestProCon(pro producer.Producer, con consumer.Consumer) {
	sum := 0
	var wg sync.WaitGroup
	var mu sync.Mutex

	for i := 0; i < 101; i++ {
		wg.Add(1)
		go pro.Produce(i)
	}

	for i := 0; i < 101; i++ {
		go func() {
			defer wg.Done()
			mu.Lock()
			sum += con.Consume()
			mu.Unlock()
		}()
	}

	wg.Wait()
	fmt.Println(sum)
}
