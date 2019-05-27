package main

import (
	"fmt"
	"regexp"
	"io/ioutil"
	"net/http"
	"bytes"
	"strings"
	"os"
	"io"
)

//全局变量
var (
	neturl string
	imgurl []string
	imgfile string
 )

//获取目标网址
func Getneturl () {
	fmt.Scanln(&neturl)
	if neturl == "" {
		fmt.Println("你输入的内容为空,请重新输入")
		Getneturl()
	}
	res,_ := regexp.MatchString(`(https?|http)://[-A-Za-z0-9+&@#/%?=~_|!:,.;]+[-A-Za-z0-9+&@#/%=~_|]`, neturl)
	if res != true {
		fmt.Println("网址格式不正确,请重新输入")
		Getneturl()
	}
	fmt.Printf("你输入的网址为:%s \n", neturl)
}
//发起http请求匹配图片
func Sendhttp() {
	res, err := http.Get(neturl)
	if err != nil {
		fmt.Println("发起请求失败,请重新开始操作")
	}
	defer res.Body.Close()
	body,_ := ioutil.ReadAll(res.Body)
	// fmt.Printf("%s",body)
	// fmt.Println(string(body))
	//定义正则对象
	r, _ := regexp.Compile("(<img.*src\\s*=\\s*(.*?)[^>]*?>)")
	u, _ := regexp.Compile("(http:|https:)\"?(.*?)(\"|>|\\s+)")
	// fmt.Println(r.FindAllString(string(body), -1))
	img := (r.FindAllString(string(body), -1))
	for _, v := range img {
		url := u.FindAllString(v, -1) 
		if url != nil {
		   var body string = ""
           imgurl = append(imgurl, strings.Join(url,body))
		}
	}
	if imgurl != nil {
		fmt.Printf("搜索内容: %s\n", imgurl)
		for _, v := range imgurl {
			defer func() {
				if err := recover(); err != nil {
					os.Exit(2)
				}
			}()
			download(v)
		}
	} else{
		fmt.Printf("未搜索到相关内容 \n")
	}
}

//下载图片
func download(imgPath string) {
	imagPath := imgPath
	reg, _ := regexp.Compile(`(\w|\d|_)*.jpg`)
	name := reg.FindStringSubmatch(imagPath)[0]
	fmt.Printf("正在下载图片:%s \n", name)
	resp, _ := http.Get(imagPath)
	body, _ := ioutil.ReadAll(resp.Body)
	out, _ := os.Create(name)
	io.Copy(out, bytes.NewReader(body))	
	fmt.Printf("下载完成:%s \n", name)
}

func main() {
	fmt.Println("请输入爬取的网址: ")
	Getneturl()
	Sendhttp()
}