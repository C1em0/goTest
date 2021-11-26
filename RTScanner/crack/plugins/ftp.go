package plugins

import (
	"fmt"
	"github.com/jlaffaye/ftp"
	"projectv1/RTScanner/config"
	"projectv1/RTScanner/crack/models"
)

func ScanFtp(s models.Service) (result models.ScanResult, err error) {
	result.Service = s
	conn, err := ftp.DialTimeout(fmt.Sprintf("%v:%v", s.IP, s.Port), config.CrackTimeOut)
	if err == nil {
		err = conn.Login(s.Username, s.Password)
		if err == nil {
			defer func() {
				err = conn.Logout()
			}()
			result.Result = true
		}
	}
	return result, err
}
