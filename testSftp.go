package main

import (
	"fmt"
	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)

func main() {
	conn, err := sshConntion()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	//var conn *ssh.Client
	// open an SFTP session over an existing ssh connection.
	sftp, err := sftp.NewClient(conn)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer sftp.Close()

	// walk a directory
	w := sftp.Walk("/apps")
	for w.Step() {
		if w.Err() != nil {
			continue
		}
		fmt.Println(w.Path())
	}

	// leave your mark
	f, err := sftp.Create("hello")
	if err != nil {
		fmt.Println(err)
	}
	if _, err := f.Write([]byte("Hello world!")); err != nil {
		fmt.Println(err)
	}

	// check it's there
	fi, err := sftp.Lstat("hello")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(fi)
}

func sshConntion() (*ssh.Client, error) {
	var hostKey ssh.PublicKey
	config := &ssh.ClientConfig{
		User: "root",
		Auth: []ssh.AuthMethod{
			ssh.Password("Axon@2016"),
		},
		HostKeyCallback: ssh.FixedHostKey(hostKey),
	}
	// Dial your ssh server.
	conn, err := ssh.Dial("tcp", "10.10.141.71:22", config)
	if err != nil {
		fmt.Println("unable to connect: ", err)
		return nil, err
	}
	return conn, err
}
