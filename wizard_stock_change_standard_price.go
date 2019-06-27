package stock_account

import (
	"github.com/hexya-erp/hexya/src/models"
	"github.com/hexya-erp/pool/h"
	"github.com/hexya-erp/pool/q"
)

func init() {
	h.StockChangeStandardPrice().DeclareModel()

	h.StockChangeStandardPrice().AddFields(map[string]models.FieldDefinition{
		"NewPrice": models.FloatField{
			String: "Price",
			//digits=dp.get_precision('Product Price')
			Required: true,
			Help: "If cost price is increased, stock variation account will" +
				"be debited and stock output account will be credited with" +
				"the value = (difference of amount * quantity available)." +
				"If cost price is decreased, stock variation account will" +
				"be creadited and stock input account will be debited.",
		},
		"CounterpartAccountId": models.Many2OneField{
			RelationModel: h.AccountAccount(),
			String:        "Counter-Part Account",
			Filter:        q.Deprecated().Equals(False),
			Required:      true,
		},
	})
	h.StockChangeStandardPrice().Methods().DefaultGet().Extend(
		`DefaultGet`,
		func(rs m.StockChangeStandardPriceSet, fields interface{}) {
			//        res = super(StockChangeStandardPrice, self).default_get(fields)
			//        product_or_template = self.env[self._context['active_model']].browse(
			//            self._context['active_id'])
			//        if 'new_price' in fields and 'new_price' not in res:
			//            res['new_price'] = product_or_template.standard_price
			//        if 'counterpart_account_id' in fields and 'counterpart_account_id' not in res:
			//            res['counterpart_account_id'] = product_or_template.property_account_expense_id.id or product_or_template.categ_id.property_account_expense_categ_id.id
			//        return res
		})
	h.StockChangeStandardPrice().Methods().ChangePrice().DeclareMethod(
		` Changes the Standard Price of Product and creates an account
move accordingly. `,
		func(rs m.StockChangeStandardPriceSet) {
			//        self.ensure_one()
			//        if self._context['active_model'] == 'product.template':
			//            products = self.env['product.template'].browse(
			//                self._context['active_id']).product_variant_ids
			//        else:
			//            products = self.env['product.product'].browse(
			//                self._context['active_id'])
			//        products.do_change_standard_price(
			//            self.new_price, self.counterpart_account_id.id)
			//        return {'type': 'ir.actions.act_window_close'}
		})
}
