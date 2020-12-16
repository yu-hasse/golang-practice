package main

import (
	"fmt"
	"os/user"
	"strconv"
	"strings"
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

	// 文字列型
	fmt.Println("Hello" + "World!")       // > Hello World!
	fmt.Println(string("Hello World"[0])) // > H

	var str string = "Hello World"
	fmt.Println(strings.Replace(str, "H", "X", 1)) // Xello World
	fmt.Println(strings.Contains(str, "World"))    // true

	fmt.Println(`test
	test`) // バッククォートだと改行が入っている

	fmt.Println("\"")
	fmt.Println(`"`)

	//論理値型
	bt, bf := true, false
	fmt.Printf("%T %v %t\n", bt, bt, bt) // bool true
	fmt.Printf("%T %v %t\n", bf, bf, bf) // bool false

	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(false && true)
	fmt.Println(false && false)

	fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Println(false || true)
	fmt.Println(false || false)

	fmt.Println(!true)
	fmt.Println(!false)

	//型変換
	var xx int = 1
	fx := float64(xx)
	fmt.Printf("%T %v %f\n", fx, fx, fx) // float64 1 1.00000

	var yy float64 = 1.2
	fy := int(yy)
	fmt.Printf("%T %v %d\n", fy, fy, fy) // int 1 1

	var ss string = "14"
	z, _ := strconv.Atoi(ss)
	fmt.Printf("%T %v", z, z) // int 14

	//配列型
	var arr [2]int
	arr[0] = 100
	arr[1] = 200
	fmt.Println(arr) //[100, 200]

	/*
	   var b [2]int = [2]int{100, 200}
	   b = append(b, 300)
	   fmt.Println(b) // 配列はアペンドでサイズ変更不可のためError
	*/

	var b []int = []int{100, 200}
	b = append(b, 300)
	fmt.Println(b) // [100, 200, 300] →スライス

	// スライス
	n := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(n)      // [1 2 3 4 5 6]
	fmt.Println(n[2])   // 3
	fmt.Println(n[2:4]) // [3 4]
	fmt.Println(n[:2])  // [1 2]
	fmt.Println(n[2:])  // [3 4 5 6]
	fmt.Println(n[:])   // [1 2 3 4 5 6]

	n[2] = 100
	fmt.Println(n) // [1 2 100 4 5 6]

	var board = [][]int{
		[]int{0, 1, 2},
		[]int{3, 4, 5},
		[]int{6, 7, 8},
	}
	fmt.Println(board) // [[0 1 2][3 4 5][6 7 8]]

	//スライスのmakeとcap
	nn := make([]int, 3, 5)
	fmt.Printf("len=%d cap=%d value=%v\n", len(nn), cap(nn), nn) // len=3 cap=5 [0 0 0]
	nn = append(nn, 0, 0)
	fmt.Printf("len=%d cap=%d value=%v\n", len(nn), cap(nn), nn) // len=5 cap=5 [0 0 0 0 0]
	nn = append(nn, 1, 2, 3, 4, 5)
	fmt.Printf("len=%d cap=%d value=%v\n", len(nn), cap(nn), nn) // len=10 cap=10 [0 0 0 0 0 1 2 3 4 5]

	aa := make([]int, 3)
	fmt.Printf("len=%d cap=%d value=%v\n", len(aa), cap(aa), aa) // len=3 cap=3 [0 0 0] →両方3で初期化

	bb := make([]int, 0)
	var c []int
	fmt.Printf("len=%d cap=%d value=%v\n", len(bb), cap(bb), bb) // len=0 cap=0 [] →0というスライスをメモリに確保
	fmt.Printf("len=%d cap=%d value=%v\n", len(c), cap(c), c)    // len=0 cap=0 [] →こっちはnil

	//cc := make([]int, 5)
	cc := make([]int, 0, 5)
	for i := 0; i < 5; i++ {
		fmt.Println(cc)
	}
	fmt.Println(cc)

	// map
	m := map[string]int{"apple": 100, "banana": 200}
	fmt.Println(m) // map[apple:100 banana:200]
	fmt.Println(m["apple"]) // 100
	m["banana"] = 300
	fmt.Println(m) // map[apple: 100 banana: 300]
	m["new"] = 500
	fmt.Println(m) // map[apple: 100 banana: 300 new: 500]

	fmt.Println(m["nothing"]) // 0

	value, isExist := m["apple"]
	fmt.Println(value, isExist) // 100 true

	value2, isExist2 := m["nothing"]
	fmt.Println(value2, isExist2) // 0 false

	m2 := make(map[string]int)
	m2["pc"] = 5000
	fmt.Println(m2) // map[pc:5000]

	var sss = []int
	if s == nil{
		fmt.Println("Nil")
	} // Nil →var宣言では初期値はnil。メモリ確保してない。
	
}
