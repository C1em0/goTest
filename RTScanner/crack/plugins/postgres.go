package plugins

import (
	"database/sql"
	"fmt"
	"projectv1/RTScanner/crack/models"
)

func ScanPostgres(service models.Service) (result models.ScanResult, err error) {
	result.Service = service

	dataSourceName := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=%v", service.Username,
		service.Password, service.IP, service.Port, "postgres", "disable")
	db, err := sql.Open("postgres", dataSourceName)

	if err == nil {
		defer func() {
			err = db.Close()
		}()
		err = db.Ping()
		if err == nil {
			result.Result = true
		}
	}

	return result, err
}
