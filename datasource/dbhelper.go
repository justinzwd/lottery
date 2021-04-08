package datasource

import (
	"fmt"
	"github.com/go-xorm/xorm"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"lottery/conf"
	"sync"
)

//互斥锁
var dbLock sync.Mutex
var masterInstance *xorm.Engine

//单例模式
func InstanceDbMaster() *xorm.Engine {
	if masterInstance != nil {
		return masterInstance
	}
	dbLock.Lock()

	defer dbLock.Unlock()

	if masterInstance != nil {
		return masterInstance
	}
	return NewDbMaster()
}

//获取一个数据库的操作引擎
func NewDbMaster() *xorm.Engine {
	sourceName := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8",
		conf.DbMaster.User,
		conf.DbMaster.Pwd,
		conf.DbMaster.Host,
		conf.DbMaster.Port,
		conf.DbMaster.Database)

	//两个参数
	//1.驱动类型 conf.DriverName，2.上面一长串的配置信息 sourceName
	instance,err := xorm.NewEngine(conf.DriverName, sourceName)
	if err != nil {
		log.Fatal("dbhelper.NewDbMaster NewEngine error", err)
		return nil
	}

	//用来调试用的，展示SQL以及执行时间
	instance.ShowSQL(true)

	masterInstance = instance
	return instance
}