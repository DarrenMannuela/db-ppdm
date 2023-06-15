package dto

type Land_sale_work_bid struct {
	Jurisdiction          string   `json:"jurisdiction" default:"source"`
	Land_sale_number      string   `json:"land_sale_number" default:"source"`
	Land_sale_offering_id string   `json:"land_sale_offering_id" default:"source"`
	Land_offering_bid_id  string   `json:"land_offering_bid_id" default:"source"`
	Work_bid_id           string   `json:"work_bid_id" default:"source"`
	Active_ind            *string  `json:"active_ind" default:""`
	Critical_date         *string  `json:"critical_date" default:""`
	Currency_conversion   *float64 `json:"currency_conversion" default:""`
	Currency_ouom         *string  `json:"currency_ouom" default:""`
	Description           *string  `json:"description" default:""`
	Effective_date        *string  `json:"effective_date" default:""`
	Expiry_date           *string  `json:"expiry_date" default:""`
	Ppdm_guid             *string  `json:"ppdm_guid" default:""`
	Remark                *string  `json:"remark" default:""`
	Source                *string  `json:"source" default:""`
	Work_bid_count        *int     `json:"work_bid_count" default:""`
	Work_bid_count_uom    *string  `json:"work_bid_count_uom" default:""`
	Work_bid_type         *string  `json:"work_bid_type" default:""`
	Work_bid_value        *float64 `json:"work_bid_value" default:""`
	Row_changed_by        *string  `json:"row_changed_by" default:""`
	Row_changed_date      *string  `json:"row_changed_date" default:""`
	Row_created_by        *string  `json:"row_created_by" default:""`
	Row_created_date      *string  `json:"row_created_date" default:""`
	Row_effective_date    *string  `json:"row_effective_date" default:""`
	Row_expiry_date       *string  `json:"row_expiry_date" default:""`
	Row_quality           *string  `json:"row_quality" default:""`
}
