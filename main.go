package main

import (
	"fmt"
	"log/slog"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	slog.Info("Handling endpoint: /")
	fmt.Fprintln(w, "{\"value\": \"Hello world\"}")
}

func main() {
	http.HandleFunc("/", helloHandler)

	port := ":8080"
	fmt.Println("Starting server on port", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		fmt.Println("Server failed:", err)
	}
}
