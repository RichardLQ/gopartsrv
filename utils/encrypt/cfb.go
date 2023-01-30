package encrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
)

func AesEncryptCFB(origData []byte, key []byte) (encrypted []byte) {
	block, err := aes.NewCipher(key)
	if err != nil {
		//panic(err)
	}
	encrypted = make([]byte, aes.BlockSize+len(origData))
	iv := encrypted[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		//panic(err)
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
func main1() {
	source := "imslimsl"
	fmt.Println("原字符：", source)
	source1 := "e07454996f04fbdccfc2acddbf1a7507bc61a4373e2cd039"
	fmt.Println("加密字符：", source)
	key := "DrvlKOAbcxbxqMKb" //16位
	encryptCode := AesEncryptCFB([]byte(source), []byte(key))
	fmt.Println("密文：", hex.EncodeToString(encryptCode))
	str, _ := hex.DecodeString(source1)
	decryptCode := AesDecryptCFB(str, []byte(key))
	fmt.Println("解密：", string(decryptCode))
}
