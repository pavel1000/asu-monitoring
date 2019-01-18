package web

import (
	"encoding/json"
	"log"
	"net/http"
)

const (
	everythingIsNotOK int = iota
	somethingIsNotOK
	allIsOK
)

var hosts = []string{"https://asu.ru", "http://umc22.asu.ru", "http://journal.asu.ru", "http://support.asu.ru"}

// Web struct
type Web struct {
	Name string
}

// WebsitesInfo keeps: urls, that give different status from 200
// and status of all websites in general.
type WebsitesInfo struct {
	BadURLs      []string `json:"bad_urls"`
	GlobalStatus int      `json:"status"`
}

// Check verifies working of the ASU's websites
func (Web) Check() []byte {
	web := WebsitesInfo{}
	for _, host := range hosts {
		resp, err := http.Get(host)
		if err != nil {
			log.Printf("unable to get %v. %v\n", host, err)
		}
		if resp.StatusCode != http.StatusOK {
			log.Printf("%v вернула код %v\n", host, resp.StatusCode)
			web.BadURLs = append(web.BadURLs, host)
		}
	}

	if len(web.BadURLs) == len(hosts) {
		log.Println("Все странички не доступны!")
		web.GlobalStatus = everythingIsNotOK
	}
	if len(web.BadURLs) > 0 && len(web.BadURLs) < len(hosts) {
		log.Println("Некоторые странички не доступны!")
		web.GlobalStatus = somethingIsNotOK
	}
	if len(web.BadURLs) == 0 {
		log.Println("Все странички успешно работают!")
		web.GlobalStatus = allIsOK
	}

	b, err := json.Marshal(web)
	if err != nil {
		log.Fatal(err)
	}
	return b
}

// GetName returns "web" always
func (w Web) GetName() string {
	return w.Name
}
