package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"regexp"
)

var (
	Reset  = "\033[0m"
	Black  = "\033[30m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
	Purple = "\033[35m"
	Cyan   = "\033[36m"
	Gray   = "\033[37m"
	White  = "\033[97m"
)

const (
	IPRegex    = `\b(?:\d{1,3}\.){3}\d{1,3}\b$`
	RangeRegex = `\-`
)

func main() {
	flag.Usage = usage
	flag.Parse()
	_, err := os.Stdin.Stat()

	if err != nil {
		log.Fatal(err)
	}
	args := os.Args[1:]

	if len(args) > 0 {
		var subnets []string
		subnets = args
		for _, subnet := range subnets {
			displayIPs(subnet)
		}
	} else {
		flag.Usage()
	}
}

func isIPRange(subnet string) bool {
	match, _ := regexp.MatchString(RangeRegex, subnet)
	return match
}

func isIPAddr(subnet string) bool {
	match, _ := regexp.MatchString(IPRegex, subnet)
	return match
}

func displayIPs(subnet string) {
	var addresses []string

	if isIPRange(subnet) && isIPAddr(subnet) {
		fmt.Println(subnet)
		return
	} else if isIPAddr(subnet) {
		fmt.Println(subnet)
		return
	}

	ipAddr, ipNet, err := net.ParseCIDR(subnet)
	if err != nil {
		fmt.Println(subnet)
		log.Print(err)
		return
	}

	for ip := ipAddr.Mask(ipNet.Mask); ipNet.Contains(ip); increment(ip) {
		addresses = append(addresses, ip.String())
	}

	for _, ip := range addresses[0:len(addresses)] {
		fmt.Println(ip)
	}
}

func increment(ip net.IP) {
	for i := len(ip) - 1; i >= 0; i-- {
		ip[i]++
		if ip[i] != 0 {
			break
		}
	}
}

func usage() {
	fmt.Fprintf(os.Stderr, "._______________________________________________________________.\n")
	fmt.Fprintf(os.Stderr, "|"+Green+"Instructions: "+Yellow+"$"+White+"./Subnet2IPList <subnet>"+Reset+"                        |\n")
	fmt.Fprintf(os.Stderr, "|         "+Gray+"or"+Reset+"                                                    |\n")
	fmt.Fprintf(os.Stderr, "|         "+Yellow+"$"+White+"cat subnet-list.txt | ./Subnet2IPList"+Reset+"                |\n")
	fmt.Fprintf(os.Stderr, "|         "+Gray+"else"+Reset+"                                                  |\n")
	fmt.Fprintf(os.Stderr, "|         "+Yellow+"$"+White+"cat subnet-list.txt | ./Subnet2IPList > out.txt"+Reset+"      |\n")
	fmt.Fprintf(os.Stderr, "`---------------------------------------------------------------`\n")
	fmt.Fprintf(os.Stderr, Cyan+"nota bene:"+Reset+"No extra whitespaces, IPv6 or DNS allowed >:(\n")
}
