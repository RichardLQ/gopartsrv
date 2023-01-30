package encrypt

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"fmt"
)

//加密
func aesCtrCrypt(plainText []byte, key []byte) ([]byte, error) {

	//1. 创建cipher.Block接口
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	//2. 创建分组模式，在crypto/cipher包中
	iv := bytes.Repeat([]byte("1"), block.BlockSize())
	stream := cipher.NewCTR(block, iv)
	//3. 加密
	dst := make([]byte, len(plainText))
	stream.XORKeyStream(dst, plainText)

	return dst, nil
}

func main3() {
	source := "hello world"
	fmt.Println("原字符：", source)

	key := "1443flfsaWfdasds"
	encryptCode, _ := aesCtrCrypt([]byte(source), []byte(key))
	fmt.Println("密文：", string(encryptCode))

	decryptCode, _ := aesCtrCrypt(encryptCode, []byte(key))

	fmt.Println("解密：", string(decryptCode))
}
