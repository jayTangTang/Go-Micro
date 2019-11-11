/*
@Time : 19-11-10 下午4:11 
@Author : itcast
@File : model
@Software: GoLand
*/

package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"fmt"
)

type Stu struct {
	gorm.Model
	Name     string
	PassWord string
}

var GlobalDB *gorm.DB

func InitModel() {

	//打开数据库   驱动名 连接字符串   数据库连接池
	db, err := gorm.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/ihome?parseTime=true")
	if err != nil {
		fmt.Println("连接数据库失败")
		return
	}

	//连接池设置
	db.DB().SetMaxIdleConns(20)
	db.DB().SetMaxOpenConns(30)
	db.DB().SetConnMaxLifetime(60 * 30)

	//设置表名为单数形式
	db.SingularTable(true)
	GlobalDB = db

	//自动迁移  在gorm中建表默认是负数形式
	db.AutoMigrate(new(Stu))
}

//插入数据函数
func InsertData() {
	//有一个赋值完的结构体对象
	var stu Stu
	stu.Name = "timor"
	stu.PassWord = "123456"

	if err := GlobalDB.Create(&stu).Error; err != nil {
		fmt.Println("创建数据失败")
		return
	}
	fmt.Println(stu)
}

func SearchData() {
	var stu Stu
	//查所有  取第一条   select * from user where name = bj5q and passWord=123
	if err := GlobalDB.Where("name = ?", "timor").Where("pass_word = ?", "123456").First(&stu).Error; err != nil {
		fmt.Println("查询错误", err)
		return
	}
	fmt.Println(stu)
}

//更新数据
func UpdateData() {
	var stu Stu
	stu.Name = "wushuang"
	stu.PassWord = "111222"
	/*
	//按照条件来更新
	if err := GlobalDB.Save(&stu).Error ; err != nil {
		fmt.Println("更新数据失败",err)
		return
	}

	if err := GlobalDB.Model(&stu).Where("neme = ?","wushuang").Update("pass_word","00000").Error;err != nil {
		fmt.Println("更新密码失败",err)
		return
	}*/

	GlobalDB.Model(&stu).Where("id = 1").Update(map[string]interface{}{"name": "ttt", "pass_word": "333"})
	fmt.Println(stu)
}

//删除数据  软删除/逻辑删除  数据是无价的
func DeleteData() {
	var stu Stu

	if err := GlobalDB.Where("id = 1").Unscoped().Delete(&stu).Error; err != nil {
		fmt.Println("删除失败", err)
		return
	}
}
