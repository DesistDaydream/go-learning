package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"io"
	"log"
)

type AesCrypto struct {
	block cipher.Block
}

func NewAesCrypto(key []byte) *AesCrypto {
	// 实例化一个 cipher.Block，i.e. Block cipher(分组密码系统)。
	// 参数 key 是后续用于加密/解密的密钥，可以是 128 bit、192 bit、256 bit。
	block, err := aes.NewCipher(key)
	if err != nil {
		log.Fatalf("%v", err)
	}

	return &AesCrypto{
		block: block,
	}
}

// generateKey 用于确保 key 的长度满足 AES 的要求。即保证密钥长度在 128、192、256 bits 这三个其中一个。
// key 长度不足 keyLen 时用 0 补足、key 长度超过 keyLen 时截取前 keyLen 个字符。
func generateKey(key []byte, keyLen int) (genKey []byte) {
	if keyLen != 16 && keyLen != 24 && keyLen != 32 {
		keyLen = 16
	}
	genKey = make([]byte, keyLen)
	copy(genKey, key)
	for i := keyLen; i < len(key); {
		for j := 0; j < keyLen && i < len(key); j, i = j+1, i+1 {
			genKey[j] ^= key[i]
		}
	}
	return genKey
	// 为什么要使用这个函数:
	// aes.NewCipher() 方法中限制了输入 k 的长度必须为16, 24或者32，若长度不对，将会报错，并且后续代码也会出现 panic: runtime error: invalid memory address or nil pointer dereference 报错
	// 为了防止由于密钥长度不足导致的情况，我们可以手动将密钥截取或补全。这样，就算输入时密钥长度错误，依然可以正确执行后续代码。
	// 若使用随机数生成器生成指定长度的密钥的话，就不用再通过这种方式来保证密钥长度了。
	//
	// 下面是一些使用 generateKey 其他的优势
	// - 为了代码的通用性。虽然AES支持多种密钥长度,但生成一个固定长度的密钥可以使得接口更统一,使用更简单。
	// - 避免使用弱密钥。虽然密钥长度合法,但某些特定值的密钥其实是存在弱点的,容易被攻破。generateKey可以起到“混淆”的作用,避免使用明显的弱密钥。
	// - 分散密钥信息。将原始密钥进行处理后,可以使得密钥信息更加分散,一定程度上提高破解难度。
	// - 后续的兼容性。如果后续调整为只支持16字节密钥,那么有了generateKey的处理就可以平滑过渡。
	// - 最佳实践。虽然aes包已支持多密钥长度,但在实际应用中强制一个固定的密钥长度也是较好的安全编程实践。
}

func main() {
	// https://blog.csdn.net/mirage003/article/details/87868999
	msg := `{"msg":"Hello World"}`
	origData := []byte(msg)           // 待加密的数据
	key := []byte("ABCDEFGHIJKLMNOP") // 加密明文所用的密钥！！！注意: 通常我们不手动设置密钥并确保长度，而是通过随机数功能，生成指定长度的随机数作为密钥
	log.Println("原文：", string(origData))

	// 生成标准长度的密钥。避免因密钥长度错误导致的代码异常
	stdKey := generateKey(key, 16)
	log.Printf("密钥: %v; 新密钥: %v", key, stdKey)

	// 实例化一个 AES 加密系统
	aesCrypto := NewAesCrypto(stdKey)

	log.Println("------------------ ECB模式 --------------------")
	encryptedECB := aesCrypto.AesEncryptECB(origData, stdKey)
	log.Println("加密结果(hex)：", hex.EncodeToString(encryptedECB))
	log.Println("加密结果(base64)：", base64.StdEncoding.EncodeToString(encryptedECB))
	decryptedECB := aesCrypto.AesDecryptECB(encryptedECB, stdKey)
	log.Println("解密结果：", string(decryptedECB))

	log.Println("------------------ CBC模式 --------------------")
	// 没有使用标准密钥，当密钥长度不符合 AES 标准时，将会 panic
	encryptedCBC := AesEncryptCBC(origData, key)
	log.Println("加密结果(hex)：", hex.EncodeToString(encryptedCBC))
	log.Println("加密结果(base64)：", base64.StdEncoding.EncodeToString(encryptedCBC))
	decryptedCBC := AesDecryptCBC(encryptedCBC, key)
	log.Println("解密结果：", string(decryptedCBC))

	log.Println("------------------ CFB模式 --------------------")
	// 没有使用标准密钥，当密钥长度不符合 AES 标准时，将会 panic
	encryptedCFB := AesEncryptCFB(origData, key)
	log.Println("加密结果(hex)：", hex.EncodeToString(encryptedCFB))
	log.Println("加密结果(base64)：", base64.StdEncoding.EncodeToString(encryptedCFB))
	decryptedCFB := AesDecryptCFB(encryptedCFB, key)
	log.Println("解密结果：", string(decryptedCFB))
}

// =================== ECB ======================
// go 标准库中，不提供 ECB 模式的加密/解密逻辑，需要自己从头实现

// ECB 模式加密
func (aesCrypto *AesCrypto) AesEncryptECB(origData []byte, key []byte) (encrypted []byte) {
	length := (len(origData) + aes.BlockSize) / aes.BlockSize
	plain := make([]byte, length*aes.BlockSize)
	copy(plain, origData)
	pad := byte(len(plain) - len(origData))
	for i := len(origData); i < len(plain); i++ {
		plain[i] = pad
	}
	encrypted = make([]byte, len(plain))
	// 分组分块加密
	for bs, be := 0, aesCrypto.block.BlockSize(); bs <= len(origData); bs, be = bs+aesCrypto.block.BlockSize(), be+aesCrypto.block.BlockSize() {
		aesCrypto.block.Encrypt(encrypted[bs:be], plain[bs:be])
	}

	return encrypted
}

// ECB 模式解密
func (aesCrypto *AesCrypto) AesDecryptECB(encrypted []byte, key []byte) (decrypted []byte) {
	decrypted = make([]byte, len(encrypted))
	//
	for bs, be := 0, aesCrypto.block.BlockSize(); bs < len(encrypted); bs, be = bs+aesCrypto.block.BlockSize(), be+aesCrypto.block.BlockSize() {
		aesCrypto.block.Decrypt(decrypted[bs:be], encrypted[bs:be])
	}

	trim := 0
	if len(decrypted) > 0 {
		trim = len(decrypted) - int(decrypted[len(decrypted)-1])
	}

	return decrypted[:trim]
}

// =================== CBC ======================
func AesEncryptCBC(origData []byte, key []byte) (encrypted []byte) {
	block, _ := aes.NewCipher(key)
	blockSize := block.BlockSize()                              // 获取块的长度
	origData = pkcs5Padding(origData, blockSize)                // 补全码
	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize]) // 加密模式
	encrypted = make([]byte, len(origData))                     // 创建数组
	blockMode.CryptBlocks(encrypted, origData)                  // 加密
	return encrypted
}

func AesDecryptCBC(encrypted []byte, key []byte) (decrypted []byte) {
	block, _ := aes.NewCipher(key)
	blockSize := block.BlockSize()                              // 获取块的长度
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize]) // 加密模式
	decrypted = make([]byte, len(encrypted))                    // 创建数组
	blockMode.CryptBlocks(decrypted, encrypted)                 // 解密
	decrypted = pkcs5UnPadding(decrypted)                       // 去除补全码
	return decrypted
}

func pkcs5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func pkcs5UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

// =================== CFB ======================
func AesEncryptCFB(origData []byte, key []byte) (encrypted []byte) {
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	encrypted = make([]byte, aes.BlockSize+len(origData))
	iv := encrypted[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}
	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(encrypted[aes.BlockSize:], origData)
	return encrypted
}

func AesDecryptCFB(encrypted []byte, key []byte) (decrypted []byte) {
	block, _ := aes.NewCipher(key)
	if len(encrypted) < aes.BlockSize {
		panic("ciphertext too short")
	}
	iv := encrypted[:aes.BlockSize]
	encrypted = encrypted[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(encrypted, encrypted)
	return encrypted
}
