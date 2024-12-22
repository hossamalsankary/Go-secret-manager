package vault

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"os"
	"strings"

	Json "Secret-manager/pkg/json"
)

// GetVaultToken handles the GET request to retrieve a Vault token from your cache.
// It restricts access based on IP allowlist via environment variable VAULT_IPS.
func GetVaultToken(w http.ResponseWriter, r *http.Request) {
	// Restrictive access: only GET method allowed
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Get allowed IPs from environment variable
	allowedIPs := os.Getenv("VAULT_IPS")
	ipList := strings.Split(allowedIPs, ",")

	// ----------------------------
	// Extract the client IP
	// ----------------------------
	// 1. Check if X-Forwarded-For header is set (the first IP is typically the client IP).
	xForwardedFor := r.Header.Get("X-Forwarded-For")

	var clientIP string
	if xForwardedFor != "" {
		// If multiple IPs are present, split on comma and take the first one.
		forwardedIPs := strings.Split(xForwardedFor, ",")
		clientIP = strings.TrimSpace(forwardedIPs[0])
	} else {
		// Fallback to r.RemoteAddr if X-Forwarded-For is empty
		remoteAddr := r.RemoteAddr
		var err error
		clientIP, _, err = net.SplitHostPort(remoteAddr)
		if err != nil {
			fmt.Printf("Error parsing client IP from RemoteAddr: %v\n", err)
		}
		// Trim any surrounding whitespace
		clientIP = strings.TrimSpace(clientIP)
	}

	fmt.Println("Client IP:", clientIP)

	// ----------------------------
	// Check if the client IP is in the allowed list
	// ----------------------------
	var allowed bool
	for _, ip := range ipList {
		if clientIP == strings.TrimSpace(ip) {
			allowed = true
			break
		}
	}

	if !allowed {
		w.WriteHeader(http.StatusForbidden)
		errMsg := fmt.Sprintf("Forbidden access from IP: %s", clientIP)
		http.Error(w, errMsg, http.StatusForbidden)
		return
	}

	// ----------------------------
	// Return the cached JSON content
	// ----------------------------
	cachedJSON := Json.VaultJson

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(cachedJSON); err != nil {
		http.Error(w, "Failed to encode JSON response", http.StatusInternalServerError)
		fmt.Printf("Error encoding response: %v\n", err)
	}
}
