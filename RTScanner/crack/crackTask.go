package crack

import (
	"fmt"
	"os"
	"projectv1/RTScanner/config"
	"projectv1/RTScanner/crack/models"
	"projectv1/RTScanner/crack/plugins"
	"strings"
	"sync"
	"time"
)

func InitTask(ipInfo models.IpAddr) (tasks []models.Service, taskNum int) {

	username := config.Users[ipInfo.Protocol]
	password := config.Passwords

	for _, user := range username {
		for _, pass := range password {
			service := models.Service{
				IP:       ipInfo.IP,
				Port:     ipInfo.Port,
				Protocol: ipInfo.Protocol,
				Username: user,
				Password: pass,
			}
			tasks = append(tasks, service)
		}
	}

	return tasks, len(tasks)
}

//tasks []models.Service, taskNum int
func RunTasks(tasks []models.Service, taskNum int, wgCheck *sync.WaitGroup) {
	wg := &sync.WaitGroup{}
	taskChan := make(chan models.Service, taskNum)

	for _, task := range tasks {
		wg.Add(1)
		taskChan <- task
	}

	for i := 0; i < taskNum; i++ {
		go crackPassword(taskChan, wg)
	}

	close(taskChan)
	waitTimeOut(wg, config.CrackTimeOut)

	wgCheck.Done()
}

func crackPassword(taskChan chan models.Service, wg *sync.WaitGroup) {
	for task := range taskChan {
		var k string
		protocol := strings.ToUpper(task.Protocol)

		if protocol == "REDIS" {
			k = fmt.Sprintf("%v-%v-%v", task.IP, task.Port, task.Protocol)
		} else {
			k = fmt.Sprintf("%v-%v-%v", task.IP, task.Port, task.Username)
		}

		h := MakeTaskHash(k)
		if CheckTaskHash(h) {
			wg.Done()
			continue
		}

		fn := plugins.ScanFuncMap[protocol]
		result, _ := fn(task)

		if result.Result {
			//lime := chalk.Green.NewStyle().WithTextStyle(chalk.Bold).Style
			//fmt.Printf(lime("[+] IP:%v:%v  Protocol:%v Cracked!, USERNAME:%v, PASSWORD:%v"), task.IP, task.Port, task.Protocol, task.Username, task.Password)
			//fmt.Printf("\n")
			fmt.Println("[+]Weak pass cracked!")
			savaResult(task.IP, task.Port, task.Protocol, task.Username, task.Password)
		} else {
			continue
		}
		//models.SaveResult(fn(task))
		wg.Done()
	}
}

func waitTimeOut(wg *sync.WaitGroup, timeout time.Duration) bool {
	c := make(chan struct{})
	go func() {
		defer close(c)
		wg.Wait()
	}()
	select {
	case <-c:
		return false // completed normally
	case <-time.After(timeout):
		return true // timed out
	}
}

func savaResult(ip string, port int, protocol, username, pass string) {
	crackFile := "./crackedPass.txt"
	file, err := os.OpenFile(crackFile, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0777)
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

	fmt.Fprintf(file, "[%v] IP:%v:%v  Protocol:%v Cracked!, USERNAME:%v, PASSWORD:%v\n", protocol, ip, port, protocol, username, pass)

}
