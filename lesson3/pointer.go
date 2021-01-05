package main

import "fmt"

func one(x *int) {
	*x = 1

}
func main() {
	var n int = 100
	one(&n)
	fmt.Println(n)  // 1
	fmt.Println(n)  // 100
	fmt.Println(&n) // 100をいれたメモリのアドレス

	var p *int = &n
	fmt.Println(p) // 100をいれたメモリのアドレス

	fmt.Println(*p) // 100 (アドレスが指す中身)
}
