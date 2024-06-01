package json

import "ssh+/app/crypt"

func SetCryptData(c Connect) Connect {
	c.Alias = crypt.Encrypt(c.Alias)
	c.Address = crypt.Encrypt(c.Address)
	c.Login = crypt.Encrypt(c.Login)
	c.Password = crypt.Encrypt(c.Password)

	return c
}

// сделать логику которая изменит сткруктуру превратив ее зашифрованные данные в дешиврованные и при обращении к данной структуре мы просто переопределим данные
func GetEncryptData(c Connections) Connect {
	return crypt.Decrypt(c.Alias)
}
