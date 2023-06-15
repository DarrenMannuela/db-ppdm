package dto

type Rate_schedule struct {
	Rate_schedule_id       string  `json:"rate_schedule_id" default:"source"`
	Active_ind             *string `json:"active_ind" default:""`
	Area_id                *string `json:"area_id" default:""`
	Area_type              *string `json:"area_type" default:""`
	Change_notice          *string `json:"change_notice" default:""`
	Contract_id            *string `json:"contract_id" default:""`
	Effective_date         *string `json:"effective_date" default:""`
	Expiry_date            *string `json:"expiry_date" default:""`
	Owner_ba_add_obs_no    *int    `json:"owner_ba_add_obs_no" default:""`
	Owner_ba_add_source    *string `json:"owner_ba_add_source" default:""`
	Owner_ba_id            *string `json:"owner_ba_id" default:""`
	Pay_method             *string `json:"pay_method" default:""`
	Ppdm_guid              *string `json:"ppdm_guid" default:""`
	Preferred_currency_uom *string `json:"preferred_currency_uom" default:""`
	Provision_id           *string `json:"provision_id" default:""`
	Rate_schedule_type     *string `json:"rate_schedule_type" default:""`
	Remark                 *string `json:"remark" default:""`
	Review_frequency       *string `json:"review_frequency" default:""`
	Source                 *string `json:"source" default:""`
	Source_document_id     *string `json:"source_document_id" default:""`
	Spatial_description_id *string `json:"spatial_description_id" default:""`
	Spatial_obs_no         *int    `json:"spatial_obs_no" default:""`
	Row_changed_by         *string `json:"row_changed_by" default:""`
	Row_changed_date       *string `json:"row_changed_date" default:""`
	Row_created_by         *string `json:"row_created_by" default:""`
	Row_created_date       *string `json:"row_created_date" default:""`
	Row_effective_date     *string `json:"row_effective_date" default:""`
	Row_expiry_date        *string `json:"row_expiry_date" default:""`
	Row_quality            *string `json:"row_quality" default:""`
}
