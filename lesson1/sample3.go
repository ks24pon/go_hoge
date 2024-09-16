package main

func main3() {
	m := map[string]int{"x": 10, "y": 20}

	// キーを指定アクセス
	println(m["x"])

	// キーを指定して入力
	m["z"] = 30

	// 存在を確認
	n, ok := m["z"]
	println(n, ok)

	// キーを指定して削除
	delete(m, "z")

	// 削除されていることを確認
	n, ok = m["z"]
	println(n, ok)
}
