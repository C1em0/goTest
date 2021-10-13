package main

import (
	"fmt"
	"os"
	"projectv1/goScannerV0.2/scanner"
	"projectv1/goScannerV0.2/util"
	"runtime"
	"time"
)

func main() {
	if len(os.Args) == 3 {
		ipList := os.Args[1]
		portList := os.Args[2]

		ips, _ := util.GetIpList(ipList)
		ports, _ := util.GetPorts(portList)

		tasks, _ := scanner.GenerateTask(ips, ports)

		start := time.Now()
		scanner.TaskChunk(tasks)
		long := time.Since(start)

		defer fmt.Printf("The duration is %v \n", long)
		scanner.StdoutResult()
	} else {
		fmt.Printf("%v iplist port\n", os.Args[0])
	}
}

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}