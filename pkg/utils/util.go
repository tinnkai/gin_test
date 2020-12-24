package utils

import "gin_test/pkg/setting"

// Setup Initialize the util
func Setup() {
	// jwt secret
	jwtSecret = []byte(setting.AppSetting.JwtSecret)
}
