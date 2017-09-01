package main

import "fmt"
import "net"
import "os"

func main() {
	host, err := os.Hostname()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(host)
	ifaces, err := net.Interfaces()
	for _, iface := range ifaces {
		fmt.Printf("Interface: %s\n", iface.Name)
		addrs, err := iface.Addrs()
		if err != nil {
			fmt.Println(err)
			return
		}
		for _, addr := range addrs {
			switch v := addr.(type) {
			case *net.IPNet:
				fmt.Printf("\tAddress: %s\n", v.IP.String())
			case *net.IPAddr:
				fmt.Printf("\tAddress: %s\n", v.IP.String())
			}
		}
	}
}
