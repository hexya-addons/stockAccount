package stock_account

import (
	"github.com/hexya-erp/hexya/src/models"
	"github.com/hexya-erp/hexya/src/models/types/dates"
	"github.com/hexya-erp/pool/h"
)

func init() {
	h.WizardValuationHistory().DeclareModel()

	h.WizardValuationHistory().AddFields(map[string]models.FieldDefinition{
		"ChooseDate": models.BooleanField{
			String: "Inventory at Date",
		},
		"Date": models.DateTimeField{
			String:   "Date",
			Default:  func(env models.Environment) interface{} { return dates.Now() },
			Required: true,
		},
	})
	h.WizardValuationHistory().Methods().OpenTable().DeclareMethod(
		`OpenTable`,
		func(rs m.WizardValuationHistorySet) {
			//        self.ensure_one()
			//        ctx = dict(
			//            self._context,
			//            history_date=self.date,
			//            search_default_group_by_product=True,
			//            search_default_group_by_location=True)
			//        action = self.env['ir.model.data'].xmlid_to_object(
			//            'stock_account.action_stock_history')
			//        if not action:
			//            action = {
			//                'view_type': 'form',
			//                'view_mode': 'tree,graph,pivot',
			//                'res_model': 'stock.history',
			//                'type': 'ir.actions.act_window',
			//            }
			//        else:
			//            action = action[0].read()[0]
			//        action['domain'] = "[('date', '<=', '" + self.date + "')]"
			//        action['name'] = _('Stock Value At Date')
			//        action['context'] = ctx
			//        return action
		})
}
