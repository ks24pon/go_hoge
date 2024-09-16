package main

func swap(x, y int) (x2, y2 int) {
	y2, x2 = x, y
	return
}

func main() {
	x, y := swap(10, 20)
	println(x, y)
}