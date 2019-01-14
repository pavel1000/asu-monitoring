package web

import (
	"log"
	"net/http"
)

const method = "GET"

var hosts = []string{"https://asu.ru", "http://umc22.asu.ru"}

// Check verifies working of the ASU's websites
func Check() bool {
	badReqs := 0
	for _, host := range hosts {
		resp, err := http.Get(host)
		if err != nil {
			log.Printf("unable to get %v. %v\n", host, err)
		}
		if resp.StatusCode != http.StatusOK {
			log.Printf("%v вернула код %v\n", host, resp.StatusCode)
			badReqs++
		}
	}
	if badReqs == len(hosts) {
		return false
	}
	log.Println("Все странички успешно откликаются!")
	return true
}
