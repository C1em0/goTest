package scanner

import (
	"fmt"
	"net"
	"time"
)

/*
	利用 Dial 提供的 socks 编程方式来连接目标，如果超时就证明未开放
*/

func Connect(ip string, port int) (string, int, error) {

	//var msg [128]byte
	//func DialTimeout(network, address string, timeout time.Duration) (Conn, error)
	conn, err := net.DialTimeout("tcp", fmt.Sprintf("%v:%v", ip, port), 1*time.Second)

	//defer conn.Close()
	defer func() {
		if conn != nil {
			_ = conn.Close()
		}
	}()

	//portBanner, err := conn.Read(msg[0:])
	//fmt.Println(portBanner)
	return ip, port, err
}

//待优化：
//	1. 建立一个tcp连接的时长理论上是 ping * 3，后期优化下根据ping值来计算超时时间
