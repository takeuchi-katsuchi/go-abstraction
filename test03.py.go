package main

import "fmt"

type Hoge interface {
	M()
	N()
}
type fuga struct{ Hoge }

func (f fuga) M() {
	fmt.Println("Hi")
	f.Hoge.M() // 元のメソッドを呼ぶ
}
func HiHoge(h Hoge) Hoge {
	return fuga{h} // 構造体作る
}
