package main

import (
	"fmt"
	"os"
	"projectv1/goScannerV1.0/scanner"
	"projectv1/goScannerV1.0/util"
	"runtime"
)

func main() {
	if len(os.Args) == 3 {
		//网络数据包的部分需要root权限
		util.CheckRoot()

		ipList := os.Args[1]
		portList := os.Args[2]
		ips, err := util.GetIpList(ipList)
		ports, err := util.GetPorts(portList)
		_ = err
		tasks, _ := scanner.GenerateTask(ips, ports)
		scanner.RunTask(tasks)
		scanner.PrintResult()
	} else {
		fmt.Printf("%v iplist port\n", os.Args[0])
	}
}

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}
