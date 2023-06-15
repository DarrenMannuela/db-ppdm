package dto

type Land_termination struct {
	Lr_termination_id      string  `json:"lr_termination_id" default:"source"`
	Lr_termination_seq_no  int     `json:"lr_termination_seq_no" default:"1"`
	Active_ind             *string `json:"active_ind" default:""`
	Business_associate_id  *string `json:"business_associate_id" default:""`
	Description            *string `json:"description" default:""`
	Effective_date         *string `json:"effective_date" default:""`
	Expiry_date            *string `json:"expiry_date" default:""`
	Fulfilled_date         *string `json:"fulfilled_date" default:""`
	Fulfilled_user         *string `json:"fulfilled_user" default:""`
	Jurisdiction           *string `json:"jurisdiction" default:""`
	Land_request_id        *string `json:"land_request_id" default:""`
	Land_right_id          *string `json:"land_right_id" default:""`
	Land_right_subtype     *string `json:"land_right_subtype" default:""`
	Land_sale_number       *string `json:"land_sale_number" default:""`
	Land_sale_offering_id  *string `json:"land_sale_offering_id" default:""`
	Mineral_zone_id        *string `json:"mineral_zone_id" default:""`
	Notification_id        *string `json:"notification_id" default:""`
	Ppdm_guid              *string `json:"ppdm_guid" default:""`
	Remark                 *string `json:"remark" default:""`
	Source                 *string `json:"source" default:""`
	Spatial_description_id *string `json:"spatial_description_id" default:""`
	Spatial_obs_no         *int    `json:"spatial_obs_no" default:""`
	Termination_reqmt      *string `json:"termination_reqmt" default:""`
	Termination_type       *string `json:"termination_type" default:""`
	Row_changed_by         *string `json:"row_changed_by" default:""`
	Row_changed_date       *string `json:"row_changed_date" default:""`
	Row_created_by         *string `json:"row_created_by" default:""`
	Row_created_date       *string `json:"row_created_date" default:""`
	Row_effective_date     *string `json:"row_effective_date" default:""`
	Row_expiry_date        *string `json:"row_expiry_date" default:""`
	Row_quality            *string `json:"row_quality" default:""`
}
