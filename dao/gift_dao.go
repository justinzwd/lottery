package dao

import (
	"github.com/go-xorm/xorm"
	"log"
	"lottery/models"
)

//相当于类的结构
type GiftDao struct {
	engine *xorm.Engine
}

//提供一个公共的方法
//实例化，获取到一个GiftDao
//数据库相关的操作对象
func NewGiftDao(engine *xorm.Engine) *GiftDao {
	return &GiftDao{engine:engine}
}

//根据主键id返回数据模型
func (d *GiftDao) Get(id int) *models.LtGift {
	//Id:id 把参数传进去
	data := &models.LtGift{Id:id}
	//这里data是一个指针
	//下面这个查询语句，如果查到数据，会将数据写到data里面
	ok,err := d.engine.Get(data)

	if ok && err == nil {
		return data
	} else {
		//如果没有查到数据，需要将数据清空再返回
		//否则返回的结果中id不为空，而其他数据都为空，会给调用者造成异常
		//或者返回 nil 也可以
		data.Id = 0
	}
	return data
}

func (d *GiftDao) GetAll() []models.LtGift {

	datalist := make([]models.LtGift,0)
	err := d.engine.  //做一个排序，这些colNames都是数据库的字段
		Asc("sys_status").
		Asc("displayorder").
		Find(&datalist)
	if err != nil {
		log.Println("gift_dao.GetAll error=",err)
		return datalist
	}
	return datalist
}

func (d *GiftDao) CountAll() int64 {
	num,err := d.engine.
		Count(&models.LtGift{})
	if err != nil {
		return 0
	} else {
		return num
	}
}

func (d *GiftDao) Delete(id int) error {
	data := &models.LtGift{
		Id:        id,
		//不做物理删除，这里其实用的是更新操作，这里的SysStatus为1代表删除状态
		SysStatus: 1,
	}
	_,err := d.engine.Id(data.Id).
		Update(data)
	return err
}

//这些 columns是强制要求更新的
func (d *GiftDao) Update(data *models.LtGift, columns []string) error {
	//如果传进来的对象对应的 columns 是空的，那么去更新的时候会报错
	//这里MustCols就是强制更新到数据库
	_, err := d.engine.Id(data.Id).
		MustCols(columns...).Update(data)
	return err
}

func (d *GiftDao) Create(data *models.LtGift) error {
	_,err := d.engine.Insert(data)
	return err
}