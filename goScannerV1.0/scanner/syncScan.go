package scanner

import (
	"fmt"
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
	"log"
	"net"
	_ "time"
)

/*
	TCP 报文 ===>  IP 首部 + TCP首部 + TCP数据部分

	gopacket 子包：
			layers: 用于解码数据包协议
			pcap: 使用libpcap来读取数据包
			pfring: 使用PF_RING来读取数据包
			afpacket: 使用Linux’s AF_PACKET来读取数据包
			tcpassembly: TCP流重组
*/

//使用一次udp连接来获取本地模拟发包的时候的源地址IP和port
func localIPPort(dstip net.IP) (net.IP, int, error) {
	serverAddr, err := net.ResolveUDPAddr("udp", dstip.String()+":54321")
	if err != nil {
		return nil, 0, err
	}
	//并发发送数据需要大量随机端口，并且在多网卡条件下不需要自己确定出口ip
	if con, err := net.DialUDP("udp", nil, serverAddr); err == nil {
		if udpaddr, ok := con.LocalAddr().(*net.UDPAddr); ok {
			//fmt.Printf("[Debug]:srcIP should use : %v, srcPort should use : %v\n", udpaddr.IP, udpaddr.Port)
			return udpaddr.IP, udpaddr.Port, nil
		}
	}
	return nil, -1, err
}


func SynScan(dstIp string, dstPort int) (string, int, error) {

	var handle *pcap.Handle
	//var router routing.Router

	srcIp, srcPort, err := localIPPort(net.ParseIP(dstIp))
	dstAddrs, err := net.LookupIP(dstIp)
	if err != nil {
		return dstIp, 0, err
	}

	dstip := dstAddrs[0].To4()

	dstport := layers.TCPPort(dstPort)
	srcport := layers.TCPPort(srcPort)
	fmt.Printf("[Debug]:dstport is %v,type is %T\n", dstport, dstport)
	fmt.Printf("[Debug]:srcport is %v\n", srcport)

	// 设置IP头
	ip := &layers.IPv4{
		SrcIP:    srcIp,
		DstIP:    dstip,
		Protocol: layers.IPProtocolTCP,
	}
	// 设置TCP头
	tcp := &layers.TCP{
		SrcPort: srcport,
		DstPort: dstport,
		SYN:     true,
	}
	err = tcp.SetNetworkLayerForChecksum(ip)

	buf := gopacket.NewSerializeBuffer()
	opts := gopacket.SerializeOptions{
		ComputeChecksums: true,
		FixLengths:       true,
	}

	iface := "en15"
	//iface, _, _, err := router.Route(dstip)
	//fmt.Println(iface)
	if err != nil {
		return dstIp, 0, err
	}

	handle, err = pcap.OpenLive(iface, 65536, true, pcap.BlockForever)
	if err != nil {
		return dstIp, 0, err
	}

	defer handle.Close()
	//if err := gopacket.SerializeLayers(buf, opts, tcp); err != nil {
	//	fmt.Printf("Serialize failed.\n")
	//	return dstIp, 0, err
	//}

	errHandle := func (handle *pcap.Handle) error {
		if err := gopacket.SerializeLayers(buf, opts, tcp); err != nil {
			fmt.Printf("Serialize failed.\n")
			return err
		}
		return handle.WritePacketData(buf.Bytes())
	}(handle)
	if errHandle != nil {
		return "", 0, errHandle
	} else {
		fmt.Println("[Debug]: Write success.")
	}

	//ln, err := net.Listen("tcp", "0.0.0.0")
	//if err != nil {
	//	panic(err)
	//}
	//conn, err := ln.Accept()

	//conn, err := net.ListenPacket("ip4:tcp", "0.0.0.0")
	//if err != nil {
	//	panic(err)
	//	//fmt.Printf("ListenPacket start failed.\n")
	//	return dstIp, 0, err
	//}
	//defer conn.Close()
	//
	////发送数据包
	//if n, err := conn.WriteTo(buf.Bytes(), &net.IPAddr{IP: dstip}); err != nil {
	//	panic(err)
	//	//fmt.Printf("Send Failed.\n")
	//	return dstIp, 0, err
	//} else {
	//	fmt.Printf("[Debug]:The packet len is %d\n", n)
	//}
	//
	//if err := conn.SetWriteDeadline(time.Now().Add(3 * time.Second)); err != nil {
	//	fmt.Printf("Write OverTime.\n")
	//	return dstIp, 0, err
	//}
	//
	//if err := conn.SetReadDeadline(time.Now().Add(3 * time.Second)); err != nil {
	//	fmt.Printf("Read OverTime.\n")
	//	return dstIp, 0, err
	//}

	// 设置超时
	//if err := conn.SetDeadline(time.Now().Add(3 * time.Second)); err != nil {
	//	fmt.Printf("OverTime.\n")
	//	return dstIp, 0, err
	//}
	//handle.WritePacketData(buf.Bytes())
	//ipFlow := gopacket.NewFlow(layers.EndpointIPv4, dstip, srcIp)

	for {

		//b := make([]byte, 4096)
		// ReadFrom方法从连接读取一个数据包，并将有效信息写入b
		// ReadFrom方法可能会在超过某个固定时间限制后超时返回错误，该错误的Timeout()方法返回真
		// 返回写入的字节数和该数据包的来源地址
		//ReadFrom(b []byte) (n int, addr Addr, err error)
		//n, addr, err := conn.ReadFrom(b)
		//stream := string(b)
		data, _, err := handle.ReadPacketData()
		//fmt.Printf("[Debug]: This is packet stream: %v\n", addr)
		if err == pcap.NextErrorTimeoutExpired {
			panic(err)
		} else if err != nil {
			log.Printf("error reading packet: %v", err)
			panic(err)
		}

		if err != nil {
			panic(err)
			return dstIp, 0, err
		} else {
			// Decode a packet
			packet := gopacket.NewPacket(data, layers.LayerTypeTCP, gopacket.Default)

			// Get the TCP layer from this packet
			if tcpLayer := packet.Layer(layers.LayerTypeTCP); tcpLayer != nil {
				tcp, ok := tcpLayer.(*layers.TCP)
				if !ok {
					panic("tcp layer is not tcp layer :-/")
				}
				//fmt.Printf("[Debug]: tcp layer is follows: %v\n", tcp)
				//返回的数据包中的DstPort是我们TCP发送包的源IP
				if tcp.DstPort == dstport {
					if tcp.SYN && tcp.ACK {
						fmt.Printf("%v:%d is OPEN\n", dstIp, dstport)
						return dstIp, dstPort, err
					} else {
						return dstIp, 0, err
					}
				}
			} else {
				fmt.Printf("tcpLayer is nil.")
				return dstIp, 0, err
			}
		}
	}
}
