package main

import (
	"golang.org/x/crypto/ssh"
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

	session, err := client.NewSession()
	if session != nil {
		defer session.Close()
	}

	if err != nil {
		log.Fatal("Failed to create session. " + err.Error())
	}

	session.Stdin = os.Stdin
	session.Stdout = os.Stdout
	session.Stderr = os.Stdout

	err = session.Run("ls -l ~/")
	if err != nil {
		log.Fatal("Command execution error. " + err.Error())
	}
}
