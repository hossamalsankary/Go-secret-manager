package main

import (
	"fmt"
	"net/http"
	vault "telnetapp/pkg/vault"
)

func main() {

	http.HandleFunc("/", vault.GetVaultToken)

	fmt.Printf("Server starting on port 8080... \n")
	http.ListenAndServe(":8080", nil)

}
