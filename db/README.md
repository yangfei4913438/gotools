### beego使用mysql模块

##### 1、路由

```golang
//api路由,一级
ns := beego.NewNamespace("/pay_private",
    //api路由,二级
    beego.NSNamespace("/api",
        //api路由,三级
        beego.NSNamespace(
            "/v1",
            //测试接口
            beego.NSRouter("/default", &controllers.API{}, "get:Welcome"),

            //执行sql
            beego.NSRouter("/exec_mysql", &controllers.API{}, "post:ExecMysql"),

            //批量执行sql
            beego.NSRouter("/batch_mysql", &controllers.API{}, "post:BatchMysql"),

            //查询MySQL接口
            beego.NSRouter("/select_mysql", &controllers.API{}, "post:SelectMysql"),
        ),
    ),
)
//注册自定义namespace
beego.AddNamespace(ns)
```

##### 2、控制器

```golang
package controllers

import (
	"xxxxxx/models"
	"encoding/json"
	"github.com/astaxie/beego"
)

func (api *API) SelectMysql() {
	var receive string
	var send interface{}

	beego.Notice("SelectMysql，收到的数据：" + string(api.Ctx.Input.RequestBody))

	if err := json.Unmarshal(api.Ctx.Input.RequestBody, &receive); err != nil {
		beego.Error(err)
		//自定义HTTP状态码
		api.Ctx.ResponseWriter.WriteHeader(200)
		send = &SendMessage{false, "获取JSON数据出错: " + err.Error()}
	} else {
		//流程处理
		result, err := models.SelectMysql(receive)
		if err != nil {
			beego.Error("查询mysql出错: " + err.Error())
			//自定义HTTP状态码
			api.Ctx.ResponseWriter.WriteHeader(500)
			send = &SendMessage{false, "查询mysql出错: " + err.Error()}
		} else {
			//自定义HTTP状态码
			api.Ctx.ResponseWriter.WriteHeader(200)
			send = result
		}
	}
	//定义返回JSON
	api.Data["json"] = send

	//返回数据
	api.ServeJSON()

}

/*
post: http://127.0.0.1:9004/pay_private/api/v1/select_mysql
提交内容:
"select id, name from user where name = 'william'"
*/

func (api *API) ExecMysql() {
	var receive string
	var send interface{}

	beego.Notice("ExecMysql，收到的数据：" + string(api.Ctx.Input.RequestBody))

	if err := json.Unmarshal(api.Ctx.Input.RequestBody, &receive); err != nil {
		beego.Error(err)
		api.Ctx.ResponseWriter.WriteHeader(500)
		send = &SendMessage{false, "获取JSON数据出错: " + err.Error()}
	} else {
		//流程处理
		result, err := models.ExecMysql(&receive)
		if err != nil {
			beego.Error("执行SQL出错: " + err.Error())
			api.Ctx.ResponseWriter.WriteHeader(500)
			send = &SendMessage{false, "执行SQL出错: " + err.Error()}
		} else {
			api.Ctx.ResponseWriter.WriteHeader(200)
			send = result
		}
	}
	//定义返回JSON
	api.Data["json"] = send

	//返回数据
	api.ServeJSON()
}

/*
post: http://127.0.0.1:9004/pay_private/api/v1/exec_mysql
提交内容:
"insert into user (name, age) values ('william5', 25)"
*/

func (api *API) BatchMysql() {
	var receive []string
	var send interface{}

	beego.Notice("BatchMysql，收到的数据：" + string(api.Ctx.Input.RequestBody))

	if err := json.Unmarshal(api.Ctx.Input.RequestBody, &receive); err != nil {
		beego.Error(err)
		api.Ctx.ResponseWriter.WriteHeader(500)
		send = &SendMessage{false, "获取JSON数据出错: " + err.Error()}
	} else {
		//流程处理
		result, err := models.BatchMysql(&receive)
		if err != nil {
			beego.Error("批量执行SQL出错: " + err.Error())
			api.Ctx.ResponseWriter.WriteHeader(500)
			send = &SendMessage{false, "批量执行SQL出错: " + err.Error()}
		} else {
			api.Ctx.ResponseWriter.WriteHeader(200)
			send = result
		}
	}

	//定义返回JSON
	api.Data["json"] = send

	//返回数据
	api.ServeJSON()
}

/*
post: http://127.0.0.1:9004/pay_private/api/v1/batch_mysql
提交内容(JSON):
[
 "insert into user (name, age) values ('william3', 23)",
 "insert into user (name, age) values ('william4', 24)"
]
*/

##### 3、模块

```golang
package models

import (
	"xxxxxx/db"
	"encoding/json"
    "github.com/astaxie/beego"
)

//接收的结构体，必须都是string类型，接受完成后，再进行类型转换
type User struct {
	Test       string `json:"test"`
	Age        string `json:"age"`
	Id         string `json:"id"`
	Money      string `json:"money"`
	Name       string `json:"name"`
	Updatetime string `json:"updatetime"`
}

func SelectMysql(sql string) (*[]User, error) {
	var u []User

	res, err := db.MysqlDB.DoQuery(sql)
	if err != nil {
		beego.Error(err)
		return nil, err
	}

	if err := json.Unmarshal(res, &u); err != nil {
		beego.Error(err)
		return nil, err
	}

	return &u, nil
}

func ExecMysql(sql *string) (bool, error) {
	return db.MysqlDB.DoExec(*sql)
}

func BatchMysql(sqls *[]string) (bool, error) {
	return db.MysqlDB.DoExecBatch(*sqls)
}

```