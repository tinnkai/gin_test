package routers

import (
	"github.com/xxjwxc/ginrpc"
)

func init() {
	ginrpc.SetVersion(1608475288)
	ginrpc.AddGenOne("HongdongController.Detail", "/hongdong/detail", []string{"post", "get"})
	ginrpc.AddGenOne("HongdongController.List", "/hongdong/list", []string{"post", "get"})
}
