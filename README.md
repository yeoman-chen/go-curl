# go-curl
go版本的curl请求库

# 安装
go get github.com/yeoman-chen/go-curl

# 使用
package main
import (
	"fmt"
	"go-curl"
)

func main() {
	req := go_curl.NewRequest()
	url := "http://ylo.yii2cms.com/test.php"
	headers := map[string]string{
		"User-Agent":   "Sublime",
		"Content-Type": "application/json",
	}
	cookies := map[string]string{
		"id": "12",
	}
	queries := map[string]string{
		"id":   "1",
		"page": "1",
	}
	resp, err := req.SetHeaders(headers).SetCookies(cookies).SetQueries(queries).Get(url)
	if err != nil {
		panic(err)
	} else {
		fmt.Println(resp.Proto)
		fmt.Println(resp.ProtoMajor)
		fmt.Println(resp.ProtoMinor)
		if resp.IsOk() {
			fmt.Println(resp.Body)
		} else {
			fmt.Println(resp.Raw)
		}
	}
}

# php接收请求

<?php
$reqMethod = $_SERVER["REQUEST_METHOD"];
$contentType = isset($_SERVER['CONTENT_TYPE']) ? $_SERVER['CONTENT_TYPE'] : $_SERVER['HTTP_ACCEPT'];
$params = $_REQUEST;

$res = [];
$res["req_method"] = $reqMethod;
$res["content_type"] = $contentType;
$res["req_params"] = $params;
echo json_encode($res);
sleep(6);
return ;
