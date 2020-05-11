package jsontool

import (
	"encoding/json"
	"fmt"
	errorhandle "funcproject/error"
	"net/http"
	"os"
	"strconv"
)

// Person struct
type Person struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// IndexHandler index handler
func IndexHandler(w http.ResponseWriter,
	r *http.Request) {
	fmt.Fprint(w, "hello world")
}

// PersonHandler person handler
func PersonHandler(w http.ResponseWriter,
	r *http.Request) {
	defer r.Body.Close() // 処理の最後にBodyを閉じる

	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	errorhandle.ErrorHandler(err)
	fmt.Println(id)

	if r.Method == "POST" {
		// リクエストボディをJSONに変換
		var person Person
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&person)
		errorhandle.ErrorHandler(err)

		// ファイル名を {id}.txtとする
		filename := fmt.Sprintf("%d.txt", person.ID)
		file, err := os.Create(filename) // ファイルを生成
		errorhandle.ErrorHandler(err)
		defer file.Close()

		// ファイルにNameを書き込む
		_, err = file.WriteString(person.Name)
		errorhandle.ErrorHandler(err)

		// レスポンスとしてステータスコード201を送信
		w.WriteHeader(http.StatusCreated)
	}
}

// RequestCall requestCall
func RequestCall(url1 string, url2 string, port string) {
	http.HandleFunc(url1, IndexHandler)
	http.HandleFunc(url2, PersonHandler)
	http.ListenAndServe(port, nil)
}
