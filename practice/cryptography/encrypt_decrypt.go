package main

import (
	"crypto/rand"
	"fmt"

	"crypto/rsa"

	"crypto/x509"

	"encoding/pem"
)

// Encrypt RSA 加密, 公钥加密
func Encrypt(plainText []byte, publicKeyFile string) []byte {
	// 解码 pem 格式的密钥文件，若密钥文件格式错误，将会 panic
	block, _ := pem.Decode(GetKeyByte(publicKeyFile))
	if block == nil {
		fmt.Println("解码 pem 格式文件错误")
		return nil
	}
	// 1. 解析 PKCS1 格式的公钥
	rsaPublicKey, err := x509.ParsePKCS1PublicKey(block.Bytes)
	// 2. 使用公钥加密
	cipherText, err := rsa.EncryptPKCS1v15(rand.Reader, rsaPublicKey, plainText)
	if err != nil {
		panic(err)
	}
	return cipherText
}

// Decrypt RSA 解密
func Decrypt(cipherText []byte, privateKeyFile string) []byte {
	// 解码 pem 格式的密钥文件，并
	block, _ := pem.Decode(GetKeyByte(privateKeyFile))
	if block == nil {
		fmt.Println("解码 pem 格式文件错误")
		return nil
	}
	// 1. 解析 PKCS1 格式的私钥
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
func EncryptAndDecrypt(message []byte, r *RSA) {
	// 使用公钥加密
	cipherText := Encrypt(message, "./practice/cryptography/public.pem")
	// 使用私钥解密
	plainText := Decrypt(cipherText, "./practice/cryptography/private.pem")

	// 输出解密后的内容
	fmt.Printf("解密后的字符串为：%v\n", string(plainText))
}
