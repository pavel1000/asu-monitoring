package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"

	"github.com/SimePel/asu-monitoring/mail"
	"github.com/SimePel/asu-monitoring/mx"
	"github.com/SimePel/asu-monitoring/vpn"
	"github.com/SimePel/asu-monitoring/proxy"
	"github.com/SimePel/asu-monitoring/proxyclass"
	"github.com/SimePel/asu-monitoring/proxykc"
	"github.com/SimePel/asu-monitoring/proxydc"
	"github.com/SimePel/asu-monitoring/proxysc"
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

// Statuses of the services
type Statuses struct {
	Proxy      string `json:"proxy"`
	ProxyClass string `json:"proxy_class"`
	ProxyKC    string `json:"proxy_kc"`
	ProxyDC    string `json:"proxy_dc"`
	ProxySC    string `json:"proxy_sc"`
	Mail       string `json:"mail"`
	MX         string `json:"mx"`
	Web        string `json:"web"`
	VPN		   string `json:"vpn"`
}

func servicesHandler(w http.ResponseWriter, r *http.Request) {
	mc := memcache.New("localhost:11211")
	items, err := mc.GetMulti([]string{"proxy", "proxy_dc", "proxy_sc", "proxy_kc", "proxy_class", "mail", "mx", "web", "vpn"})
	if err != nil {
		log.Fatal("unable to get items from memcached. ", err)
	}

	if len(items) == 0 {
		proxyItem := memcache.Item{
			Key:        "proxy",
			Value:      proxy.Check(),
			Expiration: 600,
		}
		proxyClassItem := memcache.Item{
			Key:        "proxy_class",
			Value:      proxyclass.Check(),
			Expiration: 600,
		}
		proxyKCItem := memcache.Item{
			Key:        "proxy_kc",
			Value:      proxykc.Check(),
			Expiration: 600,
		}
		proxyDCItem := memcache.Item{
			Key:        "proxy_dc",
			Value:      proxydc.Check(),
			Expiration: 600,
		}
		proxySCItem := memcache.Item{
			Key:        "proxy_sc",
			Value:      proxysc.Check(),
			Expiration: 600,
		}
		mailItem := memcache.Item{
			Key:        "mail",
			Value:      mail.Check(),
			Expiration: 600,
		}
		mxItem := memcache.Item{
			Key:        "mx",
			Value:      mx.Check(),
			Expiration: 600,
		}
		webItem := memcache.Item{
			Key:        "web",
			Value:      web.Check(),
			Expiration: 600,
		}
		vpnItem := memcache.Item{
			Key:        "vpn",
			Value:      vpn.Check(),
			Expiration: 600,
		}
		mc.Set(&proxyItem)
		mc.Set(&proxyClassItem)
		mc.Set(&proxyKCItem)
		mc.Set(&proxyDCItem)
		mc.Set(&proxySCItem)
		mc.Set(&mailItem)
		mc.Set(&mxItem)
		mc.Set(&webItem)
		mc.Set(&vpnItem)		
	}

	items, _ = mc.GetMulti([]string{"proxy", "proxy_dc", "proxy_sc", "proxy_kc", "proxy_class", "mail", "mx", "web", "vpn"})
	json.NewEncoder(w).Encode(
		Statuses{
			Proxy:      string(items["proxy"].Value),
			ProxyClass: string(items["proxy_class"].Value),
			ProxyKC:    string(items["proxy_kc"].Value),
			ProxyDC:    string(items["proxy_dc"].Value),
			ProxySC:    string(items["proxy_sc"].Value),
			Mail:       string(items["mail"].Value),
			MX:       	string(items["mx"].Value),
			Web:        string(items["web"].Value),
			VPN:        string(items["vpn"].Value),
		})
}
