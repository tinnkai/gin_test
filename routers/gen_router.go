package routers

import (
	"github.com/xxjwxc/ginrpc"
)

func init() {
	ginrpc.SetVersion(1611072059)
	ginrpc.AddGenOne("HongdongController.BirthdayPackageInfo", "/hongdong/birthdayPackageInfo", []string{"get"})
	ginrpc.AddGenOne("HongdongController.Detail", "/hongdong/detail", []string{"get"})
	ginrpc.AddGenOne("HongdongController.List", "/hongdong/list", []string{"get"})
}
