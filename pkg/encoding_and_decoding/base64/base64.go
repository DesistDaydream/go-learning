package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"math/big"
	"os"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

// randString generates a random string of the given length using crypto/rand.
func randString(n int) (string, error) {
	b := make([]byte, n)
	for i := range b {
		idx, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			return "", fmt.Errorf("generate random index: %w", err)
		}
		b[i] = charset[idx.Int64()]
	}
	return string(b), nil
}

func tmpEncode() {
	raw := make([]byte, 6072)
	_, err := rand.Read(raw)
	if err != nil {
		fmt.Printf("rand.Read error: %v\n", err)
		return
	}

	// fmt.Printf("%v\n", raw)
	fmt.Printf("待面板信息长度：%d 字节\n", len(raw))

	encoded := base64.StdEncoding.EncodeToString(raw)

	// fmt.Println(string(encoded))
	fmt.Printf("编码后长度：%d 字节\n", len(encoded))
}

func demo() {
	var result float64
	lengths := []int{3000, 3072}
	// lengths := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, l := range lengths {
		s, err := randString(l)
		if err != nil {
			fmt.Printf("error: %v\n", err)
			return
		}

		// Base64 编码后，长度变为原始长度的 4/3 倍
		// TODO: 梳理编码后的
		encoded := base64.StdEncoding.EncodeToString([]byte(s))
		// decoded, _ := base64.StdEncoding.DecodeString(encoded)

		fmt.Printf("原始长度 %d\n", l)
		// fmt.Printf("  原始: %s\n", s)
		// fmt.Printf("  编码: %s\n", encoded)
		// fmt.Printf("  解码: %s\n", decoded)
		// TODO: 把膨胀率改为 /3 的余数
		fmt.Printf("  编码长度: %d, 膨胀率: %.2f%%\n\n",
			len(encoded), float64(len(encoded)-l)/float64(l)*100)

		result += float64(len(encoded))
	}

	fmt.Printf("已编码的总长度: %v\n", result)
	fmt.Println(result * 3 / 4)
}

func tmpDemo2() {
	var (
		// 编码后的内容列表
		// encodedContents []string
		resultEncoded float64
		resultDecoded float64
	)
	files, err := os.ReadDir("./test_files/binaries")
	if err != nil {
		fmt.Printf("os.ReadDir error: %v\n", err)
		return
	}
	for _, v := range files {
		if v.IsDir() {
			continue
		}
		data, err := os.ReadFile("./test_files/binaries/" + v.Name())
		if err != nil {
			fmt.Printf("os.ReadFile error: %v\n", err)
			return
		}
		fmt.Printf("%v 已编码的字符串长度: %d（有问题日志中的内容）\n", v.Name(), len(data))
		fmt.Println(float64(len(data)) * 3 / 4)
		resultEncoded += float64(len(data))

		decoded, err := base64.StdEncoding.DecodeString(string(data))
		if err != nil {
			fmt.Printf("base64.StdEncoding.DecodeString error: %v\n", err)
			return
		}
		fmt.Printf("%v 已解码的字节长度: %d\n", v.Name(), len(decoded))
		fmt.Println(float64(len(decoded)) * 4 / 3)
		resultDecoded += float64(len(decoded))
	}

	fmt.Printf("已编码的总长度: %v\n", resultEncoded)
	fmt.Println(resultEncoded * 3 / 4)
	fmt.Printf("已解码的总长度: %v\n", resultDecoded)
	fmt.Println(resultDecoded * 4 / 3)
}

func tmpDemo3() {
	var totalStrings string
	files, err := os.ReadDir("./test_files/binaries")
	if err != nil {
		fmt.Printf("os.ReadDir error: %v\n", err)
		return
	}
	for _, v := range files {
		if v.IsDir() {
			continue
		}
		data, err := os.ReadFile("./test_files/binaries/" + v.Name())
		if err != nil {
			fmt.Printf("%v error: %v\n", v.Name(), err)
			return
		}
		totalStrings += string(data)
	}
	fmt.Printf("已编码的字符串长度: %v\n", len(totalStrings))
	decoded, _ := base64.StdEncoding.DecodeString(totalStrings)
	fmt.Printf("已解码的字节长度: %v\n", len(decoded))
}

func main() {
	// tmpEncode()
	fmt.Println("--------")
	demo()
	fmt.Println("--------")
	// tmpDemo2()
	fmt.Println("--------")
	// tmpDemo3()
}
