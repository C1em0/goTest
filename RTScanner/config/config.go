package config

import (
	"fmt"
	"sync"
	"time"
)

var (
	ThreadNum = 10000 //考虑到快速扫描B段
	Result    *sync.Map

	Host         string
	Port         = "22,3389,3306"
	Mode         = "tcp"
	Timeout      = 2
	CrackIpFile  = "../target/target"
	CrackTimeOut = 2 * time.Second
	SuccessHash  sync.Map
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

                                              RTScan version: 2.0.1



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

var PortsFlag = map[string]int{
	"ftp":     21,
	"ssh":     22,
	"findnet": 135,
	"netbios": 139,
	"smb":     445,
	"mssql":   1433,
	"mysql":   3306,
	"psql":    5432,
	"redis":   6379,
	"fcgi":    9000,
	"mem":     11211,
	"mgo":     27017,
	"oracle":  1521,
}

var Users = map[string][]string{
	"ftp":        {"ftp", "admin", "www", "web", "root", "db", "wwwroot", "data"},
	"mysql":      {"root", "mysql"},
	"mssql":      {"sa", "sql"},
	"smb":        {"administrator", "admin", "guest"},
	"postgresql": {"postgres", "admin"},
	"ssh":        {"dem0", "root", "admin", "Oracle"},
	"mongodb":    {"root", "admin"},
	"redis":      {""},
	"oracle":     {"oracle", "sys", "system", "Oracle"},
}

var Passwords = []string{
	"123456@iie",
	"123456",
	"admin",
	"admin123",
	"root",
	"oracle",
	"Oracle",
	"",
	"pass123",
	"pass@123",
	"password",
	"123123",
	"654321",
	"111111",
	"123",
	"1",
	"admin@123",
	"Admin@123",
	"admin123!@#",
	"{user}",
	"{user}1",
	"{user}111",
	"{user}123",
	"{user}@123",
	"{user}_123",
	"{user}#123",
	"{user}@111",
	"{user}@2019",
	"{user}@123#4",
	"P@ssw0rd!",
	"P@ssw0rd",
	"Passw0rd",
	"qwe123",
	"12345678",
	"test",
	"test123",
	"123qwe!@#",
	"123456789",
	"123321",
	"666666",
	"a123456.",
	"123456~a",
	"123456!a",
	"000000",
	"1234567890",
	"8888888",
	"!QAZ2wsx",
	"1qaz2wsx",
	"abc123",
	"abc123456",
	"1qaz@WSX",
	"a11111",
	"a12345",
	"Aa1234",
	"Aa1234.",
	"Aa12345",
	"a123456",
	"a123123",
	"Aa123123",
	"Aa123456",
	"Aa12345.",
	"sysadmin",
	"system",
	"1qaz!QAZ",
	"2wsx@WSX",
	"qwe123!@#",
	"Aa123456!",
	"A123456s!",
	"sa123456",
	"1q2w3e",
}
