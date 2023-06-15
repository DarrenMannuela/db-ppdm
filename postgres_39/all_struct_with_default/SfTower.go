package dto

type Sf_tower struct {
	Sf_subtype              string   `json:"sf_subtype" default:"source"`
	Tower_id                string   `json:"tower_id" default:"source"`
	Active_ind              *string  `json:"active_ind" default:""`
	Communication_freq      *float64 `json:"communication_freq" default:""`
	Communication_freq_desc *string  `json:"communication_freq_desc" default:""`
	Effective_date          *string  `json:"effective_date" default:""`
	Expiry_date             *string  `json:"expiry_date" default:""`
	Ppdm_guid               *string  `json:"ppdm_guid" default:""`
	Remark                  *string  `json:"remark" default:""`
	Source                  *string  `json:"source" default:""`
	Tower_type              *string  `json:"tower_type" default:""`
	Row_changed_by          *string  `json:"row_changed_by" default:""`
	Row_changed_date        *string  `json:"row_changed_date" default:""`
	Row_created_by          *string  `json:"row_created_by" default:""`
	Row_created_date        *string  `json:"row_created_date" default:""`
	Row_effective_date      *string  `json:"row_effective_date" default:""`
	Row_expiry_date         *string  `json:"row_expiry_date" default:""`
	Row_quality             *string  `json:"row_quality" default:""`
}
