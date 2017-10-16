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
	var receive models.SelectSql
	var send *SendMessage

	beego.Notice("SelectMysql，收到的数据：" + string(api.Ctx.Input.RequestBody))

	if err := json.Unmarshal(api.Ctx.Input.RequestBody, &receive); err != nil {
		beego.Error(err)
		send = &SendMessage{false, "获取JSON数据出错: " + err.Error()}
	} else {
		//流程处理
		result, err := models.SelectMysql(&receive)
		if err != nil {
			beego.Error("查询mysql出错: " + err.Error())
			send = &SendMessage{false, "查询mysql出错: " + err.Error()}
		} else {
			send = &SendMessage{true, result}
		}
	}
	//定义返回JSON
	api.Data["json"] = send

	//返回数据
	api.ServeJSON()

}

/*
post: http://127.0.0.1:9004/pay_private/api/v1/select_mysql
提交内容(JSON):
{
	"table_name": "",
	"fields": ["id","name"],
	"sql": "select id, name from user where name = 'william'"
}
*/

func (api *API) ExecMysql() {
	var receive models.SelectSql
	var send *SendMessage

	beego.Notice("ExecMysql，收到的数据：" + string(api.Ctx.Input.RequestBody))

	if err := json.Unmarshal(api.Ctx.Input.RequestBody, &receive); err != nil {
		beego.Error(err)
		send = &SendMessage{false, "获取JSON数据出错: " + err.Error()}
	} else {
		//流程处理
		result, err := models.ExecMysql(&receive.Sql)
		if err != nil {
			beego.Error("执行SQL出错: " + err.Error())
			send = &SendMessage{false, "执行SQL出错: " + err.Error()}
		} else {
			send = &SendMessage{true, result}
		}
	}
	//定义返回JSON
	api.Data["json"] = send

	//返回数据
	api.ServeJSON()

}

/*
post: http://127.0.0.1:9004/pay_private/api/v1/exec_mysql
提交内容(JSON):
{
	"sql": "insert into user (name, age) values ('william5', 25)"
}
*/

func (api *API) BatchMysql() {
	var receive []string
	var send *SendMessage

	beego.Notice("BatchMysql，收到的数据：" + string(api.Ctx.Input.RequestBody))

	if err := json.Unmarshal(api.Ctx.Input.RequestBody, &receive); err != nil {
		beego.Error(err)
		send = &SendMessage{false, "获取JSON数据出错: " + err.Error()}
	} else {
		//流程处理
		result, err := models.BatchMysql(&receive)
		if err != nil {
			beego.Error("批量执行SQL出错: " + err.Error())
			send = &SendMessage{false, "批量执行SQL出错: " + err.Error()}
		} else {
			send = &SendMessage{true, result}
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

```

##### 3、模块

```golang
package models

import (
	"xxxxxx/db"
	"errors"
	"strings"
)

type SelectSql struct {
	TableName string   `json:"table_name,omitempty"`
	Fields    []string `json:"fields,omitempty"`
	Sql       string   `json:"sql"`
}

func SelectMysql(ss *SelectSql) (results []map[string]string, err error) {
	var res, tableInfo map[int]map[int]string

	// 先判断，是否有表名参数
	if len(ss.TableName) > 0 {
		if !strings.Contains(ss.Sql, "*") {
			return nil, errors.New("按表查询时，必须使用*替代列名！如果是手动指定字段名方式，进行全部字段查询的，请勿指定表名! ")
		}

		tableInfo, err = db.MysqlDB.DoQuery("desc " + strings.ToLower(ss.TableName))
		if err != nil {
			return nil, err
		}

		res, err = db.MysqlDB.DoQuery(ss.Sql)
		if err != nil {
			return nil, err
		}

		for _, vv := range res {
			result := make(map[string]string)
			for k, v := range tableInfo {
				key := v[0]
				value := vv[k]
				result[key] = value
			}
			results = append(results, result)
		}

	} else {
		if len(ss.Fields) == 0 {
			return nil, errors.New("如果不指定表名查询，请将sql语句中查询的字段名，按顺序填写到fields参数中! ")
		}

		if strings.Contains(ss.Sql, "*") {
			return nil, errors.New("按字段查询时，不能使用*替代列名！")
		}

		var tmp string
		for _, v := range ss.Fields {
			tmp += "," + v
		}

		//去掉空格，用于检查字段是否一致
		checkStr := strings.Replace(strings.ToLower(ss.Sql), " ", "", -1)

		//匹配的字符串，需要去掉第一个逗号字符，所以是取了一个切片
		if !strings.Contains(checkStr, strings.ToLower(tmp[1:])) {
			return nil, errors.New("按字段查询时，sql语句中查询的字段名称，必须和fields中提供的字段名称一致，且顺序相同！")
		}

		res, err = db.MysqlDB.DoQuery(ss.Sql)
		if err != nil {
			return nil, err
		}

		for _, vv := range res{
			result := make(map[string]string)
			for k, v :=range ss.Fields {
				result[v] = vv[k]
			}
			results = append(results, result)
		}
	}

	return results, nil
}

func ExecMysql(sql *string) (bool, error) {
	return db.MysqlDB.DoExec(*sql)
}

func BatchMysql(sqls *[]string) (bool, error) {
	return db.MysqlDB.DoExecBatch(*sqls)
}

```