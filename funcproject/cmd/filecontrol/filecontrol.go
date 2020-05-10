package filecontrol

import (
	"fmt"
	errControl "funcproject/error"
	"io/ioutil"
	"os"
)

// CreateFile file create
func CreateFile(filePath string) *os.File {
	fmt.Println("createfile--------------")

	// ファイルを生成
	file, err := os.Create(filePath)
	errControl.ErrorHandler(err)
	// プログラムが終わったらファイルを閉じる
	defer file.Close()
	return file
}

// ReadFile file read
func ReadFile(filePath string) []byte {
	fmt.Println("readfile--------------")
	// ファイルの中身をすべて読み出す
	file, _ := os.Open(filePath)
	message, err := ioutil.ReadAll(file)
	errControl.ErrorHandler(err)
	// プログラムが終わったらファイルを閉じる
	defer file.Close()
	return message
}

// WriteFile file write
// func WriteFile(filePath string, message [2]string) {
func WriteFile(filePath string) {
	var str [2]string
	str[0] = `{"id":1,"name":"Gopher"}`
	str[1] = `{"id":2,"name":"Gopher"}`
	fmt.Println("writefile--------------")
	// ファイルを生成
	file, err := os.Create(filePath)
	errControl.ErrorHandler(err)

	// write
	for argNum, arg := range str {
		if argNum < len(str) {
			fmt.Fprint(file, arg)
			fmt.Fprint(file, "\n")
		}
	}
	defer file.Close()
}
