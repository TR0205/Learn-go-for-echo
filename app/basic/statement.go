package basic

import (
	"fmt"
	"log"
	"os"
	"time"
)

func Range() {
	l := []string{"python", "go", "java"}

	// 通常のfor文
	for i := 0; i < len(l); i++ {
		fmt.Println(i, l[i])
	}

	// 上と同じ
	for i, v := range l {
		fmt.Println(i, v)
	}
	// 0 python
	// 1 go
	// 2 java

	m := map[string]int{"apple": 100, "banana": 200}

	for k, v := range m {
		fmt.Println(k, v)
	}
	// apple 100
	// banana 200
}

func getOsName() string {
	return "mac"
}

func Switch() {
	// osはswitch文の中だけでしか使用できない
	switch os := getOsName(); os {
	case "mac":
		fmt.Println("mac")
	case "windows":
		fmt.Println("windows")
	default:
		fmt.Println("default")
	}
	// ここで宣言してもエラーになる
	// fmt.Println(os)

	// case文内でストラクトのメソッドを条件式に使用することも可能
	t := time.Now()
	fmt.Println(t.Hour())
	switch {
	case t.Hour() < 12:
		fmt.Println("morning")
	case t.Hour() > 17:
		fmt.Println("afternoon")
	}
}

func DeferFoo() {
	defer fmt.Println("world foo")

	fmt.Println("hello foo")
}
func Defer() {
	defer fmt.Println("world")
	fmt.Println("hello")

	DeferFoo()

	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3)

	// hello
	// hello foo
	// world foo
	// 3
	// 2
	// 1
	// world
}

func Log() {
	// Printlnのlnはライン、Printfのfはフォーマット
	log.Println("logging")

	_, err := os.Open("fiwuhiqeu")
	if err != nil {
		log.Fatalln("Exit", err)
	}
	// log.Fatalln以降の処理は実行されない
	log.Println("no output")

	// 	2024/01/07 00:26:36 logging
	// 2024/01/07 00:26:36 Exit open fiwuhiqeu: no such file or directory
	// exit status 1
}

func ErrorHandling() {
	file, err := os.Open("./lesson.go")
	if err != nil {
		log.Fatalln("Errorrrrrrrr")
	}
	defer file.Close()

	data := make([]byte, 100)
	// 上でもerrを初期化しここでも初期化をしているがエラーにはならない。
	// :=で初期化する際、最低一つでも初期化するとエラーにならない
	// 今回はcountも初期化しているのでエラーにならない
	count, err := file.Read(data)

	if err != nil {
		log.Fatalln("Err")
	}
	fmt.Println(count, string(data))

	// ここでは初期化する変数がerr一つのため、:=を使用するとエラーになる
	if err = os.Chdir("test"); err != nil {
		log.Fatalln(("Errrorrrrorrororor"))
	}
}
