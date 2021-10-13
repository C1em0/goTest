package config

import (
	"sync"
)

var (
	TreadNumber = 10000
	Result sync.Map
)

func init()  {
	Result = sync.Map{}
}
