package partChannel

import "fmt"

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

func getNotificationDemo() {
	tom := getNotification("Tom")
	tony := getNotification("Tony")

	fmt.Println(<-tom)
	fmt.Println(<-tony)
}

func Chan1_4() {
	getIntsDemo()
}
