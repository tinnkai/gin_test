package account_service

import (
	"gin_test/models/mysql_models"
	"gin_test/pkg/app"
	"gin_test/pkg/errors"
)

// 验证登录
func CheckLogin(phone int64, password string) (mysql_models.User, error) {
	user, err := mysql_models.UserRepository.GetUserInfoByPhone(phone, "id,username,password,phone,`group`")
	if err != nil {
		return user, err
	}

	// 密码是否正确
	if user.Password != password {
		return user, errors.Newf(app.ERROR_LOGIN_FAIL, "", "")
	}

	return user, nil
}
