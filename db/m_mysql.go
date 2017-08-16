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

func OpenMysql() (DBmysql *sql.DB) {
	//读取MySQL配置
	var mysql_user = beego.AppConfig.String("mysql_user")
	var mysql_password = beego.AppConfig.String("mysql_password")
	var mysql_net = beego.AppConfig.String("mysql_net")
	var mysql_host = beego.AppConfig.String("mysql_host")
	var mysql_port = beego.AppConfig.String("mysql_port")
	var mysql_db = beego.AppConfig.String("mysql_db")
	var mysql_charset = beego.AppConfig.String("mysql_charset")
	var mysql_max_life_time = beego.AppConfig.DefaultInt("mysql_max_life_time", 300)
	var mysql_max_open_conns = beego.AppConfig.DefaultInt("mysql_max_open_conns", 1000)
	var mysql_max_idle_conns = beego.AppConfig.DefaultInt("mysql_max_idle_conns", 20)

	//拼接成MySQL连接串
	var mysql_source string
	mysql_source = mysql_user + ":" + mysql_password + "@" + mysql_net + "(" + mysql_host + ":" + mysql_port + ")"
	mysql_source += "/" + mysql_db + "?" + "charset=" + mysql_charset

	var err error
	DBmysql, err = sql.Open("mysql", mysql_source)
	if err != nil {
		beego.Critical("Connect to Mysql, Error: " + err.Error())
		panic("Connect to Mysql, Error: " + err.Error())
	}

	DBmysql.SetConnMaxLifetime(time.Duration(mysql_max_life_time) * time.Second)
	DBmysql.SetMaxOpenConns(mysql_max_open_conns)
	DBmysql.SetMaxIdleConns(mysql_max_idle_conns)

	if err := DBmysql.Ping(); err != nil {
		beego.Critical("Try to ping mysql, Error: " + err.Error())
		panic("Try to ping mysql, Error: " + err.Error())
	} else {
		beego.Info("Connected to mysql successful!")
	}

	return DBmysql
}

func CloseMysql(DBmysql *sql.DB) {
	DBmysql.Close()
	beego.Info("[db closed] mysql")
}

//数据库查询
func DoQuery(sql string, DBmysql *sql.DB) (results map[int]map[int]string, err error) {
	beego.Trace("[sql]: ", sql)
	rows, err := DBmysql.Query(sql)
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
func DoExec(sql string, DBmysql *sql.DB) (bool, error) {
	beego.Trace("[sql]: ", sql)
	_, err := DBmysql.Exec(sql)
	if err != nil {
		beego.Error(err.Error())
		return false, err
	}
	return true, nil
}

// DoExecBatch 开启事务，执行批处理
func DoExecBatch(sqls []string, DBmysql *sql.DB) (bool, error) {
	tx, errBegin := DBmysql.Begin()
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
