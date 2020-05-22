package models

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	// 用户uuid
	Uuid string `orm:"cloumn(id);pk"`
	// 用户手机号
	Phone string
	// 用户openid（微信小程序）
	Openid string
	// 用户昵称
	NikeName string
	// 用户状态
	Status int
	// 用户备注
	Remark string
	// 用户创建时间
	CreateTime string

	// 用户修改时间
	ModifyTime string

	// 用户是否被软删除
	Deleted bool
}

func init() {

	// 需要在init中注册定义的model
	err := orm.RegisterDriver("mysql", orm.DRMySQL)
	if err != nil {
		fmt.Println("数据库驱动加载失败")
	}

	err = orm.RegisterDataBase(
		"default",
		"mysql",
		beego.AppConfig.String("databasesource"))
	if err != nil {
		fmt.Println("数据库注册失败")
	}

	orm.RegisterModel(new(User))
}

func (t *User) GET() User {

	o := orm.NewOrm()
	user := User{Uuid: "000646a4-59f1-4d01-be13-44f267f1794f"}
	err := o.Read(&user)
	if err == orm.ErrNoRows {
		fmt.Println("查询不到")
	} else if err == orm.ErrMissPK {
		fmt.Println("找不到主键")
	}
	return user
}
