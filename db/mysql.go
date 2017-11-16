package db

import (
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"time"
)

/*
配置文件
# mysql配置
mysql_user = "root"
mysql_password = "111111"
mysql_net = "tcp"
mysql_host = "10.0.0.252"
mysql_port = "3306"
mysql_db = "test"
mysql_charset = "utf8"
mysql_max_life_time = 300
mysql_max_open_conns = 1000
mysql_max_idle_conns = 20
*/

//mysql对象
type mysqlType struct {
	*sqlx.DB
}

// mysql对外接口
var MysqlDB *mysqlType

// 这个函数，需要放到初始化函数中调用！
func initMysql() {
	//读取MySQL配置
	var mysqlUser = beego.AppConfig.String("mysql_user")
	var mysqlPassword = beego.AppConfig.String("mysql_password")
	var mysqlNet = beego.AppConfig.String("mysql_net")
	var mysqlHost = beego.AppConfig.String("mysql_host")
	var mysqlPort = beego.AppConfig.String("mysql_port")
	var mysqlDb = beego.AppConfig.String("mysql_db")
	var mysqlCharset = beego.AppConfig.String("mysql_charset")
	var mysqlMaxLifeTime = beego.AppConfig.DefaultInt("mysql_max_life_time", 300)
	var mysqlMaxOpenConns = beego.AppConfig.DefaultInt("mysql_max_open_conns", 1000)
	var mysqlMaxIdleConns = beego.AppConfig.DefaultInt("mysql_max_idle_conns", 20)

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

//关闭MySQL连接
func (mt *mysqlType) CloseMysql() {
	mt.DB.Close()
	beego.Debug("[db closed] mysql")
}

// 查询表有多少行数据
func (mt *mysqlType) TableCount(tableName string) (int64, error) {

	// 组合查询sql
	sql := "select count(1) from " + tableName
	beego.Debug("[sql]: " + sql)

	var num int64
	err := mt.Get(&num, sql)
	if err != nil {
		beego.Error(err)
	}
	beego.Debug("[out]: ", num)
	return num, nil
}
