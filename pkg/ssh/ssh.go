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
			Host         string
			Port         int
			Password     string
			PrivateKeyID string
		}
	}
}

// Sync starts sycronization data between a remote SSH connections
// and the local dir
func (ssh *SSH) Sync() {
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
		"-av",
		"--timeout=300",
		"--delete",
		fmt.Sprintf("ssh -p %d %s", ssh.Spec.Server.Port, ssh.Spec.CustomSSHArgs),
		ssh.Spec.RemoteDir,
		backupCacheDir,
	}
	args = append(args, strings.Split(ssh.Spec.CustomSSHArgs, " ")...)

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
	cmd.Stdout = stdoutFile
	cmd.Stderr = stderrFile
	cmd.Env = os.Environ()
	err = cmd.Start()
	if err != nil {
		log.Fatal(err)
	}

	// Wait syncing finish
	cmd.Wait()

	// Show output

}
