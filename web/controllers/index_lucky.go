package controllers

import (
	"lottery/comm"
)

// localhost:8080/lucky
// 实现抽奖接口
func (c *IndexController) GetLucky() map[string]interface{} {
	rs := make(map[string]interface{})
	rs["code"] = 0
	rs["msg"] = ""
	// 1 验证登录用户
	loginuser := comm.GetLoginUser(c.Ctx.Request())
	if loginuser == nil || loginuser.Uid < 1 {
		// 用户为空
		rs["code"] = 101
		rs["msg"] = "请先登录，再来抽奖"
		return rs
	}
	ip := comm.ClientIP(c.Ctx.Request())
	api := &LuckyApi{}
	code, msg, gift := api.luckyDo(loginuser.Uid, loginuser.Username, ip)
	rs["code"] = code
	rs["msg"] = msg
	rs["gift"] = gift
	return rs
}
