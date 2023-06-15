package dto

type Pden_status_hist struct {
	Pden_subtype       string   `json:"pden_subtype" default:"source"`
	Pden_id            string   `json:"pden_id" default:"source"`
	Pden_source        string   `json:"pden_source" default:"source"`
	Status_id          string   `json:"status_id" default:"source"`
	Active_ind         *string  `json:"active_ind" default:""`
	Effective_date     *string  `json:"effective_date" default:""`
	End_time           *string  `json:"end_time" default:""`
	Expiry_date        *string  `json:"expiry_date" default:""`
	Percent_capability *float64 `json:"percent_capability" default:""`
	Ppdm_guid          *string  `json:"ppdm_guid" default:""`
	Remark             *string  `json:"remark" default:""`
	Source             *string  `json:"source" default:""`
	Start_time         *string  `json:"start_time" default:""`
	Status             *string  `json:"status" default:""`
	Status_date        *string  `json:"status_date" default:""`
	Status_type        *string  `json:"status_type" default:""`
	Timezone           *string  `json:"timezone" default:""`
	Row_changed_by     *string  `json:"row_changed_by" default:""`
	Row_changed_date   *string  `json:"row_changed_date" default:""`
	Row_created_by     *string  `json:"row_created_by" default:""`
	Row_created_date   *string  `json:"row_created_date" default:""`
	Row_effective_date *string  `json:"row_effective_date" default:""`
	Row_expiry_date    *string  `json:"row_expiry_date" default:""`
	Row_quality        *string  `json:"row_quality" default:""`
}
