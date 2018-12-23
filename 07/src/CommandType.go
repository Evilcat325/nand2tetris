package main

import (
	"bufio"
	"strings"
)

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
)

//getCommandType return the CommandType of the next line from scanner
func getCommandType(scanner *bufio.Scanner) (CommandType, []string) {
	line := strings.TrimSpace(scanner.Text())
	firstPart := strings.Split(line, "//")[0]
	valueblePart := strings.Split(firstPart, " ")
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
	}
	return Unknown, valueblePart
}
