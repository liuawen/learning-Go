package main

import (
	"flag"
	"log"
)
func main() {
	var name string
	flag.StringVar(&name, "name", "Go语言编程之旅name", "帮助信息")
	flag.StringVar(&name, "n", "Go语言编程之旅n", "帮助信息")
	flag.Parse()

	log.Printf("name: %s", name)//2020/11/29 20:59:41 name: Go语言编程之旅n
}
