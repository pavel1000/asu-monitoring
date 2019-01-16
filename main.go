package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"

	"github.com/SimePel/asu-monitoring/mail"
	"github.com/SimePel/asu-monitoring/proxy"
	"github.com/SimePel/asu-monitoring/vpn"
	"github.com/SimePel/asu-monitoring/web"
	"github.com/bradfitz/gomemcache/memcache"
)

var (
	t = template.Must(template.New("T").ParseFiles("templates/index.html"))
)

func main() {
	server := &http.Server{Addr: ":8080", Handler: nil}

	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/services", servicesHandler)

	server.ListenAndServe()
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	t.ExecuteTemplate(w, "index", nil)
}

// Statuses of the each service
type Statuses struct {
	Proxy      string `json:"proxy"`
	ProxyClass string `json:"proxy_class"`
	ProxyKC    string `json:"proxy_kc"`
	ProxyDC    string `json:"proxy_dc"`
	ProxySC    string `json:"proxy_sc"`
	Mail       string `json:"mail"`
	MX         string `json:"mx"`
	Web        string `json:"web"`
	VPN        string `json:"vpn"`
}

func servicesHandler(w http.ResponseWriter, r *http.Request) {
	mc := memcache.New("localhost:11211")
	items, err := mc.GetMulti([]string{"proxy", "proxy-dc", "proxy-sc", "proxy-kc", "proxy-class", "mail", "mx", "web", "vpn"})
	if err != nil {
		log.Fatal("unable to get items from memcached. ", err)
	}

	if len(items) == 0 {
		servicesChan := make(chan Service, 9)
		itemsChan := make(chan *memcache.Item, 9)

		for w := 1; w <= 9; w++ {
			go worker(servicesChan, itemsChan)
		}

		servicesChan <- proxy.Proxy{Name: "proxy"}
		servicesChan <- proxy.Proxy{Name: "proxy-class"}
		servicesChan <- proxy.Proxy{Name: "proxy-kc"}
		servicesChan <- proxy.Proxy{Name: "proxy-dc"}
		servicesChan <- proxy.Proxy{Name: "proxy-sc"}
		servicesChan <- mail.Mail{Name: "mail"}
		servicesChan <- mail.Mail{Name: "mx"}
		servicesChan <- web.Web{Name: "web"}
		servicesChan <- vpn.VPN{Name: "vpn"}
		close(servicesChan)

		for i := 1; i <= 9; i++ {
			mc.Set(<-itemsChan)
		}
	}

	items, _ = mc.GetMulti([]string{"proxy", "proxy-dc", "proxy-sc", "proxy-kc", "proxy-class", "mail", "mx", "web", "vpn"})
	json.NewEncoder(w).Encode(
		Statuses{
			Proxy:      string(items["proxy"].Value),
			ProxyClass: string(items["proxy-class"].Value),
			ProxyKC:    string(items["proxy-kc"].Value),
			ProxyDC:    string(items["proxy-dc"].Value),
			ProxySC:    string(items["proxy-sc"].Value),
			Mail:       string(items["mail"].Value),
			MX:         string(items["mx"].Value),
			Web:        string(items["web"].Value),
			VPN:        string(items["vpn"].Value),
		})
}

// Service has Check method that verifies working of it
// and GetName method that give you name of the service
type Service interface {
	Check() []byte
	GetName() string
}

func worker(services <-chan Service, items chan<- *memcache.Item) {
	for service := range services {
		items <- &memcache.Item{
			Key:        service.GetName(),
			Value:      service.Check(),
			Expiration: 600,
		}
	}
}
