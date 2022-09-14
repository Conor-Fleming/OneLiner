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
	var output string
	for _, val := range lines {
		//tabCount := 0
		//var tabs string
		for i, char := range val {
			if val[i+1] == ' ' && val[i-1] == ' ' {
				switch char {
				case 's', 'i':
					output += "\n" + string(char)
				}
			}
			if val[i+1] == ' ' {
				switch char {
				case 'n', 'd', 'f':
					output += "\n" + string(char)
				}
			}
			switch char {
			case '{':
				output += string(char) + "\n" + "\t"

			case '}':
				output += "\n" + string(char)

			default:
				output += string(char)
			}
		}
	}
	return output
}
