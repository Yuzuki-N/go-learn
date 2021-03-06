package main

import (
	"fmt"
	"net"
	"sync"
	"time"
)

//func main() {
//	for i:= 21; i < 120; i++ {
//		address := fmt.Sprintf("20.194.168.28:%d", i)
//		conn, err := net.Dial("tcp", address)
//		if err != nil {
//			fmt.Printf("%s closed\n", address)
//			continue
//		}
//		conn.Close()
//		fmt.Printf("%s opened\n", address)
//	}
//}
// goroutine version

func main()  {
	var wg sync.WaitGroup
	start := time.Now()
	for i:= 21; i < 120; i++ {
		wg.Add(1)
		go func(j int) {
			defer wg.Done()
			address := fmt.Sprintf("20.194.168.28:%d", j)
			conn, err := net.Dial("tcp", address)
			if err != nil {
				fmt.Printf("%s closed\n", address)
				return
			}
			conn.Close()
			fmt.Printf("%s opened\n", address)
		}(i)
	}
	wg.Wait()
	fmt.Println("seconds: ", time.Since(start).Seconds())
}