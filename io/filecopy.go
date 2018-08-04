package main

import (
	"fmt"
	"io"
	"os"
)

func copyfile(dstName, srcName string) (writen int64, err error) {
	src, err := os.Open(srcName)
	if err != nil {
		return
	}
	defer src.Close()

	dst, err := os.Create(dstName)
	if err != nil {
		return
	}
	defer dst.Close()

	return io.Copy(dst, src)
}

func main() {
	copyfile("copyPage.md", "Page.md")
	fmt.Println("Copy done!")
}
