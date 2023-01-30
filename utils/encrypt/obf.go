package encrypt

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
)

func aesEncryptOFB(data []byte, key []byte) ([]byte, error) {
	data = PKCS7Padding(data, aes.BlockSize)
	block, _ := aes.NewCipher([]byte(key))
	out := make([]byte, aes.BlockSize+len(data))
	iv := out[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}

	stream := cipher.NewOFB(block, iv)
	stream.XORKeyStream(out[aes.BlockSize:], data)
	return out, nil
}

func aesDecryptOFB(data []byte, key []byte) ([]byte, error) {
	block, _ := aes.NewCipher([]byte(key))
	iv := data[:aes.BlockSize]
	data = data[aes.BlockSize:]
	if len(data)%aes.BlockSize != 0 {
		return nil, fmt.Errorf("data is not a multiple of the block size")
	}

	out := make([]byte, len(data))
	mode := cipher.NewOFB(block, iv)
	mode.XORKeyStream(out, data)

	out = PKCS7UnPadding(out)
	return out, nil
}

//补码
//AES加密数据块分组长度必须为128bit(byte[16])，密钥长度可以是128bit(byte[16])、192bit(byte[24])、256bit(byte[32])中的任意一个。
func PKCS7Padding(ciphertext []byte, blocksize int) []byte {
	padding := blocksize - len(ciphertext)%blocksize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

//去码
func PKCS7UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}
func main4() {
	source := "hello world"
	fmt.Println("原字符：", source)
	key := "1111111111111111" //16位  32位均可
	encryptCode, _ := aesEncryptOFB([]byte(source), []byte(key))
	fmt.Println("密文：", hex.EncodeToString(encryptCode))
	decryptCode, _ := aesDecryptOFB(encryptCode, []byte(key))

	fmt.Println("解密：", string(decryptCode))
}
