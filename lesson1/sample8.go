package main

import "fmt"

func main8() {
	// fs変数のスライス作成、スライスの長さは3で中身はまだ空
	fs := make([]func(), 3)
	// fsスライスに対してforループ、iはスライスのインデックス
	for i := range fs {
		fs[i] = func() { fmt.Println(i) }
	}
	for _, f := range fs {
		f()
	}
}
