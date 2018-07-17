package main

import (
	"fmt"
	"time"
)

type Adress struct {
	Street           string
	HouseNumber      uint32
	HouseNumberAddOn string
	POBox            string
	ZipCode          string
	City             string
	Country          string
}

type VCard struct {
	FirstName string
	LastName  string
	NickName  string
	BirthDate time.Time
	Photo     string
	Address   map[string]*Adress
}

func init() {
	addr1 := &Adress{"guanggu", 18, "", "", "2600", "WuHan", "China"}
	addr2 := &Adress{"mingzu", 16, "", "", "2650", "WuHan", "China"}
	addrs := make(map[string]*Adress)
	addrs["youth"] = addr1
	addrs["now"] = addr2
	birthdt := time.Date(2000, 12, 22, 12, 54, 30, 0, time.Local)
	photo := "MyDocuments/MyPhotos/photo1.jpg"
	vcard := VCard{"Chuck", "Terry", "Turing", birthdt, photo, addrs}
	fmt.Printf("Here is the full VCard: %v\n", vcard)
	fmt.Printf("My Addresses are:\n %v\n %v", addr1, addr2)
}
