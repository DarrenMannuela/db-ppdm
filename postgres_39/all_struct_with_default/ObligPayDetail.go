package dto

type Oblig_pay_detail struct {
	Obligation_id       string   `json:"obligation_id" default:"source"`
	Obligation_seq_no   int      `json:"obligation_seq_no" default:"1"`
	Payee_payor_ba_id   string   `json:"payee_payor_ba_id" default:"source"`
	Payee_payor         string   `json:"payee_payor" default:"source"`
	Detail_id           string   `json:"detail_id" default:"source"`
	Active_ind          *string  `json:"active_ind" default:""`
	Cheque_number       *string  `json:"cheque_number" default:""`
	Currency_conversion *float64 `json:"currency_conversion" default:""`
	Currency_ouom       *string  `json:"currency_ouom" default:""`
	Deduction_ind       *string  `json:"deduction_ind" default:""`
	Effective_date      *string  `json:"effective_date" default:""`
	Expiry_date         *string  `json:"expiry_date" default:""`
	Payment_amount      *float64 `json:"payment_amount" default:""`
	Payment_date        *string  `json:"payment_date" default:""`
	Pay_detail_type     *string  `json:"pay_detail_type" default:""`
	Percent_of_payment  *float64 `json:"percent_of_payment" default:""`
	Ppdm_guid           *string  `json:"ppdm_guid" default:""`
	Rate                *float64 `json:"rate" default:""`
	Rate_ouom           *string  `json:"rate_ouom" default:""`
	Remark              *string  `json:"remark" default:""`
	Source              *string  `json:"source" default:""`
	Row_changed_by      *string  `json:"row_changed_by" default:""`
	Row_changed_date    *string  `json:"row_changed_date" default:""`
	Row_created_by      *string  `json:"row_created_by" default:""`
	Row_created_date    *string  `json:"row_created_date" default:""`
	Row_effective_date  *string  `json:"row_effective_date" default:""`
	Row_expiry_date     *string  `json:"row_expiry_date" default:""`
	Row_quality         *string  `json:"row_quality" default:""`
}
