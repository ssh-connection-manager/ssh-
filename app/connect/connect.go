package connect

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
	"os/exec"

	"ssh+/app/file"
	"ssh+/app/json"
	"ssh+/app/output"
)

func Ssh(c *json.Connections, alias string) {
	filePath, err := file.GetFullPath(viper.GetString("NameFileConnects"))
	if err != nil {
		output.GetOutError("Ошибка получения путя к файлу")
	}

	c.SerializationJson(file.ReadFile(filePath))
	c.SetDecryptData()

	for _, v := range c.Connects {
		if v.Alias == alias {
			sshConnect(v.Address, v.Login, v.Password)
			return
		}
	}

	output.GetOutError("Не найдено подключение")
}

func sshConnect(address, login, password string) {
	sshCommand := "sshpass -p '" + password + "' ssh -o StrictHostKeyChecking=no -t " + login + "@" + address
	sshCmd := exec.Command("bash", "-c", sshCommand)

	sshCmd.Stdout = os.Stdout
	sshCmd.Stderr = os.Stderr
	sshCmd.Stdin = os.Stdin

	if err := sshCmd.Run(); err != nil {
		fmt.Println("Ошибка при выполнении команды:", err)
	}
}
