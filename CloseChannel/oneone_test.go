package CloseChannel

import (
	"sync"
	"fmt"
	"strconv"
	"time"
	"testing"
	"math/rand"
)

// 一个生产者 一个消费者
// 生产者直接关闭
func TestOneOne(t *testing.T) {
	tasks := make(chan string, 5)
	wg := sync.WaitGroup{}

	wg.Add(2)

	// produce
	go func() {
		defer wg.Done()
		for {
			randi := rand.Intn(5)
			if randi == 0 {
				fmt.Println("exit")
				close(tasks)
				break
			} else {
				task := strconv.Itoa(rand.Intn(1000))
				tasks <- task
				fmt.Println("produce task " + task)
			}
			time.Sleep(time.Millisecond * 800)
		}
	}()

	//consume
	go func() {
		defer wg.Done()
		for task := range tasks {
			fmt.Println("handle task: " + task)
		}
	}()

	wg.Wait()
}
