package dto

type Well_mud_resistivity struct {
	Uwi                          string   `json:"uwi" default:"source"`
	Source                       string   `json:"source" default:"source"`
	Sample_id                    string   `json:"sample_id" default:"source"`
	Resistivity_obs_no           int      `json:"resistivity_obs_no" default:"1"`
	Active_ind                   *string  `json:"active_ind" default:""`
	Effective_date               *string  `json:"effective_date" default:""`
	Expiry_date                  *string  `json:"expiry_date" default:""`
	Mud_resistivity              *float64 `json:"mud_resistivity" default:""`
	Mud_resistivity_ouom         *string  `json:"mud_resistivity_ouom" default:""`
	Ppdm_guid                    *string  `json:"ppdm_guid" default:""`
	Remark                       *string  `json:"remark" default:""`
	Resistivity_temperature      *float64 `json:"resistivity_temperature" default:""`
	Resistivity_temperature_ouom *string  `json:"resistivity_temperature_ouom" default:""`
	Sample_type                  *string  `json:"sample_type" default:""`
	Row_changed_by               *string  `json:"row_changed_by" default:""`
	Row_changed_date             *string  `json:"row_changed_date" default:""`
	Row_created_by               *string  `json:"row_created_by" default:""`
	Row_created_date             *string  `json:"row_created_date" default:""`
	Row_effective_date           *string  `json:"row_effective_date" default:""`
	Row_expiry_date              *string  `json:"row_expiry_date" default:""`
	Row_quality                  *string  `json:"row_quality" default:""`
}
