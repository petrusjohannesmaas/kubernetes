package main

import (
	"encoding/json"
	"net/http"
)

func main() {
	http.HandleFunc("/", handleRoot)

	port := ":8080"
	println("🚀 Server is running on port", port)
	http.ListenAndServe(port, nil)
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case http.MethodGet:
		w.WriteHeader(http.StatusOK) // 200 OK
		json.NewEncoder(w).Encode(map[string]string{
			"message": "🎉 Your Go server is running",
		})
	case http.MethodPost:
		w.WriteHeader(http.StatusCreated) // 201 Created
		json.NewEncoder(w).Encode(map[string]string{
			"message": "✅ POST request received successfully",
		})
	default:
		w.WriteHeader(http.StatusMethodNotAllowed) // 405 Method Not Allowed
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Method not allowed",
		})
	}
}
