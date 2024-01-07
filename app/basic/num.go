package basic

import "fmt"

// 外部から参照される関数を定義する場合は大文字にする必要がある
func Num() {
	var i8 int8 = 127
	var f32 float32 = 0.2

	fmt.Println(i8, f32)
}

func Cast() {
	var c int = 1
	cc := float64(c)
	fmt.Printf("型：%T\n値：%v\nfloatで表示：%f\n", cc, cc, cc)

	// 出力結果
	// 型：float64
	// 値：1
	// floatで表示：1.000000
}

func Array() {
	// 配列は要素を追加できない
	var a [2]int
	a[0] = 100
	a[1] = 200
	fmt.Print(a)
	// [100 200]

	var b [3]int = [3]int{100, 200, 300}
	fmt.Println(b)
	// [100 200 300]
}

func Slice() {
	n := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(n[2:4])
	// [3 4]
	fmt.Println(n[:4])
	// [1 2 3 4]
	fmt.Println(n[4:])
	// [5 6]

	m := make([]int, 3, 5)
	fmt.Printf("len=%d cap=%d value=%v", len(m), cap(m), m)
	// len=3 cap=5 value=[0 0 0]
	// 要素が３のスライス、メモリ上は5確保している
}

func Map() {
	// PHPの連想配列、Pythonの辞書のようなもの
	m := map[string]int{"apple": 100, "banana": 200}
	fmt.Println(m)
	// map[apple:100 banana:200]

	m["banana"] = 500
	fmt.Println(m)
	// map[apple:100 banana:500]

	m["new"] = 1000
	fmt.Println(m)
	// map[apple:100 banana:500 new:1000]

	fmt.Println(m["nothing"])
	// 0

	v, ok := m["nothing"]
	fmt.Println(v, ok)
	// 0 false
	// キーがあるかの真偽値も返してくれる

	m2 := make(map[string]int)
	fmt.Println(m2)
	// map[]
	m2["pc"] = 5000
	fmt.Println(m2)
	// map[pc:5000]

	var m3 map[string]int
	m3["pc"] = 5000
	fmt.Println(m3)
	// 以下のエラーが発生
	// 変数を宣言しているが、メモリ上にmapに存在していない、

	// 	panic: assignment to entry in nil map

	// goroutine 1 [running]:
	// github.com/TR0205/go-2023/app/basic.Map()
	//         /go/src/app/app/basic/num.go:82 +0x34b
	// main.main()
	//         /go/src/app/main.go:9 +0xf
	// exit status 2
}

func Function(x, y int) (result1 int, result2 int) {
	// 引数の型が同じ場合は１つにまとめることができる
	// 戻り値が複数の場合は()をつける
	// 戻り値に変数名を指定することができる
	// その場合は変数に代入したあと、returnでOK
	// 使用例：複数の値を返す関数の場合、何を返すかがプリミティブな型だとわかりにくい、
	//        その場合変数名だと何を返すかわかりやすい

	result1, result2 = x+y, x-y
	return

	// メモ
	// func() {
	// 	fmt.Println("inner")
	// }()
	// ↑のように関数を定義し、最後に()をつけてあげると即時実行できるresult
	// 変数に入れてからfunc()みたいに書かなくてもいい
}

func incrementGenerator() func() int {
	x := 0
	// クロージャは自身が定義されたスコープ内(incrementGenerator())の変数にアクセスでき、値を保持できる。
	// そのためインクリメントされていく

	return func() int {
		x++
		return x
	}
}

// func main() {
// 	// basic.Num()

// 	counter := incrementGenerator()
// 	fmt.Println(counter())
// 	fmt.Println(counter())
// 	fmt.Println(counter())
// 	fmt.Println(counter())
// }
