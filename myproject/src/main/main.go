package main

import (
	"errors"
	f "fmt"
	s "strings"
	// _ "sub"
)

func main() {
	// n, err := div(10, 0)
	// if err != nil {
	// 	// エラーを出力しプログラムを終了する
	// 	log.Fatal(err)
	// }
	// fmt.Println(n)

	//
	sub(5)
}

func sub(j int) (string, string) {
	f.Println("------------------")

	a := s.ToUpper("aaa")
	b := s.ToUpper("bbb")
	c := s.ToUpper("cc")
	d := s.ToUpper("ddd")
	e := s.ToUpper("eee")
	result := []string{a, b, c, d, e}
	result = append(result, "fff")
	f.Println(result[len(result)-1])
	f.Println("------------------")
	f.Println("------------------")

	for {
		if j%2 == 0 {
			switch {
			case j == 0:
				f.Println("ゼロです。")
				fallthrough
			default:
				f.Println("偶数：", j)
			}
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
	return "end1", "end2"
}
func div(i, j int) (result int, err error) {
	if j == 0 {
		err = errors.New("divied by zero")
		return // return 0, errと同じ
	}
	result = i / j
	return // return result, nilと同じ
}
