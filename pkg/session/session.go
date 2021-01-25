package session

import (
	"encoding/json"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// get session
func Get(ctx *gin.Context, key string) interface{} {

	// 获取
	session := sessions.Default(ctx)
	value := session.Get(key)
	return value
}

// set session
func Set(ctx *gin.Context, key string, value interface{}) error {

	// 设置
	session := sessions.Default(ctx)
	value, err := json.Marshal(value)
	if err != nil {
		return err
	}
	session.Set(key, value)
	session.Save()

	return nil
}

// delete session
func Delete(ctx *gin.Context, key string) {

	//清除该用户登录状态的数据
	session := sessions.Default(ctx)
	session.Delete(key)
	session.Save()
}
