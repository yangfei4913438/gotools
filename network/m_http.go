package network

import (
	"bytes"
	"github.com/astaxie/beego"
	"io/ioutil"
	"net/http"
	"time"
)

// 第二个参数，没有的时候，可以使用nil传入。
// 返回值，默认需要用json反序列化处理，具体看服务器返回值
func HttpGet(url string, headers map[string]string) ([]byte, error) {
	// 定义一个网络客户端
	client := &http.Client{}

	// 定义请求对象
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		beego.Error(err)
		return nil, err
	}

	// 添加头文件
	for k, v := range headers {
		req.Header.Add(k, v)
	}

	// 定义超时时间为10秒
	client.Timeout = time.Second * 10

	// 正式请求
	result, err := client.Do(req)
	defer result.Body.Close()

	beego.Trace("response code: ", result.StatusCode)

	send, err := ioutil.ReadAll(result.Body)
	if err != nil {
		beego.Error(err)
		return nil, err
	}

	return send, nil
}

/*
使用范例

auth_url := "http://127.0.0.1:9000/api/v1/auth"

var headers = make(map[string]string)
headers["x-us-authtype"] = "1"
headers["x-us-token"] = "1111111"
headers["accept-language"] = "zh-cn"
headers["time-zone"] = "-8"

res, err := HttpGet(auth_url, headers)
if err != nil {
	beego.Error(err)
}

// 打印
beego.Notice(string(res))

// 正常的用法：
type User struct {
	ID int `json:id`
    Name string `json:name`
}

var u User

if err := json.Unmarshal(res, &u); err != nil {
	beego.Error(err)
}

*/

// POST请求，第二个参数没有可以传nil, 第三个参数使用的时候，需要用json序列化之后再传入
// 返回的对象，可以认为是一个被JSON序列化的对象，直接使用JSON反序列化就可以用了。
func HttpPost(url string, headers map[string]string, data []byte) ([]byte, error) {
	// 定义一个网络客户端
	client := &http.Client{}

	// 定义请求对象
	req, err := http.NewRequest("POST", url, bytes.NewReader(data))
	if err != nil {
		beego.Error(err)
		return nil, err
	}

	// 添加头文件
	for k, v := range headers {
		req.Header.Add(k, v)
	}

	// 定义超时时间为10秒
	client.Timeout = time.Second * 10

	// 正式请求
	result, err := client.Do(req)
	defer result.Body.Close()

	beego.Trace("response code: ", result.StatusCode)

	send, err := ioutil.ReadAll(result.Body)
	if err != nil {
		beego.Error(err)
		return nil, err
	}

	return send, nil
}

/*
和上面GET的用法差不多，唯一差别就是POST，传入的参数，是需要JSON序列化的对象。
*/
