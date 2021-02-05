package main

import (
	"crypto/rand"
	"os"

	"crypto/rsa"

	"crypto/x509"

	"encoding/pem"
)

// GetKeyByte 读取密钥文件并转换为二进制流。该行为用于在 加密/解密，签名/验签 中。
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

// RSA 是公钥和私钥两个组成一组的密钥对，以及使用 X509 格式化后的密钥对。
type RSA struct {
	pemPrivateKey []byte
	pemPublicKey  []byte
	rsaPrivateKey *rsa.PrivateKey
	rsaPublicKey  *rsa.PublicKey
}

// NewRSA 生成密钥对
func NewRSA(bits int) *RSA {
	// 随机生成一个给定大小的 RSA 密钥对。可以使用 crypto 包中的 rand.Reader 来随机。
	privateKey, _ := rsa.GenerateKey(rand.Reader, bits)
	// 从私钥中，获取公钥
	publicKey := privateKey.PublicKey
	// 生成 X509 格式的密钥对
	pemPrivateKey, pemPublicKey := GenerateX509Key(privateKey, &publicKey)
	return &RSA{
		pemPrivateKey: pemPrivateKey,
		pemPublicKey:  pemPublicKey,
		rsaPrivateKey: privateKey,
		rsaPublicKey:  &publicKey,
	}
}

// GenerateX509Key 生成 x509 格式密钥对，并为其创建文件
func GenerateX509Key(rsaPrivateKey *rsa.PrivateKey, rsaPublicKey *rsa.PublicKey) ([]byte, []byte) {
	// 1. 以 x509 标准格式编码密钥对。将密钥对转换为 PKCS#1 ASN.1 DER 的形式
	x509Private := x509.MarshalPKCS1PrivateKey(rsaPrivateKey)
	x509Public := x509.MarshalPKCS1PublicKey(rsaPublicKey)
	// privateByte, _ := x509.MarshalPKCS8PrivateKey(rsaPrivateKey)
	// publicByte, _ := x509.MarshalPKIXPublicKey(rsaPublicKey)

	// 组织 pem 格式内容，设定要添加的页眉和页脚
	blockPrivate := pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509Private,
	}
	blockPublic := pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: x509Public,
	}
	// [可选]为 x509 格式的密钥对生成文件
	GenerateKeyFile(blockPrivate, blockPublic)

	// 基于 DER 格式编码的密钥对，再进行 base64 编码。将密钥对转换为 PEM 格式
	pemPrivateKey := pem.EncodeToMemory(&blockPrivate)
	pemPublicKey := pem.EncodeToMemory(&blockPublic)
	// fmt.Println(string(x509RsaPrivateKey))
	// fmt.Println(string(X509RsaPublicKey))

	return pemPrivateKey, pemPublicKey
}

// GenerateKeyFile 为 x509 格式密钥对生成文件
func GenerateKeyFile(blockPrivate, blockPublic pem.Block) {
	privateKeyFile, err := os.Create("./practice/cryptography/private.pem")
	if err != nil {
		panic(err)
	}
	defer privateKeyFile.Close()
	pem.Encode(privateKeyFile, &blockPrivate)
	publicKeyFile, err := os.Create("./practice/cryptography/public.pem")
	if err != nil {
		panic(err)
	}
	defer publicKeyFile.Close()
	pem.Encode(publicKeyFile, &blockPublic)
}
