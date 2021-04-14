package main

import "fmt"

func consts()  {
	const(
		b = 1 << (10 * iota)
		kb
		mb
		gb
	)
	fmt.Println(b, kb, mb, gb)
}

func main() {
	consts()
}
