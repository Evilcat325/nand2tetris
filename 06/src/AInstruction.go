package main

import (
	"fmt"
	"strconv"
	"strings"
)

func translateAInstruction(instruction string, m map[string]int, nextSymbolAddr *int) string {
	var addr string

	symbol := strings.TrimPrefix(instruction, "@")

	if n, err := strconv.Atoi(symbol); err == nil {
		addr = strconv.FormatInt(int64(n), 2)
	} else {
		if val, ok := m[symbol]; ok {
			addr = strconv.FormatInt(int64(val), 2)
		} else {
			m[symbol] = *nextSymbolAddr
			addr = strconv.FormatInt(int64(*nextSymbolAddr), 2)
			*nextSymbolAddr++
		}
	}
	return fmt.Sprintf("%016s", addr)
}
