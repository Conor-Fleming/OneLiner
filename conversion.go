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
		tabCount := 0
		for i, char := range val {
			switch {
			case char == 's' && val[i+1] == ' ' && val[i-1] == ' ':
				output += "\n"
				output = addTabs(output, tabCount, string(char))

			case char == 'i' && val[i+1] == ' ' && val[i-1] == ' ':
				output += "\n"
				output = addTabs(output, tabCount, string(char))

			case char == 'f' && val[i+1] == ' ':
				output += "\n"
				output = addTabs(output, tabCount, string(char))

			case char == 'n' && val[i+1] == ' ':
				output += "\n"
				output = addTabs(output, tabCount, string(char))

			case char == 'd' && val[i+1] == ' ':
				output += "\n"
				output = addTabs(output, tabCount, string(char))

			case char == '{':
				tabCount++
				output += string(char) + "\n"
				output = addTabs(output, tabCount, string(char))

			case char == '}':
				tabCount--
				output += "\n"
				output = addTabs(output, tabCount, string(char))
			default:
				output += string(char)
			}
		}
	}
	return output
}

func addTabs(output string, tabCount int, char string) string {
	for i := 0; i < tabCount; i++ {
		output += "\t"
	}
	if char != "{" {
		output += char
	}
	return output
}
