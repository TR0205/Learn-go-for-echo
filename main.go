package main

import (
	"github.com/TR0205/go-2023/app/goroutine"
)

func main() {
	// e := echo.New()
	// e.GET("/", func(c echo.Context) error {
	// 	return c.String(http.StatusOK, "Hello, World!")
	// })
	// e.Logger.Fatal(e.Start(":8080"))
	goroutine.Channel2()
}
