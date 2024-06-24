package crypt

import "ssh+/app/output"

func Encrypt(str string) string {
	ciphertext, err := encrypt(GetKey(), str)
	if err != nil {
		output.GetOutError("Encryption error")
	}

	return ciphertext
}
