package main

import (
	vault "Secret-manager/pkg/vault"
	"encoding/json"
	"fmt"
	"net/http"
)

type CustomErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Path    string `json:"path"`
}

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Check if the request path matches "/vault-token"
		if r.URL.Path == "/" {
			vault.GetVaultToken(w, r)
			return
		}

		// For all other routes, respond with "404 Not Found"
		sendCustomError(w, http.StatusNotFound, "Oops! Lost? This is not the path you're looking for (404).", r.URL.Path)

	})
	fmt.Printf("Server starting on port 8080... \n")
	http.ListenAndServe(":8080", nil)

}
func sendCustomError(w http.ResponseWriter, statusCode int, message, path string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	errorResponse := CustomErrorResponse{
		Code:    statusCode,
		Message: message,
		Path:    path,
	}

	if err := json.NewEncoder(w).Encode(errorResponse); err != nil {
		fmt.Printf("Error encoding custom error response: %v\n", err)
	}
}
