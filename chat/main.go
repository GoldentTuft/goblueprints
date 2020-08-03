package main

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"sync"
)

// templ は1つのテンプレートを表します
type templateHandler struct {
	once     sync.Once
	filename string
	templ    *template.Template
}

// ServeHTTP はHTTPリクエストを処理します
func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func() {
		t.templ =
			template.Must(template.ParseFiles(filepath.Join("templates",
				t.filename)))
	})
	t.templ.Execute(w, nil) // 戻り値はチェックすべき
}

func main() {
	http.Handle("/", &templateHandler{filename: "chat.html"})
	// Webサーバーを開始します
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
