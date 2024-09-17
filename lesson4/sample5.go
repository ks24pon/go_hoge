package main

import "fmt"
// Hogeと言う名前の構造体を定義
type Hoge struct{ N int}
// Fugaと言う名前の構造体を定義
type Fuga struct{ Hoge }

func main() {
	// fはFugaのインスタンスであり、同時に埋め込まれていたHoge構造体を持つ
	f := Fuga{Hoge{N: 100}}
	fmt.Println(f.N)
	// 型名を指定してアクセスできる
	fmt.Println(f.Hoge.N)
}