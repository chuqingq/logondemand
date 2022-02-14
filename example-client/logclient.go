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
	count := 0
	for {
		logger.Printf("%v", count)
		log.Printf("%v", count)
		count += 1
		time.Sleep(1 * time.Second)
	}
}
