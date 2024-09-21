package json

import (
	"github.com/spf13/viper"
	"github.com/ssh-connection-manager/crypt"
	"ssh+/app/output"
)

// TODO сделать через цикл чтобы шифровал каждое значение и так же правильн обрабатывал ошибки

func SetCryptData(c Connect) Connect {
	var err error

	pathConf := viper.GetString("FullPathConfig")
	fileNameKey := viper.GetString("NameFileCryptKey")

	c.Alias, err = crypt.Encrypt(c.Alias, pathConf, fileNameKey)
	c.Address, err = crypt.Encrypt(c.Address, pathConf, fileNameKey)
	c.Login, err = crypt.Encrypt(c.Login, pathConf, fileNameKey)
	c.Password, err = crypt.Encrypt(c.Password, pathConf, fileNameKey)

	c.CreatedAt, err = crypt.Encrypt(c.CreatedAt, pathConf, fileNameKey)
	c.UpdatedAt, err = crypt.Encrypt(c.UpdatedAt, pathConf, fileNameKey)

	if err != nil {
		output.GetOutError("err crypt data")
	}

	return c
}

// TODO сделать через цикл чтобы шифровал каждое значение и так же правильн обрабатывал ошибки

func decryptData(c Connect) Connect {
	pathConf := viper.GetString("FullPathConfig")
	fileNameKey := viper.GetString("NameFileCryptKey")

	c.Alias, _ = crypt.Decrypt(c.Alias, pathConf, fileNameKey)
	c.Address, _ = crypt.Decrypt(c.Address, pathConf, fileNameKey)
	c.Login, _ = crypt.Decrypt(c.Login, pathConf, fileNameKey)
	c.Password, _ = crypt.Decrypt(c.Password, pathConf, fileNameKey)

	c.CreatedAt, _ = crypt.Decrypt(c.CreatedAt, pathConf, fileNameKey)
	c.UpdatedAt, _ = crypt.Decrypt(c.UpdatedAt, pathConf, fileNameKey)

	return c
}

func (c *Connections) SetDecryptData() {
	for key, connect := range c.Connects {
		c.Connects[key] = decryptData(connect)
	}
}

func (c *Connections) SetCryptAllData() {
	for key, connect := range c.Connects {
		c.Connects[key] = SetCryptData(connect)
	}
}
