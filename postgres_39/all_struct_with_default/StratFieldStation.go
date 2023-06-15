package dto

type Strat_field_station struct {
	Field_station_id   string  `json:"field_station_id" default:"source"`
	Active_ind         *string `json:"active_ind" default:""`
	Air_photo_num      *string `json:"air_photo_num" default:""`
	Area_id            *string `json:"area_id" default:""`
	Area_type          *string `json:"area_type" default:""`
	Complete_date      *string `json:"complete_date" default:""`
	Description        *string `json:"description" default:""`
	Effective_date     *string `json:"effective_date" default:""`
	Expiry_date        *string `json:"expiry_date" default:""`
	Field_station_type *string `json:"field_station_type" default:""`
	Ppdm_guid          *string `json:"ppdm_guid" default:""`
	Remark             *string `json:"remark" default:""`
	Source             *string `json:"source" default:""`
	Source_document_id *string `json:"source_document_id" default:""`
	Start_date         *string `json:"start_date" default:""`
	Row_changed_by     *string `json:"row_changed_by" default:""`
	Row_changed_date   *string `json:"row_changed_date" default:""`
	Row_created_by     *string `json:"row_created_by" default:""`
	Row_created_date   *string `json:"row_created_date" default:""`
	Row_effective_date *string `json:"row_effective_date" default:""`
	Row_expiry_date    *string `json:"row_expiry_date" default:""`
	Row_quality        *string `json:"row_quality" default:""`
}
