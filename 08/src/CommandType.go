package main

import (
	"bufio"
	"strings"
)

// TranslatorState keeps info for label generation
type TranslatorState struct {
	fileName      string
	functionName  string
	staticCounter int
	returnCounter int
	jumpCounter   int
	scanner       *bufio.Scanner
	writer        *bufio.Writer
	staticVarMap  map[string]int
}

// CommandType map line string to hack language type
type CommandType int

const (
	//Comment Comment or Empty string
	Comment CommandType = iota
	//Unknown undefined command
	Unknown
	//Add +
	Add
	//Sub -
	Sub
	//Neg * -1
	Neg
	//Eq ==
	Eq
	//Gt >
	Gt
	//Lt <
	Lt
	//And &
	And
	//Or |
	Or
	//Not !
	Not
	//Pop pop from stack to M
	Pop
	//Push push from M to stack
	Push
	//Label create a label for branching
	Label
	//IfGoto go to the label when last element on stack is true
	IfGoto
	//Goto got to the label unconditional
	Goto
	//Function declare a function
	Function
	//Call Jump to other function
	Call
	//Return signal return to the caller
	Return
)

//getCommandType return the CommandType of the next line from scanner
func getCommandType(scanner *bufio.Scanner) (CommandType, []string) {
	line := strings.TrimSpace(scanner.Text())
	firstPart := strings.Split(line, "//")[0]
	valueblePart := strings.Split(firstPart, " ")
	for i, str := range valueblePart {
		valueblePart[i] = strings.TrimSpace(str)
	}
	if len(valueblePart) <= 0 {
		return Comment, valueblePart
	} else if valueblePart[0] == "add" {
		return Add, valueblePart
	} else if valueblePart[0] == "sub" {
		return Sub, valueblePart
	} else if valueblePart[0] == "neg" {
		return Neg, valueblePart
	} else if valueblePart[0] == "eq" {
		return Eq, valueblePart
	} else if valueblePart[0] == "gt" {
		return Gt, valueblePart
	} else if valueblePart[0] == "lt" {
		return Lt, valueblePart
	} else if valueblePart[0] == "and" {
		return And, valueblePart
	} else if valueblePart[0] == "or" {
		return Or, valueblePart
	} else if valueblePart[0] == "not" {
		return Not, valueblePart
	} else if valueblePart[0] == "pop" {
		return Pop, valueblePart
	} else if valueblePart[0] == "push" {
		return Push, valueblePart
	} else if valueblePart[0] == "label" {
		return Label, valueblePart
	} else if valueblePart[0] == "if-goto" {
		return IfGoto, valueblePart
	} else if valueblePart[0] == "goto" {
		return Goto, valueblePart
	} else if valueblePart[0] == "function" {
		return Function, valueblePart
	} else if valueblePart[0] == "call" {
		return Call, valueblePart
	} else if valueblePart[0] == "return" {
		return Return, valueblePart
	}
	return Unknown, valueblePart
}
