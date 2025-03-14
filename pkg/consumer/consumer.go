package consumer

type Consumer struct {
	ch chan int
}

func NewConsumer(ch chan int) *Consumer {
	consumer := new(Consumer)
	consumer.ch = ch
	return consumer
}

func (c *Consumer) Consume(f func(*int, int), sum *int) int {
	f(sum, <-c.ch)
	return *sum
}
