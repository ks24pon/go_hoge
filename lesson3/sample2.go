// package main

// import (
// 	"fmt"
// 	"path/filepath"
// )

// func main() {
// 	// パスを結合する
// 	path := filepath.Join("dir", "main.go")
// 	fmt.Println(path)
// 	// 拡張子を取る
// 	fmt.Println(filepath.Ext(path))
// 	// ファイル名を取得
// 	fmt.Println(filepath.Base(path))
// 	// ディレクトリ名を取得
// 	fmt.Println(filepath.Dir(path))
// }

package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	// パスを結合
	path := filepath.Join("dir", "main.go")
	fmt.Println(path)
	// 拡張子を取る
	fmt.Println(filepath.Ext(path))
	// ファイル名を取得
	fmt.Println(filepath.Base(path))
	// ディレクトリ名を取得
	fmt.Println(filepath.Dir(path))
}