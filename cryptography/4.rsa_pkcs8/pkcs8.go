package main

import (
	"bytes"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
)

type RsaKey struct {
	stringPrivateKey string
	stringPublicKey  string
	rsaPrivateKey    *rsa.PrivateKey
	rsaPublicKey     *rsa.PublicKey
}

func NewRsaKey(bits int) *RsaKey {
	rsaPrivateKey, _ := rsa.GenerateKey(rand.Reader, bits)

	bytePrivateKey, _ := x509.MarshalPKCS8PrivateKey(rsaPrivateKey)
	bytePublicKey, _ := x509.MarshalPKIXPublicKey(&rsaPrivateKey.PublicKey)

	var (
		bufPemPrivateKey bytes.Buffer
		bufPemPublicKey  bytes.Buffer
	)
	// 2.然后，使用 pem 包将密钥对编码为 PEM 格式的数据
	// 因为 bytes.Buffer 这个类型的结构体实现了 pem.Encode 第一个参数的 io.Writer 接口，所以可以通过 pem 包，将 PEM 的标签和二进制类型的编码内容，再编码为 PEM 格式的数据
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
		stringPrivateKey: bufPemPrivateKey.String(),
		stringPublicKey:  bufPemPublicKey.String(),
		rsaPrivateKey:    rsaPrivateKey,
		rsaPublicKey:     &rsaPrivateKey.PublicKey,
	}
}

//Encrypt 加密
func (r *RsaKey) RsaPemEncrypt(plaintext []byte) ([]byte, error) {
	// blockLength = 密钥长度 = 一次能加密的明文长度
	// "/8" 将bit转为bytes
	// "-11" 为 PKCS#1 建议的 padding 占用了 11 个字节
	blockLength := r.rsaPublicKey.N.BitLen()/8 - 11
	//如果明文长度不大于密钥长度，可以直接加密
	if len(plaintext) <= blockLength {
		//对明文进行加密
		return rsa.EncryptPKCS1v15(rand.Reader, r.rsaPublicKey, []byte(plaintext))
	}
	//否则分段加密
	//创建一个新的缓冲区
	buffer := bytes.NewBufferString("")
	pages := len(plaintext) / blockLength //切分为多少块
	//循环加密
	for i := 0; i <= pages; i++ {
		start := i * blockLength
		end := (i + 1) * blockLength
		if i == pages { //最后一页的判断
			if start == len(plaintext) {
				continue
			}
			end = len(plaintext)
		}
		//分段加密
		chunk, err := rsa.EncryptPKCS1v15(rand.Reader, r.rsaPublicKey, plaintext[start:end])
		if err != nil {
			return nil, err
		}
		//写入缓冲区
		buffer.Write(chunk)
	}
	//读取缓冲区内容并返回，即返回加密结果
	return buffer.Bytes(), nil
}

//Decrypt 解密
func (r *RsaKey) RsaPemDecrypt(ciphertext []byte) ([]byte, error) {
	//加密后的密文长度=密钥长度。如果密文长度大于密钥长度，说明密文非一次加密形成
	//1、获取密钥长度
	blockLength := r.rsaPublicKey.N.BitLen() / 8
	if len(ciphertext) <= blockLength { //一次形成的密文直接解密
		return rsa.DecryptPKCS1v15(rand.Reader, r.rsaPrivateKey, ciphertext)
	}

	buffer := bytes.NewBufferString("")
	pages := len(ciphertext) / blockLength
	for i := 0; i <= pages; i++ { //循环解密
		start := i * blockLength
		end := (i + 1) * blockLength
		if i == pages {
			if start == len(ciphertext) {
				continue
			}
			end = len(ciphertext)
		}
		chunk, err := rsa.DecryptPKCS1v15(rand.Reader, r.rsaPrivateKey, ciphertext[start:end])
		if err != nil {
			return nil, err
		}
		buffer.Write(chunk)
	}
	return buffer.Bytes(), nil
}

//Sign 签名
func (r *RsaKey) RsaPemSign(plaintext []byte, sHash crypto.Hash) ([]byte, error) {
	hash := sHash.New()
	hash.Write(plaintext)

	// 由于这次要通过 PEM 格式编码的公钥进行加密，所以需要先解码 PEM 格式，再将解码后的数据转换为 *rsa.PublicKey 类型
	block, _ := pem.Decode([]byte(r.stringPrivateKey))
	// 之前在编码时，使用了 x509 进行了编码，所以同样，需要使用 x509 解码以获得 *rsa.PublicKey 类型的公钥
	rsaPrivateKeyInterface, _ := x509.ParsePKCS8PrivateKey(block.Bytes)
	// 使用类型断言，判断 PCKS8 格式解码出来的私钥是什么加密类型的
	rsaPrivateKey := rsaPrivateKeyInterface.(*rsa.PrivateKey)

	sign, err := rsa.SignPKCS1v15(rand.Reader, rsaPrivateKey, sHash, hash.Sum(nil))
	if err != nil {
		return nil, err
	}
	return sign, nil
}

//Verify 验签
func (r *RsaKey) RsaPemVerify(plaintext []byte, sign []byte, sHash crypto.Hash) bool {
	h := sHash.New()
	h.Write(plaintext)
	return rsa.VerifyPKCS1v15(r.rsaPublicKey, sHash, h.Sum(nil), sign) == nil
}
