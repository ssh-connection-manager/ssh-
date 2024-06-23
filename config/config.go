package config

import (
	"os/user"

	"ssh+/app/file"
	"ssh+/app/output"

	"github.com/spf13/viper"
)

func getHomeDir() string {
	usr, err := user.Current()
	if err != nil {
		output.GetOutError("Ошибка получения данных пользователя")
	}

	return usr.HomeDir + DirectionApp
}

func existOrCreateConfig(configPath string) {
	err := viper.ReadInConfig()
	if err != nil {
		file.CreateFile(configPath)

		err = viper.ReadInConfig()
		if err != nil {
			output.GetOutError("Ошибка создания файла")
		}
	}
}

func setConfigVariable() {
	viper.Set("NameFileConnects", NameFileConnects)
	viper.Set("NameFileCryptKey", NameFileCryptKey)
	viper.Set("FullPathConfig", getHomeDir())
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
	viper.AddConfigPath(getHomeDir())

	err := viper.ReadInConfig()
	if err != nil {
		configPath := getHomeDir() + FullNameFileConfig
		existOrCreateConfig(configPath)
		setConfigVariable()
	}
}
