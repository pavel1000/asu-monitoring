package vpn

import (
	"log"
	"net/http"
	"os/exec"
)

// Check verifies working of the vpn.asu.ru
func Check() []byte {
	cmd := exec.Command("sudo", "pon", "asuvpn")
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		cmd := exec.Command("sudo", "poff", "asuvpn")
		err := cmd.Run()
		if err != nil {
			log.Fatal(err)
		}
	}()

	hosts := []string{"https://ya.ru", "http://asu.ru"}

	for _, host := range hosts {
		resp, err := http.Get(host)
		if err != nil {
			log.Printf("unable to get %v. %v\n", host, err)
		}
		if resp.StatusCode != http.StatusOK {
			log.Printf("%v вернула код %v\n", host, resp.StatusCode)
			return []byte("false")
		}
	}

	log.Println("vpn.asu.ru работает в штатном режиме!")
	return []byte("true")
}
