package dto

type Land_sale_offering struct {
	Jurisdiction          string   `json:"jurisdiction" default:"source"`
	Land_sale_number      string   `json:"land_sale_number" default:"source"`
	Land_sale_offering_id string   `json:"land_sale_offering_id" default:"source"`
	Active_ind            *string  `json:"active_ind" default:""`
	Cancel_reason         *string  `json:"cancel_reason" default:""`
	Effective_date        *string  `json:"effective_date" default:""`
	Expiry_date           *string  `json:"expiry_date" default:""`
	Gross_size            *float64 `json:"gross_size" default:""`
	Gross_size_ouom       *string  `json:"gross_size_ouom" default:""`
	Inactivation_date     *string  `json:"inactivation_date" default:""`
	Land_offering_status  *string  `json:"land_offering_status" default:""`
	Land_offering_type    *string  `json:"land_offering_type" default:""`
	Ppdm_guid             *string  `json:"ppdm_guid" default:""`
	Remark                *string  `json:"remark" default:""`
	Source                *string  `json:"source" default:""`
	Status_date           *string  `json:"status_date" default:""`
	Term_duration         *float64 `json:"term_duration" default:""`
	Term_duration_ouom    *string  `json:"term_duration_ouom" default:""`
	Row_changed_by        *string  `json:"row_changed_by" default:""`
	Row_changed_date      *string  `json:"row_changed_date" default:""`
	Row_created_by        *string  `json:"row_created_by" default:""`
	Row_created_date      *string  `json:"row_created_date" default:""`
	Row_effective_date    *string  `json:"row_effective_date" default:""`
	Row_expiry_date       *string  `json:"row_expiry_date" default:""`
	Row_quality           *string  `json:"row_quality" default:""`
}
