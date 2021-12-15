# SnowPearDNS
![GoReport](https://goreportcard.com/badge/github.com/arryboom/snowpeardns)  
![SnowPearDNS](https://github.com/Arryboom/SnowPearDNS/blob/master/ico/pear_128px.png)  
A simple tool to convert your dns request into http dns request so we can avoid some dns hijack attack.
## Introduction

> ISP's DNS hijack In China has been drive me nuts for a long time,and I've been tried for lots of dns encryption projects,their encryption and RFC standard extend was really good,but due to most of these projects doesn't have DNS server in China,so it always makes my daily browsing slowly.

  Finally,I found that DNSpod have a simple Httpdns API server in China and currently there has no evidence showed which ISP already hijack the http dns request between user's network and DNSpod server, so I developed this tiny tool to make it work.
  > by version 1.5 now this tool will use standard DOH which will access https://doh.pub/dns-query for results,which considered as a better security performance.

- Current Version:   **1.5**
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
Size: 8172544 bytes
Modified: 2021-12-15, 12:03:25
MD5: 9A7DAB4F8D2327495DDEBBDCFF273C15
SHA1: 5E8AD32FFA3F54FC2B0DA3CD32B210CE5659D63E
CRC32: 3FD44751

File: snowpeardns_32.exe
Size: 7299072 bytes
Modified: 2021-12-15, 12:03:21
MD5: 72BFF01ADD8CCF2D1D9855F2A36B1C15
SHA1: 97C5D6089981D5D9B48260458048DBA89A2FB453
CRC32: EF39664A



File: snowpear_64  
Size: 8581391 bytes
Modified: 2021-12-15, 11:51:33
MD5: 4AE9BB5414BC5D46C3437DDD75A60B98
SHA1: 6894EEFD101C19CA1C413FA10E39C6920626B419
CRC32: 52E22173

File: snowpear_32  
Size: 7307992 bytes
Modified: 2021-12-15, 11:51:33
MD5: B58158F1077E9A8FAFB648DB9675EC59
SHA1: E0F1A11E002EE3C1A4D5AAD23FBD0ABB5949D207
CRC32: 8552DB29




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
  >For security reason it will not using your local hosts file by default,but you can enable it by running it with **-hosts** params.
  >by version 1.5 now this tool will use standard DOH which will access https://doh.pub/dns-query for results,which consider better security performance.

## Thanks
- miekg  
  https://github.com/miekg/dns
- fangdingjun  
  https://github.com/fangdingjun/gdns
- muesli  
  https://github.com/muesli/cache2go
