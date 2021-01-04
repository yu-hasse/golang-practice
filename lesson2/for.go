package main

import "fmt"

func main() {
	for i := 0; i <= 10; i++ {
		if i == 3 {
			continue
		}
		fmt.Println(i)
	}
	//ループ条件省略
	sum := 1
	for sum < 10 {
		sum += sum
		fmt.Println(sum)
	}
	fmt.Println(sum)
}
