package main

import "fmt"
func main() {
	var s1 string = "hello"
	var s2 string = "world"
	fmt.Println("s1+s2=", s1+s2)
	s1+=s2
	fmt.Println("s1=", s1)
}
