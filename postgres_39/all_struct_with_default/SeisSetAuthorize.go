package dto

type Seis_set_authorize struct {
	Seis_set_subtype    string   `json:"seis_set_subtype" default:"source"`
	Seis_set_id         string   `json:"seis_set_id" default:"source"`
	Authorize_id        string   `json:"authorize_id" default:"source"`
	Active_ind          *string  `json:"active_ind" default:""`
	Authorized_by       *string  `json:"authorized_by" default:""`
	Authorize_type      *string  `json:"authorize_type" default:""`
	Currency_conversion *float64 `json:"currency_conversion" default:""`
	Currency_ouom       *string  `json:"currency_ouom" default:""`
	Effective_date      *string  `json:"effective_date" default:""`
	Expiry_date         *string  `json:"expiry_date" default:""`
	Limit_desc          *string  `json:"limit_desc" default:""`
	Limit_type          *string  `json:"limit_type" default:""`
	Partner_ba_id       *string  `json:"partner_ba_id" default:""`
	Ppdm_guid           *string  `json:"ppdm_guid" default:""`
	Price_per_length    *float64 `json:"price_per_length" default:""`
	Reason_type         *string  `json:"reason_type" default:""`
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
