package dto

type Ba_contact_info struct {
	Business_associate_id  string  `json:"business_associate_id" default:"source"`
	Location_id            string  `json:"location_id" default:"source"`
	Active_ind             *string `json:"active_ind" default:""`
	Address_obs_no         *int    `json:"address_obs_no" default:""`
	Address_source         *string `json:"address_source" default:""`
	Ba_organization_id     *string `json:"ba_organization_id" default:""`
	Ba_organization_seq_no *int    `json:"ba_organization_seq_no" default:""`
	Contact_loc_type       *string `json:"contact_loc_type" default:""`
	Effective_date         *string `json:"effective_date" default:""`
	Expiry_date            *string `json:"expiry_date" default:""`
	Location_name          *string `json:"location_name" default:""`
	Order_level            *int    `json:"order_level" default:""`
	Ppdm_guid              *string `json:"ppdm_guid" default:""`
	Preferred_ind          *string `json:"preferred_ind" default:""`
	Remark                 *string `json:"remark" default:""`
	Source                 *string `json:"source" default:""`
	Row_changed_by         *string `json:"row_changed_by" default:""`
	Row_changed_date       *string `json:"row_changed_date" default:""`
	Row_created_by         *string `json:"row_created_by" default:""`
	Row_created_date       *string `json:"row_created_date" default:""`
	Row_effective_date     *string `json:"row_effective_date" default:""`
	Row_expiry_date        *string `json:"row_expiry_date" default:""`
	Row_quality            *string `json:"row_quality" default:""`
}
