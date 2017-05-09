package main

import (
	"fmt"
	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/agent"
	"net"
	"os"
)

func main() {
	var auths []ssh.AuthMethod
	if aconn, err := net.Dial("unix", os.Getenv("SSH_AUTH_SOCK")); err == nil {
		auths = append(auths, ssh.PublicKeysCallback(agent.NewClient(aconn).Signers))
	}
	auths = append(auths, ssh.Password(""))
	addr := fmt.Sprintf("%s:%d", "", 22)
	config := ssh.ClientConfig{
		User:            "root",
		Auth:            auths,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	conn, err := ssh.Dial("tcp", addr, &config)
	if err != nil {
		fmt.Printf("Dial run err! err: %s \n", err.Error())
		return
	}
	defer conn.Close()
	c, err := sftp.NewClient(conn, sftp.MaxPacket(1<<15))
	if err != nil {
		fmt.Printf("NewClient run err! err: %s \n", err.Error())
		return
	}
	files, err := c.ReadDir("/apps/service")
	if err != nil {
		fmt.Printf("ReadDir run err! err: %s \n", err.Error())
		return
	}
	for _, value := range files {
		fmt.Println(value.Name())
	}
	c.Close()
}
