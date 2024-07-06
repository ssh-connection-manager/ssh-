package json

import (
	"ssh+/app/crypt"
)

func SetCryptData(c Connect) Connect {
	c.Alias = crypt.Encrypt(c.Alias)
	c.Address = crypt.Encrypt(c.Address)
	c.Login = crypt.Encrypt(c.Login)
	c.Password = crypt.Encrypt(c.Password)

	c.CreatedAt = crypt.Encrypt(c.CreatedAt)
	c.UpdatedAt = crypt.Encrypt(c.UpdatedAt)

	return c
}

func decryptData(c Connect) Connect {
	c.Alias = crypt.Decrypt(c.Alias)
	c.Address = crypt.Decrypt(c.Address)
	c.Login = crypt.Decrypt(c.Login)
	c.Password = crypt.Decrypt(c.Password)

	c.CreatedAt = crypt.Decrypt(c.CreatedAt)
	c.UpdatedAt = crypt.Decrypt(c.UpdatedAt)

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
