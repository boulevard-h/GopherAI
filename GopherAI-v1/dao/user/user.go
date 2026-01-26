package user

import (
	"GopherAI/common/mysql"
	"GopherAI/model"
	"GopherAI/utils"
	"context"

	"gorm.io/gorm"
)

const (
	CodeMsg     = "GopherAI验证码如下(验证码仅限于2分钟有效): "
	UserNameMsg = "GopherAI的账号如下，请保留好，后续可以用账号/邮箱进行登录 "
)

var ctx = context.Background()

// 登录标识支持 username / email
func IsExistUser(username string) (bool, *model.User) {
	// 1) 先按 username 查
	u, err := mysql.GetUserByUsername(username)
	if err == nil && u != nil {
		return true, u
	}
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, nil
	}

	// 2) username 不存在时，尝试按 email 查（支持邮箱号登录）
	u, err = mysql.GetUserByEmail(username)
	if err == nil && u != nil {
		return true, u
	}
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, nil
	}

	return false, nil
}

func Register(username, email, password string) (*model.User, bool) {
	if user, err := mysql.InsertUser(&model.User{
		Email:    email,
		Name:     username,
		Username: username,
		Password: utils.MD5(password),
	}); err != nil {
		return nil, false
	} else {
		return user, true
	}
}
