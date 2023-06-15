package dto

type Land_xref struct {
	Parent_land_right_subtype string   `json:"parent_land_right_subtype" default:"source"`
	Child_land_right_subtype  string   `json:"child_land_right_subtype" default:"source"`
	Parent_land_right_id      string   `json:"parent_land_right_id" default:"source"`
	Child_land_right_id       string   `json:"child_land_right_id" default:"source"`
	Lr_xref_seq_no            int      `json:"lr_xref_seq_no" default:"1"`
	Active_ind                *string  `json:"active_ind" default:""`
	Effective_date            *string  `json:"effective_date" default:""`
	Expiry_date               *string  `json:"expiry_date" default:""`
	Land_xref_type            *string  `json:"land_xref_type" default:""`
	Percent_allocation        *float64 `json:"percent_allocation" default:""`
	Ppdm_guid                 *string  `json:"ppdm_guid" default:""`
	Remark                    *string  `json:"remark" default:""`
	Source                    *string  `json:"source" default:""`
	Row_changed_by            *string  `json:"row_changed_by" default:""`
	Row_changed_date          *string  `json:"row_changed_date" default:""`
	Row_created_by            *string  `json:"row_created_by" default:""`
	Row_created_date          *string  `json:"row_created_date" default:""`
	Row_effective_date        *string  `json:"row_effective_date" default:""`
	Row_expiry_date           *string  `json:"row_expiry_date" default:""`
	Row_quality               *string  `json:"row_quality" default:""`
}
