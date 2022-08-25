package main

import "fmt"

func main() {
	// 生成rsa的密钥对
	r := NewRSA(4096)

	// 该消息有两个作用：
	// 1. 使用公钥加密的的信息
	// 2. 验证签名时所用的消息。当该消息用于签名时，通常还需要将该消息，以及用私钥签名后的消息一起发送给对方。以便对方可以根据该消息验证签名的有效性。
	messages := []byte("你好 DesistDaydream！...这是一串待加密的字符串，如果你能看到，那么说明功能实现了！")

	// 使用公钥加密，私钥解密
	encryptedMessages := r.RSAEncrypt(messages)
	decryptedMessages := r.RSADecrypt(encryptedMessages)
	fmt.Printf("解密后的字符串为：%v\n", string(decryptedMessages))

	// 使用私钥签名，公钥验签
	// 注意，验证签名需要使用签名时发送的消息作为对比，只有消息一致，才算验证通过
	signed := r.RSASign(messages)
	if r.RSAVerify(messages, signed) {
		fmt.Println("验证成功")
	}
}
