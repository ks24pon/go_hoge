package main

import (
	"fmt"

	"golang.org/x/exp/slices"
)

func main() {
	ns := []int{10, 20, 30, 40, 50}

	// 削除
	ns = slices.Delete(ns, 1, 3)
	fmt.Println(ns)

	// 挿入
	ns = slices.Insert(ns, 1, 60, 70)
	fmt.Println(ns)

	//要素があるか
	ok := slices.Contains(ns, 70)
	fmt.Println(ok)

	// ソート
	slices.Sort(ns)
	fmt.Println(ns)
}
