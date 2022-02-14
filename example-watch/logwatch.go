package main

import (
	"log"

	"github.com/chuqingq/logondemand"
)

func main() {
	filename := "logondemand_test"
	err := logondemand.Cat(filename)
	if err != nil {
		log.Fatalf("Cat error: %v", err)
	}
}
