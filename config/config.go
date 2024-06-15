package config

import (
	"os"
	"path/filepath"

	"ssh+/app/file"
	"ssh+/app/output"

	"github.com/spf13/viper"
)

func existOrCreateConfig(configPath string) {
	if os.Geteuid() != IdRootUser {
		output.GetOutError("Запустите команду с помощью 'sudo'")
	}

	file.CreateFile(configPath)

	err := viper.ReadInConfig()
	if err != nil {
		output.GetOutError("Ошибка создания файла")
	}
}

func setConfigVariable() {
	viper.Set("NameFileConnects", NameFileConnects)
	viper.Set("NameFileCryptKey", NameFileCryptKey)
	viper.Set("Separator", Separator)
	viper.Set("Space", Space)

	err := viper.WriteConfig()
	if err != nil {
		output.GetOutError("Ошибка записи в конфигурационный файл")
	}
}

func Generate() {
	err := viper.ReadInConfig()
	if err != nil {
		configHome := FullPathConfig
		configName := NameFileConfig
		configType := FileTypeConfig

		configPath := filepath.Join(configHome, configName+"."+configType)

		viper.AddConfigPath(configHome)
		viper.SetConfigName(configName)
		viper.SetConfigType(configType)

		existOrCreateConfig(configPath)
		setConfigVariable()
	}
}
