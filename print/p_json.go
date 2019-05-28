package print

import (
	"encoding/json"
	"fmt"
)

// 打印 json 对象的结构体
func JsonStruct(o interface{}) {
	data, _ := json.Marshal(o)
	fmt.Println(string(data))
}

/*
应用场景：
1、http 请求，接收到了数据，通过 json 进行了 Unmarshal 操作，这个时候，想要打印一下，当前的数据，就可以用到这个方法了。
*/
