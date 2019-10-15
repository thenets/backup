package ssh

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/thenets/backup/utils"
)

// SSH data struct for ssh config file
type SSH struct {
	Kind     string
	Metadata struct {
		ID   string
		Name string
	}

	Spec struct {
		RemoteDir           string
		LocalDestinationDir string
		CustomRsyncArgs     string
		CustomSSHArgs       string

		Server struct {
			Host     string
			Port     int
			User     string
			Password string
			SecretID string
		}
	}
}

// Sync starts sycronization data between a remote SSH connections
// and the local dir
func (ssh *SSH) Sync() {
	var debug = true

	// Create cache dir
	var backupCacheDir = utils.GetCacheDir() + ssh.Metadata.ID + "/"
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
		"-e",
		fmt.Sprintf("\"ssh -p %d %s\"", ssh.Spec.Server.Port, ssh.Spec.CustomSSHArgs),
		"-av",
		"--timeout=300",
		"--recursive",
		"--delete",
		fmt.Sprintf("%s@%s:%s", ssh.Spec.Server.User, ssh.Spec.Server.Host, ssh.Spec.RemoteDir),
		backupCacheDir,
	}
	args = append(args, strings.Split(ssh.Spec.CustomSSHArgs, " ")...)
	if debug {
		log.Println("CMD", commandName)
		log.Println("ARG", args)
	}

	// Set log files
	stdoutFile, err := os.Create(utils.GetLogsPath() + ssh.Metadata.ID)
	if err != nil {
		panic(err)
	}
	defer stdoutFile.Close()
	stderrFile, err := os.Create(utils.GetLogsPath() + ssh.Metadata.ID + ".err")
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

}
