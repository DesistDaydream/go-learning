package main

import (
	"crypto/rand"
	"fmt"

	"crypto/rsa"

	"crypto/x509"

	"encoding/pem"

	"os"
)

// RSAEncrypt RSA 加密, 公钥加密
func RSAEncrypt(plainText []byte, fileName string) []byte {
	// 1. 打开文件, 并且读出文件内容
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	fileInfo, err := file.Stat()
	if err != nil {
		panic(err)
	}
	buf := make([]byte, fileInfo.Size())
	file.Read(buf)

	// 2. pem解码
	block, _ := pem.Decode(buf)
	rsaPublicKey, err := x509.ParsePKCS1PublicKey(block.Bytes)
	// 3. 使用公钥加密
	cipherText, err := rsa.EncryptPKCS1v15(rand.Reader, rsaPublicKey, plainText)
	if err != nil {
		panic(err)
	}
	return cipherText
}

// RSADecrypt RSA 解密
func RSADecrypt(cipherText []byte, fileName string) []byte {
	// 1. 打开文件, 并且读出文件内容
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	fileInfo, err := file.Stat()
	if err != nil {
		panic(err)
	}
	buf := make([]byte, fileInfo.Size())
	file.Read(buf)
	// 2. pem解码
	block, _ := pem.Decode(buf)
	privKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		panic(err)
	}
	// 3. 使用私钥解密
	plainText, err := rsa.DecryptPKCS1v15(rand.Reader, privKey, cipherText)
	if err != nil {
		panic(err)
	}
	return plainText
}

// EncryptAndDecrypt 加密与解密
func EncryptAndDecrypt(message []byte) {
	// 使用公钥加密
	cipherText := RSAEncrypt(message, "./cryptography/public.pem")
	// 使用私钥解密
	plainText := RSADecrypt(cipherText, "./cryptography/private.pem")

	// 输出解密后的内容
	fmt.Printf("解密后的字符串为：%v\n", string(plainText))
}
