package main

import "fmt"

type IPAddr [4]byte

// TODO: Add a "String() string" method to IPAddr.
func (i IPAddr) String() string {
	ip := ""
	for _, b := range i {
		ip += fmt.Sprint(b) + "."
	}
	return ip
}

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
		fmt.Printf("%v\n", fmt.Sprint(ip))
	}
}
