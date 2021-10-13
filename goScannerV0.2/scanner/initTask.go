package scanner

import (
	"fmt"
	"net"
	"strings"
	"sync"
)

var ThreadNumber = 100000
var Result sync.Map

func GenerateTask(ipList []net.IP, ports []int) ([]map[string]int, int) {
	//map类型的切片
	tasks := make([]map[string]int, 0)
	for _, ip := range ipList {
		for _, port := range ports {
			ipPort := map[string]int{ip.String():port}
			tasks = append(tasks, ipPort)
		}
	}

	return tasks, len(tasks)
}

func TaskChunk(tasks []map[string]int)  {
	var wg sync.WaitGroup

	//这里模拟生产者消费者模型将任务切片取到一个channel中
	taskCh := make(chan map[string]int, ThreadNumber*2)
	for _, task := range tasks {
		wg.Add(1)
		taskCh <- task
	}
	close(taskCh)

	for i := 0; i <= ThreadNumber; i++ {
		go InitScan(taskCh, &wg)
	}

	wg.Wait()
}

func InitScan(taskCh chan map[string]int, wg *sync.WaitGroup)  {
	for task := range taskCh {
		for ip, port := range task {
			_ = SaveResult(Connect(ip,port))
			//_ = err
			wg.Done()
		}
	}
}

func SaveResult(ip string, port int, err error) error {
	//fmt.Printf("ip:%v, port: %v, goruntineNum: %v\n", ip, port, runtime.NumGoroutine())
	if err != nil {
		return err
	}

	//Map.Load() 方法根据 key 值来取 value，这里key是ip，value是port
	v, ok := Result.Load(ip)
	if ok {
		ports, ok1 := v.([]int)
		//fmt.Printf("v.([]int) is %v\n", ports)
		if ok1 {
			ports = append(ports, port)
			Result.Store(ip, ports)
		}
	} else {
		ports := make([]int, 0)
		ports = append(ports, port)
		Result.Store(ip, ports)
	}
	return err
}

func StdoutResult()  {
	//遍历 sync.Map
	Result.Range(func(key, value interface{}) bool {
		fmt.Printf("ip:%v\t", key)
		fmt.Printf("ports:%v\n", value)
		fmt.Println(strings.Repeat("=", 50))
		return true
	})
}
