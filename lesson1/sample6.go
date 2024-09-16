package main

func main6() {
	msg := "Hello, 世界"
	func() {
		println(msg)
	}()
}