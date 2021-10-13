package util

import (
	"fmt"
	"github.com/malfunkt/iprange"
	"net"
	"os"
	"strconv"
	"strings"
)

/*
	格式化用户传入的ip和端口信息，例如：

	172.16.234.0/24,192.168.12.0/24, 80,445,1-65545

*/

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
	return os.Getegid() == 0
}

func CheckRoot()  {
	if !IsRoot() {
		fmt.Println("Must run with 'ROOT'")
		os.Exit(0)
	}
}
