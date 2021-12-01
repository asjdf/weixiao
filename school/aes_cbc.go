package school

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
)

func (s *School)aesEncrypt(text []byte) (string,error) {
	key := []byte(s.conf.AppKey)
	iv := []byte(s.conf.AppSecret[:16])
	//生成cipher.Block 数据块
	block, err := aes.NewCipher(key)
	if err != nil {
		return "",err
	}
	//填充内容，如果不足16位字符
	blockSize := block.BlockSize()
	originData := pad(text,blockSize)
	//加密方式
	blockMode := cipher.NewCBCEncrypter(block,iv)
	//加密，输出到[]byte数组
	crypted := make([]byte,len(originData))
	blockMode.CryptBlocks(crypted,originData)
	return base64.StdEncoding.EncodeToString(crypted) , nil
}

func pad(ciphertext []byte, blockSize int) []byte{
	padding := blockSize - len(ciphertext) % blockSize
	padtext := bytes.Repeat([]byte{byte(padding)},padding)
	return append(ciphertext,padtext...)
}

func (s *School)aesDecrypt(text string) (string,error){
	key := []byte(s.conf.AppKey)
	iv := []byte(s.conf.AppSecret[:16])

	decodeData,err := base64.StdEncoding.DecodeString(text)
	if err != nil {
		return "",nil
	}
	//生成密码数据块cipher.Block
	block,_ := aes.NewCipher(key)
	//解密模式
	blockMode := cipher.NewCBCDecrypter(block,iv)
	//输出到[]byte数组
	originData := make([]byte,len(decodeData))
	blockMode.CryptBlocks(originData, decodeData)
	//去除填充,并返回
	return string(unpad(originData)),nil
}

func unpad(ciphertext []byte) []byte{
	length := len(ciphertext)
	//去掉最后一次的padding
	unpadding := int(ciphertext[length - 1])
	return ciphertext[:(length - unpadding)]
}