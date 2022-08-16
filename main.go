package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	// file to be read and converted to one line will be given as a command line argument
	inputFile := os.Args[1]

	//need to setup so you can input text file as command line argument
	content, err := os.Open(inputFile)
	if err != nil {
		log.Fatalf("Error while reading %v", err)
	}

	fileScanner := bufio.NewScanner(content)
	fileScanner.Split(bufio.ScanLines)

	var lines []string
	for fileScanner.Scan() {
		lines = append(lines, fileScanner.Text())
	}
	content.Close()

	output := Conversion(lines)

	err = os.WriteFile("result.txt", []byte(output), 0777)
	if err != nil {
		log.Fatalf("%v", err)
	}
}
