package main

import "fmt"

func main()  {
	var i int = 10
	fmt.Println(i)
	var s string = "hello"
	//s变量要使用
	fmt.Println(s)
	var s2 = "world"
	fmt.Println(s2)
	var (
		a = 10
		b = 100
	)
	fmt.Println(a + b)
	var (
		c int = 100
		d int = 1000
	)
	fmt.Println(c + d)
	//fmt.Println("i的值" + string(i))
	/*
	xxx
	*/
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}