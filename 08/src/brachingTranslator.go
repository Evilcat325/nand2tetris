package main

import (
	"fmt"
	"strings"
)

func brachingTranslator(command CommandType, instructions []string, state *TransaltorState) string {
	result := ""
	if command == Label {
		result = fmt.Sprintf("(%s.%s$%s)\n", state.fileName, state.functionName, instructions[1])
	} else if command == Goto {
		result = fmt.Sprintf("@%s.%s$%s\n0;JMP\n", state.fileName, state.functionName, instructions[1])
	} else {
		result =
			`@SP
			M=M-1
			A=M
			D=M
			@LABEL_` + instructions[1] + `
			D;JNE
			`
	}
	return strings.Replace(result, "\t", "", -1)
}
