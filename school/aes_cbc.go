package school

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"fmt"
)

// CBCEncrypter CBC加密
func (s *School) CBCEncrypter(text []byte) string {
	key := []byte(s.conf.AppKey)
	iv := []byte(s.conf.AppSecret[:16])
	block, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println(err)
	}
	// 填充
	paddText := padding(text, block.BlockSize())

	blockMode := cipher.NewCBCEncrypter(block, iv)

	// 加密
	result := make([]byte, len(paddText))
	blockMode.CryptBlocks(result, paddText)
	// 返回密文
	return hex.EncodeToString(result)
}

// CBCDecrypter CBC解密
func (s *School) CBCDecrypter(encrypterStr string) (string, error) {
	encrypter, err := hex.DecodeString(encrypterStr)
	if err != nil {
		return "", err
	}
	key := []byte(s.conf.AppKey)
	iv := []byte(s.conf.AppSecret[:16])

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	blockMode := cipher.NewCBCDecrypter(block, iv)
	result := make([]byte, len(encrypter))
	blockMode.CryptBlocks(result, encrypter)
	// 去除填充
	result = unPadding(result)
	result = bytes.TrimRight(result, "\x00")
	return string(result), nil
}

/*
	padding 填充模式
	text：明文内容
	blockSize：分组块大小
*/
func padding(plainText []byte, blockSize int) []byte {
	padText := bytes.Repeat([]byte{0}, blockSize-(len(plainText)%blockSize))
	newText := append(plainText, padText...)
	return newText
}

/*
	去除 padding 填充的数据
	text 待去除填充数据的原文
*/
func unPadding(text []byte) []byte {
	// 取出填充的数据 以此来获得填充数据长度
	return text[:(len(text) - int(text[len(text)-1]))]
}
