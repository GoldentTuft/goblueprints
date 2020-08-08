package main

import (
	"flag"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sync"

	"github.com/GoldentTuft/goblueprints/trace"
	"github.com/stretchr/gomniauth"
	"github.com/stretchr/gomniauth/providers/google"
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
	t.templ.Execute(w, r) // 戻り値はチェックすべき
}

func main() {
	var addr = flag.String("addr", ":8080", "アプリケーションのアドレス")
	flag.Parse()
	skey := os.Getenv("skey")
	if skey == "" {
		log.Fatal("skey(セキュリティキー)がない")
	}
	googleCID := os.Getenv("googleCID")
	if googleCID == "" {
		log.Fatal("googleCID(クライアントID)がない")
	}
	googleSV := os.Getenv("googleSV")
	if googleSV == "" {
		log.Fatal("googleSV(秘密の値)がない")
	}
	// Gomniauthのセットアップ
	gomniauth.SetSecurityKey(skey)
	gomniauth.WithProviders(
		google.New(googleCID, googleSV,
			"http://localhost:8080/auth/callback/google"),
	)
	r := newRoom()
	r.tracer = trace.New(os.Stdout)
	http.Handle("/chat", MustAuth(&templateHandler{filename: "chat.html"}))
	http.Handle("/login", &templateHandler{filename: "login.html"})
	http.HandleFunc("/auth/", loginHandler)
	http.Handle("/room", r)
	go r.run()
	// Webサーバーを開始します
	log.Println("Webサーバーを開始します。ポート: ", *addr)
	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
