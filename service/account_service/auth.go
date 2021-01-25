package account_service

import (
	"gin_test/models/mysql_models"
	"gin_test/pkg/app"
	"gin_test/pkg/errors"
)

// 验证登录
func CheckLogin(userId int64) (mysql_models.AuthUserInfo, error) {
	authUserInfo := mysql_models.AuthUserInfo{}
	userInfo, err := mysql_models.UserRepository.GetLoginUserInfoById(userId, "id,username,password,phone,`group`")
	if err != nil {
		return authUserInfo, err
	}

	// 用户id是否一致
	if userInfo.Id != userId || userInfo.Id < 1 {
		return authUserInfo, errors.Newf(app.ERROR_LOGIN_FAIL, "", "")
	}

	authUserInfo = mysql_models.AuthUserInfo{
		UserId:   userInfo.Id,
		Username: userInfo.Username,
		Phone:    userInfo.Phone,
	}

	return authUserInfo, nil
}
