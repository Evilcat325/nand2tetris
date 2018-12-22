package main

import (
	"bufio"
	"strings"
)

// LineType map line string to hack language type
type LineType int

const (
	//Comment Comment or Empty string
	Comment LineType = iota
	//JumpSymbol (XXX) Jump
	JumpSymbol
	//AInstruction @ Instruction
	AInstruction
	//CInstruction C Instruction
	CInstruction
)

var lineTypes = [...]string{
	"Comment",
	"Jump Symbol",
	"A Instruction",
	"C Instruction",
}

func (lineType LineType) String() string {
	return lineTypes[lineType]
}

//GetLineType return the Linetype of the next line from scanner
func GetLineType(scanner *bufio.Scanner) (LineType, string) {
	line := strings.TrimSpace(scanner.Text())
	firstPart := strings.Split(line, "//")[0]
	valueblePart := strings.TrimSpace(firstPart)
	if len(valueblePart) <= 0 {
		return Comment, valueblePart
	} else if valueblePart[0] == '(' {
		return JumpSymbol, valueblePart
	} else if valueblePart[0] == '@' {
		return AInstruction, valueblePart
	} else {
		return CInstruction, valueblePart
	}
}
