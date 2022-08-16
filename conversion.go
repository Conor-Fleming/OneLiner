package main

import (
	"strings"
)

func Conversion(input []string) string {
	if len(input) == 1 {
		return ToMany(input)
	}
	return ToOne(input)
}

func ToOne(lines []string) string {
	output := ""
	for i, val := range lines {
		val = strings.TrimSpace(val)
		check := val[len(val)-1:]
		if i+1 < len(lines) {
			if lines[i+1] != "}" && check != "{" {
				val += " "
			}
		}
		output += val
	}
	return output
}

func ToMany(lines []string) string {
	var checker int = 0
	var output string
	for _, val := range lines {
		for _, char := range val {
			if char == ' ' {
				checker++
			}
			if checker == 2 {
				output += string(char) + "\n"
				checker = 0
				continue
			}
			output += string(char)
		}
	}
	return output
}
