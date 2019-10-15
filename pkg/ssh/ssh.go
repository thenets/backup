package ssh

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

// // Run sycronize data from a remote SSH connections and compress it
// func Run(configFilePath string) {
// 	var sshConfig = getConfig(configFilePath)
// 	sync(sshConfig)
// 	compress(sshConfig)
// }

// func sync(sshConfig Data) {
// 	// Create cache dir
// 	var backupCacheDir = helpers.GetCacheDir() + sshConfig.ID + "/"
// 	if !helpers.IsDirectory(backupCacheDir) {
// 		os.MkdirAll(backupCacheDir, 0755)
// 	}
// 	if !helpers.IsDirectory(backupCacheDir) {
// 		panic("backup cache dir can't be created: " + backupCacheDir)
// 	}

// 	// Create command line
// 	var commandName = "/usr/bin/rsync"
// 	var args []string
// 	args = []string{
// 		"-av",
// 		"--timeout=300",
// 		"--delete",
// 		fmt.Sprintf("ssh -p %d %s", sshConfig.Spec.Server.Port, sshConfig.Spec.CustomSSHArgs),
// 		sshConfig.Spec.RemoteDir,
// 		backupCacheDir,
// 	}
// 	args = append(args, strings.Split(sshConfig.Spec.CustomSSHArgs, " ")...)

// 	// Set log files
// 	stdoutFile, err := os.Create(helpers.GetLogsPath() + sshConfig.ID)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer stdoutFile.Close()
// 	stderrFile, err := os.Create(helpers.GetLogsPath() + sshConfig.ID + ".err")
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer stderrFile.Close()

// 	// Run command
// 	cmd := exec.Command(commandName, args...)
// 	cmd.Stdout = stdoutFile
// 	cmd.Stderr = stderrFile
// 	cmd.Env = os.Environ()
// 	err = cmd.Start()
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// Wait syncing finish
// 	cmd.Wait()

// 	// Show output

// }

// // TODO
// func compress() {

// }

// func getConfig(configFilePath string) Data {
// 	configFile, err := config.Loads("samples/minecraft-dir.yml")
// 	helpers.Check(err, "[ERROR] SSH connector: Config file can't be loaded")
// 	sshConfig, err := configFile.SSH()
// 	helpers.Check(err, "[ERROR] SSH connector: Config struct couldn't be created")
// 	// fmt.Printf("%#v\n", sshConfig)

// 	return sshConfig
// }
