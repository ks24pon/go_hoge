package main

import "fmt"

func main() {
	var v interface{}
	v = 100
	fmt.Println(v)
	v = "hoge"
	fmt.Println(v)
}
