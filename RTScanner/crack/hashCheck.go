package crack

import (
	"crypto/md5"
	"fmt"
	"io"
	"projectv1/RTScanner/config"
)

func MakeTaskHash(k string) string {
	hash := MD5(k)
	return hash
}

func CheckTaskHash(hash string) bool {
	_, ok := config.SuccessHash.Load(hash)
	return ok
}

func MD5(s string) (m string) {
	h := md5.New()
	_, _ = io.WriteString(h, s)
	return fmt.Sprintf("%x", h.Sum(nil))
}
