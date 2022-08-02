package main

import (
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	// file to be read and converted to one line will be given as a command line argument
	inputFile := os.Args[1]

	//need to setup so you can input text file as command line argument
	content, err := ioutil.ReadFile(inputFile)
	if err != nil {
		log.Fatalf("Error while reading %v", err)
	}

	//will call function "Compress Extension" that will take 'content' variable read from the input file and compress that onto a single line
	//that can then be given to the Write file function for output to result.txt
	input := CompressExtension(content)

	err = os.WriteFile("result.txt", input, 0777)
	if err != nil {
		log.Fatalf("%v", err)
	}
}

func CompressExtension(input []byte) []byte {
	input = []byte(strings.Replace(string(input), "\n", " ", -1))
	input = []byte(strings.Replace(string(input), "\r", "", -1))

	//Use Regex
	//remove space from right side of '{'____ and left side of ____'}'

	//Positive Lookahead regex \s(?=[^}\s]*) (matches first space behind '}')

	//\s(?=[])

	return input
}
