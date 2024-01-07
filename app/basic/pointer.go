package basic

import (
	"fmt"
)

func Pointer() {
	// アドレスは&、アドレスの中身を参照するときは*
	var n int = 100
	fmt.Println(n)
	// 100

	// 変数が格納されているメモリのアドレスを&で参照する
	fmt.Println(&n)
	// 0xc00009c000

	// アスタリスクをつけるとアドレスの型になる
	var p *int = &n
	fmt.Println(p)
	// 0xc00009c000

	// アドレスの型にアスタリスクをつけることで、メモリに格納されている値を参照
	fmt.Println(*p)
	// 100
}

func one(x int) {
	x = 1
}

func Pointer1() {
	var n int = 100
	one(n)
	fmt.Println(n)
	// 100
}

func two(x *int) {
	// アドレスの中身を受け取って中身に1を入れる(アスタリスク)
	*x = 1
}

func Pointer2() {
	var n int = 200
	// nのアドレスを渡す
	two(&n)
	fmt.Println(n)
	// 1
}

func NewMake() {
	var p *int = new(int)
	fmt.Println(p)
	// 0xc000012080
	// メモリ上に確保されている

	var p2 *int
	fmt.Println(p2)
	// <nil>

	s := make([]int, 0)
	fmt.Printf("%T\n", s)

	m := make(map[string]int)
	fmt.Printf("%T\n", m)

	ch := make(chan int)
	fmt.Printf("%T\n", ch)

	var po *int = new(int)
	fmt.Printf("%T\n", po)

	var st = new(struct{})
	fmt.Printf("%T\n", st)

	// newはポインタ型を返す、makeは返さない
	// []int
	// map[string]int
	// chan int
	// *int
	// *struct {}
}
