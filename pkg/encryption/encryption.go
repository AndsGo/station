package encryption

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io"
)

// AES加密函数
func AESEncrypt(plainText, key string) (string, error) {
	// 创建 AES 加密器
	block, err := aes.NewCipher(createHash(key))
	if err != nil {
		return "", err
	}

	// 使用 GCM 模式
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	// 创建 nonce
	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	// 加密数据
	cipherText := gcm.Seal(nonce, nonce, []byte(plainText), nil)
	return base64.StdEncoding.EncodeToString(cipherText), nil
}

// AES解密函数
func AESDecrypt(cipherText, key string) (string, error) {
	// 将密文解码
	data, err := base64.StdEncoding.DecodeString(cipherText)
	if err != nil {
		return "", err
	}

	// 创建 AES 解密器
	block, err := aes.NewCipher(createHash(key))
	if err != nil {
		return "", err
	}

	// 使用 GCM 模式
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonceSize := gcm.NonceSize()
	if len(data) < nonceSize {
		return "", fmt.Errorf("ciphertext too short")
	}
	// 解密
	nonce, cipherTextByte := data[:nonceSize], data[nonceSize:]
	plainText, err := gcm.Open(nil, nonce, cipherTextByte, nil)
	if err != nil {
		return "", err
	}

	return string(plainText), nil
}

// 创建哈希函数
func createHash(key string) []byte {
	hash := sha256.New()
	hash.Write([]byte(key))
	return hash.Sum(nil)[:32] // AES-256 密钥需要32字节
}
