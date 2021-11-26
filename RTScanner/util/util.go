package util

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"

	"github.com/malfunkt/iprange"
	"github.com/urfave/cli"

	"projectv1/RTScanner/config"
	_ "projectv1/RTScanner/scanner"
)

func GetPorts(selection string) ([]int, error) {
	ports := make([]int, 0)
	//检测若是用户没有输入扫描端口那就直接返回
	if selection == "" {
		return ports, nil
	}
	//单个端端口以 "," 进行分割，然后处理以 "-" 分割的范围端口
	ranges := strings.Split(selection, ",")
	for _, r := range ranges {
		r = strings.TrimSpace(r)
		if strings.Contains(r, "-") {
			parts := strings.Split(r, "-")
			if len(parts) != 2 { // "1-65545" ==> ["1" "65545"] 切片长度为2
				return nil, fmt.Errorf("Invalid port selection segment: '%s'\n", r)
			}
			//strconv 提供基本数据类型和其对应字符串表示的相互转换
			part1, err := strconv.Atoi(parts[0])
			if err != nil {
				return nil, fmt.Errorf("Invalid port number: '%s'\n", parts[0])
			}

			part2, err := strconv.Atoi(parts[1])
			if err != nil {
				return nil, fmt.Errorf("Invalid port number: '%s'\n", parts[1])
			}

			if part1 > part2 {
				return nil, fmt.Errorf("Invalid port range: '%d-%d'\n", part1, part2)
			}

			for i := part1; i <= part2; i++ {
				ports = append(ports, i)
			}
		} else {
			port, err := strconv.Atoi(r)
			if err != nil {
				return nil, fmt.Errorf("Invalid port number: '%s'\n", r)
			} else {
				ports = append(ports, port)
			}
		}
	}
	return ports, nil
}

func GetIpList(ips string) ([]net.IP, error) {

	//ParseList() 方法接收 10.0.0.1, 10.0.0.5-10, 192.168.1.*, 192.168.10.0/24 格式ip地址
	addressList, err := iprange.ParseList(ips)
	if err != nil {
		return nil, err
	}

	//Expand() 方法返回一个 []net.IP的切片
	list := addressList.Expand()
	return list, err
}

func IsRoot() bool {
	return os.Geteuid() == 0
}

func CheckRoot() {
	if !IsRoot() {
		fmt.Println("must run with root")
		os.Exit(0)
	}
}

//获取用户输入初始化配置文件
func Scan(ctx *cli.Context) {
	if ctx.IsSet("iplist") {
		config.Host = ctx.String("iplist")
	}

	if ctx.IsSet("port") {
		config.Port = ctx.String("port")
	}

	if ctx.IsSet("mode") {
		config.Mode = ctx.String("mode")
	}

	if ctx.IsSet("timeout") {
		config.Timeout = ctx.Int("timeout")
	}

	if ctx.IsSet("thread") {
		config.ThreadNum = ctx.Int("thread")
	}

	if strings.ToLower(config.Mode) == "syn" {
		CheckRoot()
	}
}

func Crack(ctx *cli.Context) {
	if ctx.IsSet("target") {
		config.CrackIpFile = ctx.String("target")
	}
}
