package main

import (
	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
	"io"
	"log"
	"os"
)

func main() {
	const SSH_ADDRESS = "0.0.0.0:3022"
	const SSH_USERNAME = "garfield"
	const SSH_PASSWORD = "garfield"

	sshConfig := &ssh.ClientConfig{
		User:            SSH_USERNAME,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Auth: []ssh.AuthMethod{
			ssh.Password(SSH_PASSWORD),
		},
	}

	client, err := ssh.Dial("tcp", SSH_ADDRESS, sshConfig)
	if client != nil {
		defer client.Close()
	}

	if err != nil {
		log.Fatal("Failed to dial " + err.Error())
	}

	sftpClient, err := sftp.NewClient(client)
	if err != nil {
		log.Fatal("Failed create client sftp client " + err.Error())
	}

	fDestination, err := sftpClient.Create("~/Documents/test.txt")
	if err != nil {
		log.Fatal("failed to create destination file. " + err.Error())
	}

	fSource, err := os.Open("./test.txt")
	if err != nil {
		log.Fatal("Failed to read source file. " + err.Error())
	}

	_, err = io.Copy(fDestination, fSource)
	if err != nil {
		log.Fatal("Failed copy source file into. " + err.Error())
	}

	log.Println("file copied")
}
