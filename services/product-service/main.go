package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, `
    <div style="text-align: center; margin-top: 50px; font-family: sans-serif;">
      <h1 style="color: #00BCD4;">Product Service</h1>
      <p>Status: 🟢 <strong>Online</strong></p>
      <p>Version: 1.0</p>
      <p>Technology: Go (Golang)</p>
    </div>
	`)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8081", nil)
}
