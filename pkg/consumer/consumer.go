package consumer

import "fmt"

type Consumer struct {
	ch    chan int
	value int
}

func NewConsumer(ch chan int) *Consumer {
	consumer := new(Consumer)
	consumer.ch = ch
	consumer.value = 0
	return consumer
}

func (c *Consumer) Consume() {
	c.value = <-c.ch
	fmt.Println("Consumed", c.value)
}
