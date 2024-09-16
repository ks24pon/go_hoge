package main

type T int

func (t *T) f() { println("hi") }

func main12() {
	var v T
	(&v).f()
	v.f()
}