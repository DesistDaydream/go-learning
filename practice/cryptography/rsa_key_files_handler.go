package main

import (
	"crypto/rand"
	"fmt"
	"os"

	"crypto/rsa"

	"crypto/x509"

	"encoding/pem"
)

// GenerateRsaKey 生成rsa的密钥对, 并且保存到磁盘文件中
func GenerateRsaKey(keySize int) {
	// ===============================
	// ========= 生成私钥文件 =========
	// ===============================
	// 1. 使用rsa中的GenerateKey方法生成私钥
	privateKey, err := rsa.GenerateKey(rand.Reader, keySize)
	if err != nil {
		panic(err)
	}
	// 2. 以PKCS#1格式整理私钥。将RSA私钥转换为PKCS＃1，ASN.1 DER形式。该形式符合 x509 标准
	derText := x509.MarshalPKCS1PrivateKey(privateKey)
	// 3. 要组织一个pem.Block(base64编码)
	block := pem.Block{
		Type:  "RSA PRIVATE KEY", // 这个地方写个字符串就行
		Bytes: derText,
	}
	// 4. pem编码
	file, err := os.Create("./cryptography/private.pem")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	pem.Encode(file, &block)

	// ===============================
	// ========== 生成公钥文件 ========
	// ===============================
	// 1. 从私钥中取出公钥
	publicKey := privateKey.PublicKey
	// 2. 以PKCS#1格式整理公钥。将RSA公钥转换为PKCS＃1，ASN.1 DER形式。该形式符合 x509 标准
	derstream := x509.MarshalPKCS1PublicKey(&publicKey)
	if err != nil {
		panic(err)
	}
	// 3. 将得到的数据放到pem.Block中
	block = pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: derstream,
	}
	// 4. pem编码
	file, err = os.Create("./cryptography/public.pem")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	pem.Encode(file, &block)
}

// GetKeyByte 读取密钥文件并转换为二进制流
func GetKeyByte(fileName string) []byte {
	// 1. 打开私钥文件, 并且读出文件内容
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	fileInfo, err := file.Stat()
	if err != nil {
		panic(err)
	}
	fileByte := make([]byte, fileInfo.Size())
	file.Read(fileByte)
	return fileByte
}

// RSAFileEncrypt RSA 加密, 公钥加密
func RSAFileEncrypt(messages []byte, fileName string) []byte {
	// 解码 pem 格式的密钥文件，若密钥文件格式错误，将会 panic
	block, _ := pem.Decode(GetKeyByte(fileName))
	if block == nil {
		fmt.Println("解码 pem 格式文件错误")
		return nil
	}
	// 1. 解析 PKCS1 格式的公钥
	rsaPublicKey, err := x509.ParsePKCS1PublicKey(block.Bytes)
	// 2. 使用公钥加密
	encryptedMessages, err := rsa.EncryptPKCS1v15(rand.Reader, rsaPublicKey, messages)
	if err != nil {
		panic(err)
	}
	return encryptedMessages
}

// RSAFileDecrypt RSA 解密
func RSAFileDecrypt(encryptedMessages []byte, fileName string) []byte {
	// 解码 pem 格式的密钥文件，并
	block, _ := pem.Decode(GetKeyByte(fileName))
	// 1. 解析 PKCS1 格式的私钥
	privKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		panic(err)
	}
	// 3. 使用私钥解密
	plainText, err := rsa.DecryptPKCS1v15(rand.Reader, privKey, encryptedMessages)
	if err != nil {
		panic(err)
	}
	return plainText
}
