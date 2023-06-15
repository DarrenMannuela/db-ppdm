package dto

type Oblig_deduction struct {
	Obligation_id         string   `json:"obligation_id" default:"source"`
	Obligation_seq_no     int      `json:"obligation_seq_no" default:"1"`
	Deduction_id          string   `json:"deduction_id" default:"source"`
	Active_ind            *string  `json:"active_ind" default:""`
	Allow_deduction_id    *string  `json:"allow_deduction_id" default:""`
	Currency_conversion   *float64 `json:"currency_conversion" default:""`
	Currency_ouom         *string  `json:"currency_ouom" default:""`
	Deduction_amount      *float64 `json:"deduction_amount" default:""`
	Deduction_percent     *float64 `json:"deduction_percent" default:""`
	Deduct_type           *string  `json:"deduct_type" default:""`
	Effective_date        *string  `json:"effective_date" default:""`
	Expiry_date           *string  `json:"expiry_date" default:""`
	Max_deduction_allowed *float64 `json:"max_deduction_allowed" default:""`
	Ppdm_guid             *string  `json:"ppdm_guid" default:""`
	Remark                *string  `json:"remark" default:""`
	Royalty_amount        *float64 `json:"royalty_amount" default:""`
	Source                *string  `json:"source" default:""`
	Row_changed_by        *string  `json:"row_changed_by" default:""`
	Row_changed_date      *string  `json:"row_changed_date" default:""`
	Row_created_by        *string  `json:"row_created_by" default:""`
	Row_created_date      *string  `json:"row_created_date" default:""`
	Row_effective_date    *string  `json:"row_effective_date" default:""`
	Row_expiry_date       *string  `json:"row_expiry_date" default:""`
	Row_quality           *string  `json:"row_quality" default:""`
}
