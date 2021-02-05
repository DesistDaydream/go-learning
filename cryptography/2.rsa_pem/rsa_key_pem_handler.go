package main

import (
	"bytes"
	"crypto/rand"
	"crypto/sha256"
	"crypto/x509"
	"fmt"
	"os"

	"crypto/rsa"

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

// RsaKey 是公钥和私钥两个组成一组的密钥对，以及该密钥的二进制格式。
type RsaKey struct {
	bytePrivateKey []byte
	bytePublicKey  []byte
	rsaPrivateKey  *rsa.PrivateKey
	rsaPublicKey   *rsa.PublicKey
}

// NewRsaKey 生成密钥对
func NewRsaKey(bits int) *RsaKey {
	// 随机生成一个给定大小的 RSA 密钥对。可以使用 crypto 包中的 rand.Reader 来随机。
	rsaPrivateKey, _ := rsa.GenerateKey(rand.Reader, bits)
	// 从私钥中，获取公钥
	rsaPublicKey := rsaPrivateKey.PublicKey

	// 将密钥转为二进制流，以便使用 PEM 包将其编码。
	// bytePrivateKey := rsaPrivateKey.D.Bytes()
	// bytePublicKey := rsaPublicKey.N.Bytes()
	// 为什么不能直接转换而必须使用 X.509 呢，因为 go 的标准库中，无法直接将 []byte 格式的数据转换为 *rsa.PrivateKey 和 *rsa.PublicKey 类型
	// 这样会导致后面在使用 rsa 加密/解密，签名/验签时，无法正确获取使用 *rsa.PrivateKey 和 *rsa.PublicKey 类型的密钥。
	// 因为 go 没有函数或方法，可以将密钥从 []byte 类型转换为 *rsa.PrivateKey 或 *rsa.PublicKey 类型
	// =================================================================================================================
	// ！！！！！注意！！！！！在正常情况下，无需将密钥先编码为 DER 格式。这是 go 语言的强制要求。*****也是 PKCS 这个标准搞的*****
	// go 在处理密钥时，需要先将密钥先转换成 DER 格式再使用 PEM 编码的，是 go 的需求！！！！！！！！！！！！！！！！注意！！！！！！！！！！！！！！！！！！！！！！
	// ==================================================================================================================
	// 因此，只能先使用 x509 包中的方法，将密钥对转换为 PKCS#1,ASN.1 DER 的形式，并以 []byte 的数据类型保存，以供 PEM 包将其编码。
	bytePrivateKey := x509.MarshalPKCS1PrivateKey(rsaPrivateKey)
	bytePublicKey := x509.MarshalPKCS1PublicKey(&rsaPublicKey)

	// 将密钥对转换为 PEM 格式，并在密钥对中添加用于标识类型的页眉与页脚
	var (
		bufPemPrivateKey bytes.Buffer // PEM 编码后的结果需要先放到实现了 io.Writer 接口的结构体中。然后再对该结构体中的数据进行处理，以获取 字符串 或 二进制 等类型的数据
		bufPemPublicKey  bytes.Buffer // PEM 编码后的结果需要先放到实现了 io.Writer 接口的结构体中。然后再对该结构体中的数据进行处理，以获取 字符串 或 二进制 等类型的数据
	)
	// 使用 pem 包将密钥对编码为 PEM 格式的数据
	pem.Encode(&bufPemPrivateKey, &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: bytePrivateKey,
	})
	pem.Encode(&bufPemPublicKey, &pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: bytePublicKey,
	})

	// 这里就可以看到平时看到的带页眉页脚的 PEM 格式的编码后的密钥内容了。
	fmt.Printf("======== PEM 格式私钥内容：========\n%s", bufPemPrivateKey.String())
	fmt.Printf("======== PEM 格式公钥内容：========\n%s", bufPemPublicKey.String())

	return &RsaKey{
		bytePrivateKey: bufPemPrivateKey.Bytes(), // 返回 PEM 格式的二进制私钥
		bytePublicKey:  bufPemPublicKey.Bytes(),  // 返回 PEM 格式的二进制公钥
		rsaPrivateKey:  rsaPrivateKey,
		rsaPublicKey:   &rsaPublicKey,
	}
}

// RsaPemEncrypt 使用 RSA 算法，加密指定明文，其中私钥是 PEM 编码后的格式
func (r *RsaKey) RsaPemEncrypt(plaintext []byte) []byte {
	// 由于这次要通过 PEM 格式编码的公钥进行加密，所以需要先解码 PEM 格式，再将解码后的数据转换为 *rsa.PublicKey 类型
	block, _ := pem.Decode(r.bytePublicKey)
	// 之前在编码时，使用了 x509 进行了编码，所以同样，需要使用 x509 解码以获得 *rsa.PublicKey 类型的公钥
	rsaPublicKey, _ := x509.ParsePKCS1PublicKey(block.Bytes)
	// 使用公钥加密 plaintext(明文，也就是准备加密的消息)。并返回 ciphertext(密文)
	// 其中 []byte("DesistDaydream") 是加密中的标签，解密时标签需与加密时的标签相同，否则解密失败
	ciphertext, err := rsa.EncryptOAEP(sha256.New(), rand.Reader, rsaPublicKey, plaintext, []byte("DesistDaydream"))
	if err != nil {
		panic(err)
	}
	return ciphertext
}

// RsaPemDecrypt 使用 RSA 算法，解密指定密文，其中公钥是 PEM 编码后的格式
func (r *RsaKey) RsaPemDecrypt(ciphertext []byte) []byte {
	// 使用私钥解密 ciphertext(密文，也就是加过密的消息)。并返回 plaintext(明文)
	// 其中 []byte("DesistDaydream") 是加密中的标签，解密时标签需与加密时的标签相同，否则解密失败
	block, _ := pem.Decode(r.bytePrivateKey)
	rsaPrivateKey, _ := x509.ParsePKCS1PrivateKey(block.Bytes)
	plaintext, err := rsa.DecryptOAEP(sha256.New(), rand.Reader, rsaPrivateKey, ciphertext, []byte("DesistDaydream"))
	if err != nil {
		panic(err)
	}
	return plaintext
}
