package main

import (
	"fmt"
	"golang-learnging/httpes"
	"net/url"
	"regexp"
	"strings"
)

func main() {
	client := httpes.Http{"https://github.com/"}

	rep, _ := client.Get("login")
	var token = ""
	if rep.StatusCode == 200 {
		html := string(rep.Body[:])
		reg := regexp.MustCompile("<input type=\"hidden\" name=\"authenticity_token\" value=\"(\\S*)\" />")
		regStr := reg.FindString(html)
		str := strings.Replace(regStr, "<input type=\"hidden\" name=\"authenticity_token\" value=\"", "", -1)
		token = strings.Replace(str, "\" />", "", -1)
	} else {
		fmt.Println(rep.StatusCode)
		return
	}

	fmt.Println(token)

	repSession, _ := client.PostForm("session", url.Values{"authenticity_token": {token}, "login": {"dota2mm"}, "password": {"PASSWORD"}})
	if repSession.StatusCode == 200 {
		fmt.Printf("Home page: %s", repSession.Body)
	} else {
		fmt.Printf("Eroor code: %d", repSession.StatusCode)
	}
}
