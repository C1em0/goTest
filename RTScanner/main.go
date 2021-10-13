package main

import (
	"fmt"
	"os"
	"projectv1/RTScanner/config"
	"projectv1/RTScanner/scanner"
	"projectv1/RTScanner/util"
	"runtime"
	"time"

	"github.com/urfave/cli"

	"projectv1/RTScanner/cmd"
)

func main() {
	config.Banner()
	app := cli.NewApp()
	app.Name = "RTScan for Internal"
	app.Author = "Dem0"
	app.Email = "hjmail1996@gmail.com"
	app.Version = "2021/9/27"
	app.Usage = "TCP / SYN port scanner"
	app.Commands = []cli.Command{cmd.Scan}
	app.Flags = append(app.Flags, cmd.Scan.Flags...)
	err := app.Run(os.Args)
	//scanner.PrintResult()

	_ = err
	ips, err := util.GetIpList(config.Host)
	ports, err := util.GetPorts(config.Port)
	tasks, n := scanner.GenerateTask(ips, ports)
	_ = n
	start := time.Now()
	scanner.RunTask(tasks)
	long := time.Since(start)

	//defer fmt.Println("=========================")
	//fmt.Printf("[Done] The duration is %v \n", long)
	scanner.PrintResult()
	fmt.Println("===================================")
	fmt.Printf("[Done] The duration is %v \n", long)
}

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}
