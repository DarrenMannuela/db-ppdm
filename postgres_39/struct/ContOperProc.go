package dto

import (
	"time"
)

type Cont_oper_proc struct {
	Contract_id              string     `json:"contract_id"`
	Operating_procedure_id   string     `json:"operating_procedure_id"`
	Active_ind               *string    `json:"active_ind"`
	Adjust_13_month_desc     *string    `json:"adjust_13_month_desc"`
	Adjust_13_month_ind      *string    `json:"adjust_13_month_ind"`
	Casing_point_election    *string    `json:"casing_point_election"`
	Claim_amount_limit       *float64   `json:"claim_amount_limit"`
	Claim_amount_limit_ouom  *string    `json:"claim_amount_limit_ouom"`
	Dev_penalty_cost         *float64   `json:"dev_penalty_cost"`
	Dev_penalty_cost_ouom    *string    `json:"dev_penalty_cost_ouom"`
	Dev_penalty_percent      *float64   `json:"dev_penalty_percent"`
	Effective_date           *time.Time `json:"effective_date"`
	Equip_penalty_cost       *float64   `json:"equip_penalty_cost"`
	Equip_penalty_cost_ouom  *string    `json:"equip_penalty_cost_ouom"`
	Equip_penalty_percent    *float64   `json:"equip_penalty_percent"`
	Expiry_date              *time.Time `json:"expiry_date"`
	Expl_penalty_cost        *float64   `json:"expl_penalty_cost"`
	Expl_penalty_cost_ouom   *string    `json:"expl_penalty_cost_ouom"`
	Expl_penalty_percent     *float64   `json:"expl_penalty_percent"`
	Inflation_adjustment_ind *string    `json:"inflation_adjustment_ind"`
	Insurance_election       *string    `json:"insurance_election"`
	Mktg_fee_elect_cost      *float64   `json:"mktg_fee_elect_cost"`
	Mktg_fee_elect_cost_ouom *string    `json:"mktg_fee_elect_cost_ouom"`
	Mktg_fee_elect_percent   *float64   `json:"mktg_fee_elect_percent"`
	Nat_currency_conversion  *float64   `json:"nat_currency_conversion"`
	Nat_currency_uom         *string    `json:"nat_currency_uom"`
	Operating_procedure_type *string    `json:"operating_procedure_type"`
	Ppdm_guid                string     `json:"ppdm_guid"`
	Recog_on_assignment_ind  *string    `json:"recog_on_assignment_ind"`
	Remark                   *string    `json:"remark"`
	Rofr_ind                 *string    `json:"rofr_ind"`
	Rofr_reply_term          *float64   `json:"rofr_reply_term"`
	Rofr_reply_term_uom      *string    `json:"rofr_reply_term_uom"`
	Source                   *string    `json:"source"`
	Source_document_id       *string    `json:"source_document_id"`
	Statute_limit_period     *float64   `json:"statute_limit_period"`
	Statute_limit_period_uom *string    `json:"statute_limit_period_uom"`
	Surrender_notice         *float64   `json:"surrender_notice"`
	Surrender_notice_uom     *string    `json:"surrender_notice_uom"`
	Title_preserve_desc      *string    `json:"title_preserve_desc"`
	Unauthorized_amount      *float64   `json:"unauthorized_amount"`
	Unauthorized_amount_ouom *string    `json:"unauthorized_amount_ouom"`
	Row_changed_by           *string    `json:"row_changed_by"`
	Row_changed_date         *time.Time `json:"row_changed_date"`
	Row_created_by           *string    `json:"row_created_by"`
	Row_created_date         *time.Time `json:"row_created_date"`
	Row_effective_date       *time.Time `json:"row_effective_date"`
	Row_expiry_date          *time.Time `json:"row_expiry_date"`
	Row_quality              *string    `json:"row_quality"`
}
