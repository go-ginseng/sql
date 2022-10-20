package encrypt

import (
	"crypto/aes"
	"crypto/cipher"
)

type AES interface {
	Encrypt(raw string) []byte
	Decrypt(encrypted []byte) string
}

type aesEncryption struct {
	Key string
}

func NewAES(encryptKey string) AES {
	return &aesEncryption{
		Key: encryptKey,
	}
}

// Encrypt text with AES
func (a *aesEncryption) Encrypt(raw string) []byte {
	src := []byte(raw)
	cipher := getCipher(a.Key)
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

	return encrypted
}

// Decrypt blob with AES
func (a *aesEncryption) Decrypt(encrypted []byte) string {
	cipher := getCipher(a.Key)

	decrypted := make([]byte, len(encrypted))
	for bs, be := 0, cipher.BlockSize(); bs < len(encrypted); bs, be = bs+cipher.BlockSize(), be+cipher.BlockSize() {
		cipher.Decrypt(decrypted[bs:be], encrypted[bs:be])
	}

	trim := 0
	if len(decrypted) > 0 {
		trim = len(decrypted) - int(decrypted[len(decrypted)-1])
	}

	return string(decrypted[:trim])
}

func getCipher(key string) cipher.Block {
	encryptKeyNotEmpty(key)
	cipher, _ := aes.NewCipher(generateKey([]byte(key)))
	return cipher
}

func encryptKeyNotEmpty(key string) {
	if key == "" {
		panic("aes key is empty")
	}
}

func generateKey(key []byte) (genKey []byte) {
	genKey = make([]byte, 16)
	copy(genKey, key)
	for i := 16; i < len(key); {
		for j := 0; j < 16 && i < len(key); j, i = j+1, i+1 {
			genKey[j] ^= key[i]
		}
	}
	return genKey
}
