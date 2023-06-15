package dto

type Oblig_payment_instr struct {
	Payee_ba_id            string   `json:"payee_ba_id" default:"source"`
	Payment_instruction_id string   `json:"payment_instruction_id" default:"source"`
	Aba_number             *string  `json:"aba_number" default:""`
	Active_ind             *string  `json:"active_ind" default:""`
	Bank_address_obs_no    *int     `json:"bank_address_obs_no" default:""`
	Bank_address_source    *string  `json:"bank_address_source" default:""`
	Bank_ba_id             *string  `json:"bank_ba_id" default:""`
	Bank_fee               *float64 `json:"bank_fee" default:""`
	Currency_conversion    *float64 `json:"currency_conversion" default:""`
	Currency_ouom          *string  `json:"currency_ouom" default:""`
	Depository_num         *string  `json:"depository_num" default:""`
	Effective_date         *string  `json:"effective_date" default:""`
	Expiry_date            *string  `json:"expiry_date" default:""`
	Invalid_date           *string  `json:"invalid_date" default:""`
	Pay_method             *string  `json:"pay_method" default:""`
	Ppdm_guid              *string  `json:"ppdm_guid" default:""`
	Remark                 *string  `json:"remark" default:""`
	Source                 *string  `json:"source" default:""`
	Suspend_payment_ind    *string  `json:"suspend_payment_ind" default:""`
	Suspend_payment_reason *string  `json:"suspend_payment_reason" default:""`
	Row_changed_by         *string  `json:"row_changed_by" default:""`
	Row_changed_date       *string  `json:"row_changed_date" default:""`
	Row_created_by         *string  `json:"row_created_by" default:""`
	Row_created_date       *string  `json:"row_created_date" default:""`
	Row_effective_date     *string  `json:"row_effective_date" default:""`
	Row_expiry_date        *string  `json:"row_expiry_date" default:""`
	Row_quality            *string  `json:"row_quality" default:""`
}
