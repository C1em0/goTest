package plugins

import (
	"database/sql"
	"fmt"
	"projectv1/RTScanner/crack/models"
)

func ScanOracle(service models.Service) (result models.ScanResult, err error) {
	result.Service = service
	dataSourceName := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v", service.Username,
		service.Password, service.IP, service.Port, "orcl")

	db, err := sql.Open("ora", dataSourceName)
	defer db.Close()
	if err == nil {
		err = db.Ping()
		if err == nil {
			result.Result = true
		}
	}

	return result, err
}
