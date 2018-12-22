package main

import (
	"strings"
)

var compMap = map[string]string{
	"0":   "101010",
	"1":   "111111",
	"-1":  "111010",
	"D":   "001100",
	"A":   "110000",
	"!D":  "001101",
	"!A":  "110001",
	"-D":  "001111",
	"-A":  "110011",
	"D+1": "011111",
	"A+1": "110111",
	"D-1": "001110",
	"A-1": "110010",
	"D+A": "000010",
	"D-A": "010011",
	"A-D": "000111",
	"D&A": "000000",
	"D|A": "010101",
}
var destMap = map[string]string{
	"NULL": "000",
	"M":    "001",
	"D":    "010",
	"MD":   "011",
	"A":    "100",
	"AM":   "101",
	"AD":   "110",
	"AMD":  "111",
}
var jumpMap = map[string]string{
	"NULL": "000",
	"JGT":  "001",
	"JEQ":  "010",
	"JGE":  "011",
	"JLT":  "100",
	"JNE":  "101",
	"JLE":  "110",
	"JMP":  "111",
}

func translateCInstruction(instruction string, m map[string]int) string {
	var destAndComp []string
	var assignPart = instruction
	var a = "0"
	var comp string
	var jump string

	if strings.Contains(instruction, ";") {
		parts := strings.Split(instruction, ";")
		assignPart = parts[0]
		jump = strings.Replace(parts[1], " ", "", -1)
	} else {
		jump = "NULL"
	}

	if strings.Contains(instruction, "=") {
		destAndComp = strings.Split(assignPart, "=")
	} else {
		destAndComp = strings.Split("NULL="+assignPart, "=")
	}

	dest := strings.Replace(destAndComp[0], " ", "", -1)
	compwitha := strings.Replace(destAndComp[1], " ", "", -1)
	if strings.Contains(compwitha, "M") {
		a = "1"
	}
	comp = strings.Replace(compwitha, "M", "A", 1)
	return "111" + a + compMap[comp] + destMap[dest] + jumpMap[jump]
}
