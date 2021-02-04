package main

import (
	"crypto/rand"

	"crypto/rsa"

	"crypto/sha256"

	"crypto/x509"

	"encoding/hex"

	"encoding/pem"

	"fmt"

	"os"
)

// GenerateRsaKey 生成rsa的密钥对, 并且保存到磁盘文件中
func GenerateRsaKey(keySize int) {
	// ============ 生成私钥 ==========
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
	file.Close()

	// ============ 生成公钥 ==========
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
	pem.Encode(file, &block)
}

// GetKeyByte 生成密钥文件的二进制流
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

func main() {
	// 生成rsa的密钥对, 并且保存到磁盘文件中
	// GenerateRsaKey(4096)

	// 待加密的信息
	message := []byte("你好 DesistDaydream！...这是一串待加密的字符串，如果你能看到，那么说明功能实现了！")

	// 使用公钥加密，私钥解密
	// EncryptAndDecrypt(message)

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
