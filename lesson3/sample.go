package main

import (
	"fmt"
	"os"
)

func main() {
	// 標準エラー出力に出力
	fmt.Fprintln(os.Stderr, "エラー")
	fmt.Fprintln(os.Stderr, "Hello")
}