package stock_account

import (
	"github.com/hexya-erp/hexya-base/web/webdata"
	"github.com/hexya-erp/hexya/src/models"
	"github.com/hexya-erp/pool/h"
)

func init() {
	h.StockHistory().DeclareModel()

	h.StockHistory().AddFields(map[string]models.FieldDefinition{
		"MoveId": models.Many2OneField{
			RelationModel: h.StockMove(),
			String:        "Stock Move",
			Required:      true,
		},
		"LocationId": models.Many2OneField{
			RelationModel: h.StockLocation(),
			String:        "Location",
			Required:      true,
		},
		"CompanyId": models.Many2OneField{
			RelationModel: h.Company(),
			String:        "Company",
		},
		"ProductId": models.Many2OneField{
			RelationModel: h.ProductProduct(),
			String:        "Product",
			Required:      true,
		},
		"ProductCategId": models.Many2OneField{
			RelationModel: h.ProductCategory(),
			String:        "Product Category",
			Required:      true,
		},
		"Quantity": models.FloatField{
			String: "Product Quantity",
		},
		"Date": models.DateTimeField{
			String: "Operation Date",
		},
		"PriceUnitOnQuant": models.FloatField{
			String: "Value",
		},
		"InventoryValue": models.FloatField{
			String:   "Inventory Value",
			Compute:  h.StockHistory().Methods().ComputeInventoryValue(),
			ReadOnly: true,
		},
		"Source": models.CharField{
			String: "Source",
		},
		"ProductTemplateId": models.Many2OneField{
			RelationModel: h.ProductTemplate(),
			String:        "Product Template",
			Required:      true,
		},
		"SerialNumber": models.CharField{
			String:   "Lot/Serial Number",
			Required: true,
		},
	})
	h.StockHistory().Methods().ReadGroup().Extend(
		`ReadGroup`,
		func(rs m.StockHistorySet, domain webdata.ReadGroupParams, fields interface{}, groupby interface{}, offset interface{}, limit interface{}, orderby interface{}, lazy interface{}) {
			//        res = super(StockHistory, self).read_group(
			//            domain, fields, groupby, offset=offset, limit=limit, orderby=orderby, lazy=lazy)
			//        if 'inventory_value' in fields:
			//            groupby_list = groupby[:1] if lazy else groupby
			//            date = self._context.get('history_date', fieldsDatetime.now())
			//
			//            # Step 2: retrieve the stock history lines. The result contains the 'expanded'
			//            # version of the read_group. We build the query manually for performance reason
			//            # (and avoid a costly 'WHERE id IN ...').
			//            fields_2 = set(
			//                ['id', 'product_id', 'price_unit_on_quant',
			//                    'company_id', 'quantity'] + groupby_list
			//            )
			//            query = self._where_calc(domain)
			//            self._apply_ir_rules(query, 'read')
			//            tables, where_clause, where_clause_params = query.get_sql()
			//            select = "SELECT %s FROM %s WHERE %s "
			//            query = select % (','.join(fields_2), tables, where_clause)
			//            self._cr.execute(query, where_clause_params)
			//
			//            # Step 3: match the lines retrieved at step 2 with the aggregated results of step 1.
			//            # In other words, we link each item of the read_group result with the corresponding
			//            # lines.
			//            stock_history_data = {}
			//            stock_histories_by_group = {}
			//            for line in self._cr.dictfetchall():
			//                stock_history_data[line['id']] = line
			//                key = tuple(line.get(g) or False for g in groupby_list)
			//                stock_histories_by_group.setdefault(key, [])
			//                stock_histories_by_group[key] += [line['id']]
			//
			//            histories_dict = {}
			//            not_real_cost_method_products = self.env['product.product'].browse(
			//                record['product_id'] for record in stock_history_data.values()
			//            ).filtered(lambda product: product.cost_method != 'real')
			//            if not_real_cost_method_products:
			//                self._cr.execute("""SELECT DISTINCT ON (product_id, company_id) product_id, company_id, cost
			//                    FROM product_price_history
			//                    WHERE product_id in %s AND datetime <= %s
			//                    ORDER BY product_id, company_id, datetime DESC, id DESC""", (tuple(not_real_cost_method_products.ids), date))
			//                for history in self._cr.dictfetchall():
			//                    histories_dict[(history['product_id'],
			//                                    history['company_id'])] = history['cost']
			//
			//            for line in res:
			//                inv_value = 0.0
			//                # Build the same keys than above, but need to take into account Many2one are tuples
			//                key = tuple(
			//                    line[g] if g in line else False for g in groupby_list)
			//                key = tuple(k[0] if isinstance(k, tuple) else k for k in key)
			//                for stock_history in self.env['stock.history'].browse(stock_histories_by_group[key]):
			//                    history_data = stock_history_data[stock_history.id]
			//                    product_id = history_data['product_id']
			//                    if self.env['product.product'].browse(product_id).cost_method == 'real':
			//                        price = history_data['price_unit_on_quant']
			//                    else:
			//                        price = histories_dict.get(
			//                            (product_id, history_data['company_id']), 0.0)
			//                    inv_value += price * history_data['quantity']
			//                line['inventory_value'] = inv_value
			//        return res
		})
	h.StockHistory().Methods().ComputeInventoryValue().DeclareMethod(
		`ComputeInventoryValue`,
		func(rs h.StockHistorySet) h.StockHistoryData {
			//        if self.product_id.cost_method == 'real':
			//            self.inventory_value = self.quantity * self.price_unit_on_quant
			//        else:
			//            self.inventory_value = self.quantity * self.product_id.get_history_price(
			//                self.company_id.id, date=self._context.get('history_date', fields.Datetime.now()))
		})
	h.StockHistory().Methods().Init().DeclareMethod(
		`Init`,
		func(rs m.StockHistorySet) {
			//        tools.drop_view_if_exists(self._cr, 'stock_history')
			//        self._cr.execute("""
			//            CREATE VIEW stock_history AS (
			//              SELECT MIN(id) as id,
			//                move_id,
			//                location_id,
			//                company_id,
			//                product_id,
			//                product_categ_id,
			//                product_template_id,
			//                SUM(quantity) as quantity,
			//                date,
			//                COALESCE(SUM(price_unit_on_quant * quantity) / NULLIF(SUM(quantity), 0), 0) as price_unit_on_quant,
			//                source,
			//                string_agg(DISTINCT serial_number, ', ' ORDER BY serial_number) AS serial_number
			//                FROM
			//                ((SELECT
			//                    stock_move.id AS id,
			//                    stock_move.id AS move_id,
			//                    dest_location.id AS location_id,
			//                    dest_location.company_id AS company_id,
			//                    stock_move.product_id AS product_id,
			//                    product_template.id AS product_template_id,
			//                    product_template.categ_id AS product_categ_id,
			//                    quant.qty AS quantity,
			//                    stock_move.date AS date,
			//                    quant.cost as price_unit_on_quant,
			//                    stock_move.origin AS source,
			//                    stock_production_lot.name AS serial_number
			//                FROM
			//                    stock_quant as quant
			//                JOIN
			//                    stock_quant_move_rel ON stock_quant_move_rel.quant_id = quant.id
			//                JOIN
			//                    stock_move ON stock_move.id = stock_quant_move_rel.move_id
			//                LEFT JOIN
			//                    stock_production_lot ON stock_production_lot.id = quant.lot_id
			//                JOIN
			//                    stock_location dest_location ON stock_move.location_dest_id = dest_location.id
			//                JOIN
			//                    stock_location source_location ON stock_move.location_id = source_location.id
			//                JOIN
			//                    product_product ON product_product.id = stock_move.product_id
			//                JOIN
			//                    product_template ON product_template.id = product_product.product_tmpl_id
			//                WHERE quant.qty>0 AND stock_move.state = 'done' AND dest_location.usage in ('internal', 'transit')
			//                AND (
			//                    not (source_location.company_id is null and dest_location.company_id is null) or
			//                    source_location.company_id != dest_location.company_id or
			//                    source_location.usage not in ('internal', 'transit'))
			//                ) UNION ALL
			//                (SELECT
			//                    (-1) * stock_move.id AS id,
			//                    stock_move.id AS move_id,
			//                    source_location.id AS location_id,
			//                    source_location.company_id AS company_id,
			//                    stock_move.product_id AS product_id,
			//                    product_template.id AS product_template_id,
			//                    product_template.categ_id AS product_categ_id,
			//                    - quant.qty AS quantity,
			//                    stock_move.date AS date,
			//                    quant.cost as price_unit_on_quant,
			//                    stock_move.origin AS source,
			//                    stock_production_lot.name AS serial_number
			//                FROM
			//                    stock_quant as quant
			//                JOIN
			//                    stock_quant_move_rel ON stock_quant_move_rel.quant_id = quant.id
			//                JOIN
			//                    stock_move ON stock_move.id = stock_quant_move_rel.move_id
			//                LEFT JOIN
			//                    stock_production_lot ON stock_production_lot.id = quant.lot_id
			//                JOIN
			//                    stock_location source_location ON stock_move.location_id = source_location.id
			//                JOIN
			//                    stock_location dest_location ON stock_move.location_dest_id = dest_location.id
			//                JOIN
			//                    product_product ON product_product.id = stock_move.product_id
			//                JOIN
			//                    product_template ON product_template.id = product_product.product_tmpl_id
			//                WHERE quant.qty>0 AND stock_move.state = 'done' AND source_location.usage in ('internal', 'transit')
			//                AND (
			//                    not (dest_location.company_id is null and source_location.company_id is null) or
			//                    dest_location.company_id != source_location.company_id or
			//                    dest_location.usage not in ('internal', 'transit'))
			//                ))
			//                AS foo
			//                GROUP BY move_id, location_id, company_id, product_id, product_categ_id, date, source, product_template_id
			//            )""")
		})
}
