package string

import (
	"fmt"
	"unicode/utf8"
)

func StringFunc() {
	s := "Hello World"
	//printBytes(s)
	//printChars(s)

	//在 UTF-8 编码中，一个代码点可能会占用超过一个字节的空间。
	//rune 能帮我们解决这个难题。
	t := "Señor"
	//printBytes(t)
	//printChars(t)

	//Go 给我们提供了一种更简单的方法来做到这一点：
	// 使用 for range 循环。
	printCharsAndBytes(t)
	printCharsAndBytes(s)

	fmt.Printf("length of %s is %d\n", s, utf8.RuneCountInString(s))
	fmt.Printf("length of %s is %d\n", t, utf8.RuneCountInString(t))

	r := []rune(s)
	r[0] = 'Y'
	fmt.Println(r)
}

func printBytes(s string) {
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
}

func printChars(s string) {
	runes := []rune(s)
	for i := 0; i < len(runes); i++ {
		fmt.Printf("%c ", runes[i])
	}
}

func printCharsAndBytes(s string) {
	for index, rune := range s {
		fmt.Printf("%c starts at byte %d\n", rune, index)
	}
}
