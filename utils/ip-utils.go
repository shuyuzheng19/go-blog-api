package utils

import (
	"net/http"
	"strings"
)

func GetIPAddress(request *http.Request) string {
	ipAddress := request.Header.Get("X-Forwarded-For")
	if ipAddress == "" || strings.ToLower(ipAddress) == "unknown" {
		ipAddress = request.Header.Get("Proxy-Client-IP")
	}
	if ipAddress == "" || strings.ToLower(ipAddress) == "unknown" {
		ipAddress = request.Header.Get("WL-Proxy-Client-IP")
	}
	if ipAddress == "" || strings.ToLower(ipAddress) == "unknown" {
		ipAddress = request.RemoteAddr
	}
	return ipAddress
}
