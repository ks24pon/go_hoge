package main

import "fmt"

// Stringerインターフェイスを定義
type Stringer interface {
	String() string
}

// Stringer実装、Hex型はintの新しい型
type Hex int

func (h Hex) String() string {
	return fmt.Sprintf("%x", int(h))
}

// Hex2もStringerを実装
type Hex2 struct{ Hex }

func main() {
	var s Stringer
	h := Hex(100)
	s = h
	fmt.Println(s.String())

	h2 := Hex2{h}
	s = h2
	fmt.Println(s.String())
}
