package dto

type Seis_patch_desc struct {
	Patch_id           string   `json:"patch_id" default:"source"`
	Patch_desc_id      string   `json:"patch_desc_id" default:"source"`
	Active_ind         *string  `json:"active_ind" default:""`
	Effective_date     *string  `json:"effective_date" default:""`
	End_channel        *float64 `json:"end_channel" default:""`
	End_station_id     *string  `json:"end_station_id" default:""`
	End_x_offset       *float64 `json:"end_x_offset" default:""`
	End_y_offset       *float64 `json:"end_y_offset" default:""`
	End_z_offset       *float64 `json:"end_z_offset" default:""`
	Expiry_date        *string  `json:"expiry_date" default:""`
	Offset_ouom        *string  `json:"offset_ouom" default:""`
	Ppdm_guid          *string  `json:"ppdm_guid" default:""`
	Recorded_line      *string  `json:"recorded_line" default:""`
	Remark             *string  `json:"remark" default:""`
	Source             *string  `json:"source" default:""`
	Start_channel      *float64 `json:"start_channel" default:""`
	Start_station_id   *string  `json:"start_station_id" default:""`
	Start_x_offset     *float64 `json:"start_x_offset" default:""`
	Start_y_offset     *float64 `json:"start_y_offset" default:""`
	Start_z_offset     *float64 `json:"start_z_offset" default:""`
	Row_changed_by     *string  `json:"row_changed_by" default:""`
	Row_changed_date   *string  `json:"row_changed_date" default:""`
	Row_created_by     *string  `json:"row_created_by" default:""`
	Row_created_date   *string  `json:"row_created_date" default:""`
	Row_effective_date *string  `json:"row_effective_date" default:""`
	Row_expiry_date    *string  `json:"row_expiry_date" default:""`
	Row_quality        *string  `json:"row_quality" default:""`
}
