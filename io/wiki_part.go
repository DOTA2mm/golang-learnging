package main

import (
	"fmt"
	"io/ioutil"
)

type Page struct {
	FileName string
	Content  []byte
}

func (this *Page) save() (err error) {
	return ioutil.WriteFile(this.FileName, this.Content, 0666)
}
func (this *Page) load(fileName string) (err error) {
	this.FileName = fileName
	this.Content, err = ioutil.ReadFile(this.FileName)
	return err
}

func main() {
	page := Page{
		"Page.md",
		[]byte("# Page\n## Section1\nThis is section1."), // string -> []byte， 此处不可省略 ,
	}

	page.save()

	// load from Page.md
	var new_page Page
	new_page.load("Page.md")
	fmt.Println(string(new_page.Content)) // []byte -> string
}
