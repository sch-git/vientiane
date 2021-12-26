package utils

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
	"vientiane/server/models"
)

/**
 * AES 加解密算法
 */

func AESEncrypt(val string) (mobile string, err error) {
	fun := "Utils.AESEncrypt"

	block, err := aes.NewCipher([]byte(models.AESSalt))
	if nil != err {
		err = fmt.Errorf("%s new cipher by salt: %s err", fun, models.AESSalt)
		return
	}

	data := pKCS7Padding([]byte(val), block.BlockSize())
	crypted := make([]byte, len(data))
	encrypter := cipher.NewCBCEncrypter(block, []byte(models.AESSalt[:block.BlockSize()]))
	encrypter.CryptBlocks(crypted, data)

	mobile = base64.StdEncoding.EncodeToString(crypted)
	return
}

func AESDecrypt(val string) (mobile string, err error) {
	fun := "Utils.AESDecrypt"

	phone, err := base64.StdEncoding.DecodeString(val)
	if nil != err {
		err = fmt.Errorf("%s base64 decoding err", fun)
		return
	}

	block, err := aes.NewCipher([]byte(models.AESSalt))
	if nil != err {
		err = fmt.Errorf("%s new cipher by salt: %s err", fun, models.AESSalt)
		return
	}

	crypted := make([]byte, len(phone))
	decrypter := cipher.NewCBCDecrypter(block, []byte(models.AESSalt[:block.BlockSize()]))

	decrypter.CryptBlocks(crypted, phone)
	crypted = pKCS7UnPadding(crypted)

	mobile = string(crypted)
	return
}

func pKCS7Padding(cipherText []byte, blockSize int) []byte {
	padding := blockSize - len(cipherText)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(cipherText, padText...)
}

func pKCS7UnPadding(origData []byte) []byte {
	length := len(origData)
	unPadding := int(origData[length-1])
	return origData[:(length - unPadding)]
}
