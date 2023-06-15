package dto

type Land_sale_rest_remark struct {
	Restriction_id          string  `json:"restriction_id" default:"source"`
	Restriction_version     int     `json:"restriction_version" default:"1"`
	Jurisdiction            string  `json:"jurisdiction" default:"source"`
	Land_sale_number        string  `json:"land_sale_number" default:"source"`
	Land_sale_offering_id   string  `json:"land_sale_offering_id" default:"source"`
	Restriction_remark_id   string  `json:"restriction_remark_id" default:"source"`
	Active_ind              *string `json:"active_ind" default:""`
	Effective_date          *string `json:"effective_date" default:""`
	Expiry_date             *string `json:"expiry_date" default:""`
	Ppdm_guid               *string `json:"ppdm_guid" default:""`
	Remark                  *string `json:"remark" default:""`
	Restriction_remark_type *string `json:"restriction_remark_type" default:""`
	Source                  *string `json:"source" default:""`
	Row_changed_by          *string `json:"row_changed_by" default:""`
	Row_changed_date        *string `json:"row_changed_date" default:""`
	Row_created_by          *string `json:"row_created_by" default:""`
	Row_created_date        *string `json:"row_created_date" default:""`
	Row_effective_date      *string `json:"row_effective_date" default:""`
	Row_expiry_date         *string `json:"row_expiry_date" default:""`
	Row_quality             *string `json:"row_quality" default:""`
}
