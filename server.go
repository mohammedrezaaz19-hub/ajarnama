package main

import (
	"fmt"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "به سایت اجرنما خوش آمدید!")
}

func main() {
	http.HandleFunc("/", handler)
	
	// دریافت پورت از سیستم (مهم برای هاست رایگان)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Println("Server starting on port", port)
	http.ListenAndServe(":"+port, nil)
}
package main

import (
	"fmt"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "به سایت اجرنما خوش آمدید!")
}

func main() {
	http.HandleFunc("/", handler)
	
	// دریافت پورت از سیستم (مهم برای هاست رایگان)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Println("Server starting on port", port)
	http.ListenAndServe(":"+port, nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
switch r.URL.Path {
fmt.Fprintln(w, "<h1>سلام! 👋</h1><p>به وب‌سرور Go من خوش آمدی.</p>")
case "/about":
fmt.Fprintln(w, "<h1>درباره</h1><p>این یک وب‌سرور ساده با زبان Go است.</p>")
default:
http.NotFound(w, r)
}
}

func main() {
http.HandleFunc("/", handler)

fmt.Println("Server starting at http://127.0.0.1:8080")
if err := http.ListenAndServe(":8080", nil); err != nil {
fmt.Println("Error:", err)
}
}
