package main

import (
	"crypto"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"os"

	"crypto/rsa"
)

// RSA 是公钥和私钥两个组成一组的密钥对
type RSA struct {
	rsaPrivateKey *rsa.PrivateKey
	rsaPublicKey  *rsa.PublicKey
}

// NewRSA 生成密钥对
func NewRSA(bits int) *RSA {
	// 随机生成一个给定大小的 RSA 密钥对。可以使用 crypto 包中的 rand.Reader 来随机。
	privateKey, _ := rsa.GenerateKey(rand.Reader, bits)
	// 从私钥中，获取公钥
	publicKey := privateKey.PublicKey
	return &RSA{
		rsaPrivateKey: privateKey,
		rsaPublicKey:  &publicKey,
	}
}

// RSAEncrypt 使用 RSA 算法，加密指定明文
func (r *RSA) RSAEncrypt(plaintext []byte) []byte {
	// 使用公钥加密 plaintext(明文，也就是准备加密的消息)。并返回 ciphertext(密文)
	// 其中 []byte("DesistDaydream") 是加密中的标签，解密时标签需与加密时的标签相同，否则解密失败
	ciphertext, err := rsa.EncryptOAEP(sha256.New(), rand.Reader, r.rsaPublicKey, plaintext, []byte("DesistDaydream"))
	if err != nil {
		panic(err)
	}
	return ciphertext
}

// RSADecrypt 使用 RSA 算法，解密指定密文
func (r *RSA) RSADecrypt(ciphertext []byte) []byte {
	// 使用私钥解密 ciphertext(密文，也就是加过密的消息)。并返回 plaintext(明文)
	// 其中 []byte("DesistDaydream") 是加密中的标签，解密时标签需与加密时的标签相同，否则解密失败
	plaintext, err := rsa.DecryptOAEP(sha256.New(), rand.Reader, r.rsaPrivateKey, ciphertext, []byte("DesistDaydream"))
	if err != nil {
		panic(err)
	}
	return plaintext
}

// RSASign RSA 签名
func (r *RSA) RSASign(plaintext []byte) []byte {
	// 只有小消息可以直接签名； 因此，对消息的哈希进行签名，而不能对消息本身进行签名。
	// 这要求哈希函数必须具有抗冲突性。 SHA-256是编写本文时(2016年)应使用的最低强度的哈希函数。
	hashed := sha256.Sum256(plaintext)
	// 使用私钥签名，必须要将明文hash后才可以签名，当验证时，同样需要对明文进行hash运算。签名于验签并不用于加密消息或消息传递，仅仅作为验证传递消息方的真实性。
	signed, err := rsa.SignPKCS1v15(rand.Reader, r.rsaPrivateKey, crypto.SHA256, hashed[:])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error from signing: %s\n", err)
		return nil
	}
	fmt.Printf("已签名的消息为: %x\n", signed)
	return signed
}

// RSAVerify RSA 验签
func (r *RSA) RSAVerify(plaintext []byte, signed []byte) bool {
	// 与签名一样，只可以对 hash 后的消息进行验证。
	hashed := sha256.Sum256(plaintext)
	// 使用公钥、已签名的信息，验证签名的真实性
	err := rsa.VerifyPKCS1v15(r.rsaPublicKey, crypto.SHA256, hashed[:], signed)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error from verification: %s\n", err)
		return false
	}
	return true
}
