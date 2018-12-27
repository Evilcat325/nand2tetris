package main

import (
	"fmt"
	"strconv"
	"strings"
)

var returnPoint int
var functionName string

func functionTranslator(command CommandType, instructions []string, state *TranslatorState) string {
	result := ""
	if command == Function {
		// Set current functionName, init return count
		state.functionName = instructions[1]
		state.returnCounter = 0
		// Format label
		functionLabel := fmt.Sprintf("%s.%s", state.fileName, state.functionName)
		nArgs, err := strconv.Atoi(instructions[2])
		check(err)
		// Set function label
		result += "(" + functionLabel + ")\n"
		// Set local segment
		for i := 0; i < nArgs; i++ {
			// initialize to 0
			result +=
				`@SP
				M=M+1
				A=M-1
				M=0
				`
		}
	} else if command == Call {
		callee := instructions[1]
		returnLabel := fmt.Sprintf("%s.%s$ret.%d", state.fileName, state.functionName, state.returnCounter)
		state.returnCounter++
		nArgs, err := strconv.Atoi(instructions[2])
		check(err)
		// save SP to R13
		result +=
			`@SP
			D=M
			@R13
			M=D
			`
		// Save return address
		result +=
			`@` + returnLabel + `
			D=A
			@SP
			A=M
			M=D
			@SP
			M=M+1
			`
		// Save LCL ARG THIS THAT on the stack
		result +=
			`@LCL
			D=M
			@SP
			A=M
			M=D
			@SP
			M=M+1
			`
		result +=
			`@ARG
			D=M
			@SP
			A=M
			M=D
			@SP
			M=M+1
			`
		result +=
			`@THIS
			D=M
			@SP
			A=M
			M=D
			@SP
			M=M+1
			`
		result +=
			`@THAT
			D=M
			@SP
			A=M
			M=D
			@SP
			M=M+1
			`
		// Set ARG
		result +=
			`@R13
			D=M
			`
		for i := 0; i < nArgs; i++ {
			result +=
				`D=D-1
				`
		}
		result +=
			`@ARG
			M=D
			`
		// Set LCL
		result +=
			`@SP
			D=M
			@LCL
			M=D
			`
		// Jump to excute function
		result +=
			`@` + callee + `
			0; JMP
			(` + returnLabel + `)
			`
	} else if command == Return {
		// Copy return value to argument 0
		result +=
			`@SP
			A=M-1
			D=M
			@ARG
			A=M
			M=D
			`
		// Save ARG to R13
		result +=
			`@ARG
			D=M
			@R13
			M=D
			`
		// Use R14 as current LCL pointer
		result +=
			`@LCL
			D=M
			@R14
			M=D
			`
		// Resotres the segment pointers of the caller
		result +=
			`@R14
			M=M-1
			A=M
			D=M
			@THAT
			M=D
			`
		result +=
			`@R14
			M=M-1
			A=M
			D=M
			@THIS
			M=D
			`
		result +=
			`@R14
			M=M-1
			A=M
			D=M
			@ARG
			M=D
			`
		result +=
			`@R14
			M=M-1
			A=M
			D=M
			@LCL
			M=D
			`
		// Set SP to ARG/R13 + 1
		result +=
			`@R13
			D=M+1
			@SP
			M=D
			`
		// Jump to return to caller
		result +=
			`@R14
			A=M
			0;JMP
			`
	}
	return strings.Replace(result, "\t", "", -1)
}
