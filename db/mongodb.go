package db

import (
	"github.com/astaxie/beego"
	"gopkg.in/mgo.v2"
)

/*
# mongodb配置
mongo.host = 10.0.0.253:27017
*/

//mongo对外接口
var MongoDB *mgo.Session

func init() {
	initMongoDB()
}

func initMongoDB() {
	mgo_url := beego.AppConfig.String("mongo.host")
	var err error

	//通过赋值对外的接口，确保了接口正常使用。。。
	MongoDB, err = mgo.Dial(mgo_url)
	if err != nil {
		panic("MongoDB Server:" + mgo_url + " Connect failed: " + err.Error() + "!")
	} else {
		beego.Info("Connect MongoDB Server(" + mgo_url + ") to successful!")
	}

	MongoDB.SetMode(mgo.Monotonic, true)
}
