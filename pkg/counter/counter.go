package counter

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"sync"
)

type Counter struct {
	mu  sync.Mutex
	sum int
}

func (counter *Counter) Add(value int) {
	counter.mu.Lock()
	counter.sum += value
	counter.mu.Unlock()
}

func (counter *Counter) GetSum() int {
	return counter.sum
}

func (counter *Counter) StringSum() string {
	return fmt.Sprintf("%d", counter.sum)
}

func (counter *Counter) ReadFile(path string) error {
	data, err := os.Open(path)
	var wg sync.WaitGroup

	if err != nil {
		return err
	}

	scanner := bufio.NewScanner(data)

	// scanner.Scan() is basically reading line by line
	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text()) // Converting string to number
		if err != nil {
			continue
		}

		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.Add(num)
		}()
	}

	wg.Wait()
	return nil
}
