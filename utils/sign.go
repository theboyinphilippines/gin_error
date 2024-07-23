package utils

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"encoding/base64"
)

// 私钥签名过程
func RSASign(data []byte, filename string) (string, error) {
	// 1、选择hash算法,对需要签名的数据进行hash运算
	myhash := crypto.SHA256
	hashInstance := myhash.New()
	hashInstance.Write(data)
	hashed := hashInstance.Sum(nil)
	// 2、读取私钥文件,解析出私钥对象
	privateKey, err := ReadParsePrivaterKey(filename)
	if err != nil {
		return "", err
	}
	// 3、RSA数字签名(参数是随机数、私钥对象、哈希类型、签名文件的哈希串),生成base64编码的签名字符串
	bytes, err := rsa.SignPKCS1v15(rand.Reader, privateKey, myhash, hashed)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(bytes), nil
}

// 公钥验证签名过程
func RSAVerify(data []byte, base64Sig, filename string) error {
	// 1、对base64编码的签名内容进行解码,返回签名字节
	bytes, err := base64.StdEncoding.DecodeString(base64Sig)
	if err != nil {
		return err
	}
	// 2、选择hash算法,对需要签名的数据进行hash运算
	myhash := crypto.SHA256
	hashInstance := myhash.New()
	hashInstance.Write(data)
	hashed := hashInstance.Sum(nil)
	// 3、读取公钥文件,解析出公钥对象
	publicKey, err := ReadParsePublicKey(filename)
	if err != nil {
		return err
	}
	// 4、RSA验证数字签名(参数是公钥对象、哈希类型、签名文件的哈希串、签名后的字节)
	return rsa.VerifyPKCS1v15(publicKey, myhash, hashed, bytes)
}
