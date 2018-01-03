package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
)

var key32Text = "12345678901234567890123456789012"
var commonIV = []byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f}

// Encrypt the source string using AES 256 algorithm
func Encrypt(source string) string {
	c, err := aes.NewCipher([]byte(key32Text))
	if err != nil {
		fmt.Printf("Error: NewCipher(%d bytes) = %s", len(key32Text), err)
	}
	cfb := cipher.NewCFBEncrypter(c, commonIV)

	srouceByte := []byte(source)
	ciphertext := make([]byte, len(srouceByte))
	cfb.XORKeyStream(ciphertext, srouceByte)
	//fmt.Printf("%s=>%x\n", source, ciphertext)
	return base64.URLEncoding.EncodeToString(ciphertext)
}

// Decrypt the source string using AES 256 algorithm
func Decrypt(source string) string {
	c, err := aes.NewCipher([]byte(key32Text))
	if err != nil {
		fmt.Printf("Error: NewCipher(%d bytes) = %s", len(key32Text), err)
	}
	cfbdec := cipher.NewCFBDecrypter(c, commonIV)

	srouceByte, _ := base64.URLEncoding.DecodeString(source)
	decryptCopy := make([]byte, len(srouceByte))
	cfbdec.XORKeyStream(decryptCopy, srouceByte)
	//fmt.Printf("%x=>%s\n", source, decryptCopy)
	return string(decryptCopy)
}
