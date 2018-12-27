package main

import (
	"bufio"
	"strings"
)

func addBootingInstruction(writer *bufio.Writer) {
	booting :=
		`@256
		D=A
		@SP
		M=D
		@Sys.init
		0;JMP
		`
	writer.WriteString(strings.Replace(booting, "\t", "", -1))
}
