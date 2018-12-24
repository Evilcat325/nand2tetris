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

func main() {
	files := os.Args[1:]
	for _, filePath := range files {
		file, err := os.Open(filePath)
		check(err)
		asmFile, err := os.Create(strings.Replace(filePath, ".vm", ".asm", 1))
		check(err)
		defer file.Close()
		defer asmFile.Close()

		jumpCounter := 0
		scanner := bufio.NewScanner(file)
		scanner.Split(bufio.ScanLines)
		writer := bufio.NewWriter(asmFile)
		for scanner.Scan() {
			command, instructions := getCommandType(scanner)
			if Add <= command && command <= Not {
				fmt.Fprint(writer, arithmeticTranslator(command, instructions, &jumpCounter))
			} else if Pop <= command && command <= Push {
				fmt.Fprint(writer, memoryAccessTranslator(command, instructions))
			} else if Label <= command && command <= Goto {
				fmt.Fprint(writer, brachingTranslator(command, instructions))
			}
		}
		writer.Flush()
	}
}
