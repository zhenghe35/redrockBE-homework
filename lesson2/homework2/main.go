package main

import (
	"fmt"
	"sync"
	"time"
)

func download(filename string, wg *sync.WaitGroup, result chan<- string) {
	defer wg.Done()

	time.Sleep(time.Second)
	result <- filename + " 下载完成"
}

func main() {
	file := []string{"file.zip", "file.pdf", "file.mp4"}
	var wg sync.WaitGroup
	result := make(chan string, 3)
	for _, f := range file {
		wg.Add(1)
		go download(f, &wg, result)
	}
	go func() {
		wg.Wait()

		close(result)
	}()
	for msg := range result {
		fmt.Println(msg)
	}
	fmt.Println("所有文件完成")

}
