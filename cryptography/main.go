package main

import "fmt"

func main() {
	// 生成rsa的密钥对, 并且保存到磁盘文件中
	// GenerateRsaKey(4096)
	r := NewRSA()

	// 待加密的信息
	messages := []byte("你好 DesistDaydream！...这是一串待加密的字符串，如果你能看到，那么说明功能实现了！")

	// 使用公钥加密，私钥解密
	// EncryptAndDecrypt(r, message)
	// encryptedMessages := r.RSAEncrypt(messages)
	// decryptedMessages := r.RSADecrypt(encryptedMessages)
	// fmt.Printf("解密后的字符串为：%v\n", string(decryptedMessages))

	// 使用私钥签名，公钥验签
	signature := r.RSASign(messages)
	if r.RSAVerify(messages, signature) {
		fmt.Println("验证成功")
	}

}
