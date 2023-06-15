package dto

type Sf_vehicle struct {
	Sf_subtype             string   `json:"sf_subtype" default:"source"`
	Vehicle_id             string   `json:"vehicle_id" default:"source"`
	Active_ind             *string  `json:"active_ind" default:""`
	Current_owner          *string  `json:"current_owner" default:""`
	Effective_date         *string  `json:"effective_date" default:""`
	Engine_size            *float64 `json:"engine_size" default:""`
	Engine_size_ouom       *string  `json:"engine_size_ouom" default:""`
	Expiry_date            *string  `json:"expiry_date" default:""`
	License_expiry_date    *string  `json:"license_expiry_date" default:""`
	License_num            *string  `json:"license_num" default:""`
	License_register_ba_id *string  `json:"license_register_ba_id" default:""`
	Ppdm_guid              *string  `json:"ppdm_guid" default:""`
	Remark                 *string  `json:"remark" default:""`
	Source                 *string  `json:"source" default:""`
	Vehicle_name           *string  `json:"vehicle_name" default:""`
	Vehicle_type           *string  `json:"vehicle_type" default:""`
	Row_changed_by         *string  `json:"row_changed_by" default:""`
	Row_changed_date       *string  `json:"row_changed_date" default:""`
	Row_created_by         *string  `json:"row_created_by" default:""`
	Row_created_date       *string  `json:"row_created_date" default:""`
	Row_effective_date     *string  `json:"row_effective_date" default:""`
	Row_expiry_date        *string  `json:"row_expiry_date" default:""`
	Row_quality            *string  `json:"row_quality" default:""`
}
