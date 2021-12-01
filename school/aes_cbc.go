package school

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

func (s *School) aesEncrypt(text []byte) (string, error) {
	key := []byte(s.conf.AppKey)
	iv := []byte(s.conf.AppSecret[:16])
	//生成cipher.Block 数据块
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	//填充内容，如果不足16位字符
	blockSize := block.BlockSize()
	originData := pad(text, blockSize)
	//加密方式
	blockMode := cipher.NewCBCEncrypter(block, iv)
	//加密，输出到[]byte数组
	crypted := make([]byte, len(originData))
	blockMode.CryptBlocks(crypted, originData)
	fmt.Println(string(crypted))
	return base64.StdEncoding.EncodeToString(crypted), nil
}

func pad(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func (s *School) aesDecrypt(text string) (string, error) {
	key := []byte(s.conf.AppKey)
	iv := []byte(s.conf.AppSecret[:16])

	decodeData, err := base64.StdEncoding.DecodeString(text)
	if err != nil {
		return "", nil
	}
	//生成密码数据块cipher.Block
	block, _ := aes.NewCipher(key)
	//解密模式
	blockMode := cipher.NewCBCDecrypter(block, iv)
	//输出到[]byte数组
	originData := make([]byte, len(decodeData))
	blockMode.CryptBlocks(originData, decodeData)
	//去除填充,并返回
	return string(unpad(originData)), nil
}

func unpad(ciphertext []byte) []byte {
	length := len(ciphertext)
	//去掉最后一次的padding
	unpadding := int(ciphertext[length-1])
	return ciphertext[:(length - unpadding)]
}

// CBCEncrypter CBC加密
func (s *School) CBCEncrypter(text []byte) string {
	key := []byte(s.conf.AppKey)
	iv := []byte(s.conf.AppSecret[:16])

	block, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println(err)
	}
	// 填充
	paddText := PKCS7Padding(text, block.BlockSize())

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
	blockMode.CryptBlocks(result, []byte(encrypter))
	// 去除填充
	result = UnPKCS7Padding(result)
	return string(result), nil
}

/*
	PKCS7Padding 填充模式
	text：明文内容
	blockSize：分组块大小
*/
func PKCS7Padding(text []byte, blockSize int) []byte {
	// 计算待填充的长度
	padding := blockSize - len(text)%blockSize
	var paddingText []byte
	if padding == 0 {
		// 已对齐，填充一整块数据，每个数据为 blockSize
		paddingText = bytes.Repeat([]byte{byte(blockSize)}, blockSize)
	} else {
		// 未对齐 填充 padding 个数据，每个数据为 padding
		paddingText = bytes.Repeat([]byte{byte(padding)}, padding)
	}
	return append(text, paddingText...)
}

/*
	去除 PKCS7Padding 填充的数据
	text 待去除填充数据的原文
*/
func UnPKCS7Padding(text []byte) []byte {
	// 取出填充的数据 以此来获得填充数据长度
	unPadding := int(text[len(text)-1])
	return text[:(len(text) - unPadding)]
}
