package main

import "fmt"

func main() {
	var a [2]int
	a[0] = 100
	a[1] = 200
	fmt.Println(a)

	// var b = [2]int{100, 200}
	// b = append(b, 300)
	// fmt.Println(b)

	var b = []int{100, 200} //配列はリサイズできないから[2]のように指定しない
	b = append(b, 300)
	fmt.Println(b)
}
