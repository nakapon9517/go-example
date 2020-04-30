package main

import (
	f "fmt"
	. "strings"
	// _ "sub"
)

func main() {
	a := "aaa"
	b := "bbb"
	c := "ccc"
	d := "ddd"
	e := "eee"
	f.Println(ToUpper(a), ToUpper(b), ToUpper(c), ToLower(d), ToUpper(e))
	// f.Println(j)

	// if j == 0 {
	// 	f.Println("aa")
	// }

	j := 25
	for {
		if j == 0 {
			f.Println("ゼロです。")
		}
		if j%2 == 0 {
			f.Println("偶数：", j)
		}
		if j%2 == 1 {
			f.Println("奇数：", j)
		}
		if j == 0 {
			break
		} else {
			j--
			continue
		}
	}
}
