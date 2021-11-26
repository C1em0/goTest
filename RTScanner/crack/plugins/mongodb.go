package plugins

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"projectv1/RTScanner/config"
	"projectv1/RTScanner/crack/models"
)

func ScanMongodb(s models.Service) (result models.ScanResult, err error) {
	result.Service = s
	url := fmt.Sprintf("mongodb://%v:%v@%v:%v/%v", s.Username, s.Password, s.IP, s.Port, "test")
	session, err := mgo.DialWithTimeout(url, config.CrackTimeOut)

	if err == nil {
		defer session.Close()
		err = session.Ping()
		if err == nil {
			result.Result = true
		}
	}

	return result, err
}
