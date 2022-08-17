package main

import (
	"fmt"
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
	var output string
	for _, val := range lines {
		for i, char := range val {
			switch {
			case char == 's' && val[i+1] == ' ' && val[i-1] == ' ':
				output += "\n" + string(char)

			case char == 'i' && val[i+1] == ' ' && val[i-1] == ' ':
				output += "\n" + string(char)

			case char == 'n' && val[i+1] == ' ':
				output += "\n" + string(char)

			case char == 'd' && val[i+1] == ' ':
				output += "\n" + string(char)

			case char == '{':
				output += string(char) + "\n" + "\t"

			case char == '}':
				output += "\n" + string(char)

			default:
				output += string(char)
			}
		}
	}
	fmt.Println(output)
	return output
}
