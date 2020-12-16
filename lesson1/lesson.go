package main

import (
	"fmt"
	"os/user"
	"time"
)

const Pi = 3.14

const (
	Username = "test_user"
	Password = "test_pass"
)

func main() {
	fmt.Println("Hello world!!", time.Now())
	fmt.Println(user.Current())

	// 変数宣言について
	//var宣言は関数の外でも宣言実行できる
	var (
		i    int     = 1
		f64  float64 = 1.2
		s    string  = "test"
		t, f bool    = true, false
	)
	fmt.Println(i, f64, s, t, f)

	xi := 1 //宣言時点でint型も決まる
	xf64 := 1.2
	xs := "test"
	xt, xf := true, false
	fmt.Println(xi, xf64, xs, xt, xf)
	fmt.Printf("%T\n", xf64)
	fmt.Printf("%T\n", xi)

	fmt.Println(Pi, Username, Password)

	//数値型
	var (
		u8  uint8     = 255
		i8  int8      = 127
		f32 float32   = 0.2
		c64 complex64 = -5 + 12i
	)
	fmt.Println(u8, i8, f32, c64)

	fmt.Printf("type = %T value = %v", u8, u8)

	fmt.Println("1 + 1 = ", 1+1)
	fmt.Println("10 - 1 = ", 10-1)
	fmt.Println("10 / 2 = ", 10/2)
	fmt.Println("10 / 3 = ", 10/3)
	fmt.Println("10.0 / 3 = ", 10.0/3)
	fmt.Println("10 / 3.0 = ", 10/3.0)
	fmt.Println("10 % 2 = ", 10%2)
	fmt.Println("10 % 3 = ", 10%3)

	//初期化宣言ではスペースを空ける慣例
	x := 1 + 1
	fmt.Println(x)

	//gofmt -w ファイル名で自動補正

	y := 0
	fmt.Println(y)
	y++
	fmt.Println(y)
	y--
	fmt.Println(y)

	// シフト
	fmt.Println(1 << 0) //0001 0001 > 1
	fmt.Println(1 << 1) //0001 0010 > 2
	fmt.Println(1 << 2) //0001 0100 > 4

}
