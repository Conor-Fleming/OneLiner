package main

import "strings"

func conversion(input []string) string {
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
	return ""
}
