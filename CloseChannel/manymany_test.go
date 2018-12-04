package CloseChannel

import (
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"testing"
	"time"
)

//多个生产者 多个消费者
//由于有多个消费者  所以无法像单个消费者那样直接关闭某个通知channel
//这里由一个中立协程去关闭这个通知的channel
//注意：这里也不能保证队列中的消息全部被消费完
func TestGraceful(t *testing.T) {
	taskChan := make(chan string, 5)
	stopper := make(chan string)
	exit := make(chan bool)
	wg := sync.WaitGroup{}

	//中立协程
	go func() {
		who := <-stopper
		fmt.Println("stopped by " + who)
		close(exit)
	}()

	// producer
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for {

				task := strconv.Itoa(rand.Intn(100))
				//模拟生产者关闭条件
				if task == "99" {
					select {
					case stopper <- "producer#" + strconv.Itoa(id):
					default:
					}
					return
				}

				select {
				case <-exit:
					return
				default:
				}

				select {
				case <-exit:
					return
				case taskChan <- task:
					fmt.Println("produce " + task)
				}
			}

		}(i)
	}

	//consumer
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for {
				select {
				case <-exit:
					return
				case task := <-taskChan:
					//模拟消费者遇到关闭条件
					if task == "0" {
						select {
						case stopper <- "consumer#" + strconv.Itoa(id):
						default:
						}
						return
					}
					fmt.Println("consume " + task)
				}
			}
		}(i)
	}

	wg.Wait()

}



// 这种能保证channel中的信息都被消费完
func TestCloseWithOnce(t *testing.T) {

	taskChan := make(chan string, 5)
	stop := make(chan bool)
	wg := sync.WaitGroup{}
	once := sync.Once{}

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			j := 0
			for {
				select {
				case <-stop:
					once.Do(func() {
						close(taskChan)
					})
					return
				default:
					task := strconv.Itoa(id) + "-" + strconv.Itoa(j)
					taskChan <- task
					time.Sleep(time.Millisecond * 500)
				}
				j++
			}
		}(i)
	}

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for task := range taskChan {
				fmt.Println("consumer#" + strconv.Itoa(id) + "=>" + task)
			}
		}(i)
	}

	go func() {
		time.Sleep(time.Second * 3)
		close(stop)
	}()

	wg.Wait()

}
