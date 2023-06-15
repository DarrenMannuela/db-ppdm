package dto

type Seis_set struct {
	Seis_set_subtype           string   `json:"seis_set_subtype" default:"source"`
	Seis_set_id                string   `json:"seis_set_id" default:"source"`
	Acqtn_design_id            *string  `json:"acqtn_design_id" default:""`
	Active_ind                 *string  `json:"active_ind" default:""`
	Area_id                    *string  `json:"area_id" default:""`
	Area_size                  *float64 `json:"area_size" default:""`
	Area_size_ouom             *string  `json:"area_size_ouom" default:""`
	Area_type                  *string  `json:"area_type" default:""`
	Coord_acquisition_id       *string  `json:"coord_acquisition_id" default:""`
	Current_seis_status        *string  `json:"current_seis_status" default:""`
	Effective_date             *string  `json:"effective_date" default:""`
	End_date                   *string  `json:"end_date" default:""`
	Expiry_date                *string  `json:"expiry_date" default:""`
	Finance_id                 *string  `json:"finance_id" default:""`
	Geographic_coord_system_id *string  `json:"geographic_coord_system_id" default:""`
	Local_coord_system_id      *string  `json:"local_coord_system_id" default:""`
	Max_latitude               *float64 `json:"max_latitude" default:""`
	Max_longitude              *float64 `json:"max_longitude" default:""`
	Min_latitude               *float64 `json:"min_latitude" default:""`
	Min_longitude              *float64 `json:"min_longitude" default:""`
	Ppdm_guid                  *string  `json:"ppdm_guid" default:""`
	Preferred_name             *string  `json:"preferred_name" default:""`
	Remark                     *string  `json:"remark" default:""`
	Source                     *string  `json:"source" default:""`
	Start_date                 *string  `json:"start_date" default:""`
	Xy_coord_system_id         *string  `json:"xy_coord_system_id" default:""`
	Row_changed_by             *string  `json:"row_changed_by" default:""`
	Row_changed_date           *string  `json:"row_changed_date" default:""`
	Row_created_by             *string  `json:"row_created_by" default:""`
	Row_created_date           *string  `json:"row_created_date" default:""`
	Row_effective_date         *string  `json:"row_effective_date" default:""`
	Row_expiry_date            *string  `json:"row_expiry_date" default:""`
	Row_quality                *string  `json:"row_quality" default:""`
}
