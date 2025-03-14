package main

import (
	"fmt"
	"sync"

	"github.com/ShardNguyen/GolangCounter/pkg/consumer"
	"github.com/ShardNguyen/GolangCounter/pkg/counter"
	"github.com/ShardNguyen/GolangCounter/pkg/producer"
)

// This is taken from the counter's method for the purpose of testing
func TestProCon(pro *producer.Producer, con *consumer.Consumer, count *counter.Counter) {
	var wg sync.WaitGroup

	for i := 0; i < 1001; i++ {
		wg.Add(1)
		go pro.Produce(i)
	}

	for i := 0; i < 1001; i++ {
		con.Consume(count.Add)
		wg.Done()
	}

	wg.Wait()
	fmt.Println(count.StringSum())
}
