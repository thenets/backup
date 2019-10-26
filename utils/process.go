package utils

import (
	"log"
	"os"
	"os/exec"
	"strings"
)

// RunCmd starts a new command and wait until finish
func RunCmd(commandName string, args []string, id string) error {
	var debug = false
	var runOverBash = true

	var err error

	if debug {
		log.Println("CMD", commandName)
		log.Println("ARG", args)
	}

	// Run over bash
	if runOverBash {
		var newCommandName = "/bin/sh"
		var newArgs = []string{"-c"}
		newArgs = append(newArgs, strings.Join(append([]string{commandName}, args...), " "))

		commandName = newCommandName
		args = newArgs
	}

	// Set log files
	stdoutFile, err := os.Create(GetLogsPath() + id)
	if err != nil {
		panic(err)
	}
	defer stdoutFile.Close()
	stderrFile, err := os.Create(GetLogsPath() + id + ".err")
	if err != nil {
		panic(err)
	}
	defer stderrFile.Close()

	// Run command
	cmd := exec.Command(commandName, args...)
	if debug {
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stdout
	} else {
		cmd.Stdout = stdoutFile
		cmd.Stderr = stderrFile
	}
	cmd.Env = os.Environ()
	err = cmd.Start()
	if err != nil {
		log.Fatal(err)
	}

	// Wait syncing finish
	cmd.Wait()

	// Show output

	return err
}
