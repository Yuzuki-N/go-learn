package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	// ex1.1
	fmt.Println(os.Args[0])
	// ex2.2
	for i, args := range os.Args[1:] {
		fmt.Println(i, ": ", args)
	}
	// ex1.3
	start := time.Now()
	fmt.Println(strings.Join(os.Args[1:], " "))
	fmt.Println("strings.Join cost time: ", time.Since(start).Seconds())

	s, sep := "", ""
	start = time.Now()
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println("for cost time: ", time.Since(start).Seconds())
}
