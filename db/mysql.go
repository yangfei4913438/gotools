package db

import (
	"database/sql"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
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

func init() {
	initMysql()
}

type mysqlType struct {
	DB *sql.DB
}

// mysql对外接口
var MysqlDB *mysqlType

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
	openMysql, err := sql.Open("mysql", mysqlSource)
	if err != nil {
		beego.Critical("Connect to Mysql, Error: " + err.Error())
		panic("Connect to Mysql, Error: " + err.Error())
	}

	MysqlDB = &mysqlType{openMysql}

	//SetConnMaxLifetime连接的最大空闲时间(可选)
	MysqlDB.DB.SetConnMaxLifetime(time.Duration(mysqlMaxLifeTime) * time.Second)
	//SetMaxOpenConns用于设置最大打开的连接数，默认值为0表示不限制。
	MysqlDB.DB.SetMaxOpenConns(mysqlMaxOpenConns)
	//SetMaxIdleConns用于设置闲置的连接数。
	MysqlDB.DB.SetMaxIdleConns(mysqlMaxIdleConns)

	if err := MysqlDB.DB.Ping(); err != nil {
		beego.Critical("Try to ping mysql, Error: " + err.Error())
		panic("Try to ping mysql, Error: " + err.Error())
	} else {
		beego.Info("Connected to mysql successful!")
	}

}

//关闭MySQL连接
func (mt *mysqlType) CloseMysql() {
	mt.DB.Close()
	beego.Info("[db closed] mysql")
}

//数据库查询
func (mt *mysqlType) DoQuery(sql string) (results map[int]map[int]string, err error) {
	beego.Trace("[sql]: ", sql)
	rows, err := mt.DB.Query(sql)
	if err != nil {
		beego.Error(err.Error())
		return nil, err
	}
	cols, _ := rows.Columns()
	values := make([][]byte, len(cols))
	scans := make([]interface{}, len(cols))
	for i := range values {
		scans[i] = &values[i]
	}
	results = make(map[int]map[int]string)
	i := 0
	for rows.Next() {
		if err = rows.Scan(scans...); err != nil {
			beego.Error(err.Error())
			return nil, err
		}
		row := make(map[int]string) //每行数据
		for k, v := range values {  //每行数据是放在values里面，现在把它挪到row里
			row[k] = string(v)
		}
		results[i] = row //装入结果集中
		i++
	}
	rows.Close()
	return results, nil
}

//单一sql执行
func (mt *mysqlType) DoExec(sql string) (bool, error) {
	beego.Trace("[sql]: ", sql)
	_, err := mt.DB.Exec(sql)
	if err != nil {
		beego.Error(err.Error())
		return false, err
	}
	return true, nil
}

// DoExecBatch 开启事务，执行批处理
func (mt *mysqlType) DoExecBatch(sqls []string) (bool, error) {
	tx, errBegin := mt.DB.Begin()
	if errBegin != nil {
		beego.Error(errBegin.Error())
		return false, errBegin
	}

	var errExec error
	for _, sql_txt := range sqls {
		beego.Trace("[sql]: ", sql_txt)
		_, errExec = tx.Exec(sql_txt)
		if errExec != nil {
			if errRollback := tx.Rollback(); errRollback != nil {
				beego.Error(errRollback.Error())
				return false, errRollback
			}
			beego.Error(errExec.Error())
			return false, errExec
		}
	}

	if errCommit := tx.Commit(); errCommit != nil {
		if errRollback := tx.Rollback(); errRollback != nil {
			beego.Error(errRollback.Error())
			return false, errRollback
		}
		beego.Error(errCommit.Error())
		return false, errCommit
	}

	return true, nil
}
