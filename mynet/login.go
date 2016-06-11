/*
* @Author: buyi
* @Date:   2016-06-11 17:15:08
* @Last Modified by:   buyi
* @Last Modified time: 2016-06-11 17:15:08
 */

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

func sayhello(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)

}

func upload(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		currenttime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(currenttime, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))
		t, _ := template.ParseFiles("login.gtpl")
		t.Execute(w, token)
	} else {

		r.ParseMultipartForm(32 << 20)

		file, handler, err := r.FormFile("uploadfile")

		if err != nil {
			fmt.Println(err)
			return
		}

		defer file.Close()

		fmt.Fprintln(w, "%v", handler.Header)

		f, err := os.OpenFile(handler.Filename, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
		if err != nil {
			fmt.Println("openfile:", err)
			return
		}
		defer f.Close()

		io.Copy(f, file)
	}
}

func login(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println("method:", r.Method)
	fmt.Println("form:", r.Form)

	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}

	if r.Method == "GET" {
		//		url.Values

		currenttime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(currenttime, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))
		t, _ := template.ParseFiles("login.gtpl")
		t.Execute(w, token)
	} else {
		fmt.Fprintln(w, "Hello~~")

		//		r.Form["uname"] = []string{"<script>alert()</script>"}

		fmt.Println("uname", r.Form["uname"])
		fmt.Println("passwd", r.Form["passwd"])
		//		r.Form["uname"] = []string{"<script>alert()</script>"}
		if len(r.Form["uname"][0]) == 0 {
			fmt.Println("uname's length is zero")
		} else {
			fmt.Println("uname's length is not zero")
		}

		slice1 := []string{"apple", "pear", "banana"}

		for _, v := range slice1 {
			if v == r.Form.Get("fruit") {
				fmt.Println("fruit's value is valid")
			}
		}

		//		slice2 := []int{1, 2}

		//		for _, v := range slice2 {
		//			if v == r.Form.Get("gender") {
		//				//				if v == 1 {
		//				//					fmt.Println("fruit's value is valid")
		//				//				} else {
		//				//					fmt.Println("fruit's value is valid")
		//				//				}
		//			}
		//		}

	}
}

func main() {
	//	http.HandleFunc("/", sayhello)
	http.HandleFunc("/login", login)
	http.HandleFunc("/upload", upload)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
