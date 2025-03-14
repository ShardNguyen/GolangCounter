package main

import (
	"github.com/ShardNguyen/GolangCounter/pkg/consumer"
	"github.com/ShardNguyen/GolangCounter/pkg/counter"
	"github.com/ShardNguyen/GolangCounter/pkg/producer"
)

func main() {
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
