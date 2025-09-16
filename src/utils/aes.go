package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
)

const (
	NonceSize = 16
)

var (
	ErrKeyError     = errors.New("E001: 密钥错误")
	ErrDecryptError = errors.New("E002: 解密失败")
	ErrFormatError  = errors.New("E003: 数据格式错误")
	ErrChannelError = errors.New("E004: 渠道缺失或无效,注意要传自己产品的编码号")
)

// Encrypt 加密数据
func Encrypt(data []byte, base64Key string) (string, error) {
	// 解码密钥
	key, err := base64.StdEncoding.DecodeString(base64Key)
	if err != nil {
		return "", ErrKeyError
	}

	// 生成随机nonce
	nonce := make([]byte, NonceSize)
	if _, err := rand.Read(nonce); err != nil {
		return "", err
	}

	// 创建加密器
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", ErrKeyError
	}

	// 加密数据
	ciphertext := make([]byte, len(data))
	stream := cipher.NewCTR(block, nonce)
	stream.XORKeyStream(ciphertext, data)

	// 拼接结果: nonce + 密文
	combined := make([]byte, 0, len(nonce)+len(ciphertext))
	combined = append(combined, nonce...)
	combined = append(combined, ciphertext...)

	return base64.StdEncoding.EncodeToString(combined), nil
}

// Decrypt 解密数据
func Decrypt(encryptedData string, base64Key string) ([]byte, error) {
	// 解码密钥
	key, err := base64.StdEncoding.DecodeString(base64Key)
	if err != nil {
		return nil, ErrKeyError
	}

	// Base64解码数据
	combined, err := base64.StdEncoding.DecodeString(encryptedData)
	if err != nil {
		return nil, ErrFormatError
	}

	// 检查数据长度
	if len(combined) < NonceSize {
		return nil, ErrFormatError
	}

	// 提取nonce
	nonce := combined[:NonceSize]

	// 提取密文
	ciphertext := combined[NonceSize:]

	// 创建解密器
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, ErrKeyError
	}

	// 解密数据
	plaintext := make([]byte, len(ciphertext))
	stream := cipher.NewCTR(block, nonce)
	stream.XORKeyStream(plaintext, ciphertext)

	return plaintext, nil
}
