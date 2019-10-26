package ssh

import (
	"fmt"
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

// TestConnection returns nil if connection was established
func (ssh *SSH) TestConnection() error {
	var err error

	// # Add or re-add key to the know_hosts
	// ssh-keygen -f "~/.ssh/known_hosts" -R ${SERVER_IP} 1>/dev/null 2>/dev/null || true
	// ssh-keyscan -p ${SSH_PORT} -H ${SERVER_IP} >> ~/.ssh/known_hosts 2>/dev/null || true

	// # HACK first one makes new "warning" be ignored
	// TEST_OUT=$(ssh -p ${SSH_PORT} -T ${SERVER_CONN} echo ok 2>&1)

	return err
}

// Sync starts sycronization data between a remote SSH connections
// and the local dir
func (ssh *SSH) Sync() {
	id := ssh.Metadata.ID

	// Create cache dir
	var backupCacheDir, err = utils.CreateCacheDir(id)
	utils.Check(err, "Can't create SSH sync cache dir")

	// Create command line
	var commandName = "/usr/bin/rsync"
	var args []string
	args = []string{
		"-av",
		"--timeout=300",
		"--recursive",
		"--delete",
		"-e",
		fmt.Sprintf("'ssh -p %d %s'", ssh.Spec.Server.Port, ssh.Spec.CustomSSHArgs),
		fmt.Sprintf("%s@%s:%s", ssh.Spec.Server.User, ssh.Spec.Server.Host, ssh.Spec.RemoteDir),
		backupCacheDir,
	}
	args = append(args, strings.Split(ssh.Spec.CustomRsyncArgs, " ")...)

	utils.RunCmd(commandName, args, id)
}
