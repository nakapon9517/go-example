package main

import (
	"errors"
	f "fmt"
	"log"
	s "strings"
)

func main() {

	// 配列１
	target := []int{1, 2, 3}
	// 配列２
	target2 := []int{4, 5, 6}

	// 配列１，２を引数に
	// 返り値は正常メッセージとエラー情報
	errIndex := 5
	result, err := sub1(errIndex, target, target2)

	// エラー情報が存在する場合
	if err != nil {
		// エラー出力＆処理終了
		log.Fatal(err)
	}

	// 正常メッセージ出力
	f.Println(result)
}

// 引数はint型の配列
// 返り値指定（result, err)
func sub1(errIndex int, num ...[]int) (result string, err error) {

	// 配列結合
	outputStr := append(num[0], num[1]...)

	// 配列の回数処理実行
	for i, str := 0, outputStr; i < len(str); i++ {

		if i != errIndex {
			// 配列出力
			f.Println(outputStr[i])
		} else {
			// 自作エラー作成＆return
			err = errors.New("error")
			return
		}
	}
	result = "正常終了"
	return
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
