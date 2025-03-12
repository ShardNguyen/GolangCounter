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
	wg  sync.WaitGroup
	sum int
}

func (counter *Counter) Add(value int) {
	counter.mu.Lock()
	counter.sum += value
	counter.mu.Unlock()
}

func (counter *Counter) StringSum() string {
	return fmt.Sprintf("%d", counter.sum)
}

func (counter *Counter) ReadFile(path string) error {
	data, err := os.Open(path)

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

		counter.wg.Add(1)
		go func() {
			defer counter.wg.Done()
			counter.Add(num)
		}()
	}
	counter.wg.Wait()
	return nil
}
