package main

import (
	"log"
	"os"
)

func main() {
	input := []byte("test test test test \n test again")
	err := os.WriteFile("result.txt", input, 0777)
	if err != nil {
		log.Fatalf("%v", err)
	}
}
