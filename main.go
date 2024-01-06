package main

import (
	"fmt"
)

func incrementGenerator() func() int {
	x := 0
	// クロージャは自身が定義されたスコープ内(incrementGenerator())の変数にアクセスでき、値を保持できる。
	// そのためインクリメントされていく

	return func() int {
		x++
		return x
	}
}

func main() {
	// basic.Num()

	counter := incrementGenerator()
	fmt.Println(counter())
	fmt.Println(counter())
	fmt.Println(counter())
	fmt.Println(counter())
}
