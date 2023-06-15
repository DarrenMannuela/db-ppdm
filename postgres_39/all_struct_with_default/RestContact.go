package dto

type Rest_contact struct {
	Restriction_id        string  `json:"restriction_id" default:"source"`
	Restriction_version   int     `json:"restriction_version" default:"1"`
	Business_associate_id string  `json:"business_associate_id" default:"source"`
	Contact_obs_no        int     `json:"contact_obs_no" default:"1"`
	Active_ind            *string `json:"active_ind" default:""`
	Address_obs_no        *int    `json:"address_obs_no" default:""`
	Address_source        *string `json:"address_source" default:""`
	Effective_date        *string `json:"effective_date" default:""`
	Expiry_date           *string `json:"expiry_date" default:""`
	Phone_num             *string `json:"phone_num" default:""`
	Phone_num_id          *string `json:"phone_num_id" default:""`
	Ppdm_guid             *string `json:"ppdm_guid" default:""`
	Primary_contact_ind   *string `json:"primary_contact_ind" default:""`
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
