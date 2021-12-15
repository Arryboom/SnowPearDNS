# SnowPearDNS
![GoReport](https://goreportcard.com/badge/github.com/arryboom/snowpeardns)  
![SnowPearDNS](https://github.com/Arryboom/SnowPearDNS/blob/master/ico/pear_128px.png)  
A simple tool to convert your dns request into http dns request so we can avoid some dns hijack attack.
## Introduction

> ISP's DNS hijack In China has been drive me nuts for a long time,and I've been tried for lots of dns encryption projects,their encryption and RFC standard extend was really good,but due to most of these projects doesn't have DNS server in China,so it always makes my daily browsing slowly.

  finally,I found that DNSpod have a simple Httpdns API server in China and currently there has no evidence showed which ISP already hijack the http dns request between user's network and DNSpod server, so I developed this tiny tool to make it work.
  > by version 1.5 now this tool will use standard DOH which will access https://doh.pub/dns-query for results,which considered as a better security performance.

- Current Version:   **1.6**
- Language:golang  

## Usage
**1.** Clone this project to your local disk.  
**2.** Open this project's folder and Compile SnowPearDNS.go to executables:
  ```
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
Size: 8179712 bytes  
Modified: 2021-12-15, 14:20:24  
MD5: 7AE99B3D4D2418703355C589AD81709E  
SHA1: 6CB6F452CC1A8C1C539ACE5D6829676AFDE25E16  
CRC32: 35456F76  

File: snowpeardns_32.exe   
Size: 7304704 bytes  
Modified: 2021-12-15, 14:20:20  
MD5: FF90C6A7CC712E7760E8F1CEB7B339EE  
SHA1: B2DC431442AC15AF1F722D67EAC6919BDBF46EEC  
CRC32: 0FF8BDC9  




File: snowpear_64   
Size: 8587245 bytes  
Modified: 2021-12-15, 14:18:00  
MD5: D7E44AB18F15492B8C89888E0718E300  
SHA1: C2187EB31ACC242984BE858EF5F1D1151533D0DB  
CRC32: 60F2029A   

File: snowpear_32  
Size: 7309386 bytes  
Modified: 2021-12-15, 14:18:01  
MD5: AC57333E5FE3D3FB2EE55D6789D6A2B6  
SHA1: E2B881388C5EA88C403DF4A69F933A1A6D73AE2E  
CRC32: 39C3F3BB  




## Tips

1. Do you have a dns record cache in this tool?  
  >Yes,since sometimes http request will lost and DNSpod has limit on concurrency request made the http dns request very slow,I've add a memcache by using muesli's cache2go,default expiration time of one record haven't been hited is One day(60\*60\*24\*seconds).  Since it's cached in memory so everytime you restart this process it will force clear cache.   
2. How can I know DNSpod's server works good for me,what if there is a high latency to DNSpod server from my network?  
  >Well,you can using ping tool to check the latency between your network and DNSpod server by:  
  ```ping 119.29.29.29```  
  remember that the ICMP echo back latency in ms may not the real latency compare to do one http request to dnspod on your network.If it's really too slow,you may got a vps which you can trust and has no dns hijack in the cloud network and running a httpdns server like dnspod,and change the dns server url in SnowPearDNS.go,or you can do some encryption modify to improve the security performence.But The better way here I think you can try some other dns encryption project like
  https://github.com/bitbeans/SimpleDnsCrypt
  This opensource project has been provided a lots of dns encrypted server all over the world.  
3. How's the security of this tool?  
  >As I said at first "currently there has no evidence showed that which ISP already hijack the http dns request between user's network and DNSpod server In China",so if you got wrong dns record or you found your ISP hijacked the httpdns request,please let me know and I'll consider to develop a better dns tool with encryption which has it's own server and client.   
  >For security reason it will not using your local hosts file by default,but you can enable it by running it with **-hosts** params.
  >by version 1.5 now this tool will use standard DOH which will access https://doh.pub/dns-query for results,which consider better security performance.
4. on windows there need at least 2 dns ipaddress set on ethernet settings?
  >the tool itself by default will listen on ```0.0.0.0```,so simple put main DNS as 127.0.0.1,another to your LAN address.(which like 192.168.1.x 172.0.1.x...)  
  >also by this you can serve the other teminal in your LAN if you like.

## Thanks
- miekg  
  https://github.com/miekg/dns
- fangdingjun  
  https://github.com/fangdingjun/gdns
- muesli  
  https://github.com/muesli/cache2go
