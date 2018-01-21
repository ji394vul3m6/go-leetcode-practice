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
	ret := Questions[problemID].(func() bool)()
	fmt.Printf("Run result: %t\n", ret)
}
