package main

import "fmt"

func main() {
	msg := "Hello......World!"
	fmt.Println(msg)
	bytes := addBytes(msg)
	fmt.Println(bytes)
}

func addBytes(s string) int {

	var res int

	for _, v := range s {
		res = res + int(v)
	}

	return res

}
