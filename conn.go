package netutils

import (
	"fmt"
	"net"
	"net/url"
	"time"

	"github.com/jpillora/backoff"
)

// getHostnameFromString will return a hostname if a url is passed
// and attempt to add the correct ports
func getHostnameFromString(potentialHost string) (string, error) {
	parsedURL, err := url.Parse(potentialHost)
	if err != nil {
		return potentialHost, nil
	}
	if parsedURL.Scheme != "" && parsedURL.Host != "" {
		port := parsedURL.Port()
		if port == "" {
			// net requires a port
			switch parsedURL.Scheme {
			case "https":
				port = "443"
			case "http":
				port = "80"
			default:
				return "", fmt.Errorf("host appears to be url, but could not find port for scheme")
			}
		}
		return fmt.Sprintf("%s:%s", parsedURL.Hostname(), port), nil
	}
	return potentialHost, nil
}

// WaitForConnectTimeout will wait for a connection
// on a port progressively backing off
// returns false if we couldn't establish a connection by timeout
// after 10 timeouts
func WaitForConnectTimeout(host string, timeout time.Duration) bool {
	connected := false
	host, err := getHostnameFromString(host)
	if err != nil {
		return false
	}

	b := &backoff.Backoff{
		Factor: 2,
		Jitter: true,
		Min:    timeout / 10,
		Max:    timeout,
	}
	for {
		b.Attempt()
		conn, err := net.Dial("tcp", host)
		if err != nil {
			d := b.Duration()
			if d == timeout && b.Attempt() > 10 {
				return false
			}
			time.Sleep(d)
			continue
		}
		connected = true
		b.Reset()
		_ = conn.Close()
		break
	}
	return connected
}

// WaitForConnect on a port progressively backing off
// returns false if we couldn't establish a connection
// uses default timeout of 5 seconds
func WaitForConnect(host string) bool {
	return WaitForConnectTimeout(host, time.Second*5)
}
