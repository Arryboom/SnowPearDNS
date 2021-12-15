package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"github.com/arryboom/go-hostsparser"
	"github.com/miekg/dns"
	"github.com/muesli/cache2go"
)

var (
	//server_url string = "http://119.29.29.29/d?dn=%s"
	server_url string = "https://doh.pub/dns-query?name=%s&type=%s"
	//server_url string
	version string = "1.6"
	//cache_time :=60*60*24
	dnsAcache  = cache2go.Cache("DNACACHE")
	dnsCcache  = cache2go.Cache("DNCCACHE")
	hostsflag  *bool
	cache_time = 60 * 60 * 24 * time.Second
)

type DOH_Answer struct {
	Name     string
	Thattype int `json:"type"`
	TTL      int
	Expires  string
	Data     string
}

/*{
	"name": "www.baidu.com.",
	"type": 5,
	"TTL": 1200,
	"Expires": "Wed, 15 Dec 2021 00:18:23 UTC",
	"data": "www.a.shifen.com."
}*/
type DOH_Request struct {
	NAME    string
	Thetype int `json:"type"`
}
type DOH_Response struct {
	Status             int
	TC                 bool          `json:"-"`
	RD                 bool          `json:"-"`
	RA                 bool          `json:"-"`
	AD                 bool          `json:"-"`
	CD                 bool          `json:"-"`
	Question           []DOH_Request `json:"-"`
	Answer             []DOH_Answer
	edns_client_subnet string `json:"-"`
}

//https://vsupalov.com/go-json-omitempty/ ignore json parse
func init_dohip() bool {
	var initurl string = "http://119.29.29.29/d?dn=doh.pub"
	//current resolve to 175.24.219.66,may change.
	var c_buf string
	r, err := http.Get(initurl)
	if err != nil {
		fmt.Println(err)
		return false
	}
	defer r.Body.Close()

	buf, err := ioutil.ReadAll(r.Body)
	//fmt.Println(string(buf))
	if err != nil {
		fmt.Println(err)
		return false
	}
	c_buf = byteString(buf)
	ips := strings.Split(c_buf, ";")
	randomIndex := rand.Intn(len(ips))
	pick := ips[randomIndex]
	if pick != "" {
		dnsAcache.Add("doh.pub.", 0, pick)
		fmt.Println(curtime() + " ###Init DOH resolve success.###")
		return true
	}
	return false
}
func get_a(domain string) []string {
	//Here we add cache
	var c_buf string
	ip := []string{}
	dncres, dncerr := dnsAcache.Value(domain)
	if dncerr == nil {
		//fmt.Println("Found value in cache:", dncres.Data().(string))
		//found cache
		c_buf = dncres.Data().(string)
	} else {
		//Didn't found cache
		url := fmt.Sprintf(server_url, domain, "A")

		r, err := http.Get(url)

		if err != nil {
			fmt.Println(err)
			return []string{}
		}

		defer r.Body.Close()

		buf, err := ioutil.ReadAll(r.Body)
		//fmt.Println(string(buf))
		if err != nil {
			fmt.Println(err)
			return []string{}
		}
		//here we add res to cache
		//dnscache.Add(domain,cache_time*time.Second,buf)
		//var c_buf string = byteString(buf)
		var resp DOH_Response
		if err := json.Unmarshal(buf, &resp); err != nil {
			fmt.Println(err)
			return []string{}
		} else {
			if resp.Status != 0 || resp.Answer == nil {
				fmt.Println(err)
				return []string{}
			}
		}
		//fmt.Printf("%+v\n", resp)
		for _, vl := range resp.Answer {
			//fmt.Println(vl)
			if vl.Thattype == 1 {
				if c_buf != "" {
					c_buf = c_buf + ";" + vl.Data
				} else {

					c_buf = vl.Data
				}

			}
			if vl.Thattype == 5 {
				dntres, dnterr := dnsCcache.Value(domain)
				if dnterr == nil {
					//fmt.Println("Found value in cache:", dncres.Data().(string))
					//found cache
					var tempc string = dntres.Data().(string)
					tempc = tempc + ";" + vl.Data
					dnsCcache.Delete(domain)
					dnsCcache.Add(domain, cache_time, tempc)
				} else {
					dnsCcache.Add(domain, cache_time, vl.Data)
				}
			}

		}
		//fmt.Println(c_buf)
		dnsAcache.Add(domain, cache_time, c_buf)

		//	dnscache.Add(domain,5*time.Second,c_buf)
	}

	ips := strings.Split(c_buf, ";")

	for _, ii := range ips {
		ip = append(ip, string(ii))
	}
	return ip

}
func get_cname(domain string) []string {
	//Here we add cache
	var c_buf string
	ip := []string{}
	dncres, dncerr := dnsCcache.Value(domain)
	if dncerr == nil {
		//fmt.Println("Found value in cache:", dncres.Data().(string))
		//found cache
		c_buf = dncres.Data().(string)
	} else {
		//Didn't found cache
		url := fmt.Sprintf(server_url, domain, "CNAME")

		r, err := http.Get(url)

		if err != nil {
			fmt.Println(err)
			return []string{}
		}

		defer r.Body.Close()

		buf, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Println(err)
			return []string{}
		}
		//here we add res to cache
		//dnscache.Add(domain,cache_time*time.Second,buf)
		//var c_buf string = byteString(buf)
		var resp DOH_Response
		if err := json.Unmarshal(buf, &resp); err != nil {
			fmt.Println(err)
			return []string{}
		} else {
			if resp.Status != 0 || resp.Answer == nil {
				fmt.Println(err)
				return []string{}
			}
		}
		for _, vl := range resp.Answer {
			//fmt.Println(vl)
			if vl.Thattype == 5 {
				if c_buf != "" {
					c_buf = c_buf + ";" + vl.Data
				} else {

					c_buf = vl.Data
				}

			}
			if vl.Thattype == 1 {
				dntres, dnterr := dnsAcache.Value(domain)
				if dnterr == nil {
					//fmt.Println("Found value in cache:", dncres.Data().(string))
					//found cache
					var tempc string = dntres.Data().(string)
					tempc = tempc + ";" + vl.Data
					dnsAcache.Delete(domain)
					dnsAcache.Add(domain, cache_time, tempc)
				} else {
					dnsAcache.Add(domain, cache_time, vl.Data)
				}
			}

		}
		dnsCcache.Add(domain, cache_time, c_buf)
		//	dnscache.Add(domain,5*time.Second,c_buf)
	}

	ips := strings.Split(c_buf, ";")

	for _, ii := range ips {
		ip = append(ip, string(ii))
	}

	return ip

}
func curtime() string {
	return time.Now().Format("2006-01-02 15:04:05.000")
}
func handleRoot(w dns.ResponseWriter, r *dns.Msg) {
	// Only A record supported
	if r.Question[0].Qtype == dns.TypeA {
		domain := r.Question[0].Name
		fmt.Println(curtime() + "   DnsReq_A: " + domain)
		ip := get_a(domain)

		if len(ip) == 0 {
			dns.HandleFailed(w, r)
			//fmt.Println("Failed to get DNS record: %s", domain)
			fmt.Println(curtime() + fmt.Sprintf("   Failed to get DNS record: %s", domain))
			return
		}

		msg := new(dns.Msg)
		msg.SetReply(r)

		for _, ii := range ip {
			s := fmt.Sprintf("%s 3600 IN A %s",
				dns.Fqdn(domain), ii)
			rr, _ := dns.NewRR(s)
			msg.Answer = append(msg.Answer, rr)
		}

		w.WriteMsg(msg)
		return
	}
	if r.Question[0].Qtype == dns.TypeCNAME {
		domain := r.Question[0].Name
		fmt.Println(curtime() + "   DnsReq_CNAME: " + domain)
		ip := get_cname(domain)

		if len(ip) == 0 {
			dns.HandleFailed(w, r)
			fmt.Println(curtime() + fmt.Sprintf("   Failed to get DNS record: %s", domain))
			return
		}

		msg := new(dns.Msg)
		msg.SetReply(r)

		for _, ii := range ip {
			s := fmt.Sprintf("%s 3600 IN CNAME %s",
				dns.Fqdn(domain), ii)
			rr, _ := dns.NewRR(s)
			msg.Answer = append(msg.Answer, rr)
		}

		w.WriteMsg(msg)
		return
	}

	dns.HandleFailed(w, r)
	return
}
func byteString(p []byte) string {
	for i := 0; i < len(p); i++ {
		if p[i] == 0 {
			return string(p[0:i])
		}
	}
	return string(p)
}
func add_localhosts() {
	if *hostsflag {
		fmt.Println("loading Hosts file...")
		hostsMap, err := hostsparser.ParseHosts(hostsparser.ReadHostsFile())
		if err != nil {
			return
		}
		for k, v := range hostsMap {
			dnsAcache.Add(k+".", 0, v)
		}
	}
}
func main() {
	fmt.Println("SnowPearDNS version: ", version)
	fmt.Println("https://github.com/arryboom/SnowPearDNS")
	hostsflag = flag.Bool("hosts", false, "using local hosts file,default false")
	flag.Parse()
	add_localhosts()
	fmt.Println("Start Dns Server Now...")
	if !(init_dohip()) {
		fmt.Println("Failed to init DOH server's DNS resolve,pls check your network connection.")
	}
	dns.HandleFunc(".", handleRoot)
	err := dns.ListenAndServe("0.0.0.0:53", "udp", nil)
	if err != nil {
		log.Fatal(err)
		fmt.Println("Failed to bind UDP port 53,please check your appliction and network.")
	}
}
