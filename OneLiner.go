package main

import (
	"os"
)

func main() {
	ipfile, err := os.Open("test.txt")
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := ipfile.Close(); err != nil {
			panic(err)
		}
	}()

	opfile, err := os.Create("result.txt")
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := opfile.Close(); err != nil {
			panic(err)
		}
	}()
}
