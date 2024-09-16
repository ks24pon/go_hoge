package main

import (
	"fmt"
)

func main7() {
	fs := make([]func() string, 2)
	fs[0] = func() string { return "hoge" }
	fs[1] = func() string { return "fuga" }

	for _, f := range fs {
		fmt.Println(f())
	}
}