package counter

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"
)

type Counter struct {
	mu    sync.Mutex
	count int
	sum   int
}

func (counter *Counter) Add(value int) {
	counter.mu.Lock()
	counter.count++
	counter.sum += value
	counter.mu.Unlock()
}

func (counter *Counter) String() string {
	return fmt.Sprintf("%d", counter.count)
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
		fmt.Println(num)
		go func() {
			if err != nil {
				panic(err)
			} else {
				counter.Add(num)
			}
		}()
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	// fmt.Printf("File content: %s\n", data)
}
