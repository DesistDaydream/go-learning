package main

import (
	"crypto/sha256"

	"encoding/hex"

	"fmt"
)

func main() {
	// 生成rsa的密钥对, 并且保存到磁盘文件中
	// GenerateRsaKey(4096)
	// NewRSA(2048)

	// 待加密的信息
	message := []byte("你好 DesistDaydream！...这是一串待加密的字符串，如果你能看到，那么说明功能实现了！")

	// 使用公钥加密，私钥解密
	EncryptAndDecrypt(message)

	// 使用私钥签名，公钥验签
	SignatureAndVerifying(message)

	// myHash()
}

// 使用sha256

func myHash() {
	// sha256.Sum256([]byte("hello, go"))
	// 1. 创建哈希接口对象
	myHash := sha256.New()
	// 2. 添加数据
	src := []byte("你好 DesistDaydream！...这是一串待加密的字符串，如果你能看到，那么说明功能实现了！")
	myHash.Write(src)
	myHash.Write(src)
	myHash.Write(src)
	// 3. 计算结果
	res := myHash.Sum(nil)
	// 4. 格式化为16进制形式
	myStr := hex.EncodeToString(res)
	fmt.Printf("%s\n", myStr)
}
