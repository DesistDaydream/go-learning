package main

import (
	"crypto"
	"fmt"
)

func main() {
	// 生成rsa的密钥对
	r := NewRsaKey(2048)

	// 待加密的信息
	messages := []byte("你好 DesistDaydream！...这是一串待加密的字符串，如果你能看到，那么说明功能实现了！")

	// 使用公钥加密，私钥解密
	encryptedMessages, _ := r.RsaPemEncrypt(messages)
	decryptedMessages, _ := r.RsaPemDecrypt(encryptedMessages)
	fmt.Printf("解密后的字符串为：%v\n", string(decryptedMessages))

	// 使用私钥签名，公钥验签
	// 注意，验证签名需要使用签名时发送的消息作为对比，只有消息一致，才算验证通过
	signed, _ := r.RsaPemSign(messages, crypto.SHA256)
	if r.RsaPemVerify(messages, signed, crypto.SHA256) {
		fmt.Println("验证成功")
	}

}
