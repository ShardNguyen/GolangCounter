package producer

type Producer struct {
	ch chan int
}

func NewProducer(ch chan int) *Producer {
	producer := new(Producer)
	producer.ch = ch
	return producer
}

func (p *Producer) Produce(value int) {
	p.ch <- value
}
