/*
@Time : 19-11-14 下午4:52 
@Author : itcast
@File : modelFunc
@Software: GoLand
*/

package model

//获取用户信息
func GetUserInfo(userName string) (User, error) {
	//连接数据库
	var user User
	err := GlobalDB.Where("name = ?", userName).Find(&user).Error
	return user, err
}

//更新用户名
func UpdateUserName(oldName, newName string) error {
	//更新  链式调用
	return GlobalDB.Model(new(User)).
		Where("name = ?", oldName).
			Update("name", newName).Error
}

//存储用户头像   更新
func SaveUserAvatar(userName, avataeUrl string) error {
	return GlobalDB.Model(new(User)).
		Where("name = ?", userName).
			Update("avatar_url", avataeUrl).Error
}

//存储用户真实姓名
func SaveRealName(userName, realName, idcard string) error {
	return GlobalDB.Model(new(User)).
		Where("name = ?", userName).
		Update(map[string]interface{}{"real_name": realName, "id_card": idcard}).Error
}
