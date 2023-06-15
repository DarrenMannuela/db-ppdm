package dto

type Seis_point_flow_desc struct {
	Seis_set_subtype   string  `json:"seis_set_subtype" default:"source"`
	Seis_set_id        string  `json:"seis_set_id" default:"source"`
	Seis_point_id      string  `json:"seis_point_id" default:"source"`
	Flow_id            string  `json:"flow_id" default:"source"`
	Active_ind         *string `json:"active_ind" default:""`
	Description        *string `json:"description" default:""`
	Description_type   *string `json:"description_type" default:""`
	Effective_date     *string `json:"effective_date" default:""`
	Expiry_date        *string `json:"expiry_date" default:""`
	Ppdm_guid          *string `json:"ppdm_guid" default:""`
	Remark             *string `json:"remark" default:""`
	Source             *string `json:"source" default:""`
	Row_changed_by     *string `json:"row_changed_by" default:""`
	Row_changed_date   *string `json:"row_changed_date" default:""`
	Row_created_by     *string `json:"row_created_by" default:""`
	Row_created_date   *string `json:"row_created_date" default:""`
	Row_effective_date *string `json:"row_effective_date" default:""`
	Row_expiry_date    *string `json:"row_expiry_date" default:""`
	Row_quality        *string `json:"row_quality" default:""`
}
