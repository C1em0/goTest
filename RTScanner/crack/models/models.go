package models

type Service struct {
	IP       string
	Port     int
	Protocol string
	Username string
	Password string
}
type ScanResult struct {
	Service Service
	Result  bool
}

type IpAddr struct {
	IP       string
	Port     int
	Protocol string
}
