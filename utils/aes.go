package utils

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
)

// 对称加密
func SCEncrypt(originalBytes, key, iv []byte) ([]byte, error) {
	// 1、实例化密码器block(参数为密钥)
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	// 2、对明文进行填充(参数为原始字节切片和密码对象的区块个数)
	paddingBytes := PKCS5Padding(originalBytes, blockSize)
	// 3、实例化加密模式（参数为密码对象和密钥）
	blockMode := cipher.NewCBCEncrypter(block, iv)
	// 4、对填充字节后的明文进行加密(参数为加密字节切片和填充字节切片)
	cipherBytes := make([]byte, len(paddingBytes))
	blockMode.CryptBlocks(cipherBytes, paddingBytes)
	return cipherBytes, nil
}

// 对称解密
func SCDecrypt(cipherBytes, key, iv []byte) ([]byte, error) {
	// 1、实例化密码器block(参数为密钥)
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	//blockSize := block.BlockSize()
	// 2、实例化解密模式（参数为密码对象和密钥）
	blockMode := cipher.NewCBCDecrypter(block, iv)
	// 3、对密文进行解密(参数为填充字节切片和加密字节切片)
	paddingBytes := make([]byte, len(cipherBytes))
	blockMode.CryptBlocks(paddingBytes, cipherBytes)
	// 4、去除填充的字节(参数为填充切片)
	originalBytes := PKCS5UnPadding(paddingBytes)
	return originalBytes, nil
}

// 封装字符串对称加密
func SCEncryptString(originalText, key, iv string) (string, error) {
	cipherBytes, err := SCEncrypt([]byte(originalText), []byte(key), []byte(iv))
	if err != nil {
		return "", err
	}
	// base64编码(encoded)
	base64str := base64.StdEncoding.EncodeToString(cipherBytes)
	return base64str, nil
}

// 封装字符串对称解密
func SCDecryptString(cipherText, key, iv string) (string, error) {
	// base64解码(decode)
	cipherBytes, _ := base64.StdEncoding.DecodeString(cipherText)
	cipherBytes, err := SCDecrypt(cipherBytes, []byte(key), []byte(iv))
	if err != nil {
		return "", err
	}
	return string(cipherBytes), nil
}

// 末尾填充字节（原始数据与key长度不一样，CBC是块加密。blocksize为16时，原始数据不足就要末尾添加，原始数据是18，就要以32位的目标填充14个数据）
func PKCS5Padding(data []byte, blockSize int) []byte {
	// 要填充的值和个数
	padding := blockSize - len(data)%blockSize
	// 要填充的单个二进制值
	slice1 := []byte{byte(padding)}
	// 要填充的二进制数组
	slice2 := bytes.Repeat(slice1, padding)
	// 填充到数据末端
	return append(data, slice2...)
}

// 去除填充的字节
func PKCS5UnPadding(data []byte) []byte {
	// 获取二进制数组最后一个数值
	unpadding := data[len(data)-1]
	// 截取开始至总长度减去填充值之间的有效数据
	result := data[:(len(data) - int(unpadding))]
	return result
}
