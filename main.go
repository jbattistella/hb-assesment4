package main

import "fmt"

func main() {
	msg := "Hello......World!"
	fmt.Println(msg)
	info := charInfo(msg)
	fmt.Println(info)
}

func charInfo(s string) map[string]int {

	var res = make(map[string]int)
	var byteCt int

	for _, v := range s {
		byteCt = byteCt + int(v)
	}

	res["Byte Count"] = byteCt
	res["Character Count"] = len(s)

	return res

}
