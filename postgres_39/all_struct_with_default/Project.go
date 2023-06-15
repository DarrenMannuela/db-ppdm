package dto

type Project struct {
	Project_id                string  `json:"project_id" default:"source"`
	Active_ind                *string `json:"active_ind" default:""`
	Complete_date             *string `json:"complete_date" default:""`
	Confidential_ind          *string `json:"confidential_ind" default:""`
	Confidential_release_date *string `json:"confidential_release_date" default:""`
	Description               *string `json:"description" default:""`
	Effective_date            *string `json:"effective_date" default:""`
	Expiry_date               *string `json:"expiry_date" default:""`
	Field_station_ind         *string `json:"field_station_ind" default:""`
	Land_right_ind            *string `json:"land_right_ind" default:""`
	Pden_ind                  *string `json:"pden_ind" default:""`
	Ppdm_guid                 *string `json:"ppdm_guid" default:""`
	Project_name              *string `json:"project_name" default:""`
	Project_num               *string `json:"project_num" default:""`
	Project_type              *string `json:"project_type" default:""`
	Remark                    *string `json:"remark" default:""`
	Seis_set_ind              *string `json:"seis_set_ind" default:""`
	Source                    *string `json:"source" default:""`
	Start_date                *string `json:"start_date" default:""`
	Strat_column_ind          *string `json:"strat_column_ind" default:""`
	Strat_interpretation_ind  *string `json:"strat_interpretation_ind" default:""`
	Well_ind                  *string `json:"well_ind" default:""`
	Row_changed_by            *string `json:"row_changed_by" default:""`
	Row_changed_date          *string `json:"row_changed_date" default:""`
	Row_created_by            *string `json:"row_created_by" default:""`
	Row_created_date          *string `json:"row_created_date" default:""`
	Row_effective_date        *string `json:"row_effective_date" default:""`
	Row_expiry_date           *string `json:"row_expiry_date" default:""`
	Row_quality               *string `json:"row_quality" default:""`
}
