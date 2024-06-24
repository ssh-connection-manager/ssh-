package crypt

import (
	"ssh+/app/output"
)

func Decrypt(str string) string {
	decrypted, err := decrypt(GetKey(), str)
	if err != nil {
		output.GetOutError("Encryption error")
	}

	return decrypted
}
