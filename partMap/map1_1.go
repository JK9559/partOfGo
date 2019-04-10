package partMap

import (
	"fmt"
)

type person struct {
	name string
	age  int
	sex  string
}

func buildAMapWithPointer() {
	pMap := make(map[int]*person)
	pMap[0] = &person{"tom", 30, "woman"}
	fmt.Println(pMap[0])
	fmt.Println(*pMap[0])
	pMap[0].name = "Tony"
	fmt.Println(pMap[0])
	fmt.Println(*pMap[0])
}

func buildAMapWithOutPointer() {
	pMap := make(map[int]person)
	pMap[0] = person{"tony", 20, "man"}
	fmt.Println(pMap[0])
	//pMap[0].name = "kevin"
	fmt.Println(pMap[0].name)
}

func howToInitMap() {
	// 1)
	var map1 map[string]string
	map1 = make(map[string]string)
	map1["a"] = "aaa"
	map1["b"] = "bbb"

	// 2)
	map2 := make(map[string]string)
	map2["a"] = "aa"
	map2["b"] = "bb"

	map3 := map[string]string{
		"a": "aaa",
		"b": "bbb",
	}

	if v, ok := map3["a"]; ok {
		fmt.Println(v)
	} else {
		fmt.Println("Nothing")
	}

	for k, v := range map3 {
		fmt.Println(k, v)
	}

}

// golang使用map时 构造是否为指针的区别
// golang使用map时 如何初始化map
func Map1_1() {
	howToInitMap()
}
