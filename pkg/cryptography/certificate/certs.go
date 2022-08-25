package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"math/big"
	"net"
	"os"
	"time"
)

// 参考 k8s client-go 中 ./util/cert/cert.go 中 NewSelfSignedCACert() 中创建的证书模板代码
func main() {
	tmpl := x509.Certificate{
		// SerialNumber 是 CA 颁布的唯一序列号，在此使用一个大随机数来代表它
		SerialNumber: new(big.Int).SetInt64(0),
		// Subject 是一个 X.509 标准的 DN
		Subject: pkix.Name{
			Organization: []string{"GitHub"},
			CommonName:   "DesistDaydream",
		},
		// 证书的生命周期
		NotBefore: time.Now(),
		NotAfter:  time.Now().Add(365 * 24 * time.Hour),
		//KeyUsage 用来表明该证书是做什么用的
		KeyUsage: x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign,
		// 下面2个是 X509v3 extensions 字段下的 SAN
		DNSNames:    []string{"www.desistdaydream.ltd", "test.desistdaydream.ltd"},
		IPAddresses: []net.IP{net.ParseIP("127.0.0.1")},
		// 当 BasicConstraintsValid 为 true 时，系统必须信任该自建 CA 证书，否则查看证书时会报错：Expecting: TRUSTED CERTIFICATE
		// BasicConstraintsValid: true,
		// IsCA:                  true,
	}

	// 生成密钥对
	// ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	rootKey, _ := rsa.GenerateKey(rand.Reader, 2048)

	// 使用 证书模板、密钥对 创建一个证书
	certBytes, _ := x509.CreateCertificate(rand.Reader, &tmpl, &tmpl, &rootKey.PublicKey, rootKey)

	cacrt, _ := os.Create("ca.crt")
	defer cacrt.Close()
	cakey, _ := os.Create("ca.key")
	defer cakey.Close()

	// 将证书编码
	pem.Encode(cacrt, &pem.Block{
		Type:  "CERTIFICATE",
		Bytes: certBytes,
	})

	pem.Encode(cakey, &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(rootKey),
	})
}
