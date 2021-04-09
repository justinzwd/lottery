package models

// 站点中与浏览器交互的用户模型
type ObjLoginuser struct {
	Uid      int
	Username string
	Now      int
	Ip       string
	//签名
	Sign     string
}
