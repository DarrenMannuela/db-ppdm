package dto

type Seis_transaction struct {
	Seis_transaction_id string   `json:"seis_transaction_id" default:"source"`
	Transaction_type    string   `json:"transaction_type" default:"source"`
	Active_ind          *string  `json:"active_ind" default:""`
	Currency_conversion *float64 `json:"currency_conversion" default:""`
	Currency_ouom       *string  `json:"currency_ouom" default:""`
	Effective_date      *string  `json:"effective_date" default:""`
	Expiry_date         *string  `json:"expiry_date" default:""`
	Length              *float64 `json:"length" default:""`
	Length_ouom         *string  `json:"length_ouom" default:""`
	Ppdm_guid           *string  `json:"ppdm_guid" default:""`
	Price_per_length    *float64 `json:"price_per_length" default:""`
	Reference_num       *string  `json:"reference_num" default:""`
	Remark              *string  `json:"remark" default:""`
	Source              *string  `json:"source" default:""`
	Total_price         *float64 `json:"total_price" default:""`
	Transaction_date    *string  `json:"transaction_date" default:""`
	Transaction_status  *string  `json:"transaction_status" default:""`
	Row_changed_by      *string  `json:"row_changed_by" default:""`
	Row_changed_date    *string  `json:"row_changed_date" default:""`
	Row_created_by      *string  `json:"row_created_by" default:""`
	Row_created_date    *string  `json:"row_created_date" default:""`
	Row_effective_date  *string  `json:"row_effective_date" default:""`
	Row_expiry_date     *string  `json:"row_expiry_date" default:""`
	Row_quality         *string  `json:"row_quality" default:""`
}
