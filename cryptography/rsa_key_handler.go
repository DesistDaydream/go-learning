package main

import (
	"crypto/rand"
	"crypto/sha256"

	"crypto/rsa"
)

// RSA is
type RSA struct {
	privateKey    string
	publicKey     string
	rsaPrivateKey *rsa.PrivateKey
	rsaPublicKey  *rsa.PublicKey
}

// NewRSA 实例化加密解密所需的数据
func NewRSA() *RSA {
	privateKey, _ := rsa.GenerateKey(rand.Reader, 2048)
	publicKey := privateKey.PublicKey
	return &RSA{
		privateKey:    "",
		publicKey:     "",
		rsaPrivateKey: privateKey,
		rsaPublicKey:  &publicKey,
	}
}

// RSAEncrypt 使用 RSA 算法，加密指定明文
func (r *RSA) RSAEncrypt(plaintext []byte, fileName string) []byte {
	// 使用公钥加密 plaintext(明文，也就是准备加密的消息)。并返回 ciphertext(密文)
	// 其中 []byte("DesistDaydream") 是加密中的标签，解密时标签需与加密时的标签相同，否则解密失败
	ciphertext, err := rsa.EncryptOAEP(sha256.New(), rand.Reader, r.rsaPublicKey, plaintext, []byte("DesistDaydream"))
	if err != nil {
		panic(err)
	}
	return ciphertext
}

// RSADecrypt 使用 RSA 算法，解密指定密文
func (r *RSA) RSADecrypt(ciphertext []byte, fileName string) []byte {
	// 使用私钥解密 ciphertext(密文，也就是加过密的消息)。并返回 plaintext(明文)
	// 其中 []byte("DesistDaydream") 是加密中的标签，解密时标签需与加密时的标签相同，否则解密失败
	plaintext, err := rsa.DecryptOAEP(sha256.New(), rand.Reader, r.rsaPrivateKey, ciphertext, []byte("DesistDaydream"))
	if err != nil {
		panic(err)
	}
	return plaintext
}

// RSASign RSA 签名
func (r *RSA) RSASign() {

}

// RSAVerify RSA 验签
func (r *RSA) RSAVerify() {

}
