package main

import (
	"fmt"
	"net"
)

func main() {
	//network := &net.IPNet{IP: net.IPv4(10, 38, 0, 0), Mask: net.CIDRMask(16, 32)}
	//ip := net.IPv4(10, 37, 163, 87)
	//fmt.Println(network.Contains(ip))
	i := "10.85.15.151"
	n := "10.85.0.0/16"

	ip := net.ParseIP(i)

	_, network, _ := net.ParseCIDR(n)
	fmt.Println(network.Contains(ip), network)
}
