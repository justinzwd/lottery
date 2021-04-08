package conf

//驱动名称
const DriverName = "mysql"

type DbConfig struct {
	Host string
	Port int
	User string
	Pwd string
	Database string
	IsRunning bool
}

//配置不一定只有一个
var DbMasterList = []DbConfig{
	{
		Host:      "127.0.0.1",
		Port:      3306,
		User:      "root",
		Pwd:       "000000",
		Database: "lottery",
		IsRunning: true,
	},
}

//定义一个可以直接用的配置
var DbMaster DbConfig = DbMasterList[0]