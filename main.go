package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var problemID int
	var err error
	if len(os.Args) > 1 {
		problemID, err = strconv.Atoi(os.Args[1])
		if err != nil {
			println("Usage: ./%s <q-id>")
			os.Exit(-1)
		}
	} else {
		println("Usage: ./%s <q-id>")
		os.Exit(-1)
	}

	fmt.Printf("Run question: %d\n", problemID)
	if runFunc, ok := Questions[problemID]; ok {
		ret := runFunc.(func() bool)()
		fmt.Printf("Run result: %t\n", ret)
	} else {
		fmt.Printf("Run %d hasn't not finish\n", problemID)
	}
}
