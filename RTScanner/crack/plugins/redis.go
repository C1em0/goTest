package plugins

import (
	"fmt"
	"github.com/go-redis/redis"
	"projectv1/RTScanner/config"
	"projectv1/RTScanner/crack/models"
)

func ScanRedis(s models.Service) (result models.ScanResult, err error) {
	result.Service = s
	opt := redis.Options{Addr: fmt.Sprintf("%v:%v", s.IP, s.Port),
		Password: s.Password, DB: 0, DialTimeout: config.CrackTimeOut}
	client := redis.NewClient(&opt)
	_, err = client.Ping().Result()
	if err != nil {
		return result, err
	}

	result.Result = true

	defer func() {
		if client != nil {
			_ = client.Close()
		}
	}()

	return result, err
}
