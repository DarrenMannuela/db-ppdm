package dto

type Ba_permit struct {
	Business_associate_id string  `json:"business_associate_id" default:"source"`
	Jurisdiction          string  `json:"jurisdiction" default:"source"`
	Permit_type           string  `json:"permit_type" default:"source"`
	Permit_obs_no         int     `json:"permit_obs_no" default:"1"`
	Active_ind            *string `json:"active_ind" default:""`
	Effective_date        *string `json:"effective_date" default:""`
	Expiry_date           *string `json:"expiry_date" default:""`
	Permit_num            *string `json:"permit_num" default:""`
	Ppdm_guid             *string `json:"ppdm_guid" default:""`
	Remark                *string `json:"remark" default:""`
	Source                *string `json:"source" default:""`
	Row_changed_by        *string `json:"row_changed_by" default:""`
	Row_changed_date      *string `json:"row_changed_date" default:""`
	Row_created_by        *string `json:"row_created_by" default:""`
	Row_created_date      *string `json:"row_created_date" default:""`
	Row_effective_date    *string `json:"row_effective_date" default:""`
	Row_expiry_date       *string `json:"row_expiry_date" default:""`
	Row_quality           *string `json:"row_quality" default:""`
}
