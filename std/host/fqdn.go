package host

import (
	"net"
	"os"
	"strings"
)

func getHostFQDN() string {
	hostname, _ := os.Hostname()
	addrs, _ := net.LookupIP(hostname)
	for _, addr := range addrs {
		if ipv4 := addr.To4(); ipv4 != nil {
			names, err := net.LookupAddr(ipv4.String())
			if err == nil {
				for _, name := range names {
					if len(name) > 0 {
						return strings.TrimSuffix(name, ".")
					}
				}
			}
		}
	}
	return ""
}
