/**
 *
 * 抽奖中用到的锁
 */
package utils

import (
	"fmt"

	"lottery/datasource"
)

// 加锁，抽奖的时候需要用到的锁，避免一个用户并发多次抽奖
func LockLucky(uid int) bool {
	return lockLuckyServ(uid)
}

// 解锁，抽奖的时候需要用到的锁，避免一个用户并发多次抽奖
func UnlockLucky(uid int) bool {
	return unlockLuckyServ(uid)
}

func getLuckyLockKey(uid int) string {
	return fmt.Sprintf("lucky_lock_%d", uid)
}

func lockLuckyServ(uid int) bool {
	key := getLuckyLockKey(uid)
	cacheObj := datasource.InstanceCache()
	// set 命令，往key里面写入数据1，ex代表是否存在，3秒自动过期
	// 如果一直不过期的话，就容易造成思索
	// 如果时间设置太短，有些操作还没完成就释放锁了，也不可以
	rs, _ := cacheObj.Do("SET", key, 1, "EX", 3, "NX")
	if rs == "OK" {
		return true
	} else {
		return false
	}
}

func unlockLuckyServ(uid int) bool {
	key := getLuckyLockKey(uid)
	cacheObj := datasource.InstanceCache()
	//删除key
	rs, _ := cacheObj.Do("DEL", key)
	if rs == "OK" {
		return true
	} else {
		return false
	}
}
