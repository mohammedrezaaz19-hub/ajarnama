package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func main() {
	// سرو کردن فایل index.html برای مسیر اصلی
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})

	// مسیر جدید برای API
	http.HandleFunc("/api/status", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		response := map[string]string{
			"status":  "online",
			"message": "خدا قوت! سیستم با موفقیت در حال اجراست.",
		}
		json.NewEncoder(w).Encode(response)
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	fmt.Println("Server starting on port", port)
	http.ListenAndServe(":"+port, nil)
}

