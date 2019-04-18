package partChannel

import (
	"fmt"
	"math/rand"
	"time"
)

// https://studygolang.com/articles/9531

func getInts() chan int {
	ch := make(chan int)
	fmt.Println(&ch)
	go func() {
		fmt.Println("Hahaha")
		for i := 0; ; i++ {
			ch <- i
		}
	}()

	return ch
}

// 该函数的执行过程
// 1. 明确该程序有两个不同协程 一是主协程 二是由go关键字开启的 通过声明 可以看到两个协程共同使用一个共享的信道ch
// 2. 当主协程执行到go func时 匿名协程开始执行。首先打印Hahaha 然后向ch传入一个i 继续传入时 容量不够 阻塞
// 3. 主协程未被副协程中断，继续执行，读取了信道ch的值，再想读取 没值了 阻塞
// 4. 这时副协程因为主协程取走了值 而不阻塞 继续流入 同时 因为有值 主协程读取。依次循环直到主协程执行完毕，副协程随主协程执行完毕而执行完毕
func getIntsDemo() {

	gen := getInts()
	fmt.Println("pointer is : ", &gen)
	for i := 0; i < 10; i++ {
		fmt.Println(<-gen)
	}
}

func getNotification(s string) chan string {

	notifyChan := make(chan string)

	go func() {
		notifyChan <- fmt.Sprintf("Hi %s,welcome to this place!", s)
	}()

	return notifyChan
}

// 服务化
func getNotificationDemo() {
	tom := getNotification("Tom")
	tony := getNotification("Tony")

	fmt.Println(<-tom)
	fmt.Println(<-tony)
}

func complexOper(x int) int {
	time.Sleep(time.Duration(rand.Intn(10)) * time.Millisecond)
	return 100 - x
}

func branch(x int) chan int {
	ch := make(chan int)
	go func() {
		ch <- complexOper(x)
	}()
	return ch
}

func fanIn(chs ...chan int) chan int {
	ch := make(chan int)
	//for _, c := range chs {
	//	go func(c chan int) {
	//		ch <- <-c
	//	}(c)
	//}

	go func() {
		for i := 0; i < len(chs); i++ {
			select {
			case v1 := <-chs[i]:
				ch <- v1
			}
		}
	}()
	return ch
}

// 多路复用 进行一个复杂的操作 如运算。分为三路运算最后将信道的数据合并到一个信道
// 我们需要按顺序输出我们的返回值
func multiUseChannel() {
	resCh := fanIn(branch(1), branch(3), branch(200))

	for i := 0; i < 3; i++ {
		fmt.Println(<-resCh)
	}
}

func selectListen(i int) chan int {
	ch := make(chan int)
	go func() {
		time.Sleep(1 * time.Millisecond)
		ch <- i
	}()
	return ch
}

// 通过for循环和select来取信道里的值，当主协程执行完毕，无限循环的for也就停止执行了
// 通过select来监听信道
func selectListenDemo() {
	c1, c2, c3 := selectListen(1), selectListen(2), selectListen(3)
	ch := make(chan int)

	go func() {
		for {
			select {
			case v1 := <-c1:
				ch <- v1
			case v2 := <-c2:
				ch <- v2
			case v3 := <-c3:
				ch <- v3
			}
		}
	}()

	for i := 0; i < 3; i++ {
		fmt.Println(<-ch)
	}
}

// select信道的超时处理
func selectTimeOut() {
	timeout := time.After(10 * time.Millisecond)
	c1, c2, c3 := selectListen(1), selectListen(2), selectListen(3)
	ch := make(chan int)

	go func() {
		for is_timeout := false; !is_timeout; {
			select {
			case v1 := <-c1:
				ch <- v1
			case v2 := <-c2:
				ch <- v2
			case v3 := <-c3:
				ch <- v3
			case <-timeout:
				is_timeout = true
			}
		}
	}()

	for i := 0; i < 3; i++ {
		fmt.Println(<-ch)
		fmt.Println(len(ch))
		//if len(ch) <= 0 {
		//	break
		//}
	}
}

func Chan1_4() {
	//getIntsDemo()
	//multiUseChannel()
	//selectListenDemo()
	selectTimeOut()
}
