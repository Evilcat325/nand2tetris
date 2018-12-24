package main

import "fmt"

func brachingTranslator(command CommandType, instructions []string) string {
	result := ""
	if command == Label {
		result = fmt.Sprintf("(LABEL_%s)", instructions[1])
	}
	return result
}
