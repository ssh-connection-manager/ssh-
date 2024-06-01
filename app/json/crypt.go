package json

import (
	"encoding/base64"
	"fmt"
	"ssh+/app/crypt"
)

func SetCryptData(c Connect) Connect {
	c.Alias = crypt.Encrypt(c.Alias)
	c.Address = crypt.Encrypt(c.Address)
	c.Login = crypt.Encrypt(c.Login)
	c.Password = crypt.Encrypt(c.Password)

	return c
}

func decryptData(c Connect) Connect {
	c.Alias = crypt.Decrypt(c.Alias)
	c.Address = crypt.Decrypt(c.Address)
	c.Login = crypt.Decrypt(c.Login)
	c.Password = crypt.Decrypt(c.Password)

	return c
}

func (c *Connections) SetDecryptData() {
	str := "123"
	str1 := crypt.Encrypt(str)

	for _, connect := range c.Connects {
		fmt.Printf("str - %s, type str - %T\n\n alias - %s, str - %T", str1, connect.Alias, base64.StdEncoding.EncodeToString([]byte(connect.Alias)), connect.Alias)
		connect = decryptData(connect)
	}
}
