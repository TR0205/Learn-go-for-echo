package goroutine

import (
	"fmt"
	"sync"
)

func sendMessage(msg string) {
	fmt.Println(msg)
}

func Loop() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		// ここをコメントアウトすると何も出力されない
		// goroutineが実行中にも関わらずmain()が強制終了されてしまう
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println(i)
		}()
	}
	wg.Wait()
}

// root@02b3e5dd2636:/go/src/app# go run main.go
// 10
// 10
// 10
// 10
// 10
// 10
// 10
// 10
// 10
// 10

func Loop2() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		// ここが1の場合、無名関数がそれぞれ5回の合計10回しか実行されない
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println(i)
		}()
		go func() {
			defer wg.Done()
			fmt.Println("ふたつめ")
		}()
	}
	// ここをコメントアウトすると何も出力されない
	// goroutineが実行中にも関わらずmain()が強制終了されてしまう
	wg.Wait()
}

// go run main.go
// ふたつめ
// 10
// ふたつめ
// 10
// ふたつめ
// 10
// ふたつめ
// 10
// ふたつめ
// 10

func Loop3() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		// 10ばかり表示されてしまう場合の対策。一度コピーしてやる
		v := i
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println(v)
		}()
	}
	wg.Wait()
}

func Loop4() {
	n := 0
	// 複数のgoroutineで同じ変数を参照する場合
	var mu sync.Mutex
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := 0; i < 1000; i++ {
			mu.Lock()
			n++
			mu.Unlock()
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 1000; i++ {
			mu.Lock()
			n++
			mu.Unlock()
		}
	}()
	wg.Wait()
	fmt.Println(n)
}

func server(ch chan string) {
	defer close(ch)
	ch <- "one"
	ch <- "two"
	ch <- "three"
}

func Channel() {
	// var s string

	ch := make(chan string)
	go server(ch)

	// fmt.Println(ch)
	for v := range ch {
		fmt.Println(v)
	}

	// s = <-ch
	// fmt.Println(s)
	// s = <-ch
	// fmt.Println(s)
	// s = <-ch
	// fmt.Println(s)
}

func goroutine1(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
		c <- sum
	}
	// 閉じないとループが終了しているのにrangeで取りに来ようとしてエラーになる
	// close(c)
}

func Channel2() {
	s := []int{1, 2, 3, 4, 5}
	c := make(chan int, len(s))
	go goroutine1(s, c)
	for i := range c {
		fmt.Println(i)
	}
}
