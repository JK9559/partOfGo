package partChannel

import (
	"fmt"
	"runtime"
)

func getNums() chan int {
	var ch = make(chan int)

	go func() {
		for i := 2; ; i++ {
			ch <- i
		}
	}()

	return ch
}

func filter(in chan int, number int) chan int {

	fmt.Println("Haha")
	out := make(chan int)
	go func() {
		for {
			i := <-in
			fmt.Printf("Chan name is %p,Inner func val is:%d\n", &in, i)
			if i%number != 0 {
				out <- i
			}
		}
	}()
	return out
}

func daisyDemo() {
	const max = 25
	nums := getNums()
	number := <-nums

	for number < max {
		fmt.Println(number)
		nums = filter(nums, number)
		number = <-nums
	}
}

func Chan1_5() {
	runtime.GOMAXPROCS(1)
	daisyDemo()
}
