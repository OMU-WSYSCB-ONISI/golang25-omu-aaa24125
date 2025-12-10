package main

import (
	"encoding/json"
	"fmt"
	"html"
	"net/http"
	"os"
	"runtime"
)

const saveFile = "public/memo.json" //データファイルの保存先

type Memo struct {
	Text string `json:"text"`
}

func main() {
	fmt.Printf("Go version: %s\n", runtime.Version())

	http.HandleFunc("/hello", hellohandler)
	http.HandleFunc("/memo", memo)
	http.HandleFunc("/mwrite", mwrite)

	fmt.Println("Launch server...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Failed to launch server: %v", err)
	}
}

func hellohandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "こんにちは from Codespace !")
}

func memo(w http.ResponseWriter, r *http.Request) {
	// データファイルを開く
	var memos []Memo
	data, err := os.ReadFile(saveFile)
	if err == nil {
		json.Unmarshal(data, &memos)
	}

	// HTMLのフォームを返す
	s := "<html><body>"
	s += "<h1>メモ帳</h1>"
	s += "<form method='post' action='/mwrite'>" +
		"<textarea name='text' style='width:99%; height:100px;'></textarea>" +
		"<input type='submit' value='追加' /></form>"

	s += "<h2>メモ一覧</h2><ul>"
	for _, m := range memos {
		s += "<li>" + html.EscapeString(m.Text) + "</li>"
	}
	s += "</ul></body></html>"

	w.Write([]byte(s))
}

func mwrite(w http.ResponseWriter, r *http.Request) {
    // 投稿されたフォームを解析
	r.ParseForm()
	if len(r.Form["text"]) == 0 {
		w.Write([]byte("フォームから投稿してください。"))
		return
	}
	newMemo := Memo{Text: r.Form["text"][0]}

	// 既存メモを読み込み
	var memos []Memo
	data, err := os.ReadFile(saveFile)
	if err == nil {
		json.Unmarshal(data, &memos)
	}

	// 新しいメモを追加
	memos = append(memos, newMemo)

	// JSONに保存
	jsonData, _ := json.MarshalIndent(memos, "", "  ")
	os.WriteFile(saveFile, jsonData, 0644)

	fmt.Println("save: " + newMemo.Text)

	// 一覧ページへリダイレクト
	http.Redirect(w, r, "/memo", http.StatusSeeOther)
}
