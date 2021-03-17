package main

import (
	"crypto"
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/DesistDaydream/GoLearning/cryptography/other/helpers"
)

func main() {
	content := strings.Repeat("H", 245) + "q"
	data := []byte(content)

	//生成 公钥 私钥
	privateKey, publicKey := helpers.NewRsa("", "").CreatePkcs8Keys(2048)
	//privateKey, publicKey := helpers.NewRsa("", "").CreateKeys(1024)
	// fmt.Printf("公钥：%v \n 私钥: %v \n", publicKey, privateKey)

	rsaObj := helpers.NewRsa(publicKey, privateKey)
	//加密
	sData, err := rsaObj.Encrypt(data)
	if err != nil {
		fmt.Println("加密失败:", err)
	}
	//解密
	pData, err := rsaObj.Decrypt(sData)
	if err != nil {
		fmt.Println("解密失败：", err)
	}
	//签名
	sign, _ := rsaObj.Sign(data, crypto.SHA256)
	//验签
	verify := rsaObj.Verify(data, sign, crypto.SHA256)
	fmt.Printf(" 加密：%v\n 解密：%v\n 签名：%v\n 验签结果：%v\n",
		hex.EncodeToString(sData),
		string(pData),
		hex.EncodeToString(sign),
		verify,
	)
}
