package main

import (
	"fmt"
	"github.com/urfave/cli"
	"os"
	"projectv1/RTScanner/cmd"
	"projectv1/RTScanner/config"
	"projectv1/RTScanner/scanner"
	"projectv1/RTScanner/util"
	"runtime"
	"time"
)

func main() {
	config.Banner()
	model := os.Args[1]
	app := cli.NewApp()
	app.Name = "RTScan for Internal"
	app.Author = "Dem0"
	app.Email = "hjmail1996@gmail.com"
	app.Version = "2021/9/27"
	app.Usage = "Multifunction Internal scanner"
	if model == "scan" {
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

		timeFlag := config.ScanTime()
		fmt.Printf("[*] %v\n\n", timeFlag)
		//defer fmt.Println("=========================")
		//fmt.Printf("[Done] The duration is %v \n", long)
		scanner.PrintResult()
		scanner.SaveToText()
		fmt.Println("===============================================")
		fmt.Printf("[Done] The duration is %v \n", long)
	} else if model == "crack" {
		//爆破模块
	}
}

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}
