package main

import (
	"fmt"
	"strings"
)

func brachingTranslator(command CommandType, instructions []string, state *TranslatorState) string {
	result := ""
	if command == Label {
		result = fmt.Sprintf("(%s$%s)\n", state.functionName, instructions[1])
	} else if command == Goto {
		result = fmt.Sprintf("@%s$%s\n0;JMP\n", state.functionName, instructions[1])
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
