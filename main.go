package main

import (
	"fmt"
	"os"
)

func main() {
	tm, err := brubeck(os.Args[1:])
	if err != nil {
		panic(err)
	}
	fmt.Println(tm)
}
