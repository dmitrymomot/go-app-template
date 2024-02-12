package clientip

import (
	"net"
	"net/http"
	"strings"
)

// LookupFromRequest returns the client IP address from the request.
func LookupFromRequest(r *http.Request) string {
	return lookupFromRequest(r)
}

// IP header names
var ipHeaders = []string{
	"DO_Connecting-IP", // DigitalOcean
	"DO-Connecting-IP",
	"True-Client-IP",
	"X-Real-IP",
	"CF-Connecting-IP", // Cloudflare
	"Fastly-Client-IP",
	"X-Cluster-Client-IP",
	"X-Client-IP",
}

// getRealIP returns the real IP address.
func lookupFromRequest(r *http.Request) string {
	var ip string
	for _, header := range ipHeaders {
		if ip = r.Header.Get(header); ip != "" {
			break
		}
	}
	if ip == "" {
		if xff := r.Header.Get("X-Forwarded-For"); xff != "" {
			i := strings.Index(xff, ", ")
			if i == -1 {
				i = len(xff)
			}
			ip = xff[:i]
		} else {
			var err error
			ip, _, err = net.SplitHostPort(r.RemoteAddr)
			if err != nil {
				ip = r.RemoteAddr
			}
		}
	}

	return canonicalizeIP(ip)
}

// canonicalizeIP returns a form of ip suitable for comparison to other IPs.
// For IPv4 addresses, this is simply the whole string.
// For IPv6 addresses, this is the /64 prefix.
func canonicalizeIP(ip string) string {
	isIPv6 := false
	// This is how net.ParseIP decides if an address is IPv6
	// https://cs.opensource.google/go/go/+/refs/tags/go1.17.7:src/net/ip.go;l=704
	for i := 0; !isIPv6 && i < len(ip); i++ {
		switch ip[i] {
		case '.':
			// IPv4
			return ip
		case ':':
			// IPv6
			isIPv6 = true
		}
	}
	if !isIPv6 {
		// Not an IP address at all
		return ip
	}

	ipv6 := net.ParseIP(ip)
	if ipv6 == nil {
		return ip
	}

	return ipv6.Mask(net.CIDRMask(64, 128)).String()
}
