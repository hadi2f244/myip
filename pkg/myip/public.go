package myip

import (
	"io"
	"net/http"

	"github.com/rs/zerolog/log"
)

func getPublicIPv4() (string, error) {
	// conn, err := net.Dial("udp", "8.8.8.8:53")
	// if err != nil {
	// 	return "", err
	// }
	// defer conn.Close()
	// localAddr := conn.LocalAddr().String()
	// idx := strings.LastIndex(localAddr, ":")
	// ipaddr := localAddr[0:idx]
	// if !net.ParseIP(ipaddr).IsPrivate() {
	// 	return ipaddr, nil
	// }
	externalIP := ""
	// trying to get the public IP from multiple sources to see if they match.
	resp, err := http.Get("https://api.ipify.org")
	if err == nil {
		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		if err == nil {
			externalIP = string(body)
		}

		// // backup method of getting a public IP
		// if externalIP == "" {
		// 	// dig +short myip.opendns.com @208.67.222.222
		// 	dnsRes, _, err := c.dnsClient.performExternalAQuery("myip.opendns.com.", dns.TypeA)
		// 	if err != nil {
		// 		return "", err
		// 	}
		// 	externalIP = dnsRes[0].(*dns.A).A.String()
		// }

		if externalIP != "" {
			return externalIP, nil
		}
		log.Error().Msg("Could not automatically find the public IPv4 address. Please specify it in the configuration.")

	}
	return "", nil
}

func Public() string {
	ipv4, err := getPublicIPv4()
	if err == nil {
		return ipv4
	} else {
		return ""
	}
}