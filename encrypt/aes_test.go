package encrypt_test

import (
	"github.com/go-ginseng/sql/encrypt"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Aes", func() {
	It("should encrypt and decrypt", func() {
		testAesKey := "@NcRfUjXn2r5u8x/A?D*G-KaPdSgVkYp"
		testData := "encrypt me"

		encryption := encrypt.NewAesEncryption(testAesKey)
		encrypted, err := encryption.Encrypt(testData)
		Expect(err).To(BeNil())

		decrypted, err := encryption.Decrypt(encrypted)
		Expect(err).To(BeNil())

		Expect(decrypted).To(Equal(testData))
	})

	It("should not encrypt with empty key", func() {
		testAesKey := ""
		testData := "encrypt me"

		encryption := encrypt.NewAesEncryption(testAesKey)
		_, err := encryption.Encrypt(testData)
		Expect(err).ToNot(BeNil())
	})

	It("should not decrypt with empty key", func() {
		testAesKey := ""
		testData := "encrypt me"

		encryption := encrypt.NewAesEncryption(testAesKey)
		_, err := encryption.Decrypt(testData)
		Expect(err).ToNot(BeNil())
	})
})
