
package main

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"os"
)

type Quote struct {
	Text   string `json:"text"`
	Author string `json:"author"`
}

var quotes = []Quote{
	{"گر مرد رهی میان خون باید رفت / وز پای فتاده سرنگون باید رفت", "عطار"},
	{"هر که را صبح صادق آید پیش / معترف گردد او به عیب خویش", "مولانا"},
	{"در عشق تو هر چه هست و بود است / نیستی است و وجود توست", "حافظ"},
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})

	// مسیر جدید برای گرفتن شعر تصادفی
	http.HandleFunc("/api/quote", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		q := quotes[rand.Intn(len(quotes))]
		json.NewEncoder(w).Encode(q)
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	http.ListenAndServe(":"+port, nil)
}
