package dto

type Seis_proc_set struct {
	Seis_set_subtype   string   `json:"seis_set_subtype" default:"source"`
	Seis_proc_set_id   string   `json:"seis_proc_set_id" default:"source"`
	Active_ind         *string  `json:"active_ind" default:""`
	Area_id            *string  `json:"area_id" default:""`
	Area_type          *string  `json:"area_type" default:""`
	Description        *string  `json:"description" default:""`
	Effective_date     *string  `json:"effective_date" default:""`
	End_date           *string  `json:"end_date" default:""`
	Expiry_date        *string  `json:"expiry_date" default:""`
	Max_latitude       *float64 `json:"max_latitude" default:""`
	Max_longitude      *float64 `json:"max_longitude" default:""`
	Min_latitude       *float64 `json:"min_latitude" default:""`
	Min_longitude      *float64 `json:"min_longitude" default:""`
	Objective          *string  `json:"objective" default:""`
	Original_proc_ind  *string  `json:"original_proc_ind" default:""`
	Ppdm_guid          *string  `json:"ppdm_guid" default:""`
	Processed_for      *string  `json:"processed_for" default:""`
	Processing_company *string  `json:"processing_company" default:""`
	Processing_name    *string  `json:"processing_name" default:""`
	Process_status     *string  `json:"process_status" default:""`
	Proc_set_type      *string  `json:"proc_set_type" default:""`
	Remark             *string  `json:"remark" default:""`
	Reprocessed_ind    *string  `json:"reprocessed_ind" default:""`
	Source             *string  `json:"source" default:""`
	Start_date         *string  `json:"start_date" default:""`
	Row_changed_by     *string  `json:"row_changed_by" default:""`
	Row_changed_date   *string  `json:"row_changed_date" default:""`
	Row_created_by     *string  `json:"row_created_by" default:""`
	Row_created_date   *string  `json:"row_created_date" default:""`
	Row_effective_date *string  `json:"row_effective_date" default:""`
	Row_expiry_date    *string  `json:"row_expiry_date" default:""`
	Row_quality        *string  `json:"row_quality" default:""`
}
