package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func vmToAsm(state *TranslatorState) {
	for state.scanner.Scan() {
		command, instructions := getCommandType(state.scanner)
		if Add <= command && command <= Not {
			fmt.Fprint(state.writer, arithmeticTranslator(command, instructions, state))
		} else if Pop <= command && command <= Push {
			fmt.Fprint(state.writer, memoryAccessTranslator(command, instructions, state))
		} else if Label <= command && command <= Goto {
			fmt.Fprint(state.writer, brachingTranslator(command, instructions, state))
		} else if Function <= command && command <= Return {
			fmt.Fprint(state.writer, functionTranslator(command, instructions, state))
		}
		state.writer.Flush()
	}
}

func main() {
	paths := os.Args[1:]
	for _, path := range paths {
		pathInfo, err := os.Stat(path)
		check(err)
		state := TranslatorState{
			jumpCounter:   0,
			staticCounter: 0, // Counter increment per asm file rather than per vm file
			staticVarMap:  map[string]int{},
		}
		switch mode := pathInfo.Mode(); {
		case mode.IsDir():
			programName := filepath.Base(path)
			asmFile, err := os.Create(path + "/" + programName + ".asm")
			check(err)
			defer asmFile.Close()
			state.writer = bufio.NewWriter(asmFile)
			addBootingInstruction(state.writer)
			filepath.Walk(path, func(p string, f os.FileInfo, _ error) error {
				if filepath.Ext(p) == ".vm" {
					file, err := os.Open(p)
					check(err)
					defer file.Close()
					state.fileName = strings.TrimSuffix(filepath.Base(p), filepath.Ext(p))
					state.scanner = bufio.NewScanner(file)
					vmToAsm(&state)
				}
				return nil
			})

		case mode.IsRegular():
			file, err := os.Open(path)
			check(err)
			defer file.Close()
			asmFile, err := os.Create(strings.Replace(path, ".vm", ".asm", 1))
			check(err)
			defer asmFile.Close()
			state.fileName = strings.TrimSuffix(filepath.Base(path), filepath.Ext(path))
			state.staticCounter = 0
			state.scanner = bufio.NewScanner(file)
			state.writer = bufio.NewWriter(asmFile)
			vmToAsm(&state)
		}

	}
}
