package utils

import (
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/thenets/backup/utils"
)

// RunCmd starts a new command and wait until finish
func RunCmd(cmdLine string) error {
	var err error

	var debug = true
	var runOverBash = true

	var nameRef = "cafe"

	// Create cache dir
	var backupCacheDir = utils.GetCacheDir() + nameRef + "/"
	if !utils.IsDirectory(backupCacheDir) {
		os.MkdirAll(backupCacheDir, 0755)
	}
	if !utils.IsDirectory(backupCacheDir) {
		panic("backup cache dir can't be created: " + backupCacheDir)
	}

	// Create command line
	var commandName = "/usr/bin/rsync"
	var args []string
	args = []string{
		"-av",
		"--timeout=300",
		"--recursive",
		"--delete",
		"-e",
		// fmt.Sprintf("'ssh -p %d %s'", ssh.Spec.Server.Port, ssh.Spec.CustomSSHArgs),
		// fmt.Sprintf("%s@%s:%s", ssh.Spec.Server.User, ssh.Spec.Server.Host, ssh.Spec.RemoteDir),
		backupCacheDir,
	}
	// args = append(args, strings.Split(ssh.Spec.CustomRsyncArgs, " ")...)
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
	stdoutFile, err := os.Create(utils.GetLogsPath() + nameRef)
	if err != nil {
		panic(err)
	}
	defer stdoutFile.Close()
	stderrFile, err := os.Create(utils.GetLogsPath() + nameRef + ".err")
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
