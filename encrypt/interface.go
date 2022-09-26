package encrypt

type Encryption interface {
	Encrypt(plainText string) (string, error)
	Decrypt(cipherText string) (string, error)
}
