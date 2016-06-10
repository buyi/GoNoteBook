package main

import (
	"errors"
	"fmt"
)

func main() {
	var i int // i declared and not used

	// var: define a variance. syntax:var variableName type
	// const: define a unchaged variance. syntax const constantName = value
	i = 4
	fmt.Println("%d", i)

	valid := true
	fmt.Println("%b", valid)

	var c complex64 = 5 + 5i
	fmt.Println("value is %v", c)

	var s string = "hello"
	cc := []byte(s)
	cc[0] = 'c'
	s2 := string(cc)
	fmt.Println("%s", s2)

	err := errors.New("test a error")
	if err != nil {
		fmt.Println(err)
	}

	const (
		x = iota
		y
		z
		w
	)
	fmt.Println("%v", x)
	fmt.Println("%v", y)
	fmt.Println("%v", z)
	fmt.Println("%v", w)

	//集合元素 var arr [n]type
	var arr [10]int
	index := 0
	for ; index < 10; index++ {
		arr[index] = index + 100
	}
	index = 0
	for ; index < 10; index++ {
		fmt.Println("%v", arr[index])
	}

	// 切片 var arr []type 可以理解为动态数组
	aSlice := make([]int, 10, 20)
	aSlice = arr[2:4]
	fmt.Println("%v", aSlice)

	// key:value 健值对 map［keytype］valuetype
	number := make(map[string]int)
	number["1"] = 1
	number["2"] = 2
	fmt.Println("%v", number)

}
