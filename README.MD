# SnowPearDNS
![GoReport](https://goreportcard.com/badge/github.com/arryboom/snowpeardns)  
![SnowPearDNS](https://github.com/Arryboom/SnowPearDNS/blob/master/ico/pear_128px.png)  
convert your dns request into httpdns request(DoH) and log them,Provide alt DNS to LAN with customized hosts.  
## Introduction

> ISP's DNS hijack In China has been drive me nuts for a long time,and I've been tried for lots of dns encryption projects,their encryption and RFC standard extend was really good,but due to most of these projects doesn't have DNS server in China,so it always makes my daily browsing slowly.

  finally,I found that DNSpod have a simple Httpdns API server in China and currently there has no evidence showed which ISP already hijack the http dns request between user's network and DNSpod server(although under China law dnspod has some wrong dns records as trade off but still better than hijacked by ISP), so I developed this tiny tool to make it work.  
  > by version 1.5 now this tool will use standard DoH which will access https://doh.pub/dns-query for results,which considered as a better security performance.

- Current Version:   **1.8**
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

>from version 1.8,you can now load dns records from ``-c <your hosts file location>`` option.  

## Tips
  
- typical deployment:  
![example](/typical_deploy.png)    
> to avoid unnessary disk io/speed up resolve the spdhosts.conf will load into memory when startup.  

1. Do you have a dns record cache in this tool?  
  >Yes,since sometimes http request will lost and DNSpod has limit on concurrency request made the http dns request very slow,I've add a memcache by using muesli's cache2go,default expiration time of one record haven't been hited is One day(60\*60\*24\*seconds).  Since it's cached in memory so everytime you restart this process its equal to force clear cache.   
2. How can I know DNSpod's server works good for me,what if there is a high latency to DNSpod server from my network?  
  >Well,you can using ping tool to check the latency between your network and DNSpod server by:  
  ```ping 119.29.29.29```  
  remember that the ICMP echo back latency in ms may not the real latency compare to do one http request to dnspod on your network.If it's really too slow,you may got a vps which you can trust and has no dns hijack in the cloud network and running a httpdns server like dnspod,and change the dns server url in SnowPearDNS.go,or you can do some encryption modify to improve the security performence.But The better way here I think you can try some other dns encryption project like
  https://github.com/bitbeans/SimpleDnsCrypt
  This opensource project has been provided a lots of dns encrypted server all over the world.  
3. How's the security of this tool?  
  >As I said at first "currently there has no evidence showed that which ISP already hijack the http dns request between user's network and DNSpod server In China",so if you got wrong dns record or you found your ISP hijacked the httpdns request,please let me know.(to be clear,the dnspod server itself may still gave you some wrong records,but as return you got dns records most resolve to China,this improves daily broswer experience if you physically in China,and dnspod is under China law so it got some wrong records but is still better if compare to ISP's default dns,which gave you more wrong records due to law,or they think its bad,or they been hacked or someone in their company was looking for more dark profit and choose some specific area to hijack,the main goal this tool invented at first was to solve this issue that is raw udp dns requests to dnspod was hijacked by ISP coz dns requests on udp 53 was plain text and ISP's fake records always return first and real requests or response was droped)
  >For security reason it will not using your local hosts file by default,but you can enable it by running it with **-hosts** params.    
  
  >by version 1.5 now this tool will use standard DoH which will access https://doh.pub/dns-query for results,which considered better security performance.remember that due to doh.pub is aslo a domain that need to resolve before we use it as DOH server,so there still need access old http dns server 119.119.119.119 first to get it's resolve when startup.  
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
