package main

import (
	"testing"

	"github.com/ShardNguyen/GolangCounter/pkg/consumer"
	"github.com/ShardNguyen/GolangCounter/pkg/counter"
	"github.com/ShardNguyen/GolangCounter/pkg/producer"
)

// This is taken from the counter's method for the purpose of testing
func TestMain(t *testing.T) {
	// Counter
	counter := counter.NewCounter()
	/*
		var wg sync.WaitGroup

		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.ReadFile("test.txt")
		}()
		wg.Wait()
	*/

	// Producer and Consumer
	ch := make(chan int)
	pro := producer.NewProducer(ch)
	con := consumer.NewConsumer(ch)

	TestProCon(pro, con, counter)
}
