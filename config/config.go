package config

import (
	"os"
	"path/filepath"

	"ssh+/app/file"
	"ssh+/app/output"

	"github.com/spf13/viper"
)

func Generate() {
	configHome := FullPath
	configName := NameFile
	configType := FileType

	configPath := filepath.Join(configHome, configName+"."+configType)

	viper.AddConfigPath(configHome)
	viper.SetConfigName(configName)
	viper.SetConfigType(configType)

	err := viper.ReadInConfig()
	if err != nil {
		if os.Geteuid() != IdRootUser {
			output.GetOutError("Запустите команду с помощью 'sudo'")
		}

		file.CreateFile(configPath)

		err = viper.ReadInConfig()
		if err != nil {
			output.GetOutError("Ошибка создания файла")
		}
	}
}
