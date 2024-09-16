package main

func main10() {
	ns := []int{10, 20, 30}
	ns2 := ns
	ns[1] = 200

	println(ns[0], ns[1], ns[2])
	println(ns2[0], ns2[1], ns2[2])
}