# SnowPearDNS
![SnowPearDNS](https://github.com/Arryboom/SnowPearDNS/blob/master/ico/pear_128px.png)  
A simple tool to convert your dns request into http dns request so we can avoid some dns hijack attack.
## Introduction

> ISP's DNS hijack In China has been drive me nuts for a long time,and I've been tried for lots of dns encryption projects,their encryption and RFC standard extend was really good,but due to most of these projects doesn't have DNS server in China,so it always makes my daily browsing slowly.

  Finally,I found that DNSpod have a simple Httpdns API server in China and currently there has no evidence showed ISP will hijack the http dns request between user's network and DNSpod server,so I developed this tiny tool to make it work.

- Current Version:   **1.1**
- Language:golang

## Usage
**1.** Clone this project to your local disk.
**2.** Open this project's folder and Compile SnowPearDNS.go to executables:
  ```
  go get github.com/miekg/dns
  go build github.com/miekg/dns
  go get github.com/muesli/cache2go
  go build httpdns_plugs.go
  ```
  or you can download compiled executables for win if your are using Windows:  
  **[64bit_forWindows](https://github.com/Arryboom/SnowPearDNS/raw/master/release/snowpeardns_64.exe "snowpeardns64.exe")**  
  **[32bit_forWindows](https://github.com/Arryboom/SnowPearDNS/raw/master/release/snowpeardns_32.exe  "snowpeardns32.exe")**  
  or download here for linux if you are using Linux:  
  **[64bit_forLinux_ELF](https://github.com/Arryboom/SnowPearDNS/raw/master/release/snowpear_64 "snowpeardns64")**  
  **[32bit_forLinux_ELF](https://github.com/Arryboom/SnowPearDNS/raw/master/release/snowpear_32 "snowpeardns32")**  
**3.** Running the executables you got from last step and Configuration your local DNS server to 127.0.0.1.  
**4.** Enjoy your browsing!


**FileInfo:** 

File: snowpeardns_64.exe  
Size: 7316480 bytes  
File Version: 1.1.0.0  
Modified: 03-02-2018,17:28:32  
MD5: F2A1F3B99D6A8E3C063D53D0F80C83D2  
SHA1: DFE2F32B174DB42858C145254761BD78A537D39C  
CRC32: 6DB2AC56  

File: snowpeardns_32.exe  
Size: 6369086 bytes  
File Version: 1.1.0.0  
Modified: 03-02-2018,17:28:29  
MD5: ECA40487266FFD5A5D6D859C7DDF618C  
SHA1: 48A726DD94F3235B62CEC5C43D2C4B7B5B9E6F65  
CRC32: 2106D9B6  



File: snowpear_64  
Size: 7641943 bytes  
Modified: 03-02-2018,17:47:37  
MD5: 3CD5CFC87D2406943B199C37DD578CA8  
SHA1: 74C986F97935C338CA93764E5F6E010367E2AC98  
CRC32: 25391C93  

File: snowpear_32  
Size: 6638162 bytes  
Modified: 03-02-2018,17:47:03  
MD5: C36C2C2FD4EC7BA226AC5DCE200D6BF1  
SHA1: 3E19AFE19C7B7A5F6744FF9FBBCE43AA371203B1  
CRC32: D258ED16  


## Tips

1. Do you have a dns record cache in this tool?  
  >Yes,since sometimes http request will lost and DNSpod has limit on concurrency requestion made the http dns request very slow,I've add a memcache by using muesli's cache2go,default expiration time of one record haven't been hited is One day(60\*60\*24\*seconds).
2. How can I know DNSpod's server works good for me,what if there is a high latency to DNSpod server from my network?  
  >Well,you can using ping tool to check the latency between your network and DNSpod server by:  
  ```ping 119.29.29.29```  
  remember that the ICMP echo back latency in ms may not the real latency compare to do one http request to dnspod on your network.If it's really too slow,you may got a vps which you can trust and has no dns hijack in the cloud network and running a httpdns server like dnspod,and change the dns server url in SnowPearDNS.go,or you can do some encryption modify to improve the security performence.But The better way here I think you can try some other dns encryption project like
  https://github.com/bitbeans/SimpleDnsCrypt
  This opensource project has been provided a lots of dns encrypted server all over the world.  
3. How's the security of this tool?  
  >As I said at first "currently there has no evidence showed that ISP will hijack the http dns request between user's network and DNSpod server In China",so if you got wrong dns record or you found your ISP hijacked the httpdns request,please let me know and I'll consider to develop a better dns tool with encryption which has it's own server and client. 

## Thanks
- miekg  
  https://github.com/miekg/dns
- fangdingjun  
  https://github.com/fangdingjun/gdns
- muesli  
  https://github.com/muesli/cache2go