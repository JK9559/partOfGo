package partChannel

import (
	"fmt"
)

func goRoutineA(a <-chan int) {
	val := <-a
	fmt.Println("goRouA have a Val: ", val)
}

func Chan1_1() {

}
