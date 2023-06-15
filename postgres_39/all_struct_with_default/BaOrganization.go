package dto

type Ba_organization struct {
	Business_associate_id string  `json:"business_associate_id" default:"source"`
	Organization_id       string  `json:"organization_id" default:"source"`
	Organization_seq_no   int     `json:"organization_seq_no" default:"1"`
	Active_ind            *string `json:"active_ind" default:""`
	Address_obs_no        *int    `json:"address_obs_no" default:""`
	Address_source        *string `json:"address_source" default:""`
	Area_id               *string `json:"area_id" default:""`
	Area_type             *string `json:"area_type" default:""`
	Created_date          *string `json:"created_date" default:""`
	Description           *string `json:"description" default:""`
	Effective_date        *string `json:"effective_date" default:""`
	Expiry_date           *string `json:"expiry_date" default:""`
	Organization_name     *string `json:"organization_name" default:""`
	Organization_type     *string `json:"organization_type" default:""`
	Ppdm_guid             *string `json:"ppdm_guid" default:""`
	Remark                *string `json:"remark" default:""`
	Removed_date          *string `json:"removed_date" default:""`
	Source                *string `json:"source" default:""`
	Row_changed_by        *string `json:"row_changed_by" default:""`
	Row_changed_date      *string `json:"row_changed_date" default:""`
	Row_created_by        *string `json:"row_created_by" default:""`
	Row_created_date      *string `json:"row_created_date" default:""`
	Row_effective_date    *string `json:"row_effective_date" default:""`
	Row_expiry_date       *string `json:"row_expiry_date" default:""`
	Row_quality           *string `json:"row_quality" default:""`
}
