package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	// تعریف هندلر برای صفحه اصلی
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		fmt.Fprintf(w, "به سایت اجرنما خوش آمدید! صفحه اصلی.")
	})

	// تعریف هندلر برای صفحه درباره
	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "این صفحه درباره ما است. اینجا مسیر اجرنماست!")
	})

	// تنظیم پورت برای Railway
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	
	fmt.Println("Server is running on port", port)
	http.ListenAndServe(":"+port, nil)
}
