package main

import (
	"io"
	"log"
	"os"
)

func loggingMain() {
	loggingSettings("test.log")
	log.Println("logging")
	log.Printf("%T %v", "test", "test")

	log.Fatalf("%T %v", "test", "test")
	log.Fatalln("error!")

}

func loggingSettings(logFile string) {
	logfile, _ := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	multiLogFile := io.MultiWriter(os.Stdout, logfile)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.SetOutput(multiLogFile)
}
