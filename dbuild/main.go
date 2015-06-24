package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {

	// the way to use it
	// ./executable -t myself/wordpress-mysql wordpress-mysql

	argsWithoutProg := os.Args[1:]

	commandString := "docker-machine ssh " + argsWithoutProg[3] + " -- sudo mkdir -p /var/lib/docker/stack/" + argsWithoutProg[1]

	fmt.Println(commandString)
	words := strings.Split(commandString, " ")
	cmd := exec.Command(words[0], words[1:]...)

	// Stdout buffer
	cmdOutput := &bytes.Buffer{}
	// Attach buffer to command
	cmd.Stdout = cmdOutput

	// Execute command
	// printCommand(cmd)
	err := cmd.Run() // will wait for command to return
	if err != nil {
		fmt.Println("Error:", err.Error())
	}
	// printError(err)
	// Only output the commands stdout
	// printOutput(cmdOutput.Bytes())
	fmt.Println(string(cmdOutput.Bytes()))

	//
	//
	//
	//
	//
	//

	commandString = "docker-machine ssh " + argsWithoutProg[3] + " -- sudo cp " + argsWithoutProg[2] + " /var/lib/docker/stack/" + argsWithoutProg[1]
	// cmd docker-machine ssh dev1 -- sudo cp wordpress-mysql /var/lib/docker/stack/myself/wordpress-mysql
	fmt.Println("cmd", commandString)
	words = strings.Split(commandString, " ")
	cmd = exec.Command(words[0], words[1:]...)

	// Stdout buffer
	cmdOutput = &bytes.Buffer{}
	// Attach buffer to command
	cmd.Stdout = cmdOutput

	// Execute command
	// printCommand(cmd)
	err = cmd.Run() // will wait for command to return
	if err != nil {
		fmt.Println("Error:", err.Error())
	}
	// printError(err)
	// Only output the commands stdout
	// printOutput(cmdOutput.Bytes())
	fmt.Println(string(cmdOutput.Bytes()))
}
