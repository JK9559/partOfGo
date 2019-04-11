package partChannel

import (
	"fmt"
	"runtime"
)

var quit chan int = make(chan int)

// goroutine 其实是操作系统的一个线程
// 默认地,Go所有的goroutines只能在一个线程里跑
// 现在的版本CPU默认全开
func serialGoroutine() {

}

func printLoop() {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", i)
	}
	quit <- 0
}

func printLoop1() {
	for i := 120; i < 125; i++ {
		fmt.Printf("%d ", i)
	}
	quit <- 0
}

func Chan1_3() {
	runtime.GOMAXPROCS(1)
	go printLoop()
	go printLoop1()
	go printLoop()

	for i := 0; i < 3; i++ {
		<-quit
	}
}
