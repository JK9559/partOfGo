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

func Chan1_2() {
	go blockThisChan()
	<-complete
}
