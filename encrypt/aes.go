package encrypt

import (
	"crypto/aes"
	"errors"
)

type aseEncryption struct {
	Key string
}

func NewAesEncryption(key string) Encryption {
	return &aseEncryption{Key: key}
}

func (a *aseEncryption) Encrypt(plainText string) (string, error) {
	if a.Key == "" {
		return "", errors.New("aes key is empty")
	}
	src := []byte(plainText)
	cipher, _ := aes.NewCipher(_generateKey([]byte(a.Key)))
	length := (len(src) + aes.BlockSize) / aes.BlockSize
	plain := make([]byte, length*aes.BlockSize)
	copy(plain, src)
	pad := byte(len(plain) - len(src))
	for i := len(src); i < len(plain); i++ {
		plain[i] = pad
	}
	encrypted := make([]byte, len(plain))

	for bs, be := 0, cipher.BlockSize(); bs <= len(src); bs, be = bs+cipher.BlockSize(), be+cipher.BlockSize() {
		cipher.Encrypt(encrypted[bs:be], plain[bs:be])
	}

	return string(encrypted), nil
}

func (a *aseEncryption) Decrypt(cipherText string) (string, error) {
	if a.Key == "" {
		return "", errors.New("aes key is empty")
	}
	cipher, _ := aes.NewCipher(_generateKey([]byte(a.Key)))

	encrypted := []byte(cipherText)

	decrypted := make([]byte, len(encrypted))
	for bs, be := 0, cipher.BlockSize(); bs < len(encrypted); bs, be = bs+cipher.BlockSize(), be+cipher.BlockSize() {
		cipher.Decrypt(decrypted[bs:be], encrypted[bs:be])
	}

	trim := 0
	if len(decrypted) > 0 {
		trim = len(decrypted) - int(decrypted[len(decrypted)-1])
	}

	return string(decrypted[:trim]), nil
}

func _generateKey(key []byte) (genKey []byte) {
	genKey = make([]byte, 16)
	copy(genKey, key)
	for i := 16; i < len(key); {
		for j := 0; j < 16 && i < len(key); j, i = j+1, i+1 {
			genKey[j] ^= key[i]
		}
	}
	return genKey
}
