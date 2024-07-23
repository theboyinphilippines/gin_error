package utils

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"os"
)

// 生成RSA密钥对并保存到文件（公钥和私钥）
func GenerateRSAKey() error {
	// 1、RSA生成私钥文件的核心步骤
	// 1)、生成RSA密钥对
	// 密钥长度,默认值为1024位
	bits := 1024
	privateKer, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return err
	}
	// 2)、将私钥对象转换成DER编码形式
	derPrivateKer := x509.MarshalPKCS1PrivateKey(privateKer)
	// 3)、创建私钥pem文件
	file, err := os.Create("./certs/private.pem")
	if err != nil {
		return err
	}
	// 4)、对密钥信息进行编码,写入到私钥文件中
	block := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: derPrivateKer,
	}
	err = pem.Encode(file, block)
	if err != nil {
		return err
	}
	// 2、RSA生成公钥文件的核心步骤
	// 1)、生成公钥对象
	publicKey := &privateKer.PublicKey
	// 2)、将公钥对象序列化为DER编码格式
	derPublicKey, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		return err
	}
	// 3)、创建公钥pem文件
	file, err = os.Create("./certs/public.pem")
	if err != nil {
		return err
	}
	// 4)、对公钥信息进行编码,写入到公钥文件中
	block = &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: derPublicKey,
	}
	err = pem.Encode(file, block)
	if err != nil {
		return err
	}
	return nil
}

// RSA加密字节数组,返回字节数组
func RSAEncrypt(originalBytes []byte, filename string) ([]byte, error) {
	// 1、读取公钥文件,解析出公钥对象
	publicKey, err := ReadParsePublicKey(filename)
	if err != nil {
		return nil, err
	}
	// 2、RSA加密,参数是随机数、公钥对象、需要加密的字节
	// Reader是一个全局共享的密码安全的强大的伪随机生成器
	//return rsa.EncryptPKCS1v15(rand.Reader, publicKey, originalBytes)  //这种加密不太安全
	return rsa.EncryptOAEP(sha256.New(), rand.Reader, publicKey, originalBytes, nil)
}

// RSA解密字节数组,返回字节数组
func RSADecrypt(cipherBytes []byte, filename string) ([]byte, error) {
	// 1、读取私钥文件，解析出私钥对象
	privateKey, err := ReadParsePrivaterKey(filename)
	if err != nil {
		return nil, err
	}
	// 2、ras解密,参数是随机数、私钥对象、需要解密的字节
	//return rsa.DecryptPKCS1v15(rand.Reader, privateKey, cipherBytes)
	return rsa.DecryptOAEP(sha256.New(), rand.Reader, privateKey, cipherBytes, nil)
}

// 读取公钥文件,解析出公钥对象
func ReadParsePublicKey(filename string) (*rsa.PublicKey, error) {
	// 1、读取公钥文件,获取公钥字节
	publicKeyBytes, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	// 2、解码公钥字节,生成加密块对象
	block, _ := pem.Decode(publicKeyBytes)
	if block == nil {
		return nil, errors.New("公钥信息错误！")
	}
	// 3、解析DER编码的公钥,生成公钥接口
	publicKeyInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	// 4、公钥接口转型成公钥对象
	publicKey := publicKeyInterface.(*rsa.PublicKey)
	return publicKey, nil
}

// 读取私钥文件,解析出私钥对象
func ReadParsePrivaterKey(filename string) (*rsa.PrivateKey, error) {
	// 1、读取私钥文件,获取私钥字节
	privateKeyBytes, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	// 2、对私钥文件进行编码,生成加密块对象
	block, _ := pem.Decode(privateKeyBytes)
	if block == nil {
		return nil, errors.New("私钥信息错误！")
	}
	// 3、解析DER编码的私钥,生成私钥对象
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return privateKey, nil
}

// RSA加密字符串,返回base64处理的字符串
func RSAEncryptString(originalText, filename string) (string, error) {
	cipherBytes, err := RSAEncrypt([]byte(originalText), filename)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(cipherBytes), nil
}

// RSA解密经过base64处理的加密字符串,返回加密前的明文
func RSADecryptString(cipherlText, filename string) (string, error) {
	cipherBytes, _ := base64.StdEncoding.DecodeString(cipherlText)
	originalBytes, err := RSADecrypt(cipherBytes, filename)
	if err != nil {
		return "", err
	}
	return string(originalBytes), nil
}
