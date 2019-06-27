package stock_account

import (
	"github.com/hexya-erp/pool/h"
)

//vars

var (
	//
	Stock.GroupStockManager *security.Group
	//Manage Inventory Valuation and Costing Methods
	GroupInventoryValuation *security.Group
)


//rights
func init() {
	h.Account.ModelAccountInvoice().Methods().Load().AllowGroup(GroupStockUser)
	h.Account.ModelAccountInvoice().Methods().Write().AllowGroup(GroupStockUser)
	h.Account.ModelAccountInvoice().Methods().Create().AllowGroup(GroupStockUser)
	h.Account.ModelAccountInvoiceLine().Methods().Load().AllowGroup(GroupStockUser)
	h.Account.ModelAccountInvoiceLine().Methods().Write().AllowGroup(GroupStockUser)
	h.Account.ModelAccountInvoiceLine().Methods().Create().AllowGroup(GroupStockUser)
	h.Account.ModelAccountInvoiceTax().Methods().Load().AllowGroup(GroupStockUser)
	h.Account.ModelAccountInvoiceTax().Methods().Write().AllowGroup(GroupStockUser)
	h.Account.ModelAccountInvoiceTax().Methods().Create().AllowGroup(GroupStockUser)
	h.Account.ModelAccountJournal().Methods().Load().AllowGroup(GroupStockUser)
	h.Account.ModelAccountAccount().Methods().Load().AllowGroup(GroupStockManager)
	h.StockHistory().Methods().Load().AllowGroup(GroupStockManager)
	h.Stock.ModelStockPicking().Methods().Load().AllowGroup(GroupAccountInvoice)
	h.Stock.ModelStockPicking().Methods().Write().AllowGroup(GroupAccountInvoice)
	h.Stock.ModelStockPicking().Methods().Create().AllowGroup(GroupAccountInvoice)
	h.StockMove().Methods().Load().AllowGroup(GroupAccountInvoice)
	h.StockMove().Methods().Write().AllowGroup(GroupAccountInvoice)
	h.StockMove().Methods().Create().AllowGroup(GroupAccountInvoice)
}
