package main

import (
	"strconv"
	"strings"
)

var segmentMap = map[string]int{
	"local":    1,
	"argument": 2,
	"this":     3,
	"that":     4,
	"pointer":  3,
	"temp":     5,
	"static":   16,
}

func memoryAccessTranslator(command CommandType, instructions []string, jumpCounter *int) string {
	segment := instructions[1]
	index, err := strconv.Atoi(instructions[2])
	check(err)
	result := ""
	if command == Push {
		//load value to D
		if segment == "constant" {
			result += "@" + strconv.Itoa(index) + "\n"
			result += "D=A\n"
		} else {
			result += "@" + strconv.Itoa(segmentMap[segment]) + "\n"
			if segment == "temp" || segment == "pointer" || segment == "static" {
				result += "D=A\n"
			} else {
				result += "D=M\n"
			}
			result += "@" + strconv.Itoa(index) + `
			A=D+A
			D=M
			`
		}
		result +=
			`@SP
			A=M
			M=D
			@SP
			M=M+1
			`
	} else {
		result += "@" + strconv.Itoa(segmentMap[segment]) + "\n"
		if segment == "temp" || segment == "pointer" || segment == "static" {
			result += "D=A\n"
		} else {
			result += "D=M\n"
		}
		result +=
			`@` + strconv.Itoa(index) + `
			D=D+A
			@R15
			M=D
			@SP
			M=M-1
			A=M
			D=M
			@R15
			A=M
			M=D
			`
	}
	return strings.Replace(result, "\t", "", -1)
}
