package connectors

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/thenets/backup/config"
)

// SSHRunAll sycronize data from a remote SSH connections and compress it
func SSHRunAll(configFilePath string) {
	var sshConfig = sshGetConfig(configFilePath)
	sshSync(sshConfig)
	sshCompress(sshConfig)
}

func sshSync(sshConfig config.SSHData) {
	// Create cache dir
	var backupCacheDir = getCacheDir() + sshConfig.ID + "/"
	if !isDirectory(backupCacheDir) {
		os.MkdirAll(backupCacheDir, 0755)
	}
	if !isDirectory(backupCacheDir) {
		panic("backup cache dir can't be created: " + backupCacheDir)
	}

	// Create command line
	var commandName = "/usr/bin/rsync"
	var args []string
	args = []string{
		"-av",
		"--timeout=300",
		"--delete",
		fmt.Sprintf("ssh -p %d %s", sshConfig.Spec.Server.Port, sshConfig.Spec.CustomSSHArgs),
		sshConfig.Spec.RemoteDir,
		backupCacheDir,
	}
	args = append(args, strings.Split(sshConfig.Spec.CustomSSHArgs, " ")...)

	// Set log files
	stdoutFile, err := os.Create(getLogsPath() + sshConfig.ID)
	if err != nil {
		panic(err)
	}
	defer stdoutFile.Close()
	stderrFile, err := os.Create(getLogsPath() + sshConfig.ID + ".err")
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

func sshCompress(sshConfig config.SSHData) {

}

func sshGetConfig(configFilePath string) config.SSHData {
	configFile, err := config.Loads("samples/minecraft-dir.yml")
	check(err, "[ERROR] SSH connector: Config file can't be loaded")
	sshConfig, err := configFile.SSH()
	check(err, "[ERROR] SSH connector: Config struct couldn't be created")
	// fmt.Printf("%#v\n", sshConfig)

	return sshConfig
}
