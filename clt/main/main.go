package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	//var s, sep string
	//s, sep := "", ""

	//for _, arg := range os.Args[1:] {
	//	s += sep + arg
	//	sep = " "
	//}


	//for i:= 1; i < len(os.Args); i++ {
	//	s += sep + os.Args[i]
	//	sep = " "
	//}

	fmt.Println(strings.Join(os.Args[1:], ""))
}
