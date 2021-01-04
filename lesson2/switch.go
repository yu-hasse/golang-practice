package main

import (
	"fmt"
	"time"
)

func switchMain() {
	os := "mac"

	switch os {
	// switch os := getOsName(); os{
	case "mac":
		fmt.Println("Mac!!")
	case "win":
		fmt.Println("Win!!")
	default:
		fmt.Println("Default!!")
	}

	t := time.Now()
	fmt.Println(t.Hour())
	switch {
	case t.Hour() < 12:
		fmt.Println("morning")
	case t.Hour() < 17:
		fmt.Println("Afternoon")
	}
}

func getOsName() string {
	return "mac"
}
