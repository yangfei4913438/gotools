### beego使用mysql模块

```golang
package main

import (
	"encoding/json"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"strconv"
	"time"
)

// 用户结构体。数据类型是指针，可以防止查询出来的对象是null(也就是go里面的nil)
type Users struct {
	Id   *int    `db:"id"`
	Name *string `db:"name"`
	Age  *int    `db:"age"`
	City *string `db:"city"`
}

type mysqlType struct {
	*sqlx.DB
}

var MysqlDB *mysqlType

func init() {
	//读取MySQL配置
	var mysqlUser = "root"
	var mysqlPassword = "111111"
	var mysqlNet = "tcp"
	var mysqlHost = "10.0.0.252"
	var mysqlPort = "3306"
	var mysqlDb = "cydex"
	var mysqlCharset = "utf8"
	var mysqlMaxLifeTime = 300
	var mysqlMaxOpenConns = 1000
	var mysqlMaxIdleConns = 20

	//拼接成MySQL连接串
	var mysqlSource string
	mysqlSource = mysqlUser + ":" + mysqlPassword + "@" + mysqlNet + "(" + mysqlHost + ":" + mysqlPort + ")"
	mysqlSource += "/" + mysqlDb + "?" + "charset=" + mysqlCharset

	var err error
	db, err := sqlx.Connect("mysql", mysqlSource)
	if err != nil {
		beego.Critical("Connect to Mysql, Error: " + err.Error())
		panic("Connect to Mysql, Error: " + err.Error())
	}

	//实例化一个mysql连接对象
	MysqlDB = &mysqlType{db}

	//SetConnMaxLifetime连接的最大空闲时间(可选)
	MysqlDB.SetConnMaxLifetime(time.Duration(mysqlMaxLifeTime) * time.Second)
	//SetMaxOpenConns用于设置最大打开的连接数，默认值为0表示不限制。
	MysqlDB.SetMaxOpenConns(mysqlMaxOpenConns)
	//SetMaxIdleConns用于设置闲置的连接数。
	MysqlDB.SetMaxIdleConns(mysqlMaxIdleConns)

	if err := MysqlDB.Ping(); err != nil {
		beego.Critical("Attempt to connect to MySQL failed, Error: " + err.Error())
		panic("Attempt to connect to MySQL failed, Error: " + err.Error())
	} else {
		beego.Info("Connect Mysql Server(" + mysqlHost + ":" + mysqlPort + ") to successful!")
	}

}

func main() {
	var u Users

	err := MysqlDB.Get(&u, "select * from users where id=?", 1)
	if err != nil {
		beego.Error(err)
	}
	u_str, err := json.Marshal(u)
	if err != nil {
		beego.Error(err)
	}
	beego.Notice("单行查询:", string(u_str))

	var us []Users
	err = MysqlDB.Select(&us, "select * from users where id in (?,?,?)", 1, 2, 3)
	if err != nil {
		beego.Error(err)
	}
	for k, v := range us {
		vs, err := json.Marshal(v)
		if err != nil {
			beego.Error(err)
		}
		beego.Notice("多行查询:", "[第", k+1, "行]", string(vs))
	}

	isql := "insert into users (name,age,city) values("
	isql += "'LiLei',"
	isql += IntToStr(20) + ","
	isql += "'NanJing')"

	beego.Notice("[执行sql]:", isql)
	_, err = MysqlDB.Exec(isql)
	if err != nil {
		beego.Error(err)
	}

	var users []Users
	err = MysqlDB.Select(&users, "select * from users")
	if err != nil {
		beego.Error(err)
	}
	for k, v := range users {
		vs, err := json.Marshal(v)
		if err != nil {
			beego.Error(err)
		}
		beego.Notice("多行查询:", "[第", k+1, "行]", string(vs))
	}
}

func IntToStr(x int) string {
	return strconv.Itoa(x)
}

```

#### 执行输出
```shell
2017/11/16 20:59:30 [I] Connect Mysql Server(10.0.0.252:3306) to successful!
2017/11/16 20:59:30 [N] 单行查询: {"Id":1,"Name":"tom","Age":19,"City":"beijing"}
2017/11/16 20:59:30 [N] 多行查询: [第 1 行] {"Id":1,"Name":"tom","Age":19,"City":"beijing"}
2017/11/16 20:59:30 [N] 多行查询: [第 2 行] {"Id":2,"Name":"bill","Age":18,"City":"shanghai"}
2017/11/16 20:59:30 [N] [执行sql]: insert into users (name,age,city) values('LiLei',20,'NanJing')
2017/11/16 20:59:30 [N] 多行查询: [第 1 行] {"Id":1,"Name":"tom","Age":19,"City":"beijing"}
2017/11/16 20:59:30 [N] 多行查询: [第 2 行] {"Id":2,"Name":"bill","Age":18,"City":"shanghai"}
2017/11/16 20:59:30 [N] 多行查询: [第 3 行] {"Id":3,"Name":"LiLei","Age":20,"City":"NanJing"}
```