package dto

type Seis_acqtn_survey struct {
	Seis_set_subtype     string  `json:"seis_set_subtype" default:"source"`
	Seis_acqtn_survey_id string  `json:"seis_acqtn_survey_id" default:"source"`
	Acqtn_survey_name    *string `json:"acqtn_survey_name" default:""`
	Active_ind           *string `json:"active_ind" default:""`
	Area_id              *string `json:"area_id" default:""`
	Area_type            *string `json:"area_type" default:""`
	Completed_date       *string `json:"completed_date" default:""`
	Completed_date_desc  *string `json:"completed_date_desc" default:""`
	Effective_date       *string `json:"effective_date" default:""`
	Expiry_date          *string `json:"expiry_date" default:""`
	Ppdm_guid            *string `json:"ppdm_guid" default:""`
	Remark               *string `json:"remark" default:""`
	Seis_dimension       *string `json:"seis_dimension" default:""`
	Shot_for             *string `json:"shot_for" default:""`
	Source               *string `json:"source" default:""`
	Start_date           *string `json:"start_date" default:""`
	Start_date_desc      *string `json:"start_date_desc" default:""`
	Row_changed_by       *string `json:"row_changed_by" default:""`
	Row_changed_date     *string `json:"row_changed_date" default:""`
	Row_created_by       *string `json:"row_created_by" default:""`
	Row_created_date     *string `json:"row_created_date" default:""`
	Row_effective_date   *string `json:"row_effective_date" default:""`
	Row_expiry_date      *string `json:"row_expiry_date" default:""`
	Row_quality          *string `json:"row_quality" default:""`
}
