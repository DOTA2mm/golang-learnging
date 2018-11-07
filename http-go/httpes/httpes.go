package httpes

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type Jar struct {
	cookies []*http.Cookie
}

func (jar *Jar) SetCookies(u *url.URL, cookies []*http.Cookie) {
	jar.cookies = cookies
}
func (jar *Jar) Cookies(u *url.URL) []*http.Cookie {
	return jar.cookies
}

type Http struct {
	Host string
}

type ResponesBody struct {
	Body       []byte
	StatusCode int
}

var client = &http.Client{nil, nil, new(Jar), 9999999999}

func (http *Http) Post(url, body string) (*ResponesBody, error) {
	resp, err := client.Post(http.Host+url, "application/x-www-form-urlencoded", strings.NewReader(body))
	if err == nil {
		body, _ := ioutil.ReadAll(resp.Body)
		res := &ResponesBody{Body: body, StatusCode: resp.StatusCode}
		defer resp.Body.Close()
		return res, nil
	}
	return nil, err
}

func (http *Http) Get(url string) (*ResponesBody, error) {
	resp, err := client.Get(http.Host + url)
	if err == nil {
		body, _ := ioutil.ReadAll(resp.Body)
		res := &ResponesBody{Body: body, StatusCode: resp.StatusCode}
		defer resp.Body.Close()
		return res, nil
	}
	return nil, err
}

func (http *Http) PostForm(url string, values url.Values) (*ResponesBody, error) {
	resp, err := client.PostForm(http.Host+url, values)
	if err == nil {
		body, _ := ioutil.ReadAll(resp.Body)
		res := &ResponesBody{Body: body, StatusCode: resp.StatusCode}
		defer resp.Body.Close()
		return res, nil
	}
	return nil, err
}
