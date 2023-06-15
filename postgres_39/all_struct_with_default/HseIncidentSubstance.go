package dto

type Hse_incident_substance struct {
	Incident_id        string   `json:"incident_id" default:"source"`
	Incident_substance string   `json:"incident_substance" default:"source"`
	Substance_seq_no   int      `json:"substance_seq_no" default:"1"`
	Active_ind         *string  `json:"active_ind" default:""`
	Effective_date     *string  `json:"effective_date" default:""`
	Expiry_date        *string  `json:"expiry_date" default:""`
	Ppdm_guid          *string  `json:"ppdm_guid" default:""`
	Remark             *string  `json:"remark" default:""`
	Source             *string  `json:"source" default:""`
	Substance_role     *string  `json:"substance_role" default:""`
	Volume             *float64 `json:"volume" default:""`
	Volume_ouom        *string  `json:"volume_ouom" default:""`
	Volume_uom         *string  `json:"volume_uom" default:""`
	Row_changed_by     *string  `json:"row_changed_by" default:""`
	Row_changed_date   *string  `json:"row_changed_date" default:""`
	Row_created_by     *string  `json:"row_created_by" default:""`
	Row_created_date   *string  `json:"row_created_date" default:""`
	Row_effective_date *string  `json:"row_effective_date" default:""`
	Row_expiry_date    *string  `json:"row_expiry_date" default:""`
	Row_quality        *string  `json:"row_quality" default:""`
}
