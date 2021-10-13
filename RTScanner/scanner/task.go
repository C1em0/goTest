package scanner

import (
	"fmt"
	"net"
	"os"
	"reflect"
	"strings"
	"sync"

	"projectv1/RTScanner/config"
)

func GenerateTask(ipList []net.IP, ports []int) ([]map[string]int, int) {
	tasks := make([]map[string]int, 0)

	for _, ip := range ipList {
		for _, port := range ports {
			ipPort := map[string]int{ip.String(): port}
			tasks = append(tasks, ipPort)
		}
	}

	return tasks, len(tasks)
}

func RunTask(tasks []map[string]int) {
	wg := &sync.WaitGroup{}

	// 创建一个buffer为config.threadNum 的channel
	taskChan := make(chan map[string]int, 2*config.ThreadNum)

	// 创建config.ThreadNum个协程
	for i := 0; i <= config.ThreadNum; i++ {
		go Scan(taskChan, wg)
	}

	// 生产者，不断地往taskChan channel发送数据，直接channel阻塞
	for _, task := range tasks {
		wg.Add(1)
		taskChan <- task
	}

	close(taskChan)
	wg.Wait()
}

func Scan(taskChan chan map[string]int, wg *sync.WaitGroup) {
	// 每个协程都从channel中读取数据后开始扫描并入库
	for task := range taskChan {
		for ip, port := range task {
			flag := strings.ToLower(config.Mode)
			switch {
			case flag == "tcp":
				err := SaveResult(Connect(ip, port))
				_ = err
			case flag == "syn":
				s, err := newSynScanner(ip, port)
				err = SaveResult(s.scan())
				s.close()
				_ = err
			}
			wg.Done()
		}
	}
}

func SaveResult(ip string, port int, err error) error {
	// fmt.Printf("ip:%v, port: %v, goruntineNum: %v\n", ip, port, runtime.NumGoroutine())
	if err != nil {
		return err
	}

	v, ok := config.Result.Load(ip)
	if ok {
		ports, ok1 := v.([]int)
		//fmt.Printf("v.([]int) is %v\n", ports)
		if ok1 {
			ports = append(ports, port)
			config.Result.Store(ip, ports)
		}
	} else {
		ports := make([]int, 0)
		ports = append(ports, port)
		config.Result.Store(ip, ports)
	}
	return err
}

func PrintResult() {
	config.Result.Range(func(key, value interface{}) bool {
		fmt.Printf(" --> "+"IP: %v\t", key)
		fmt.Printf("PORTS: %v\n", value)
		//fmt.Println(strings.Repeat("", 40))
		return true
	})
}

func SaveToText() {
	filePath := "./results.txt"
	targetFilePath := "./target/target"
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	targetFile, err := os.OpenFile(targetFilePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("打开文件错误。")
		return
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("文件句柄关闭失败。")
			return
		}
	}(file)

	//开始写入文件
	timeFlag := config.ScanTime()
	fmt.Fprintf(file, "[*] %v\n\n", timeFlag)
	config.Result.Range(func(key, value interface{}) bool {
		fmt.Fprintf(file, " --> "+"IP: %v\t", key)
		fmt.Fprintf(file, "PORTS: %v\n", value)
		return true
	})

	config.Result.Range(func(key, value interface{}) bool {
		ports := reflect.ValueOf(value)
		if ports.Len() > 1 {
			for i := 0; i < ports.Len(); i++ {
				fmt.Fprintf(targetFile, "%v:%v\n", key, ports.Index(i))
			}
		} else {
			fmt.Fprintf(targetFile, "%v:%v\n", key, ports.Index(0))
		}
		return true
	})
}
