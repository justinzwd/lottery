package conf

import "time"

const SysTimeform = "2006-01-12 15:04:05"
const SysTimeformShort = "2006-01-12"

//时区的定义
var SysTimeLocation, _ = time.LoadLocation("Asia/Chongqing")

//签名密钥
var SignSecret = []byte("0123456789abcdef")

//cookie密钥
var CookieSecret = "hellolottery"



