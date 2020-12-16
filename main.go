package main

import (
	"fmt"
	"github.com/Omelman/digits-interpreter/digits"
	"log"
)

func main() {
	var inputText string

	fmt.Println("Enter your decode text: ")
	inputText, err := digits.Input()
	if err != nil {
		log.Fatalf("Failed to read input: %s", err.Error())
	}
	fmt.Print("Decoded: ", digits.TextInterpreter(inputText, digits.DecodeMap))

	fmt.Println("Enter your encode text: ")
	inputText, err = digits.Input()
	if err != nil {
		log.Fatalf("Failed to read input: %s", err.Error())
	}
	fmt.Print("Encoded: ", digits.TextInterpreter(inputText, digits.EncodeMap))
}
