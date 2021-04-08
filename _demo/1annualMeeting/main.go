/**
curl http://localhost:8080/
curl http://localhost:8080/lucky
curl --data "users=zwd,zwd2" http://localhost:8080/import
 */
package main

import (
	"fmt"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"math/rand"
	"strings"
	"sync"
	"time"
)

var userList []string

//互斥锁
var mu sync.Mutex

//抽奖的控制器
type lotteryController struct {
	Ctx iris.Context
}

func newApp() *iris.Application {
	app := iris.New()
	mvc.New(app.Party("/")).Handle(&lotteryController{})
	return app
}

func main() {
	app := newApp()
	//userList = make([]string, 0)
	userList = []string{}
	mu = sync.Mutex{}

	app.Run(iris.Addr(":8080"))
}

// GET http://localhost:8080/
func (c *lotteryController) Get() string {
	count := len(userList)
	return fmt.Sprintf("当前总共参与抽奖的用户数：%d\n", count)
}

// POST http://localhost:8080/import
// params: users
//PostImport
func (c *lotteryController) PostImport() string {
	strUsers := c.Ctx.FormValue("users")
	users := strings.Split(strUsers, ",")

	mu.Lock()
	defer mu.Unlock()

	count1 := len(userList)

	for _,u := range users {
		u = strings.TrimSpace(u)
		if len(u) > 0 {
			userList = append(userList, u)
		}
	}

	count2 := len(userList)
	return fmt.Sprintf("当前总共参与抽奖的用户数：%d，成功导入的用户数：%d\n", count2, (count2-count1))
}

// GET http://localhost:8080/lucky
//GetLucky
func (c *lotteryController) GetLucky() string {
	mu.Lock()
	defer mu.Unlock()
	count := len(userList)
	if count > 1 {
		seed := time.Now().UnixNano()
		index := rand.New(rand.NewSource(seed)).Int31n(int32(count))
		user := userList[index]
		userList = append(userList[0:index], userList[index+1:]...)
		return fmt.Sprintf("当前中奖用户：%s, 剩余用户数：%d\n", user, count-1)
	} else if count == 1{
		user := userList[0]
		return fmt.Sprintf("当前中奖用户：%s, 剩余用户数：%d\n", user, count-1)
	} else {
		return fmt.Sprintf("没有任何参与用户，请先导入用户\n")
	}
}