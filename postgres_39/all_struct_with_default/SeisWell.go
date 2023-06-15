package dto

type Seis_well struct {
	Seis_set_subtype        string   `json:"seis_set_subtype" default:"source"`
	Seis_set_id             string   `json:"seis_set_id" default:"source"`
	Active_ind              *string  `json:"active_ind" default:""`
	Checkshot_survey_type   *string  `json:"checkshot_survey_type" default:""`
	Checkshot_velocity      *float64 `json:"checkshot_velocity" default:""`
	Checkshot_velocity_ouom *string  `json:"checkshot_velocity_ouom" default:""`
	Depth_datum             *string  `json:"depth_datum" default:""`
	Depth_datum_elev        *float64 `json:"depth_datum_elev" default:""`
	Depth_datum_elev_ouom   *string  `json:"depth_datum_elev_ouom" default:""`
	Dir_survey_id           *string  `json:"dir_survey_id" default:""`
	Dir_survey_source       *string  `json:"dir_survey_source" default:""`
	Dir_survey_uwi          *string  `json:"dir_survey_uwi" default:""`
	Effective_date          *string  `json:"effective_date" default:""`
	Expiry_date             *string  `json:"expiry_date" default:""`
	Ppdm_guid               *string  `json:"ppdm_guid" default:""`
	Receiver_uwi            *string  `json:"receiver_uwi" default:""`
	Remark                  *string  `json:"remark" default:""`
	Seismic_path            *string  `json:"seismic_path" default:""`
	Source                  *string  `json:"source" default:""`
	Source_uwi              *string  `json:"source_uwi" default:""`
	Survey_id               *string  `json:"survey_id" default:""`
	Survey_run_num          *string  `json:"survey_run_num" default:""`
	Vsp_type                *string  `json:"vsp_type" default:""`
	Row_changed_by          *string  `json:"row_changed_by" default:""`
	Row_changed_date        *string  `json:"row_changed_date" default:""`
	Row_created_by          *string  `json:"row_created_by" default:""`
	Row_created_date        *string  `json:"row_created_date" default:""`
	Row_effective_date      *string  `json:"row_effective_date" default:""`
	Row_expiry_date         *string  `json:"row_expiry_date" default:""`
	Row_quality             *string  `json:"row_quality" default:""`
}
