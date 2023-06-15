package dto

type Cont_oper_proc struct {
	Contract_id              string   `json:"contract_id" default:"source"`
	Operating_procedure_id   string   `json:"operating_procedure_id" default:"source"`
	Active_ind               *string  `json:"active_ind" default:""`
	Adjust_13_month_desc     *string  `json:"adjust_13_month_desc" default:""`
	Adjust_13_month_ind      *string  `json:"adjust_13_month_ind" default:""`
	Casing_point_election    *string  `json:"casing_point_election" default:""`
	Claim_amount_limit       *float64 `json:"claim_amount_limit" default:""`
	Claim_amount_limit_ouom  *string  `json:"claim_amount_limit_ouom" default:""`
	Dev_penalty_cost         *float64 `json:"dev_penalty_cost" default:""`
	Dev_penalty_cost_ouom    *string  `json:"dev_penalty_cost_ouom" default:""`
	Dev_penalty_percent      *float64 `json:"dev_penalty_percent" default:""`
	Effective_date           *string  `json:"effective_date" default:""`
	Equip_penalty_cost       *float64 `json:"equip_penalty_cost" default:""`
	Equip_penalty_cost_ouom  *string  `json:"equip_penalty_cost_ouom" default:""`
	Equip_penalty_percent    *float64 `json:"equip_penalty_percent" default:""`
	Expiry_date              *string  `json:"expiry_date" default:""`
	Expl_penalty_cost        *float64 `json:"expl_penalty_cost" default:""`
	Expl_penalty_cost_ouom   *string  `json:"expl_penalty_cost_ouom" default:""`
	Expl_penalty_percent     *float64 `json:"expl_penalty_percent" default:""`
	Inflation_adjustment_ind *string  `json:"inflation_adjustment_ind" default:""`
	Insurance_election       *string  `json:"insurance_election" default:""`
	Mktg_fee_elect_cost      *float64 `json:"mktg_fee_elect_cost" default:""`
	Mktg_fee_elect_cost_ouom *string  `json:"mktg_fee_elect_cost_ouom" default:""`
	Mktg_fee_elect_percent   *float64 `json:"mktg_fee_elect_percent" default:""`
	Nat_currency_conversion  *float64 `json:"nat_currency_conversion" default:""`
	Nat_currency_uom         *string  `json:"nat_currency_uom" default:""`
	Operating_procedure_type *string  `json:"operating_procedure_type" default:""`
	Ppdm_guid                *string  `json:"ppdm_guid" default:""`
	Recog_on_assignment_ind  *string  `json:"recog_on_assignment_ind" default:""`
	Remark                   *string  `json:"remark" default:""`
	Rofr_ind                 *string  `json:"rofr_ind" default:""`
	Rofr_reply_term          *float64 `json:"rofr_reply_term" default:""`
	Rofr_reply_term_uom      *string  `json:"rofr_reply_term_uom" default:""`
	Source                   *string  `json:"source" default:""`
	Source_document_id       *string  `json:"source_document_id" default:""`
	Statute_limit_period     *float64 `json:"statute_limit_period" default:""`
	Statute_limit_period_uom *string  `json:"statute_limit_period_uom" default:""`
	Surrender_notice         *float64 `json:"surrender_notice" default:""`
	Surrender_notice_uom     *string  `json:"surrender_notice_uom" default:""`
	Title_preserve_desc      *string  `json:"title_preserve_desc" default:""`
	Unauthorized_amount      *float64 `json:"unauthorized_amount" default:""`
	Unauthorized_amount_ouom *string  `json:"unauthorized_amount_ouom" default:""`
	Row_changed_by           *string  `json:"row_changed_by" default:""`
	Row_changed_date         *string  `json:"row_changed_date" default:""`
	Row_created_by           *string  `json:"row_created_by" default:""`
	Row_created_date         *string  `json:"row_created_date" default:""`
	Row_effective_date       *string  `json:"row_effective_date" default:""`
	Row_expiry_date          *string  `json:"row_expiry_date" default:""`
	Row_quality              *string  `json:"row_quality" default:""`
}
