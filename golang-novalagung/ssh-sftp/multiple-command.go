package main

import (
	"bytes"
	"fmt"
	"golang.org/x/crypto/ssh"
	"log"
	"strings"
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

	var stdout, stderr bytes.Buffer
	session.Stdout = &stdout
	session.Stderr = &stderr

	stdin, err := session.StdinPipe()
	if err != nil {
		log.Fatal("Error getting stdin pipe. " + err.Error())
	}

	err = session.Start("/bin/bash")
	if err != nil {
		log.Fatal("Error starting bash. " + err.Error())
	}

	commands := []string{
		"cd ~/Documents",
		"ls",
		"exit",
	}

	for _, command := range commands {
		if _, err = fmt.Fprintln(stdin, command); err != nil {
			log.Fatal(err)
		}
	}

	err = session.Wait()
	if err != nil {
		log.Fatal(err)
	}

	outputErr := stderr.String()
	fmt.Println("========= ERROR")
	fmt.Println(strings.TrimSpace(outputErr))

	outputString := stdout.String()
	fmt.Println("========= OUTPUT")
	fmt.Println(strings.TrimSpace(outputString))

}
