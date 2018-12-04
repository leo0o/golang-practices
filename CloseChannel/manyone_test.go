package CloseChannel

import (
	"sync"
	"strconv"
	"fmt"
	"time"
	"testing"
	"math/rand"
)

// 多个生产者 一个消费者
// 可以通过消费者关闭某个通知channel， 生产者同时也监听此channel，如果发现关闭则退出所有协程，当没有任何协程使用channel时，会自动关闭。
// 注意：这种情况会导致队列中还有消息未被消费完
func TestManyOne(t *testing.T) {

	taskChan := make(chan string, 5)
	exit := make(chan bool)
	wg := sync.WaitGroup{}

	//producer
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			for {
				task := strconv.Itoa(rand.Intn(20))
				select {
				case taskChan <- task:
					fmt.Println("produce task " + task)
				case <-exit:
					fmt.Println("producer" + strconv.Itoa(i) + " return")
					return
				}
				time.Sleep(time.Millisecond * 800)
			}
		}(i)
	}

	//consumer
	wg.Add(1)
	go func() {
		defer wg.Done()
		for task := range taskChan {
			if task == "1" {
				fmt.Println("exit..")
				close(exit)
				break
			}
			fmt.Println("handle task " + task)
		}
	}()

	wg.Wait()

	//for task := range taskChan {
	//	fmt.Println("rest task " + task)
	//}
}
