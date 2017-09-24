package main

import (
	"net/http"
	"fmt"
	"log"
	"text/template"
	"strings"
	"io"
	"os"
	"time"
	"crypto/md5"
	"strconv"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() //解析参数，默认是不会解析的
	fmt.Println(r.Form) //这些信息是输出到服务器端的打印信息
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello astaxie!") //这个写入到w的是输出到客户端的
}

func Upload(w http.ResponseWriter, r *http.Request){
	fmt.Println("method:", r.Method)
	if r.Method =="GET"{
		crutime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(crutime, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))
		t, _ := template.ParseFiles("upload.gtpl")
		t.Execute(w, token)
	}else{
		r.ParseMultipartForm(32 << 20)
		file, handler, err := r.FormFile("uploadfile")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()
		fmt.Fprintf(w, "%v", handler.Filename)
		f, err := os.OpenFile("/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()
		io.Copy(f, file)
	}
}
func Login(w http.ResponseWriter, r *http.Request){
	fmt.Println("method:",r.Method)
	r.ParseForm()
	if r.Method =="GET" {
		t, _ := template.ParseFiles("login.gtpl")
		t.Execute(w, nil)
	}else if r.Method =="POST"{
		fmt.Println("path", r.URL.Path)
		fmt.Println("scheme", r.URL.Scheme)
		fmt.Println(r.Form["url_long"])
		for k, v := range r.Form {
			fmt.Println("key:", k)
			fmt.Println("val:", strings.Join(v, ""))
		}
	}
}
func main(){
	//http.HandleFunc("/" , sayhelloName)
	http.HandleFunc("/login", Login)
	http.HandleFunc("/upload", Upload)
	err := http.ListenAndServe(":8080" , nil)
	if err != nil {
		log.Fatal("listenServer",err)
	}
}