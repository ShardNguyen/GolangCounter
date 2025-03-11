package counter

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"sync"
)

type Counter struct {
	mu    sync.Mutex
	count int
	sum   int
}

func (counter *Counter) IncreaseCounter() {
	counter.mu.Lock()
	counter.count++
	counter.mu.Unlock()
}

func (counter *Counter) DecreaseCounter() {
	counter.count--
}

func (counter *Counter) checkZeroCount() bool {
	counter.mu.Lock()
	check := counter.count == 0
	counter.mu.Unlock()
	return check
}

func (counter *Counter) Add(value int) {
	counter.mu.Lock()
	counter.sum += value
	counter.DecreaseCounter()
	counter.mu.Unlock()
}

func (counter *Counter) StringCount() string {
	return fmt.Sprintf("%d", counter.count)
}

func (counter *Counter) StringSum() string {
	return fmt.Sprintf("%d", counter.sum)
}

func (counter *Counter) ReadFile(path string) {
	data, err := os.Open(path)

	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(data)

	// scanner.Scan() is basically reading line by line
	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text()) // Converting string to number
		counter.IncreaseCounter()
		go func() {
			if err != nil {
				panic(err)
			} else {
				counter.Add(num)
			}
		}()
	}

	for !counter.checkZeroCount() {

	}
}
