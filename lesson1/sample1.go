package main

import "fmt"

func main1() {
	ns := []int{10, 20, 30, 40, 50}
	n, m := 2, 4

	fmt.Println(ns[n:])
	fmt.Println(ns[:m])

	ms := ns[:m:m]
	fmt.Println(cap(ms))
}