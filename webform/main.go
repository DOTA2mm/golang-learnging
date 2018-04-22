package main

import (
	"crypto/md5"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	// 解析url传递的参数，对于POST则解析响应包的主体（request body）
	r.ParseForm()
	fmt.Println(r.Form["s"])
	for k, v := range r.Form {
		fmt.Println("key: ", k)
		fmt.Println("val: ", strings.Join(v, " "))
	}

	fmt.Fprintf(w, "Hello Golang")
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Request method: ", r.Method)
	if r.Method == "GET" {
		// 生成页面直出 token
		crutime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(crutime, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))

		t, _ := template.ParseFiles("login.html")
		log.Println(t.Execute(w, token))
	} else {
		r.ParseForm()
		// 验证 token
		token := r.Form.Get("token")
		if token != "" {
			// 验证 token 合法性
		} else {
			fmt.Println("Token empty")
		}
		// 处理登录表单数据
		fmt.Println("username: ", template.HTMLEscapeString(r.Form.Get("username"))) // 转义
		fmt.Println("password: ", template.HTMLEscapeString(r.Form.Get("password")))
		template.HTMLEscape(w, []byte(r.Form.Get("username"))) // 输出到客户端
		if checkRadio(r) {
			fmt.Println("gender :", r.Form["gender"])
		} else {
			fmt.Println("Illegal gender value :", r.Form["gender"])
		}

		// t, err := template.New("foo").Parse(`{{define "T"}}Hello, {{.}}!{{end}}`)
		// err = t.ExecuteTemplate(out, "T", "<script>alert('you have been pwned')</script>")
	}
}

// 处理 /upload 逻辑
func upload(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Request path: %s, Request method: %s", r.URL.Path, r.Method)
	if r.Method == "GET" {
		crutime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(crutime, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))

		t, _ := template.ParseFiles("upload.html")
		t.Execute(w, token)
	} else {
		// 把上传的文件存储在内存和临时文件中
		r.ParseMultipartForm(32 << 20)
		// 获取文件句柄，然后对文件进行存储等处理
		file, handler, err := r.FormFile("uploadfile")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()
		// type FileHeader struct {
		// 	Filename string
		// 	Header   textproto.MIMEHeader
		// 	// contains filtered or unexported fields
		// }
		fmt.Fprintf(w, "%v", handler.Header)
		// 转存客户端上传的文件
		f, err := os.OpenFile("./upload/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()
		io.Copy(f, file)
	}
}

func checkRadio(r *http.Request) bool {
	// radio 表单验证
	slice := []string{"1", "2"}
	for _, v := range slice {
		if v == r.Form.Get("gender") {
			return true
		}
	}
	return false
}

func main() {
	// 路由
	http.HandleFunc("/", sayHello)
	http.HandleFunc("/login", login)
	http.HandleFunc("/upload", upload)
	// 监听端口
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
