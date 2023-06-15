package dto

type Ba_authority struct {
	Business_associate_id string   `json:"business_associate_id" default:"source"`
	Authority_id          string   `json:"authority_id" default:"source"`
	Active_ind            *string  `json:"active_ind" default:""`
	Authority_limit       *float64 `json:"authority_limit" default:""`
	Authority_type        *string  `json:"authority_type" default:""`
	Authorized_by         *string  `json:"authorized_by" default:""`
	Currency_conversion   *float64 `json:"currency_conversion" default:""`
	Currency_ouom         *string  `json:"currency_ouom" default:""`
	Effective_date        *string  `json:"effective_date" default:""`
	Expiry_date           *string  `json:"expiry_date" default:""`
	Ppdm_guid             *string  `json:"ppdm_guid" default:""`
	Remark                *string  `json:"remark" default:""`
	Represented_ba_id     *string  `json:"represented_ba_id" default:""`
	Source                *string  `json:"source" default:""`
	Row_changed_by        *string  `json:"row_changed_by" default:""`
	Row_changed_date      *string  `json:"row_changed_date" default:""`
	Row_created_by        *string  `json:"row_created_by" default:""`
	Row_created_date      *string  `json:"row_created_date" default:""`
	Row_effective_date    *string  `json:"row_effective_date" default:""`
	Row_expiry_date       *string  `json:"row_expiry_date" default:""`
	Row_quality           *string  `json:"row_quality" default:""`
}
