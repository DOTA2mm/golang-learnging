package main

import (
	"encoding/base64"
	"fmt"
)

func base64Encode(src []byte) []byte {
	return []byte(base64.StdEncoding.EncodeToString(src))
}

func base64Decode(src []byte) ([]byte, error) {
	return base64.StdEncoding.DecodeString(string(src))
}

func main() {
	// encode
	word := "你好，世界！ hello world"
	debyte := base64Encode([]byte(word))
	fmt.Println("%s base64decode is: %s\n", word, string(debyte))
	// decode
	enbyte, err := base64Decode(debyte)
	if err != nil {
		fmt.Println(err.Error())
	}
	if word != string(enbyte) {
		fmt.Println("word is not equal to enbyte")
	}
	fmt.Println(string(enbyte))
}
