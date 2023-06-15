package dto

type Ba_service_address struct {
	Business_associate_id string  `json:"business_associate_id" default:"source"`
	Address_obs_no        int     `json:"address_obs_no" default:"1"`
	Ba_service_type       string  `json:"ba_service_type" default:"source"`
	Ba_service_seq_no     int     `json:"ba_service_seq_no" default:"1"`
	Active_ind            *string `json:"active_ind" default:""`
	Address_source        *string `json:"address_source" default:""`
	Alias_source          *string `json:"alias_source" default:""`
	Ba_alias_id           *string `json:"ba_alias_id" default:""`
	Ba_service_source     *string `json:"ba_service_source" default:""`
	Effective_date        *string `json:"effective_date" default:""`
	Expiry_date           *string `json:"expiry_date" default:""`
	Ppdm_guid             *string `json:"ppdm_guid" default:""`
	Remark                *string `json:"remark" default:""`
	Service_quality       *string `json:"service_quality" default:""`
	Source                *string `json:"source" default:""`
	Row_changed_by        *string `json:"row_changed_by" default:""`
	Row_changed_date      *string `json:"row_changed_date" default:""`
	Row_created_by        *string `json:"row_created_by" default:""`
	Row_created_date      *string `json:"row_created_date" default:""`
	Row_effective_date    *string `json:"row_effective_date" default:""`
	Row_expiry_date       *string `json:"row_expiry_date" default:""`
	Row_quality           *string `json:"row_quality" default:""`
}
