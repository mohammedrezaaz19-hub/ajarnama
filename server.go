package main

import (
	"fmt"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		fmt.Fprintln(w, "<h1>سلام! 👋</h1><p>به وب‌سرور Go من خوش آمدی.</p>")
	case "/about":
		fmt.Fprintln(w, "<h1>درباره</h1><p>این یک وب‌سرور ساده با زبان Go است.</p>")
	default:
		http.NotFound(w, r)
	}
}

func main() {
	http.HandleFunc("/", handler)

	// دریافت پورت از سیستم (بسیار مهم برای Railway)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Println("Server starting on port", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		fmt.Println("Error:", err)
	}
}
