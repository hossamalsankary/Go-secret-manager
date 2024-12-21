package vault

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"os"
	"strings"

	Json "telnetapp/pkg/json"
)

func GetVaultToken(w http.ResponseWriter, r *http.Request) {

	// Restrictive access: only GET method allowed
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Get allowed IPs from environment variable
	ip := os.Getenv("VAULT_IPS")
	ipList := strings.Split(ip, ",")

	// Extract the client IP from the request
	remoteAddr := r.RemoteAddr
	clientIp, _, err := net.SplitHostPort(remoteAddr)
	if err != nil {
		fmt.Printf("Error parsing client IP: %v\n", err)
	}

	// Trim any surrounding whitespace
	clientIp = strings.TrimSpace(clientIp)

	// Check if the client IP is in the allowed list
	var ok bool
	for _, ip := range ipList {
		if clientIp == strings.TrimSpace(ip) {
			ok = true
			break
		}
	}

	if !ok {
		w.WriteHeader(http.StatusForbidden)
		http.Error(w, fmt.Sprintf("opps Forbidden access from this IP %s", clientIp), http.StatusForbidden)
		return
	}

	// Cache the JSON content
	cachedJSON := Json.VaultJson

	// Write the response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(cachedJSON); err != nil {
		http.Error(w, "Failed to encode JSON response", http.StatusInternalServerError)
		fmt.Printf("Error encoding response: %v\n", err)
	}
}
