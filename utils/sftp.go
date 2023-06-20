package utils

import (
	"fmt"
	log "github.com/go-admin-team/go-admin-core/logger"
	"github.com/go-admin-team/go-admin-core/sdk/pkg"
	"github.com/pkg/sftp"
	extConfig "go-admin/config"
	"golang.org/x/crypto/ssh"
)

var SshConn *ssh.Client = nil

func connectSSH() {

	sshConfig := &ssh.ClientConfig{
		User:            extConfig.ExtConfig.Vzoom.Sftp.Username,
		Auth:            []ssh.AuthMethod{ssh.Password(extConfig.ExtConfig.Vzoom.Sftp.Password)},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Timeout:         10000000000,
	}
	uri := extConfig.ExtConfig.Vzoom.Sftp.Host + ":" + extConfig.ExtConfig.Vzoom.Sftp.Port

	client, err := ssh.Dial("tcp", uri, sshConfig)
	if err != nil {
		log.Errorf(pkg.Red("SSH Server UNABLE to connect. " + err.Error()))
	} else {
		SshConn = client
		log.Info(pkg.Green("SSH Server connected."))
	}
}

func GetSshConn() *ssh.Client {
	if SshConn == nil {
		connectSSH()
	}
	return SshConn
}

func CloseShhConn() {
	if SshConn != nil {
		err := GetSshConn().Close()
		if err != nil {
			log.Errorf(pkg.Red("SSH connection fail to close. " + err.Error()))
			panic(err)
		} else {
			fmt.Println("SSH drive disconnected.")
			SshConn = nil
		}
	}

}

func GetSftpClient() (*sftp.Client, error) {
	sshConn := GetSshConn()
	clientP, err := sftp.NewClient(sshConn)
	if err != nil {
		log.Errorf(pkg.Red("SFTP Server UNABLE to connect. " + err.Error()))
		panic(err)
		return nil, err
	}
	return clientP, nil
}
