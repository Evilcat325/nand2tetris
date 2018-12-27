package main

import (
	"fmt"
	"strings"
)

func brachingTranslator(command CommandType, instructions []string, state *TranslatorState) string {
	result := ""
	if command == Label {
		if state.functionName != "" {
			result = fmt.Sprintf("(%s$%s)\n", state.functionName, instructions[1])
		} else {
			result = fmt.Sprintf("(%s)\n", instructions[1])
		}
	} else if command == Goto {
		if state.functionName != "" {
			result = fmt.Sprintf("@%s$%s", state.functionName, instructions[1])
		} else {
			result = fmt.Sprintf("@%s", instructions[1])
		}
		result += "\n0; JMP\n"
	} else if command == IfGoto {
		result =
			`@SP
			M=M-1
			A=M
			D=M
			@` + state.functionName + instructions[1] + `
			D;JNE
			`
	}
	return strings.Replace(result, "\t", "", -1)
}
