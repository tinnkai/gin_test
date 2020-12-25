package routers

import (
	"github.com/xxjwxc/ginrpc"
)

func init() {
	ginrpc.SetVersion(1608914352)
	ginrpc.AddGenOne("HongdongController.Detail", "/hongdong/detail", []string{"get"})
	ginrpc.AddGenOne("HongdongController.List", "/hongdong/list", []string{"get"})
}
