package main

import (
	"fmt"
)
// Func型を新しく定義
type Func func() string
//Func型はfmt.Stringerインターフェイスを実装
func (f Func) String() string { return f() }
func main() {
	var s fmt.Stringer = Func(func() string { return "hello"})
	fmt.Println(s)
}