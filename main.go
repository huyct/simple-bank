package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println("duckhue01")
		}()
	}
	fmt.Println("duckhue01")
}
