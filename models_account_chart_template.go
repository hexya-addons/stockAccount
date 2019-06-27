package stock_account

import (
	"github.com/hexya-erp/pool/h"
)

//import logging
//_logger = logging.getLogger(__name__)
func init() {
	h.AccountChartTemplate().DeclareModel()

	h.AccountChartTemplate().Methods().GenerateJournals().DeclareMethod(
		`GenerateJournals`,
		func(rs m.AccountChartTemplateSet, acc_template_ref interface{}, company interface{}, journals_dict interface{}) {
			//        journal_to_add = [{'name': _(
			//            'Stock Journal'), 'type': 'general', 'code': 'STJ', 'favorite': False, 'sequence': 8}]
			//        return super(AccountChartTemplate, self).generate_journals(acc_template_ref=acc_template_ref, company=company, journals_dict=journal_to_add)
		})
	h.AccountChartTemplate().Methods().GenerateProperties().DeclareMethod(
		`GenerateProperties`,
		func(rs m.AccountChartTemplateSet, acc_template_ref interface{}, company interface{}, property_list interface{}) {
			//        res = super(AccountChartTemplate, self).generate_properties(
			//            acc_template_ref=acc_template_ref, company=company)
			//        PropertyObj = self.env['ir.property']  # Property Stock Journal
			//        value = self.env['account.journal'].search(
			//            [('company_id', '=', company.id), ('code', '=', 'STJ'), ('type', '=', 'general')], limit=1)
			//        if value:
			//            field = self.env['ir.model.fields'].search([('name', '=', 'property_stock_journal'), (
			//                'model', '=', 'product.category'), ('relation', '=', 'account.journal')], limit=1)
			//            vals = {
			//                'name': 'property_stock_journal',
			//                'company_id': company.id,
			//                'fields_id': field.id,
			//                'value': 'account.journal,%s' % value.id,
			//            }
			//            properties = PropertyObj.search(
			//                [('name', '=', 'property_stock_journal'), ('company_id', '=', company.id)])
			//            if properties:
			//                # the property exist: modify it
			//                properties.write(vals)
			//            else:
			//                # create the property
			//                PropertyObj.create(vals)
			//        todo_list = [  # Property Stock Accounts
			//            'property_stock_account_input_categ_id',
			//            'property_stock_account_output_categ_id',
			//            'property_stock_valuation_account_id',
			//        ]
			//        for record in todo_list:
			//            account = getattr(self, record)
			//            value = account and 'account.account,' + \
			//                str(acc_template_ref[account.id]) or False
			//            if value:
			//                field = self.env['ir.model.fields'].search([('name', '=', record), (
			//                    'model', '=', 'product.category'), ('relation', '=', 'account.account')], limit=1)
			//                vals = {
			//                    'name': record,
			//                    'company_id': company.id,
			//                    'fields_id': field.id,
			//                    'value': value,
			//                }
			//                properties = PropertyObj.search(
			//                    [('name', '=', record), ('company_id', '=', company.id)])
			//                if properties:
			//                    # the property exist: modify it
			//                    properties.write(vals)
			//                else:
			//                    # create the property
			//                    PropertyObj.create(vals)
			//        return res
		})
}
