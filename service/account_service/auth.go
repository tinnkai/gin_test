package account_service

import (
	"gin_test/models/mysql_models"
	"gin_test/pkg/app"
	"gin_test/pkg/errors"
)

// 验证登录
func CheckLogin(userId int64) (mysql_models.User, error) {
	user, err := mysql_models.UserRepository.GetUserInfoById(userId, "id,username,password,phone,`group`")
	if err != nil {
		return user, err
	}

	// 用户id是否一致
	if user.Id != userId {
		return user, errors.Newf(app.ERROR_LOGIN_FAIL, "", "")
	}

	return user, nil
}
