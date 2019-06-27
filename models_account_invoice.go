package stock_account

import (
	"github.com/hexya-erp/pool/h"
)

//import logging
//_logger = logging.getLogger(__name__)
func init() {
	h.AccountInvoice().DeclareModel()

	h.AccountInvoice().Methods().InvoiceLineMoveLineGet().DeclareMethod(
		`InvoiceLineMoveLineGet`,
		func(rs m.AccountInvoiceSet) {
			//        res = super(AccountInvoice, self).invoice_line_move_line_get()
			//        if self.company_id.anglo_saxon_accounting and self.type in ('out_invoice', 'out_refund'):
			//            for i_line in self.invoice_line_ids:
			//                res.extend(self._anglo_saxon_sale_move_lines(i_line))
			//        return res
		})
	h.AccountInvoice().Methods().AngloSaxonSaleMoveLines().DeclareMethod(
		`Return the additional move lines for sales invoices and refunds.

        i_line: An account.invoice.line object.
        res: The move line entries produced so far by the
parent move_line_get.
        `,
		func(rs m.AccountInvoiceSet, i_line interface{}) {
			//        inv = i_line.invoice_id
			//        company_currency = inv.company_id.currency_id
			//        price_unit = i_line._get_anglo_saxon_price_unit()
			//        if inv.currency_id != company_currency:
			//            currency = inv.currency_id
			//            amount_currency = i_line._get_price(company_currency, price_unit)
			//        else:
			//            currency = False
			//            amount_currency = False
			//        return self.env['product.product']._anglo_saxon_sale_move_lines(i_line.name, i_line.product_id, i_line.uom_id, i_line.quantity, price_unit, currency=currency, amount_currency=amount_currency, fiscal_position=inv.fiscal_position_id, account_analytic=i_line.account_analytic_id, analytic_tags=i_line.analytic_tag_ids)
		})
	h.AccountInvoiceLine().DeclareModel()

	h.AccountInvoiceLine().Methods().GetAngloSaxonPriceUnit().DeclareMethod(
		`GetAngloSaxonPriceUnit`,
		func(rs m.AccountInvoiceLineSet) {
			//        self.ensure_one()
			//        if not self.product_id:
			//            return self.price_unit
			//        return self.product_id._get_anglo_saxon_price_unit(uom=self.uom_id)
		})
	h.AccountInvoiceLine().Methods().GetPrice().DeclareMethod(
		`GetPrice`,
		func(rs m.AccountInvoiceLineSet, company_currency interface{}, price_unit interface{}) {
			//        if self.invoice_id.currency_id.id != company_currency.id:
			//            price = company_currency.with_context(date=self.invoice_id.date_invoice).compute(
			//                price_unit * self.quantity, self.invoice_id.currency_id)
			//        else:
			//            price = price_unit * self.quantity
			//        return round(price, self.invoice_id.currency_id.decimal_places)
		})
	h.AccountInvoiceLine().Methods().GetInvoiceLineAccount().DeclareMethod(
		`GetInvoiceLineAccount`,
		func(rs m.AccountInvoiceLineSet, typeName interface{}, product interface{}, fpos interface{}, company interface{}) {
			//        if company.anglo_saxon_accounting and type in ('in_invoice', 'in_refund') and product and product.type == 'product':
			//            accounts = product.product_tmpl_id.get_product_accounts(
			//                fiscal_pos=fpos)
			//            if accounts['stock_input']:
			//                return accounts['stock_input']
			//        return super(AccountInvoiceLine, self).get_invoice_line_account(type, product, fpos, company)
		})
}
