package plugins

import (
	"fmt"
	"golang.org/x/crypto/ssh"
	"net"
	"projectv1/RTScanner/config"
	"projectv1/RTScanner/crack/models"
)

func ScanSsh(s models.Service) (result models.ScanResult, err error) {
	result.Service = s
	config := &ssh.ClientConfig{
		User: s.Username,
		Auth: []ssh.AuthMethod{
			ssh.Password(s.Password),
		},
		Timeout: config.CrackTimeOut,
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
	}

	client, err := ssh.Dial("tcp", fmt.Sprintf("%v:%v", s.IP, s.Port), config)
	if err != nil {
		return result, err
	}

	session, err := client.NewSession()
	if err != nil {
		return result, err

	}
	err = session.Run("echo SSH")
	if err != nil {
		return result, err
	}

	result.Result = true

	defer func() {
		if client != nil {
			_ = client.Close()
		}
		if session != nil {
			_ = session.Close()
		}
	}()

	return result, err
}
