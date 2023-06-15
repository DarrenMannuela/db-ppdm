package dto

type Well_log_job struct {
	Uwi                    string   `json:"uwi" default:"source"`
	Source                 string   `json:"source" default:"source"`
	Job_id                 string   `json:"job_id" default:"source"`
	Active_ind             *string  `json:"active_ind" default:""`
	Area_id                *string  `json:"area_id" default:""`
	Area_type              *string  `json:"area_type" default:""`
	Casing_shoe_depth      *float64 `json:"casing_shoe_depth" default:""`
	Casing_shoe_depth_ouom *string  `json:"casing_shoe_depth_ouom" default:""`
	Drilling_md            *float64 `json:"drilling_md" default:""`
	Drilling_md_ouom       *string  `json:"drilling_md_ouom" default:""`
	Effective_date         *string  `json:"effective_date" default:""`
	End_date               *string  `json:"end_date" default:""`
	Engineer               *string  `json:"engineer" default:""`
	Expiry_date            *string  `json:"expiry_date" default:""`
	Logging_company        *string  `json:"logging_company" default:""`
	Logging_unit           *string  `json:"logging_unit" default:""`
	Observer               *string  `json:"observer" default:""`
	Ppdm_guid              *string  `json:"ppdm_guid" default:""`
	Remark                 *string  `json:"remark" default:""`
	Start_date             *string  `json:"start_date" default:""`
	Row_changed_by         *string  `json:"row_changed_by" default:""`
	Row_changed_date       *string  `json:"row_changed_date" default:""`
	Row_created_by         *string  `json:"row_created_by" default:""`
	Row_created_date       *string  `json:"row_created_date" default:""`
	Row_effective_date     *string  `json:"row_effective_date" default:""`
	Row_expiry_date        *string  `json:"row_expiry_date" default:""`
	Row_quality            *string  `json:"row_quality" default:""`
}
