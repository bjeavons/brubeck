package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var args []string
	fileInfo, _ := os.Stdin.Stat()
	if fileInfo.Mode()&os.ModeNamedPipe != 0 {
		scanner := bufio.NewScanner(os.Stdin)
		var input string
		for scanner.Scan() {
			input = scanner.Text()
		}
		if err := scanner.Err(); err != nil {
			panic(err)
		}
		args = strings.Fields(input)
	} else {
		args = os.Args[1:]
	}

	tm, err := brubeck(args)
	if err != nil {
		panic(err)
	}
	fmt.Println(tm)
}
