package main

import (
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	content, err := ioutil.ReadFile("test.txt")
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
	return input
}
