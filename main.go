package main

import (
	"sync"

	"github.com/ShardNguyen/GolangCounter/pkg/consumer"
	"github.com/ShardNguyen/GolangCounter/pkg/counter"
	"github.com/ShardNguyen/GolangCounter/pkg/producer"
)

func main() {
	// Counter
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

	// Producer and Consumer
	ch := make(chan int)
	pro := producer.NewProducer(ch)
	con := consumer.NewConsumer(ch)

	wg.Add(1)
	go func() {
		pro.Produce(counter.GetSum())
	}()
	go func() {
		defer wg.Done()
		con.Consume()
	}()

	wg.Wait()
}
