package shiro

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"github.com/go-basic/uuid"
	"io"
)

func padding(plainText []byte, blockSize int) []byte {
	n := blockSize - len(plainText)%blockSize
	temp := bytes.Repeat([]byte{byte(n)}, n)
	plainText = append(plainText, temp...)
	return plainText
}

func doAESCBCEncrypt(key []byte, content []byte) string {
	block, _ := aes.NewCipher(key)
	content = padding(content, block.BlockSize())
	iv := []byte(uuid.New())[:16]
	blockMode := cipher.NewCBCEncrypter(block, iv)
	cipherText := make([]byte, len(content))
	blockMode.CryptBlocks(cipherText, content)
	return base64.StdEncoding.EncodeToString(append(iv[:], cipherText[:]...))
}

func doAESGCMEncrypt(key []byte, content []byte) string {
	block, _ := aes.NewCipher(key)
	nonce := make([]byte, 16)
	_, _ = io.ReadFull(rand.Reader, nonce)
	aesgcm, _ := cipher.NewGCMWithNonceSize(block, 16)
	ciphertext := aesgcm.Seal(nil, nonce, content, nil)
	return base64.StdEncoding.EncodeToString(append(nonce, ciphertext...))
}
