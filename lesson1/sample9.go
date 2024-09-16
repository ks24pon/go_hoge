package main

func main9() {
	p := struct {
		age  int
		name string
	}{age: 10, name: "Gopher"}

	p2 := p
	p2.age = 20

	println(p.age, p.name)
	println(p2.age, p2.name)
}
