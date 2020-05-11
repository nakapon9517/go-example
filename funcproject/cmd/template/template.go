package template

import (
	"encoding/json"
	"fmt"
	errorhandle "funcproject/error"
	jsontool "funcproject/jsontool"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"text/template"
)

var t = template.Must(template.ParseFiles("index.html"))

// PersonHandler person handler
func PersonHandler(w http.ResponseWriter,
	r *http.Request) {
	defer r.Body.Close() // 処理の最後にBodyを閉じる

	if r.Method == "POST" {
		// リクエストボディをJSONに変換
		var person jsontool.Person
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&person)
		errorhandle.ErrorHandler(err)

		// ファイル名を{id}.txtとする
		filename := fmt.Sprintf("%d.txt", person.ID)
		file, err := os.Create(filename) // ファイルを生成
		errorhandle.ErrorHandler(err)
		defer file.Close()

		// ファイルにNameを書き込む
		_, err = file.WriteString(person.Name)
		errorhandle.ErrorHandler(err)

		// レスポンスとしてステータスコード201を送信
		w.WriteHeader(http.StatusCreated)
	} else if r.Method == "GET" {
		// パラメータを取得
		id, err := strconv.Atoi(r.URL.Query().Get("id"))
		errorhandle.ErrorHandler(err)

		filename := fmt.Sprintf("%d.txt", id)
		b, err := ioutil.ReadFile(filename)
		errorhandle.ErrorHandler(err)

		// personを生成
		person := jsontool.Person{
			ID:   id,
			Name: string(b),
		}

		// レスポンスにエンコーディングしたHTMLを書き込む
		t.Execute(w, person)
	}
}

// RequestCall requestCall
func RequestCall(url1 string, url2 string, port string) {
	http.HandleFunc(url1, PersonHandler)
	http.ListenAndServe(port, nil)
}
