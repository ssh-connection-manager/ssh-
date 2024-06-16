package config

import (
	"fmt"
	"os"

	"ssh+/app/file"
	"ssh+/app/output"

	"github.com/spf13/viper"
)

func existOrCreateConfig(configPath string) {
	err := viper.ReadInConfig()
	if err != nil {
		if os.Geteuid() != IdRootUser {
			output.GetOutError("Запустите команду с помощью 'sudo'")
		}

		file.CreateFile(configPath)

		err = viper.ReadInConfig()
		fmt.Println(err)
		if err != nil {
			output.GetOutError("Ошибка создания файла")
		}
	}
}

func setConfigVariable() {
	viper.Set("NameFileConnects", NameFileConnects)
	viper.Set("NameFileCryptKey", NameFileCryptKey)
	viper.Set("FullPathConfig", FullPathConfig)
	viper.Set("Separator", Separator)
	viper.Set("Space", Space)

	err := viper.WriteConfig()
	if err != nil {
		output.GetOutError("Ошибка записи в конфигурационный файл")
	}
}

func Generate() {
	viper.New()
	viper.SetConfigName(NameFileConfig)
	viper.SetConfigType(TypeFileConfig)
	viper.AddConfigPath(FullPathConfig)

	err := viper.ReadInConfig()
	if err != nil {
		configPath := FullPathConfig + FullNameFileConfig
		existOrCreateConfig(configPath)
		setConfigVariable()
	}
}
