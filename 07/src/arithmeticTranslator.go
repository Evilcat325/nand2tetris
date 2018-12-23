package main

import (
	"strconv"
	"strings"
)

func arithmeticTranslator(command CommandType, instructions []string, jumpCounter *int) string {
	var result string
	oneVar :=
		`@SP
		A=M-1
		`
	twoVar := oneVar +
		`D=M
		@SP
		M=M-1
		A=M-1
		`
	if command == Neg || command == Not {
		result = oneVar
		if command == Neg {
			result += "M=-M\n"
		} else if command == Not {
			result += "M=!M\n"
		}
	} else {
		result = twoVar
		if command == Add {
			result += "M=M+D\n"
		} else if command == Sub {
			result += "M=M-D\n"
		} else if command == And {
			result += "M=M&D\n"
		} else if command == Or {
			result += "M=M|D\n"
		} else if command == Eq || command == Gt || command == Lt {
			cmp := strings.ToUpper(instructions[0])
			result +=
				`D=M-D
				M=-1
				@D_` + cmp + `_M_` + strconv.Itoa(*jumpCounter) + `
				D;J` + cmp + `
				@SP
				A=M-1
				M=0
				(D_` + cmp + `_M_` + strconv.Itoa(*jumpCounter) + `)
				`
			*jumpCounter++
		}
	}
	return strings.Replace(result, "\t", "", -1)
}
