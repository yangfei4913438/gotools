package network

import (
	"github.com/astaxie/beego"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

type HTTP struct {
	BaseUrl    string
	BaseHeader map[string]string
}

type HttpResult struct {
	StatusCode int    `json:"status_code"`
	Body       []byte `json:"body"`
}

func (api *HTTP) client(method string, url string, header map[string]string, body io.Reader) (*HttpResult, error) {
	// 生成请求对象
	request, err := http.NewRequest(method, url, body)
	if err != nil {
		beego.Error(err)
		return nil, err
	}

	// 添加默认头部信息
	if len(api.BaseHeader) > 0 {
		for k, v := range api.BaseHeader {
			request.Header.Add(k, v)
		}
	}

	// 添加自定义头部信息
	if len(header) > 0 {
		for k, v := range header {
			request.Header.Add(k, v)
		}
	}

	// 发起 http 请求
	res, err := http.DefaultClient.Do(request)
	defer res.Body.Close()
	if err != nil {
		beego.Error(err)
		return nil, err
	}

	// 生成返回数据
	send, err := ioutil.ReadAll(res.Body)
	if err != nil {
		beego.Error(err)
		return nil, err
	}

	// 返回结果
	return &HttpResult{res.StatusCode, send}, nil
}

func (api *HTTP) Get(url string, header map[string]string, params map[string]string) (*HttpResult, error) {
	// 生成 URL
	url = api.BaseUrl + url
	if len(params) > 0 {
		index := 0
		for k, v := range params {
			if index == 0 {
				url += "?" + k + "=" + v
			} else {
				url += "&" + k + "=" + v
			}
			index++
		}
	}

	return api.client("GET", url, header, nil)
}

func (api *HTTP) Post(url string, header map[string]string, data []byte) (*HttpResult, error) {
	// 生成 URL
	url = api.BaseUrl + url

	return api.client("POST", url, header, strings.NewReader(string(data)))
}

func (api *HTTP) Put(url string, header map[string]string, data []byte) (*HttpResult, error) {
	// 生成 URL
	url = api.BaseUrl + url

	return api.client("PUT", url, header, strings.NewReader(string(data)))
}

// 如果是从 data 参数中定义删除参数，使用这个方法
func (api *HTTP) DeleteFromData(url string, header map[string]string, data []byte) (*HttpResult, error) {
	// 生成 URL
	url = api.BaseUrl + url

	return api.client("DELETE", url, header, strings.NewReader(string(data)))
}

// 如果是从 url 参数中定义删除参数，使用这个方法
func (api *HTTP) DeleteFromParams(url string, header map[string]string, params map[string]string) (*HttpResult, error) {
	// 生成 URL
	url = api.BaseUrl + url
	if len(params) > 0 {
		index := 0
		for k, v := range params {
			if index == 0 {
				url += "?" + k + "=" + v
			} else {
				url += "&" + k + "=" + v
			}
			index++
		}
	}

	return api.client("DELETE", url, header, nil)
}

/*
func main ()  {
	var axios HTTP
	axios.BaseHeader = map[string]string{"X-Access-Token": "test"}

	// GET
	//params := map[string]string{"id": "5"}
	//res, err := axios.Get("http://127.0.0.1:8060/test/api/v1/user", nil, params)
	//fmt.Println(err)
	//type User struct {
	//	Name  string `json:"name"`
	//	Age   int64  `json:"age"`
	//	Email string `json:"email"`
	//}
	//var user User
	//_ = json.Unmarshal(res.Body, &user)
	//fmt.Println(user, res.StatusCode)

	// POST
	//var send = map[string]interface{}{
	//	"name": "李雷",
	//	"age": 19,
	//	"email":"lilei@qq.com",
	//}
	//data, _ := json.Marshal(send)
	//res, err := axios.Post("http://127.0.0.1:8060/test/api/v1/user", nil, data)
	//fmt.Println(err)
	//type Result struct {
	//	Code int64 `json:"code"`
	//	Message string `json:"message"`
	//}
	//var rt Result
	//_ = json.Unmarshal(res.Body, &rt)
	//fmt.Println(rt, res.StatusCode)

	// PUT
	//var send = map[string]interface{}{
	//	"id": 1,
	//	"name": "韩梅梅",
	//	"age": 19,
	//	"email":"hanmeimei@qq.com",
	//}
	//data, _ := json.Marshal(send)
	//res, err := axios.Put("http://127.0.0.1:8060/test/api/v1/user", nil, data)
	//fmt.Println(err)
	//type Result struct {
	//	Code int64 `json:"code"`
	//	Message string `json:"message"`
	//}
	//var rt Result
	//_ = json.Unmarshal(res.Body, &rt)
	//fmt.Println(rt, res.StatusCode)

	// DELETE
	var send = map[string]interface{}{
		"id": 5,
	}
	data, _ := json.Marshal(send)
	res, err := axios.DeleteFromData("http://127.0.0.1:8060/test/api/v1/user", nil, data)
	fmt.Println(err)
	type Result struct {
		Code int64 `json:"code"`
		Message string `json:"message"`
	}
	var rt Result
	_ = json.Unmarshal(res.Body, &rt)
	fmt.Println(rt, res.StatusCode)
}
*/
