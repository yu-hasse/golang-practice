package main

import "fmt"

func deferMain() {
	defer fmt.Println("world")
	fmt.Println("Hello")
	// Hello world
	// mainの処理の最後に実行

	fmt.Println("run")
	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3)
	fmt.Println("success")

	// run success 3 2 1

	//ファイルの読み書きのclose処理に使う
}
