package stock_account

import (
	"github.com/hexya-erp/hexya/src/models"
	"github.com/hexya-erp/hexya/src/models/types"
	"github.com/hexya-erp/pool/h"
)

func init() {
	h.StockConfigSettings().DeclareModel()

	h.StockConfigSettings().AddFields(map[string]models.FieldDefinition{
		"GroupStockInventoryValuation": models.SelectionField{
			Selection: types.Selection{
				"": "Periodic inventory valuation (recommended)",
				"": "Perpetual inventory valuation (stock move generates accounting entries)",
			},
			String: "Inventory Valuation",
			//implied_group='stock_account.group_inventory_valuation'
			Help: "Allows to configure inventory valuations on products and" +
				"product categories.",
		},
		"ModuleStockLandedCosts": models.SelectionField{
			Selection: types.Selection{
				"": "No landed costs",
				"": "Include landed costs in product costing computation",
			},
			String: "Landed Costs",
			Help: "Install the module that allows to affect landed costs on" +
				"pickings, and split them onto the different products.",
		},
	})
	h.StockConfigSettings().Methods().OnchangeLandedCosts().DeclareMethod(
		`OnchangeLandedCosts`,
		func(rs m.StockConfigSettingsSet) {
			//        if self.module_stock_landed_costs:
			//            self.group_stock_inventory_valuation = 1
		})
}
