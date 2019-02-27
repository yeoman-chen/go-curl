package goCurl

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type Response struct {
	Raw        *http.Response
	Headers    map[string]string
	Body       string
	Status     string
	StatusCode int
	Method     string
	Proto      string
	ProtoMajor int
	ProtoMinor int
}

//实例化对象
func NewResponse() *Response {
	return &Response{}
}

//结果是否OK
func (this *Response) IsOk() bool {
	return this.Raw.StatusCode == 200
}

//解析header信息
func (this *Response) parseHeaders() error {
	headers := map[string]string{}
	for k, v := range this.Raw.Header {
		headers[k] = v[0]
	}
	this.Headers = headers
	return nil
}

//解析body
func (this *Response) parseBody() error {
	fmt.Println(this.Raw.Status)
	if body, err := ioutil.ReadAll(this.Raw.Body); err != nil {
		panic(err)
	} else {
		this.Body = string(body)
	}
	return nil
}

//解析状态
func (this *Response) parseStatus() error {
	this.Status = this.Raw.Status
	return nil
}

//解析状态码
func (this *Response) parseStatusCode() error {
	this.StatusCode = this.Raw.StatusCode
	return nil
}

//解析协议
func (this *Response) parseProto() error {
	this.Proto = this.Raw.Proto
	return nil
}

//解析协议大版本号(例如:1.0中的1)
func (this *Response) parseProtoMajor() error {
	this.ProtoMajor = this.Raw.ProtoMajor
	return nil
}

//解析协议小版本号(例如:1.0中的0)
func (this *Response) parseProtoMinor() error {
	this.ProtoMinor = this.Raw.ProtoMinor
	return nil
}

//解析响应结果
func (this *Response) parseResponse() error {
	this.parseStatus()
	this.parseStatusCode()
	this.parseHeaders()
	this.parseBody()
	this.parseProto()
	this.parseProtoMajor()
	this.parseProtoMinor()
	return nil
}
