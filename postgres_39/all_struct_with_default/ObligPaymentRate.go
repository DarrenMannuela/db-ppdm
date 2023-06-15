package dto

type Oblig_payment_rate struct {
	Jurisdiction       string   `json:"jurisdiction" default:"source"`
	Pay_rate_type      string   `json:"pay_rate_type" default:"source"`
	Payment_rate_id    string   `json:"payment_rate_id" default:"source"`
	Active_ind         *string  `json:"active_ind" default:""`
	Contract_id        *string  `json:"contract_id" default:""`
	Effective_date     *string  `json:"effective_date" default:""`
	Expiry_date        *string  `json:"expiry_date" default:""`
	Payment_rate       *float64 `json:"payment_rate" default:""`
	Payment_rate_ouom  *string  `json:"payment_rate_ouom" default:""`
	Ppdm_guid          *string  `json:"ppdm_guid" default:""`
	Rate_percent       *float64 `json:"rate_percent" default:""`
	Remark             *string  `json:"remark" default:""`
	Source             *string  `json:"source" default:""`
	Row_changed_by     *string  `json:"row_changed_by" default:""`
	Row_changed_date   *string  `json:"row_changed_date" default:""`
	Row_created_by     *string  `json:"row_created_by" default:""`
	Row_created_date   *string  `json:"row_created_date" default:""`
	Row_effective_date *string  `json:"row_effective_date" default:""`
	Row_expiry_date    *string  `json:"row_expiry_date" default:""`
	Row_quality        *string  `json:"row_quality" default:""`
}
