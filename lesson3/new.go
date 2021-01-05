package main

import "fmt"

func main() {
	var p *int = new(int) // 実体がなしでポインタがさすアドレスを確保するだけのとき
	fmt.Println(p)        // pの値を入れるメモリ領域のアドレスが出力される（まだメモリ自体には値は入っていない）
	fmt.Println(*p)       // 0

	var p2 *int
	fmt.Println(p2) // nil

	s := make([]int, 0)
	fmt.Printf("%T\n", s) // []int

	m := make(map[string]int) // map[string]int
	fmt.Printf("%T\n", m)

	fmt.Printf("%T\n", p) // *int

	// newとmakeはポインタ型が返ってくるかどうかで使い分ける
}
