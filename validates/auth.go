package validates

import (
	"regexp"
	"strconv"

	"github.com/gin-gonic/gin"
	validator "github.com/go-playground/validator/v10"
)

// 验证
type Auth struct {
	Phone    int64  `form:"phone" json:"phone" validate:"required,numeric,checkMobile"`
	Password string `form:"password" json:"password" validate:"required"`
}

// 绑定错误校验
func (t *Auth) Bind(c *gin.Context) []string {
	// 绑定
	err := c.ShouldBind(t)
	if err != nil {
		return nil
	}
	// 自定义校验
	validate.RegisterValidation("checkMobile", checkMobile)

	if err := validate.Struct(t); err != nil {
		errors := err.(validator.ValidationErrors)
		sliceErrs := []string{}
		for _, e := range errors {
			sliceErrs = append(sliceErrs, e.Translate(validateTrans))
		}
		return sliceErrs
	}

	return nil
}

func checkMobile(fl validator.FieldLevel) bool {
	// 强制转化成int64
	value := fl.Field().Interface().(int64)
	// 转化成string
	phone := strconv.FormatInt(value, 10)

	reg := `^1([39][0-9])\d{8}$`
	rgx := regexp.MustCompile(reg)
	return rgx.MatchString(phone)
}
