package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	mu    sync.Mutex
	count int
}

func (c *Counter) INcrement() {
	c.mu.Lock()
	c.count++
	defer c.mu.Unlock()
}
func (c *Counter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.count
}
func main() {
	c := &Counter{}
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			fmt.Println("Goroutine:", id)
			for j := 0; j < 10; j++ {
				c.INcrement()
			}
		}(i)
	}
	wg.Wait()
	fmt.Println("最终计数", c.Value())

}
