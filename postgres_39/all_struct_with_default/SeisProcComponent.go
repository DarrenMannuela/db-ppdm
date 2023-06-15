package dto

type Seis_proc_component struct {
	Seis_set_subtype          string  `json:"seis_set_subtype" default:"source"`
	Seis_proc_set_id          string  `json:"seis_proc_set_id" default:"source"`
	Component_id              string  `json:"component_id" default:"source"`
	Active_ind                *string `json:"active_ind" default:""`
	Bin_grid_id               *string `json:"bin_grid_id" default:""`
	Bin_grid_source           *string `json:"bin_grid_source" default:""`
	Comp_seis_set_id          *string `json:"comp_seis_set_id" default:""`
	Comp_seis_set_subtype     *string `json:"comp_seis_set_subtype" default:""`
	Effective_date            *string `json:"effective_date" default:""`
	Expiry_date               *string `json:"expiry_date" default:""`
	Information_item_id       *string `json:"information_item_id" default:""`
	Info_item_subtype         *string `json:"info_item_subtype" default:""`
	Input_ind                 *string `json:"input_ind" default:""`
	Output_ind                *string `json:"output_ind" default:""`
	Ppdm_guid                 *string `json:"ppdm_guid" default:""`
	Processing_component_type *string `json:"processing_component_type" default:""`
	Remark                    *string `json:"remark" default:""`
	Seis_proc_plan_id         *string `json:"seis_proc_plan_id" default:""`
	Source                    *string `json:"source" default:""`
	Uwi                       *string `json:"uwi" default:""`
	Velocity_volume_id        *string `json:"velocity_volume_id" default:""`
	Well_log_curve_id         *string `json:"well_log_curve_id" default:""`
	Row_changed_by            *string `json:"row_changed_by" default:""`
	Row_changed_date          *string `json:"row_changed_date" default:""`
	Row_created_by            *string `json:"row_created_by" default:""`
	Row_created_date          *string `json:"row_created_date" default:""`
	Row_effective_date        *string `json:"row_effective_date" default:""`
	Row_expiry_date           *string `json:"row_expiry_date" default:""`
	Row_quality               *string `json:"row_quality" default:""`
}
