package dto

type Seis_segment struct {
	Seis_set_subtype      string  `json:"seis_set_subtype" default:"source"`
	Segment_id            string  `json:"segment_id" default:"source"`
	Acqtn_design_id       *string `json:"acqtn_design_id" default:""`
	Active_ind            *string `json:"active_ind" default:""`
	Area_id               *string `json:"area_id" default:""`
	Area_type             *string `json:"area_type" default:""`
	Business_associate_id *string `json:"business_associate_id" default:""`
	Effective_date        *string `json:"effective_date" default:""`
	Expiry_date           *string `json:"expiry_date" default:""`
	Ppdm_guid             *string `json:"ppdm_guid" default:""`
	Refraction_reflection *string `json:"refraction_reflection" default:""`
	Remark                *string `json:"remark" default:""`
	Row_audit_by          *string `json:"row_audit_by" default:""`
	Row_audit_date        *string `json:"row_audit_date" default:""`
	Seis_dimension        *string `json:"seis_dimension" default:""`
	Seis_line_id          *string `json:"seis_line_id" default:""`
	Seis_line_set_subtype *string `json:"seis_line_set_subtype" default:""`
	Seis_segment_reason   *string `json:"seis_segment_reason" default:""`
	Seis_spectrum_type    *string `json:"seis_spectrum_type" default:""`
	Seis_station_type     *string `json:"seis_station_type" default:""`
	Source                *string `json:"source" default:""`
	Test_experimental     *string `json:"test_experimental" default:""`
	Row_changed_by        *string `json:"row_changed_by" default:""`
	Row_changed_date      *string `json:"row_changed_date" default:""`
	Row_created_by        *string `json:"row_created_by" default:""`
	Row_created_date      *string `json:"row_created_date" default:""`
	Row_effective_date    *string `json:"row_effective_date" default:""`
	Row_expiry_date       *string `json:"row_expiry_date" default:""`
	Row_quality           *string `json:"row_quality" default:""`
}
