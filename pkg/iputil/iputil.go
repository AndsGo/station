package iputil

import (
	"errors"
	"net"
	"net/http"
	"net/url"
	"strings"
)

func GetIP(r *http.Request) (string, error) {
	ip := r.Header.Get("X-Forwarded-For")
	for _, i := range strings.Split(ip, ",") {
		if net.ParseIP(i) != nil {
			return i, nil
		}
	}

	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return "", err
	}

	if net.ParseIP(ip) != nil {
		return ip, nil
	}

	return "", errors.New("no valid ip found")
}

// 获取域名ip
func GetDomainIP(domain string) (string, error) {
	domain, err := extractDomain(domain)
	if err != nil {
		return "", err
	}
	ips, err := net.LookupIP(domain)
	if err != nil {
		return "", err
	}

	for _, ip := range ips {
		return ip.String(), err
	}
	return "", errors.New("no valid ip found")
}

// 提取域名
func extractDomain(u string) (string, error) {
	if strings.Contains(u, "http") {
		parsedUrl, err := url.Parse(u)
		if err != nil {
			return "", err
		}
		return parsedUrl.Hostname(), nil
	}
	return u, nil
}
