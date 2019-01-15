package proxydc

import (
	"log"
	"net/http"
	"net/url"
	"os"
	"time"
)

const method = "GET"

var (
	login = os.Getenv("PROXY_LOGIN")
	pass  = os.Getenv("PROXY_PASS")
)

// Check verifies working of proxy-dc.asu.ru
func Check() []byte {
	tr := &http.Transport{
		Proxy: func(r *http.Request) (*url.URL, error) {
			return url.Parse("http://" + login + ":" + pass + "@proxy-dc.asu.ru:3168")
		},
		DisableCompression: true,
		IdleConnTimeout:    20 * time.Second,
	}

	hosts := []string{"https://ya.ru", "http://asu.ru"}

	for _, host := range hosts {
		req, err := http.NewRequest(method, host, nil)
		if err != nil {
			log.Printf("Unable to make request for %v. %v Please try again later.", host, err)
			continue
		}
		_, err = tr.RoundTrip(req)
		if err != nil {
			log.Printf("Unable to %v %v. %v", method, host, err)
			return []byte("false")
		}
	}

	log.Println("proxy-dc.asu.ru успешно работает!")
	return []byte("true")
}
