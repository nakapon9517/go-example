package main

import (
	"fmt"
	data "funcproject/datastruct"
	jsontool "funcproject/jsontool"
)

func main() {

	// person1
	person := data.CreatePerson(1, "Gopher", "gopher@example.org", 5, "", "golang lover")
	byteData := jsontool.ConvertPersonToJSON(person)
	fmt.Println(string(byteData))

	// person2
	jsondata := []byte(`{"id":1,"name":"Gopher","age":5}`)
	person2 := jsontool.ConvertJSONToPerson(jsondata)
	fmt.Println(person2)

	// // ファイルを生成
	// file, err := os.Create("./file.txt")
	// errControl.ErrorHandler(err)
	// if err != nil { // エラー処理
	// 	log.Fatal(err)
	// }
	// // プログラムが終わったらファイルを閉じる
	// defer file.Close()
}
