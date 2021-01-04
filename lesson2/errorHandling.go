package main

import (
	"fmt"
	"log"
	"os"
)

func handlingMain() {
	file, err := os.Open("./errorHandling.go")
	if err != nil {
		log.Fatal("Error!")
	}
	defer file.Close()
	data := make([]byte, 100)
	count, err := file.Read(data)
	if err != nil {
		log.Fatal("Error")
	}
	fmt.Println(count, string(data))

	if err = os.Chdir("test"); err != nil {
		log.Fatal("Error")
	}
}
