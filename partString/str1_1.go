package partString

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"
)

// Learn from https://studygolang.com/articles/19610

func init() {
	fmt.Println("Package partString")
}

// 按照16进制打印字符串每个字节
func printBytes(s string) {
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
	fmt.Println()
}

// 打印每个字符使用rune
func printChars(s string) {
	// rune 是 int32 的别称
	runes := []rune(s)
	for i := 0; i < len(runes); i++ {
		fmt.Printf("%c ", runes[i])
	}
	fmt.Println()
}

// 使用range打印字符
func printCharsAndBytes(s string) {
	for index, r := range s {
		fmt.Printf("%c starts at %d\n", r, index)
	}
}

// 通过byte切片打印字符，16进制 10进制均可
func printByteSlice() {
	byteSlice := []byte{0x43, 0x61, 0x66, 0xC3, 0xA9}
	byteSlice1 := []byte{67, 97, 102, 195, 169}
	fmt.Println(string(byteSlice))
	fmt.Println(string(byteSlice1))
}

// 通过rune切片打印字符
func printRuneSlice() {
	runeSlice := []rune{0x0053, 0x0065, 0x00f1, 0x006f, 0x0072}
	fmt.Println(string(runeSlice))
}

// 获取字符串长度（该方法传入一个字符串参数然后返回字符串中的rune数量）
func length(s string) int {
	return utf8.RuneCountInString(s)
}

// 字符串的修改 字符串本身是不可变的 但是可以把字符串转化为一个rune切片，改变完成后再转化为一个字符串
func changeStr(s string) string {
	run := []rune(s)
	run[0] = 'a'
	return string(run)
}

func operationsOfStr() {
	fmt.Println("字符串测试")
	fmt.Println("字符串转换为int类型：")
	intVal, err := strconv.Atoi("100")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(intVal)
	}
	fmt.Println("字符串转换为float64类型：")
	fltVal, err := strconv.ParseFloat("100.564", 10)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(fltVal)
	}

	fmt.Println("int转字符串", strconv.Itoa(92))

	fmt.Println("字符串比较")
	str1 := "Hello World"
	str2 := "bcd"
	fmt.Println(strings.Compare(str1, str2))
	fmt.Println(strings.Compare("a", "a"))

	fmt.Println("字符串包含")
	fmt.Println(strings.Contains(str1, "ello"))
	fmt.Println(strings.Contains(str1, "asd"))

	fmt.Println("查找位置")
	fmt.Println(strings.Index(str1, "lo"))
	fmt.Println(strings.Index(str1, "abc"))

	// 在字符串中最后出现位置的索引 -1表示字符串str1中不包含字符串o
	fmt.Println(strings.LastIndex(str1, "o"))

	// 统计给定子串出现的次数，当查找子串为空时，返回1+字符串的长度
	fmt.Println(strings.Count("cheeseeee", "ee"))
	fmt.Println(strings.Count("five", ""))

	// 重复字符串count次，返回新生成的重复的字符串
	fmt.Println(strings.Repeat("world", 5))

	// 字符串的替换，把old替换为new，n为替换的次数，小于0为全部替换
	str3 := "/Users//Documents/GOPatch/src/MyGO/config/TestString/"
	fmt.Println(strings.Replace(str3, "/", "**", -1))
	fmt.Println(strings.Replace(str3, "/", "**", 4))

	// 删除字符串开头和结尾
	fmt.Println(strings.Trim(str3, "/"))
	fmt.Println(strings.TrimLeft(str3, "/"))
	fmt.Println(strings.TrimRight(str3, "/"))
	fmt.Println(strings.TrimSpace("   sahjj shjagg s     "))

	// 大小写
	str4 := "hello hao Hao "
	fmt.Println(strings.Title(str4))
	fmt.Println(strings.ToLower(str4))
	fmt.Println(strings.ToUpper(str4))

	// 前缀 后缀
	fmt.Println(strings.HasPrefix("Gopher", "Go"))
	fmt.Println(strings.HasSuffix("Amigo", "go"))

	// 字符串分割
	// 根据空白符分割
	fieldsStr := "	it`s a 	good idea	 "
	fieldsSlice := strings.Fields(fieldsStr)
	fmt.Println(fieldsSlice)

	for k, v := range fieldsSlice {
		fmt.Printf("%d = %s\n", k, v)
	}
	for i := 0; i < len(fieldsSlice); i++ {
		fmt.Println(fieldsSlice[i])
	}

	// 根据指定字符分割
	slice01 := strings.Split("q,w,e,r,t,y,", ",")
	fmt.Println(slice01)      //[q w e r t y ]
	fmt.Println(cap(slice01)) //7  最后多个空""
	for i, v := range slice01 {
		fmt.Printf("下标 %d 对应值 = %s \n", i, v)
	}

}

func connectStr() {
	sliceStr := []string{"hello", "world", "hahah", "yep"}
	str := strings.Join(sliceStr, ",")
	fmt.Println(str)

	fmt.Println("=========比较字符串拼接速度=========")
	var buffer bytes.Buffer
	start := time.Now()
	for i := 0; i < 100000; i++ {
		buffer.WriteString("this is a string\n")
	}
	end := time.Now()
	fmt.Println("buffer time is: ", end.Sub(start).Seconds())

	start = time.Now()
	str = ""
	for i := 0; i < 100000; i++ {
		str += "this is a string\n"
	}
	end = time.Now()
	fmt.Println("+ time is: ", end.Sub(start).Seconds())

	start = time.Now()
	var strSlice []string
	for i := 0; i < 100000; i++ {
		strSlice = append(strSlice, "this is a string\n")
	}
	strings.Join(strSlice, "")
	end = time.Now()
	fmt.Println("Join time is: ", end.Sub(start).Seconds())
}

func Str1_1() {
	// name := "Señor"
	// 获取程序运行的操作系统平台下 int 类型所占的位数
	//fmt.Println(strconv.IntSize)
	connectStr()
}
