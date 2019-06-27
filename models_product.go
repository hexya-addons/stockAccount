package stock_account

	import (
		"net/http"

		"github.com/hexya-erp/hexya/src/controllers"
		"github.com/hexya-erp/hexya/src/models"
		"github.com/hexya-erp/hexya/src/models/types"
		"github.com/hexya-erp/hexya/src/models/types/dates"
		"github.com/hexya-erp/pool/h"
		"github.com/hexya-erp/pool/q"
	)
	
func init() {
h.ProductTemplate().DeclareModel()


h.ProductTemplate().AddFields(map[string]models.FieldDefinition{
"PropertyValuation": models.SelectionField{
Selection: types.Selection{
"manual_periodic": "Periodic (manual)",
"real_time": "Perpetual (automated)",
},
String: "Inventory Valuation",
//company_dependent=True
NoCopy: false,
Default: models.DefaultValue("manual_periodic"),
Help: "If perpetual valuation is enabled for a product, the system" + 
"will automatically create journal entries corresponding" + 
"to stock moves, with product price as specified by the" + 
"'Costing Method'The inventory variation account set on" + 
"the product category will represent the current inventory" + 
"value, and the stock input and stock output account will" + 
"hold the counterpart moves for incoming and outgoing products.",
},
"Valuation": models.CharField{
Compute: h.ProductTemplate().Methods().ComputeValuationType(),
//inverse='_set_valuation_type'
},
"PropertyCostMethod": models.SelectionField{
Selection: types.Selection{
"standard": "Standard Price",
"average": "Average Price",
"real": "Real Price",
},
String: "Costing Method",
//company_dependent=True
NoCopy: false,
Help: "Standard Price: The cost price is manually updated at the" + 
"end of a specific period (usually once a year)." + 
"                Average Price: The cost price is recomputed" + 
"at each incoming shipment and used for the product valuation." + 
"                Real Price: The cost price displayed is" + 
"the price of the last outgoing product (will be use in" + 
"case of inventory loss for example).",
},
"CostMethod": models.CharField{
Compute: h.ProductTemplate().Methods().ComputeCostMethod(),
//inverse='_set_cost_method'
},
"PropertyStockAccountInput": models.Many2OneField{
RelationModel: h.AccountAccount(),
String: "Stock Input Account",
//company_dependent=True
Filter: q.Deprecated().Equals(False),
Help: "When doing real-time inventory valuation, counterpart journal" + 
"items for all incoming stock moves will be posted in this" + 
"account, unless there is a specific valuation account set" + 
"on the source location. When not set on the product, the" + 
"one from the product category is used.",
},
"PropertyStockAccountOutput": models.Many2OneField{
RelationModel: h.AccountAccount(),
String: "Stock Output Account",
//company_dependent=True
Filter: q.Deprecated().Equals(False),
Help: "When doing real-time inventory valuation, counterpart journal" + 
"items for all outgoing stock moves will be posted in this" + 
"account, unless there is a specific valuation account set" + 
"on the destination location. When not set on the product," + 
"the one from the product category is used.",
},
})
h.ProductTemplate().Methods().ComputeValuationType().DeclareMethod(
`ComputeValuationType`,
func(rs h.ProductTemplateSet) h.ProductTemplateData {
//        self.valuation = self.property_valuation or self.categ_id.property_valuation
})
h.ProductTemplate().Methods().SetValuationType().DeclareMethod(
`SetValuationType`,
func(rs m.ProductTemplateSet)  {
//        return self.write({'property_valuation': self.valuation})
})
h.ProductTemplate().Methods().ComputeCostMethod().DeclareMethod(
`ComputeCostMethod`,
func(rs h.ProductTemplateSet) h.ProductTemplateData {
//        self.cost_method = self.property_cost_method or self.categ_id.property_cost_method
})
h.ProductTemplate().Methods().SetCostMethod().DeclareMethod(
`SetCostMethod`,
func(rs m.ProductTemplateSet)  {
//        return self.write({'property_cost_method': self.cost_method})
})
h.ProductTemplate().Methods().OnchangeTypeValuation().DeclareMethod(
`OnchangeTypeValuation`,
func(rs m.ProductTemplateSet)  {
//        pass
})
h.ProductTemplate().Methods().GetProductAccounts().DeclareMethod(
` Add the stock accounts related to product to the result of super()
        @return: dictionary which contains information
regarding stock accounts and super (income+expense accounts)
        `,
func(rs m.ProductTemplateSet)  {
//        accounts = super(ProductTemplate, self)._get_product_accounts()
//        res = self._get_asset_accounts()
//        accounts.update({
//            'stock_input': res['stock_input'] or self.property_stock_account_input or self.categ_id.property_stock_account_input_categ_id,
//            'stock_output': res['stock_output'] or self.property_stock_account_output or self.categ_id.property_stock_account_output_categ_id,
//            'stock_valuation': self.categ_id.property_stock_valuation_account_id or False,
//        })
//        return accounts
})
h.ProductTemplate().Methods().GetProductAccounts().DeclareMethod(
` Add the stock journal related to product to the result of super()
        @return: dictionary which contains all needed information
regarding stock accounts and journal and super (income+expense accounts)
        `,
func(rs m.ProductTemplateSet, fiscal_pos interface{})  {
//        accounts = super(ProductTemplate, self).get_product_accounts(
//            fiscal_pos=fiscal_pos)
//        accounts.update(
//            {'stock_journal': self.categ_id.property_stock_journal or False})
//        return accounts
})
h.ProductProduct().DeclareModel()

h.ProductProduct().Methods().OnchangeTypeValuation().DeclareMethod(
`OnchangeTypeValuation`,
func(rs m.ProductProductSet)  {
//        pass
})
h.ProductProduct().Methods().DoChangeStandardPrice().DeclareMethod(
` Changes the Standard Price of Product and creates an account
move accordingly.`,
func(rs m.ProductProductSet, new_price interface{}, account_id interface{})  {
//        AccountMove = self.env['account.move']
//        quant_locs = self.env['stock.quant'].sudo().read_group(
//            [('product_id', 'in', self.ids)], ['location_id'], ['location_id'])
//        quant_loc_ids = [loc['location_id'][0] for loc in quant_locs]
//        locations = self.env['stock.location'].search([('usage', '=', 'internal'), (
//            'company_id', '=', self.env.user.company_id.id), ('id', 'in', quant_loc_ids)])
//        product_accounts = {
//            product.id: product.product_tmpl_id.get_product_accounts() for product in self}
//        for location in locations:
//            for product in self.with_context(location=location.id, compute_child=False).filtered(lambda r: r.valuation == 'real_time'):
//                diff = product.standard_price - new_price
//                if float_is_zero(diff, precision_rounding=product.currency_id.rounding):
//                    raise UserError(
//                        _("No difference between standard price and new price!"))
//                if not product_accounts[product.id].get('stock_valuation', False):
//                    raise UserError(
//                        _('You don\'t have any stock valuation account defined on your product category. You must define one before processing this operation.'))
//                qty_available = product.qty_available
//                if qty_available:
//                    # Accounting Entries
//                    if diff * qty_available > 0:
//                        debit_account_id = account_id
//                        credit_account_id = product_accounts[product.id]['stock_valuation'].id
//                    else:
//                        debit_account_id = product_accounts[product.id]['stock_valuation'].id
//                        credit_account_id = account_id
//
//                    move_vals = {
//                        'journal_id': product_accounts[product.id]['stock_journal'].id,
//                        'company_id': location.company_id.id,
//                        'line_ids': [(0, 0, {
//                            'name': _('Standard Price changed'),
//                            'account_id': debit_account_id,
//                            'debit': abs(diff * qty_available),
//                            'credit': 0,
//                            'product_id': product.id,
//                        }), (0, 0, {
//                            'name': _('Standard Price changed'),
//                            'account_id': credit_account_id,
//                            'debit': 0,
//                            'credit': abs(diff * qty_available),
//                            'product_id': product.id,
//                        })],
//                    }
//                    move = AccountMove.create(move_vals)
//                    move.post()
//        self.write({'standard_price': new_price})
//        return True
})
h.ProductProduct().Methods().AngloSaxonSaleMoveLines().DeclareMethod(
`Prepare dicts describing new journal COGS journal items
for a product sale.

        Returns a dict that should be passed to `_convert_prepared_anglosaxon_line()`
to
        obtain the creation value for the new journal items.

        :param Model product: a product.product record
of the product being sold
        :param Model uom: a product.uom record of the UoM
of the sale line
        :param Integer qty: quantity of the product being sold
        :param Integer price_unit: unit price of the product being sold
        :param Model currency: a res.currency record from
the order of the product being sold
        :param Interger amount_currency: unit price in
the currency from the order of the product being sold
        :param Model fiscal_position: a account.fiscal.position
record from the order of the product being sold
        :param Model account_analytic: a account.account.analytic
record from the line of the product being sold
        `,
func(rs m.ProductProductSet, name interface{}, product interface{}, uom interface{}, qty interface{}, price_unit interface{}, currency interface{}, amount_currency interface{}, fiscal_position interface{}, account_analytic interface{}, analytic_tags interface{})  {
//        if product.type == 'product' and product.valuation == 'real_time':
//            accounts = product.product_tmpl_id.get_product_accounts(
//                fiscal_pos=fiscal_position)
//            # debit account dacc will be the output account
//            dacc = accounts['stock_output'].id
//            # credit account cacc will be the expense account
//            cacc = accounts['expense'].id
//            if dacc and cacc:
//                return [
//                    {
//                        'type': 'src',
//                        'name': name[:64],
//                        'price_unit': price_unit,
//                        'quantity': qty,
//                        'price': price_unit * qty,
//                        'currency_id': currency and currency.id,
//                        'amount_currency': amount_currency,
//                        'account_id': dacc,
//                        'product_id': product.id,
//                        'uom_id': uom.id,
//                        'account_analytic_id': account_analytic and account_analytic.id,
//                        'analytic_tag_ids': analytic_tags and analytic_tags.ids and [(6, 0, analytic_tags.ids)] or False,
//                    },
//
//                    {
//                        'type': 'src',
//                        'name': name[:64],
//                        'price_unit': price_unit,
//                        'quantity': qty,
//                        'price': -1 * price_unit * qty,
//                        'currency_id': currency and currency.id,
//                        'amount_currency': -1 * amount_currency,
//                        'account_id': cacc,
//                        'product_id': product.id,
//                        'uom_id': uom.id,
//                        'account_analytic_id': account_analytic and account_analytic.id,
//                        'analytic_tag_ids': analytic_tags and analytic_tags.ids and [(6, 0, analytic_tags.ids)] or False,
//                    },
//                ]
//        return []
})
h.ProductProduct().Methods().GetAngloSaxonPriceUnit().DeclareMethod(
`GetAngloSaxonPriceUnit`,
func(rs m.ProductProductSet, uom interface{})  {
//        price = self.standard_price
//        if not self or not uom or self.uom_id.id == uom.id:
//            return price or 0.0
//        return self.uom_id._compute_price(price, uom)
})
h.ProductCategory().DeclareModel()

h.ProductCategory().AddFields(map[string]models.FieldDefinition{
"PropertyValuation": models.SelectionField{
Selection: types.Selection{
"manual_periodic": "Periodic (manual)",
"real_time": "Perpetual (automated)",
},
String: "Inventory Valuation",
//company_dependent=True
NoCopy: false,
Required: true,
Help: "If perpetual valuation is enabled for a product, the system" + 
"will automatically create journal entries corresponding" + 
"to stock moves, with product price as specified by the" + 
"'Costing Method'. The inventory variation account set on" + 
"the product category will represent the current inventory" + 
"value, and the stock input and stock output account will" + 
"hold the counterpart moves for incoming and outgoing products.",
},
"PropertyCostMethod": models.SelectionField{
Selection: types.Selection{
"standard": "Standard Price",
"average": "Average Price",
"real": "Real Price",
},
String: "Costing Method",
//company_dependent=True
NoCopy: false,
Required: true,
Help: "Standard Price: The cost price is manually updated at the" + 
"end of a specific period (usually once a year)." + 
"Average Price: The cost price is recomputed at each incoming" + 
"shipment and used for the product valuation." + 
"Real Price: The cost price displayed is the price of the" + 
"last outgoing product (will be used in case of inventory" + 
"loss for example).",
},
"PropertyStockJournal": models.Many2OneField{
RelationModel: h.AccountJournal(),
String: "Stock Journal",
//company_dependent=True
Help: "When doing real-time inventory valuation, this is the Accounting" + 
"Journal in which entries will be automatically posted when" + 
"stock moves are processed.",
},
"PropertyStockAccountInputCategId": models.Many2OneField{
RelationModel: h.AccountAccount(),
String: "Stock Input Account",
//company_dependent=True
Filter: q.Deprecated().Equals(False),
//oldname="property_stock_account_input_categ"
Help: "When doing real-time inventory valuation, counterpart journal" + 
"items for all incoming stock moves will be posted in this" + 
"account, unless there is a specific valuation account set" + 
"on the source location. This is the default value for all" + 
"products in this category. It can also directly be set on each product",
},
"PropertyStockAccountOutputCategId": models.Many2OneField{
RelationModel: h.AccountAccount(),
String: "Stock Output Account",
//company_dependent=True
Filter: q.Deprecated().Equals(False),
//oldname="property_stock_account_output_categ"
Help: "When doing real-time inventory valuation, counterpart journal" + 
"items for all outgoing stock moves will be posted in this" + 
"account, unless there is a specific valuation account set" + 
"on the destination location. This is the default value" + 
"for all products in this category. It can also directly" + 
"be set on each product",
},
"PropertyStockValuationAccountId": models.Many2OneField{
RelationModel: h.AccountAccount(),
String: "Stock Valuation Account",
//company_dependent=True
Filter: q.Deprecated().Equals(False),
Help: "When real-time inventory valuation is enabled on a product," + 
"this account will hold the current value of the products.",
},
})
}