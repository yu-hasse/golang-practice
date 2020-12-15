package main

import (
	"fmt"
	"os/user"
	"time"
)

func main() {
	fmt.Println("Hello world!!", time.Now())
	fmt.Println(user.Current())

	// 変数宣言について
	var (
		i int = 1
		f64 float64 = 1.2
		s string = "test"
		t, f bool = true, false
	)
	fmt.Println(i, f64, s, t, f)

	xi := 1 //宣言時点でint型も決まる
	xf64 := 1.2
	xs := "test"
	xt, xf := true, false
	fmt.Println(xi, xf64, xs, xt, xf)
	fmt.Printf("%T\n", xf64)
	fmt.Printf("%T\n", xi)

	//var宣言は関数の外でも宣言実行できる
}
