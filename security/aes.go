package main

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"os"
)

var commonIV = []byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f}

func main() {
	//需要加密的字符串
	plaintext := []byte("password")
	if len(os.Args) > 1 {
		plaintext = []byte(os.Args[1])
	}

	// aes 加密的 字符串
	keyText := "chuck!@#$%yiqiaooaiqiy%$#@!kcuhc"
	if len(os.Args) > 2 {
		keyText = os.Args[2]
	}

	fmt.Println(len(keyText))

	// 创建加密算法aes
	c, err := aes.NewCipher([]byte(keyText))
	if err != nil {
		fmt.Printf("Error: newCipher(%d bytes) = %s", len(keyText), err)
		os.Exit(-1)
	}

	// 加密字符串
	cfb := cipher.NewCFBEncrypter(c, commonIV)
	ciphertext := make([]byte, len(plaintext))
	cfb.XORKeyStream(ciphertext, plaintext)
	fmt.Printf("%s => %x\n", plaintext, ciphertext)

	// 解密字符串
	cfbdec := cipher.NewCFBDecrypter(c, commonIV)
	plaintextCopy := make([]byte, len(plaintext))
	cfbdec.XORKeyStream(plaintextCopy, ciphertext)
	fmt.Printf("%x => %s\n", ciphertext, plaintextCopy)
}
