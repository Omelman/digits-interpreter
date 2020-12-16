package main

import (
	"fmt"
	"github.com/Omelman/digits-interpreter/digits"
	"log"
)

func main() {
	fmt.Println("Enter your decode text: ")
	inputText, err := digits.Input()
	if err != nil {
		log.Fatalf("Failed to read input: %s", err.Error())
	}
	fmt.Print("Decoded: ", digits.TranslateText(inputText, digits.DecodeMap))

	fmt.Println("Enter your encode text: ")
	inputText, err = digits.Input()
	if err != nil {
		log.Fatalf("Failed to read input: %s", err.Error())
	}
	fmt.Print("Encoded: ", digits.TranslateText(inputText, digits.EncodeMap))
}
