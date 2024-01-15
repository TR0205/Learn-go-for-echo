package basic

import "fmt"

type Vertex struct {
	X, Y int
	// 小文字だと外部からアクセスできない
	// s string
}

func (v Vertex) Area() int {
	// 関数名の前にストラクトを定義
	// Goにはクラスが無いが、クラスのメソッドのように定義することも可能
	return v.X * v.Y
}

func (v *Vertex) Scale(i int) {
	// ストラクトの実体を書き換える
	v.X *= i
	v.Y *= i
}

func StructArea() {
	v := Vertex{3, 4}
	v.Scale(10)
	fmt.Println(v.Area())
}

type Vertex3D struct {
	// Embedded
	// PHPとかPythonでいう継承みたいなやつ
	// ストラクト名を書くだけでいい
	Vertex
	Z int
}

func (v Vertex3D) Area3D() int {
	return v.X * v.Y * v.Z
}

func (v *Vertex3D) Scale3D(i int) {
	v.X *= i
	v.Y *= i
	v.Z *= i
}

// コンストラクタ
// Goのパターン。パッケージはNew関数を定義し、外部から使用できるようにしてあげるらしい
func New(x, y, z int) *Vertex3D {
	return &Vertex3D{Vertex{x, y}, z}
}

// インターフェース
type Human interface {
	Say() string
}

type Person struct {
	Name string
}

func (p *Person) Say() string {
	p.Name = "Mr." + p.Name
	fmt.Println(p.Name)
	return p.Name
}

// 以下の関数のように必ずSay()を使用する場合にインターフェースで強制させる
func DriveCar(human Human) {
	if human.Say() == "Mr.Mike" {
		fmt.Println("Run")
	} else {
		fmt.Println("Get out")
	}
}

// func Exec() {
// 	var mike Human = &Person{"Mike"}
// 	var x Human = &Person{"X"}
// 	DriveCar(mike)
// 	DriveCar(x)
// 	// mike.Say()

// }

// タイプアサーション、switch type文
// どんな型でも受け付ける関数
func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Println(v * 2)
	case string:
		fmt.Println(v + "!")
	default:
		fmt.Println(v)
	}
}
