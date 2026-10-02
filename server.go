package main

import (
	"net/http"
	"fmt"
	"os"
)

func main() {
	// این خط به سرور میگه وقتی کسی وارد شد، فایل index.html رو نشون بده
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	fmt.Println("Server is running on port", port)
	http.ListenAndServe(":"+port, nil)
}
