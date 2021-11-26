package plugins

import "projectv1/RTScanner/crack/models"

//使用函数签名设计成插件化模块
type ScanFunc func(service models.Service) (result models.ScanResult, err error)

var ScanFuncMap map[string]ScanFunc

func init() {
	ScanFuncMap = make(map[string]ScanFunc)
	ScanFuncMap["FTP"] = ScanFtp
	ScanFuncMap["SSH"] = ScanSsh
	ScanFuncMap["MYSQL"] = ScanMysql
	ScanFuncMap["MSSQL"] = ScanMssql
	ScanFuncMap["REDIS"] = ScanRedis
	ScanFuncMap["POSTGRESQL"] = ScanPostgres
	ScanFuncMap["MONGODB"] = ScanMongodb
	ScanFuncMap["ORACLE"] = ScanOracle
}
