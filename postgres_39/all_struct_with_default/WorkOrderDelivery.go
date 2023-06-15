package dto

type Work_order_delivery struct {
	Work_order_id              string  `json:"work_order_id" default:"source"`
	Delivery_id                string  `json:"delivery_id" default:"source"`
	Active_ind                 *string `json:"active_ind" default:""`
	Ba_obs_no                  *int    `json:"ba_obs_no" default:""`
	Business_associate_id      *string `json:"business_associate_id" default:""`
	Delivery_ba_address_obs_no *int    `json:"delivery_ba_address_obs_no" default:""`
	Delivery_ba_address_source *string `json:"delivery_ba_address_source" default:""`
	Delivery_ba_id             *string `json:"delivery_ba_id" default:""`
	Delivery_date              *string `json:"delivery_date" default:""`
	Delivery_type              *string `json:"delivery_type" default:""`
	Effective_date             *string `json:"effective_date" default:""`
	Expiry_date                *string `json:"expiry_date" default:""`
	Ppdm_guid                  *string `json:"ppdm_guid" default:""`
	Remark                     *string `json:"remark" default:""`
	Source                     *string `json:"source" default:""`
	Row_changed_by             *string `json:"row_changed_by" default:""`
	Row_changed_date           *string `json:"row_changed_date" default:""`
	Row_created_by             *string `json:"row_created_by" default:""`
	Row_created_date           *string `json:"row_created_date" default:""`
	Row_effective_date         *string `json:"row_effective_date" default:""`
	Row_expiry_date            *string `json:"row_expiry_date" default:""`
	Row_quality                *string `json:"row_quality" default:""`
}
