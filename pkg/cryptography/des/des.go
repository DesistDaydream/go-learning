package main

import (
	"bytes"
	"crypto/des"
	"encoding/base64"
	"fmt"
)

func main() {
	// 需要加密的数据
	plaintext := []byte("desistdaydream")

	// 8字节的密钥。必须是 8 Byte
	key := []byte("mypwd123")

	// 加密
	ciphertext, err := encrypt(plaintext, key)
	if err != nil {
		fmt.Println("加密错误:", err)
		return
	}

	// Base64 编码
	encodedCiphertext := base64.StdEncoding.EncodeToString(ciphertext)
	fmt.Printf("加密并编码后的数据: %s\n", encodedCiphertext)

	// Base64 解码
	decodedCiphertext, err := base64.StdEncoding.DecodeString(encodedCiphertext)
	if err != nil {
		fmt.Println("Base64 解码错误:", err)
		return
	}

	// 解密
	decrypted, err := decrypt(decodedCiphertext, key)
	if err != nil {
		fmt.Println("解密错误:", err)
		return
	}

	fmt.Printf("解密后的数据: %s\n", decrypted)
}

func encrypt(plaintext []byte, key []byte) ([]byte, error) {
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}

	// 填充明文
	plaintext = pkcs7Padding(plaintext, block.BlockSize())

	ciphertext := make([]byte, len(plaintext))

	// ECB模式加密
	for i := 0; i < len(plaintext); i += block.BlockSize() {
		block.Encrypt(ciphertext[i:i+block.BlockSize()], plaintext[i:i+block.BlockSize()])
	}

	return ciphertext, nil
}

func decrypt(ciphertext []byte, key []byte) ([]byte, error) {
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}

	plaintext := make([]byte, len(ciphertext))

	// ECB模式解密
	for i := 0; i < len(ciphertext); i += block.BlockSize() {
		block.Decrypt(plaintext[i:i+block.BlockSize()], ciphertext[i:i+block.BlockSize()])
	}

	// 去除填充
	plaintext = pkcs7UnPadding(plaintext)

	return plaintext, nil
}

func pkcs7Padding(data []byte, blockSize int) []byte {
	padding := blockSize - len(data)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padtext...)
}

func pkcs7UnPadding(data []byte) []byte {
	length := len(data)
	unpadding := int(data[length-1])
	return data[:(length - unpadding)]
}
