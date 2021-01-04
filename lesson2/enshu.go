package main

import "fmt"

//スライスから最小値を探す

//mapの合計を求める
func main() {
	l := []int{100, 300, 23, 11, 23, 2, 4, 6, 4}

	var min int

	for i, v := range l {
		if i == 0 {
			min = v
			continue
		}
		if min >= v {
			min = v
		}
	}
	fmt.Println(min)

	m := map[string]int{
		"apple":  200,
		"banana": 300,
		"grapes": 150,
		"orange": 50,
		"papaya": 500,
		"kiwi":   90,
	}

	sum := 0

	for _, v := range m {
		sum += v
	}
	fmt.Println(sum)
}
