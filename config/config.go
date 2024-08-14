package config

import (
	"os/user"

	"ssh+/app/output"

	"github.com/spf13/viper"
	"github.com/ssh-connection-manager/file"
)

func getHomeDir() string {
	usr, err := user.Current()
	if err != nil {
		output.GetOutError("Error retrieving user data")
	}

	return usr.HomeDir + DirectionApp
}

func existOrCreateConfig(configPath string) {
	err := viper.ReadInConfig()
	if err != nil {
		file.CreateFile(configPath)

		err = viper.ReadInConfig()
		if err != nil {
			output.GetOutError("File creation error")
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
		output.GetOutError("Error writing to configuration file")
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
