package dto

type Land_right_instrument struct {
	Land_right_subtype      string  `json:"land_right_subtype" default:"source"`
	Land_right_id           string  `json:"land_right_id" default:"source"`
	Instrument_id           string  `json:"instrument_id" default:"source"`
	Lr_inst_seq_no          int     `json:"lr_inst_seq_no" default:"1"`
	Active_ind              *string `json:"active_ind" default:""`
	Discharge_date          *string `json:"discharge_date" default:""`
	Discharge_ind           *string `json:"discharge_ind" default:""`
	Discharge_instrument_id *string `json:"discharge_instrument_id" default:""`
	Discharge_num           *string `json:"discharge_num" default:""`
	Effective_date          *string `json:"effective_date" default:""`
	Expiry_date             *string `json:"expiry_date" default:""`
	Ppdm_guid               *string `json:"ppdm_guid" default:""`
	Remark                  *string `json:"remark" default:""`
	Source                  *string `json:"source" default:""`
	Row_changed_by          *string `json:"row_changed_by" default:""`
	Row_changed_date        *string `json:"row_changed_date" default:""`
	Row_created_by          *string `json:"row_created_by" default:""`
	Row_created_date        *string `json:"row_created_date" default:""`
	Row_effective_date      *string `json:"row_effective_date" default:""`
	Row_expiry_date         *string `json:"row_expiry_date" default:""`
	Row_quality             *string `json:"row_quality" default:""`
}
