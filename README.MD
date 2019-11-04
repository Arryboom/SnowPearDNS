# SnowPearDNS
![GoReport](https://goreportcard.com/badge/github.com/arryboom/snowpeardns)  
![SnowPearDNS](https://github.com/Arryboom/SnowPearDNS/blob/master/ico/pear_128px.png)  
A simple tool to convert your dns request into http dns request so we can avoid some dns hijack attack.
## Introduction

> ISP's DNS hijack In China has been drive me nuts for a long time,and I've been tried for lots of dns encryption projects,their encryption and RFC standard extend was really good,but due to most of these projects doesn't have DNS server in China,so it always makes my daily browsing slowly.

  Finally,I found that DNSpod have a simple Httpdns API server in China and currently there has no evidence showed which ISP already hijack the http dns request between user's network and DNSpod server, so I developed this tiny tool to make it work.

- Current Version:   **1.3**
- Language:golang

## Usage
**1.** Clone this project to your local disk.  
**2.** Open this project's folder and Compile SnowPearDNS.go to executables:
  ```
  go get github.com/miekg/dns
  go build github.com/miekg/dns
  go get github.com/muesli/cache2go
  go get github.com/arryboom/go-hostsparser
  go build SnowPearDNS.go
  ```
  or you can directly download compiled executables for win if your are using Windows:  
  **[64bit_forWindows](https://github.com/Arryboom/SnowPearDNS/raw/master/release/snowpeardns_64.exe "snowpeardns64.exe")**  
  **[32bit_forWindows](https://github.com/Arryboom/SnowPearDNS/raw/master/release/snowpeardns_32.exe  "snowpeardns32.exe")**  
  or download here for linux if you are using Linux:  
  **[64bit_forLinux_ELF](https://github.com/Arryboom/SnowPearDNS/raw/master/release/snowpear_64 "snowpeardns64")**  
  **[32bit_forLinux_ELF](https://github.com/Arryboom/SnowPearDNS/raw/master/release/snowpear_32 "snowpeardns32")**  
**3.** Running the executables you got from last step and Configuration your local DNS server to 127.0.0.1.  
![example](/ex.png)  
**4.** Enjoy your browsing!


**FileInfo:** 

File: snowpeardns_64.exe  
Size: 8187904 bytes  
File Version: 1.3.0.0  
Modified: 2019-11-5, 5:38:18  
MD5: 63C204A73A28177F16522797853FDB0C  
SHA1: 299C7FE2517AAE60FCFE831AB3B458AE5B8D4FFC  
CRC32: D3E8FA32  

File: snowpeardns_32.exe  
Size: 7284224 bytes  
File Version: 1.3.0.0  
Modified: 2019-11-5, 5:38:03  
MD5: C9ABF800D2F3B98AB98CE11F46C1A1BE  
SHA1: D630A7C31FB43469B79397D9E545D3CC15D27AF1  
CRC32: CF8D9E99  



File: snowpear_64  
Size: 8454488 bytes  
Modified: 2019-11-5, 5:38:31  
MD5: 597408A44FBE67C52F6394497FA63AB1  
SHA1: 58E0BBEBFFF0860080B9E1A9ED7FE23C6C1C7513  
CRC32: AFF8CA53  

File: snowpear_32  
Size: 7514642 bytes  
Modified: 2019-11-5, 5:38:51  
MD5: 6D6CF9EEDCB91CC72C987AF53115A690  
SHA1: 773ED0573CF134BD76B33E16F9A0B01A7E5037C6  
CRC32: 75758C2F  



## Tips

1. Do you have a dns record cache in this tool?  
  >Yes,since sometimes http request will lost and DNSpod has limit on concurrency request made the http dns request very slow,I've add a memcache by using muesli's cache2go,default expiration time of one record haven't been hited is One day(60\*60\*24\*seconds).  Since it's cached in memory so everytime you restart this process it will force clear cache and regenerate them.   
2. How can I know DNSpod's server works good for me,what if there is a high latency to DNSpod server from my network?  
  >Well,you can using ping tool to check the latency between your network and DNSpod server by:  
  ```ping 119.29.29.29```  
  remember that the ICMP echo back latency in ms may not the real latency compare to do one http request to dnspod on your network.If it's really too slow,you may got a vps which you can trust and has no dns hijack in the cloud network and running a httpdns server like dnspod,and change the dns server url in SnowPearDNS.go,or you can do some encryption modify to improve the security performence.But The better way here I think you can try some other dns encryption project like
  https://github.com/bitbeans/SimpleDnsCrypt
  This opensource project has been provided a lots of dns encrypted server all over the world.  
3. How's the security of this tool?  
  >As I said at first "currently there has no evidence showed that which ISP already hijack the http dns request between user's network and DNSpod server In China",so if you got wrong dns record or you found your ISP hijacked the httpdns request,please let me know and I'll consider to develop a better dns tool with encryption which has it's own server and client.   
  >For security reason it will not using your local hosts file by default,but you can enable it by running it with ``-hosts`` params.

## Thanks
- miekg  
  https://github.com/miekg/dns
- fangdingjun  
  https://github.com/fangdingjun/gdns
- muesli  
  https://github.com/muesli/cache2go
