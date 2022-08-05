package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
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
	var output string
	for fileScanner.Scan() {
		lines = append(lines, fileScanner.Text())
	}
	content.Close()

	for _, val := range lines {

		if val[len(val)-1] == '{' {
			output += strings.TrimSpace(val)
		} else {
			output += " "
			output += val
		}
	}
	fmt.Println(output)

	/*wrting to file with reulting string after processing
	err = os.WriteFile("result.txt", input, 0777)
	if err != nil {
		log.Fatalf("%v", err)
	}*/

	//return input
}
