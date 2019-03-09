package main

import "fmt"

func main() {
	for i := 0; i < 100; i++ {
		defer fmt.Println(i * i) //一番最後に実行される
		if i%2 == 1 {
			continue
		}
		fmt.Println(i)
	}
}
