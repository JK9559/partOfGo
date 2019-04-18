package partChannel

import (
	"fmt"
	"time"
)

func get01() chan int {
	ch := make(chan int)
	go func() {
		for {
			select {
			case ch <- 0:
			case ch <- 1:
			}
		}
	}()
	return ch
}

// 生成随机数
func getRandNum() {
	var ge = make(chan int)
	ge = get01()

	for i := 0; i < 5; i++ {
		fmt.Println(<-ge)
	}
}

func timer(duration time.Duration) chan bool {
	ch := make(chan bool)
	go func() {
		time.Sleep(duration)
		ch <- true
	}()
	return ch
}

// 计时器
func clock() {
	timeout := timer(10 * time.Second)

	for {
		select {
		case <-timeout:
			fmt.Println("10 second")
			return
		}
	}
}

func Chan1_6() {
	clock()
}
