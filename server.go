package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	// این هندلر برای "همه مسیرها" است
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "سلام! سایت اجرنما بالا آمد. مسیر فعلی: %s", r.URL.Path)
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Println("Server starting on port", port)
	http.ListenAndServe(":"+port, nil)
}
