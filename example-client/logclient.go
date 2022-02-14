package main

import (
	"log"
	"time"

	"github.com/chuqingq/logondemand"
)

func main() {
	filename := "logondemand_test"
	l, err := logondemand.New(filename)
	if err != nil {
		log.Fatalf("New error: %v", err)
	}
	logger := log.New(l, "", log.Flags()|log.Lshortfile|log.Lmicroseconds)
	// go func() {
	// 	err := Cat(filename)
	// 	log.Printf("Cat error: %v", err)
	// }()
	for {
		logger.Printf("123")
		time.Sleep(1 * time.Second)
	}
}
