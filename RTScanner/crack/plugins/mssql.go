package plugins

import (
	"database/sql"
	"fmt"
	"projectv1/RTScanner/crack/models"
)

func ScanMssql(service models.Service) (result models.ScanResult, err error) {
	result.Service = service

	dataSourceName := fmt.Sprintf("server=%v;port=%v;user id=%v;password=%v;database=%v", service.IP,
		service.Port, service.Username, service.Password, "master")

	db, err := sql.Open("mssql", dataSourceName)
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
