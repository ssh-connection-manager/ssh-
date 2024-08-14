package connect

import (
	"fmt"
	"os"
	"os/exec"

	"ssh+/app/json"
	"ssh+/app/output"

	"github.com/spf13/viper"
	"github.com/ssh-connection-manager/file"
)

func Ssh(c *json.Connections, alias string) {
	filePath := file.GetFullPath(
		viper.GetString("FullPathConfig"),
		viper.GetString("NameFileConnects"))

	data, err := file.ReadFile(filePath)
	if err != nil {
		output.GetOutError("Error opening a file")
	}

	c.SerializationJson(data)
	c.SetDecryptData()

	for _, v := range c.Connects {
		if v.Alias == alias {
			sshConnect(v.Address, v.Login, v.Password)
			return
		}
	}

	output.GetOutError("No connection found")
}

func sshConnect(address, login, password string) {
	sshCommand := "sshpass -p '" + password + "' ssh -o StrictHostKeyChecking=no -t " + login + "@" + address
	sshCmd := exec.Command("bash", "-c", sshCommand)

	sshCmd.Stdout = os.Stdout
	sshCmd.Stderr = os.Stderr
	sshCmd.Stdin = os.Stdin

	if err := sshCmd.Run(); err != nil {
		fmt.Println("Error while executing the command:", err)
	}
}
