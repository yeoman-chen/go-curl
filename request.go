/**
 * todo 20170920
 * golang版本的curl请求库
 * Request构造类，用于设置请求参数，发起http请求
 */

package go_curl

import (
	"bytes"
	"encoding/json"
	//"fmt"
	"io"
	"net/http"
	"time"
	//"net/url"
)

// Request构造类
type Request struct {
	cli *http.Client
	req *http.Request
	//Method string
	//Url      string
	Headers  map[string]string
	Cookies  map[string]string
	Queries  map[string]string
	PostData map[string]interface{}
	OverTime time.Duration
}

// 创建实例
func NewRequest() *Request {
	return &Request{OverTime: 0}
}

//设置请求方法
/*func (this *Request) SetMethod(method string) *Request {
	this.Method = method
	return this
}*/

//设置请求url
/*func (this *Request) SetUrl(url string) *Request {
	this.Url = url
	return this
}*/

//设置请求超时时间
func (this *Request) SetTimeout(timeout time.Duration) *Request {
	this.OverTime = timeout
	return this
}

//设置请求头
func (this *Request) SetHeaders(headers map[string]string) *Request {
	this.Headers = headers
	return this
}

// 将用户自定义请求头添加到http.Request实例上
func (this *Request) setHeaders() error {
	for k, v := range this.Headers {
		this.req.Header.Set(k, v)
	}
	return nil
}

//设置请求头cookies
func (this *Request) SetCookies(cookies map[string]string) *Request {
	this.Cookies = cookies
	return this
}

// 将用户自定义cookies添加到http.Request实例上
func (this *Request) setCookies() error {
	for k, v := range this.Cookies {
		this.req.AddCookie(&http.Cookie{Name: k, Value: v})
	}
	return nil
}

// 设置url查询参数
func (this *Request) SetQueries(queries map[string]string) *Request {
	this.Queries = queries
	return this
}

// 将用户自定义url查询参数添加到http.Request
func (this *Request) setQueries() error {
	q := this.req.URL.Query()
	for k, v := range this.Queries {
		q.Add(k, v)
	}
	this.req.URL.RawQuery = q.Encode()
	return nil
}

//设置片头请求的提交参数
func (this *Request) SetPostData(postData map[string]interface{}) *Request {
	this.PostData = postData
	return this
}

//发起get请求
func (this *Request) Get(url string) (*Response, error) {
	return this.Send(url, http.MethodGet)
}

//发起post请求
func (this *Request) Post(url string) (*Response, error) {
	return this.Send(url, http.MethodPost)
}

//发起head请求
func (this *Request) Head(url string) (*Response, error) {
	return this.Send(url, http.MethodHead)
}

//发起put请求
func (this *Request) Put(url string) (*Response, error) {
	return this.Send(url, http.MethodPut)
}

//发起patch请求
func (this *Request) Patch(url string) (*Response, error) {
	return this.Send(url, http.MethodPatch)
}

//发起delete请求
func (this *Request) Delete(url string) (*Response, error) {
	return this.Send(url, http.MethodDelete)
}

//发起Options请求
func (this *Request) Options(url string) (*Response, error) {
	return this.Send(url, http.MethodOptions)
}

func (this *Request) beforeRequetHandle() error {
	this.setHeaders()
	this.setCookies()
	this.setQueries()
	return nil
}

//发起请求
func (this *Request) Send(url string, method string) (*Response, error) {
	// 初始化Response对象
	response := NewResponse()
	if this.OverTime == 0 {
		this.OverTime = 15
	}

	//初始化http.Client对象
	this.cli = &http.Client{
		Timeout: this.OverTime * time.Second,
	}

	// 检测请求url是否设置
	if url == "" {
		panic("Lack of request url")
	}

	// 检测请求方式是否设置
	if method == "" {
		panic("Lack of request method")
	}

	// 加载用户自定义的post数据到http.Request
	var payload io.Reader
	if (method == "POST") && this.PostData != nil {
		if jData, err := json.Marshal(this.PostData); err != nil {
			panic(err)
		} else {
			payload = bytes.NewReader(jData)
		}
	} else {
		payload = nil
	}

	if req, err := http.NewRequest(method, url, payload); err != nil {
		panic(err)
	} else {
		this.req = req
	}
	// 请求前处理
	this.beforeRequetHandle()
	//this.setHeaders()
	//this.setCookies()
	//this.setQueries()

	if resp, err := this.cli.Do(this.req); err != nil {
		panic(err)
	} else {
		response.Raw = resp
	}

	defer response.Raw.Body.Close()
	//解析结果
	response.parseResponse()
	return response, nil
}
