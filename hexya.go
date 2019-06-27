package stock_account

import (
	"github.com/hexya-erp/hexya/src/server"
)

const MODULE_NAME string = "stock_account"

func init() {
	server.RegisterModule(&server.Module{
		Name:     MODULE_NAME,
		PreInit:  func() {},
		PostInit: func() {},
	})

}
