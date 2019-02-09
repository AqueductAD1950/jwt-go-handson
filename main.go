package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"jwt-go-handson/auth"
	"log"
	"net/http"
)

type Post struct {
	Title string `json:"title"`
	Tag string `json:"tag"`
	URL string `json:"url"`
}

func main() {
	r := mux.NewRouter()
	r.Handle("/public", public)
	r.Handle("/private", auth.JwtMiddleware.Handler(private))
	r.Handle("/auth", auth.GetTokenHandler)
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal("ListenAndServe Error:", err)
	}
}
var public = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	post := Post {
		Title:	"VueCLIからVue.js入門①【VueCLIで出てくるファイルを概要図で理解】",
		Tag: 	"Vue.js",
		URL: 	"https://qiita.com/po3rin/items/3968f825f3c86f9c4e21",
	}
	json.NewEncoder(w).Encode(post)
})

var private = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	post := &Post{
		Title: "VGolangとGoogle Cloud Vision APIで画像から文字認識するCLIを速攻でつくる",
		Tag:   "Go",
		URL:   "https://qiita.com/po3rin/items/bf439424e38757c1e69b",
	}
	json.NewEncoder(w).Encode(post)
})