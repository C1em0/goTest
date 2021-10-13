package config

import (
	"fmt"
	"sync"
	"time"
)

var (
	ThreadNum = 64516 //考虑到快速扫描B段
	Result    *sync.Map

	Host    string
	Port    = "22,3389,3306"
	Mode    = "tcp"
	Timeout = 2
)

func Banner() {
	banner := `
				
	         _         _          _             _             _                   _          
        /\ \      /\ \       / /\         /\ \           / /\                /\ \     _  
       /  \ \     \_\ \     / /  \       /  \ \         / /  \              /  \ \   /\_\
      / /\ \ \    /\__ \   / / /\ \__   / /\ \ \       / / /\ \            / /\ \ \_/ / /
     / / /\ \_\  / /_ \ \ / / /\ \___\ / / /\ \ \     / / /\ \ \          / / /\ \___/ / 
    / / /_/ / / / / /\ \ \\ \ \ \/___// / /  \ \_\   / / /  \ \ \        / / /  \/____/  
   / / /__\/ / / / /  \/_/ \ \ \     / / /    \/_/  / / /___/ /\ \      / / /    / / /   
  / / /_____/ / / /    _    \ \ \   / / /          / / /_____/ /\ \    / / /    / / /    
 / / /\ \ \  / / /    /_/\__/ / /  / / /________  / /_________/\ \ \  / / /    / / /     
/ / /  \ \ \/_/ /     \ \/___/ /  / / /_________\/ / /_       __\ \_\/ / /    / / /      
\/_/    \_\/\_\/       \_____\/   \/____________/\_\___\     /____/_/\/_/     \/_/       

                                              RTScan version: 1.0.1



`
	fmt.Print(banner)
}

func ScanTime() time.Time {
	timeLine := time.Date(
		time.Now().Year(),
		time.Now().Month(),
		time.Now().Day(),
		time.Now().Hour(),
		time.Now().Minute(),
		time.Now().Second(), 0, time.Local)
	//fmt.Printf("[*] %v\n\n", timeLine)
	return timeLine

}

func init() {
	Result = &sync.Map{}
}
