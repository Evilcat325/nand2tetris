package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func getJumpMap(filePath string) map[string]int {
	file, err := os.Open(filePath)
	check(err)
	defer file.Close()

	m := map[string]int{
		"R0":     0,
		"R1":     1,
		"R2":     2,
		"R3":     3,
		"R4":     4,
		"R5":     5,
		"R6":     6,
		"R7":     7,
		"R8":     8,
		"R9":     9,
		"R10":    10,
		"R11":    11,
		"R12":    12,
		"R13":    13,
		"R14":    14,
		"R15":    15,
		"SCREEN": 16384,
		"KBD":    24576,
		"SP":     0,
		"LCL":    1,
		"ARG":    2,
		"THIS":   3,
		"THAT":   4,
	}

	i := -1
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		lineType, instruction := GetLineType(scanner)
		if lineType == AInstruction || lineType == CInstruction {
			i++
		} else if lineType == JumpSymbol {
			m[strings.Trim(instruction, "()")] = i + 1
		}
	}
	return m
}

func translateToMachineCode(filePath string, m map[string]int, hackFilePath string) {
	file, err := os.Open(filePath)
	check(err)
	hackFile, err := os.Create(hackFilePath)
	check(err)
	defer file.Close()
	defer hackFile.Close()
	var nextSymbolAddr = 16

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	writer := bufio.NewWriter(hackFile)
	for scanner.Scan() {
		lineType, instruction := GetLineType(scanner)
		if lineType == AInstruction {
			fmt.Fprintln(writer, translateAInstruction(instruction, m, &nextSymbolAddr))
		} else if lineType == CInstruction {
			fmt.Fprintln(writer, translateCInstruction(instruction, m))
		}
	}
	writer.Flush()
}

func main() {
	files := os.Args[1:]
	for _, filePath := range files {
		jumpMap := getJumpMap(filePath)
		translateToMachineCode(filePath, jumpMap, strings.Replace(filePath, ".asm", ".hack", 1))
	}
}
