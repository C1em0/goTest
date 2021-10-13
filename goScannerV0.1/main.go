package main

import (
	"fmt"
	"os"
	"projectv1/goScannerV0.1/scanner"
	"projectv1/goScannerV0.1/util"
)

func main() {
	if len(os.Args) == 3 {
		ipList := os.Args[1]
		portList := os.Args[2]
		ips, _ := util.GetIpList(ipList)
		ports, _ := util.GetPorts(portList)

		for _, ip := range ips {
			for _, port := range ports {
				_, err := scanner.Connect(ip.String(), port)
				if err != nil {
					continue
				}
				fmt.Printf("ip: %v. port: %v is open\n", ip, port)
			}
		}
	} else {
		fmt.Printf("Usage: %v ips ports\n", os.Args[0])
	}
}
