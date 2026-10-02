package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

// ساختار داده برای API وضعیت سایت
type StatusResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Version string `json:"version"`
}

func main() {
	// خواندن پورت از محیط ریلی‌وی (یا پیش‌فرض 8080)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// ۱. مسیر صفحه اصلی سایت
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		http.ServeFile(w, r, "index.html")
	})

	// ۲. مسیر API برای بررسی وضعیت سرور
	http.HandleFunc("/api/status", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		response := StatusResponse{
			Status:  "online",
			Message: "سرور اجرنما با قدرت روی Go و Railway اجرا می‌شود 🚀",
			Version: "1.1.0",
		}
		json.NewEncoder(w).Encode(response)
	})

	fmt.Println("Server is running on port " + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
