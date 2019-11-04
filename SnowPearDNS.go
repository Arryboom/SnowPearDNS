package main

import (
	"strings"
	"fmt"
	"github.com/miekg/dns"
	"io/ioutil"
	"log"
	"net/http"
	"github.com/muesli/cache2go"
	"time"
	"github.com/arryboom/go-hostsparser"
	"flag"
)

var(
server_url string = "http://119.29.29.29/d?dn=%s"
version string ="1.3"
//cache_time :=60*60*24
dnscache= cache2go.Cache("DNCACHE")
hostsflag *bool
)


func get_a(domain string) []string {
//Here we add cache
	var c_buf string
	ip := []string{}
	dncres, dncerr := dnscache.Value(domain)
	if dncerr == nil {
		//fmt.Println("Found value in cache:", dncres.Data().(string))
		//found cache
		c_buf=dncres.Data().(string)
	} else {
	//Didn't found cache
	url := fmt.Sprintf(server_url, domain)

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
	var c_buf string =byteString(buf)
	dnscache.Add(domain,60*60*24*time.Second,c_buf)
//	dnscache.Add(domain,5*time.Second,c_buf)
	}

	ips := strings.Split(c_buf, ";")

	for _, ii := range ips {
		ip = append(ip, string(ii))
	}
	return ip
	
}

func handleRoot(w dns.ResponseWriter, r *dns.Msg) {
	// Only A record supported
	if r.Question[0].Qtype != dns.TypeA {
		dns.HandleFailed(w, r)
		return
	}

	domain := r.Question[0].Name
	fmt.Println("DnsReq: "+domain)
	ip := get_a(domain)

	if len(ip) == 0 {
		dns.HandleFailed(w, r)
		fmt.Println("Failed to get DNS record: %s",domain)
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
}
func byteString(p []byte) string {
        for i := 0; i < len(p); i++ {
                if p[i] == 0 {
                        return string(p[0:i])
                }
        }
        return string(p)
}
func add_localhosts(){
	if *hostsflag{
	hostsMap, err := hostsparser.ParseHosts(hostsparser.ReadHostsFile())
	if err != nil {
		return
	}
	for k,v :=range hostsMap{
	dnscache.Add(k+".",60*60*24*time.Second,v)
	}
}
}
func main() {
	fmt.Println("SnowPearDNS version: ",version)
	fmt.Println("https://github.com/arryboom/SnowPearDNS")
	hostsflag =flag.Bool("hosts",false,"using local hosts file,default false")
	flag.Parse()
	fmt.Println("Start Dns Server Now...")
	add_localhosts()
	dns.HandleFunc(".", handleRoot)
	err := dns.ListenAndServe("0.0.0.0:53", "udp", nil)
	if err != nil {
		log.Fatal(err)
		fmt.Println("Failed to bind UDP port 53,please check your appliction and network.")
	}
}
