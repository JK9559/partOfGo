package partChannel

import "fmt"

//https://studygolang.com/articles/9532

var complete chan int = make(chan int)

func blockThisChan() {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", i)
	}
	complete <- 0
}

// 判断信道是否死锁 有机会执行的信道必须既有流入也有流出
func deadLock() {
	ch := make(chan int)
	quit := make(chan int)
	// 不会发生死锁 因为有信道quit流入数据 在主线程中的quit流出了数据 不会阻塞
	go func() {
		quit <- 2
		ch <- 1
	}()
	<-quit

	// 会发生死锁 因为ch1先流入数据 并未被读出 所以阻塞，导致quit1未流入数据
	// 却在主线程中读取quit1 所以死锁
	ch1 := make(chan int)
	quit1 := make(chan int)
	go func() {
		ch1 <- 1
		quit1 <- 2
	}()
	<-quit1
}

// 关于无缓冲信道的数据的进出顺序 答案就是无顺序，因为无缓冲信道不存储数据 只负责数据的流通
func input(i int) {
	complete <- i
}

func zeroLenChannel() {
	for i := 0; i < 5; i++ {
		go input(i)
	}
	for i := 0; i < 5; i++ {
		fmt.Println(<-complete)
	}

}

// 以下的例子 缓冲信道ch可以无缓冲的流入3个元素
// 如果再流入一个元素的话 会报死锁 ，也就是缓冲信道在容量满的时候会加锁
// 而且缓冲信道是先进先出的 可以把缓冲信道看做为一个线程安全的队列
func bufferedChannel() {
	ch := make(chan int, 3)
	ch <- 1
	ch <- 2
	ch <- 3
}

// range 不等到信道关闭是不会停止读取的
// 解决方案有 1.显式的关闭信道 2.判断信道的数据量
func readChannelUnsafe() {
	ch := make(chan int, 3)
	ch <- 1
	ch <- 2
	ch <- 3
	for v := range ch {
		// if len(ch)<=0 break
		fmt.Println(v)
	}
	// 1. close(ch)
	// output:
	//1
	//2
	//fatal error: all goroutines are asleep - deadlock!
	//3
}

func Chan1_2() {
	// go blockThisChan()
	// <-complete
	// deadLock()
	// zeroLenChannel()
	// bufferedChannel()
	readChannelUnsafe()

}
