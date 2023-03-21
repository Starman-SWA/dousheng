package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
)

func login() {
	url := "http://localhost:8080/douyin/user/login/?username=123&password=111111"
	method := "POST"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("User-Agent", "Apifox/1.0.0 (https://www.apifox.cn)")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}

func publishAction() {
	uri := "http://localhost:8080/douyin/publish/action/"
	paramName := "data"
	filePath := "../resources/bear.mp4"
	//打开要上传的文件
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println(" post err=", err)
	}
	defer file.Close()
	payload := &bytes.Buffer{}
	//创建一个multipart类型的写文件
	writer := multipart.NewWriter(payload)
	//使用给出的属性名paramName和文件名filePath创建一个新的form-data头
	part, err := writer.CreateFormFile(paramName, filePath)
	if err != nil {
		fmt.Println(" post err=", err)
	}
	//将源复制到目标，将file写入到part   是按默认的缓冲区32k循环操作的，不会将内容一次性全写入内存中,这样就能解决大文件的问题
	_, err = io.Copy(part, file)
	_ = writer.WriteField("token", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NzY5NzU5MDAsImlkIjowLCJvcmlnX2lhdCI6MTY3Njk3MjMwMH0.T7Z-U6xGc1Rqm-KgEl68Kkwqts8mabIWFpca7nMQgS8")
	_ = writer.WriteField("title", "testbear")
	err = writer.Close()
	if err != nil {
		fmt.Println(" post err=", err)
	}
	request, err := http.NewRequest("POST", uri, payload)
	request.Header.Add("User-Agent", "Apifox/1.0.0 (https://www.apifox.cn)")
	//writer.FormDataContentType() ： 返回w对应的HTTP multipart请求的Content-Type的值，多以multipart/form-data起始
	request.Header.Set("Content-Type", writer.FormDataContentType())
	//设置host，只能用request.Host = “”，不能用request.Header.Add(),也不能用request.Header.Set()来添加host
	c := &http.Client{}
	res, err := c.Do(request)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))

	//url := "http://localhost:8080/douyin/publish/action/"
	//method := "POST"
	//
	//payload := &bytes.Buffer{}
	//writer := multipart.NewWriter(payload)
	//file, errFile1 := os.Open("../resources/bear.mp4")
	//defer file.Close()
	//part1,
	//	errFile1 := writer.CreateFormFile("data", filepath.Base("./data_tmp"))
	//_, errFile1 = io.Copy(part1, file)
	//if errFile1 != nil {
	//	fmt.Println(errFile1)
	//	return
	//}
	//_ = writer.WriteField("token", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NzY4OTEwMDYsImlkIjowLCJvcmlnX2lhdCI6MTY3Njg4NzQwNn0.L_-UgPFO5j6s-qEfk83xkeunErfNuuO1P_LSAguP9Cg")
	//_ = writer.WriteField("title", "testbear")
	//err := writer.Close()
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//
	//client := &http.Client{}
	//req, err := http.NewRequest(method, url, payload)
	//
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//req.Header.Add("User-Agent", "Apifox/1.0.0 (https://www.apifox.cn)")
	//
	//req.Header.Set("Content-Type", writer.FormDataContentType())
	//res, err := client.Do(req)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//defer res.Body.Close()
	//
	//body, err := ioutil.ReadAll(res.Body)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Println(string(body))
}

func main() {
	//login()
	publishAction()
	//util.GetSnapshot("../resources/bear.mp4", "bear", 0)
}
