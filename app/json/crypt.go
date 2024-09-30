package json

import (
	"ssh+/app/output"

	"github.com/spf13/viper"
	"github.com/ssh-connection-manager/crypt"
)

func SetCryptData(c Connect) Connect {
	var err error

	pathConf := viper.GetString("FullPathConfig")
	fileNameKey := viper.GetString("NameFileCryptKey")

	c.Alias, err = crypt.Encrypt(c.Alias, pathConf, fileNameKey)
	if err != nil {
		output.GetOutError("err crypt data")
	}
	c.Address, err = crypt.Encrypt(c.Address, pathConf, fileNameKey)
	if err != nil {
		output.GetOutError("err crypt data")
	}
	c.Login, err = crypt.Encrypt(c.Login, pathConf, fileNameKey)
	if err != nil {
		output.GetOutError("err crypt data")
	}
	c.Password, err = crypt.Encrypt(c.Password, pathConf, fileNameKey)
	if err != nil {
		output.GetOutError("err crypt data")
	}
	c.CreatedAt, err = crypt.Encrypt(c.CreatedAt, pathConf, fileNameKey)
	if err != nil {
		output.GetOutError("err crypt data")
	}
	c.UpdatedAt, err = crypt.Encrypt(c.UpdatedAt, pathConf, fileNameKey)
	if err != nil {
		output.GetOutError("err crypt data")
	}

	return c
}

func decryptData(c Connect) Connect {
	var err error

	pathConf := viper.GetString("FullPathConfig")
	fileNameKey := viper.GetString("NameFileCryptKey")

	c.Alias, err = crypt.Decrypt(c.Alias, pathConf, fileNameKey)
	if err != nil {
		output.GetOutError("err crypt data")
	}
	c.Address, err = crypt.Decrypt(c.Address, pathConf, fileNameKey)
	if err != nil {
		output.GetOutError("err crypt data")
	}
	c.Login, err = crypt.Decrypt(c.Login, pathConf, fileNameKey)
	if err != nil {
		output.GetOutError("err crypt data")
	}
	c.Password, err = crypt.Decrypt(c.Password, pathConf, fileNameKey)
	if err != nil {
		output.GetOutError("err crypt data")
	}
	c.CreatedAt, err = crypt.Decrypt(c.CreatedAt, pathConf, fileNameKey)
	if err != nil {
		output.GetOutError("err crypt data")
	}
	c.UpdatedAt, err = crypt.Decrypt(c.UpdatedAt, pathConf, fileNameKey)
	if err != nil {
		output.GetOutError("err crypt data")
	}

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
