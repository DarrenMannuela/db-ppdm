package dto

type Seis_channel struct {
	Seis_set_subtype       string   `json:"seis_set_subtype" default:"source"`
	Seis_set_id            string   `json:"seis_set_id" default:"source"`
	Record_id              string   `json:"record_id" default:"source"`
	Channel_id             string   `json:"channel_id" default:"source"`
	Active_ind             *string  `json:"active_ind" default:""`
	Channel_num            *string  `json:"channel_num" default:""`
	Channel_type           *string  `json:"channel_type" default:""`
	Effective_date         *string  `json:"effective_date" default:""`
	Expiry_date            *string  `json:"expiry_date" default:""`
	File_num               *string  `json:"file_num" default:""`
	Ppdm_guid              *string  `json:"ppdm_guid" default:""`
	Remark                 *string  `json:"remark" default:""`
	Seis_receiver_point_id *string  `json:"seis_receiver_point_id" default:""`
	Source                 *string  `json:"source" default:""`
	Streamer_id            *string  `json:"streamer_id" default:""`
	Vessel_config_obs_no   *int     `json:"vessel_config_obs_no" default:""`
	Vessel_id              *string  `json:"vessel_id" default:""`
	Vessel_sf_subtype      *string  `json:"vessel_sf_subtype" default:""`
	X_offset               *float64 `json:"x_offset" default:""`
	X_offset_ouom          *string  `json:"x_offset_ouom" default:""`
	Y_offset               *float64 `json:"y_offset" default:""`
	Y_offset_ouom          *string  `json:"y_offset_ouom" default:""`
	Z_offset               *float64 `json:"z_offset" default:""`
	Z_offset_ouom          *string  `json:"z_offset_ouom" default:""`
	Row_changed_by         *string  `json:"row_changed_by" default:""`
	Row_changed_date       *string  `json:"row_changed_date" default:""`
	Row_created_by         *string  `json:"row_created_by" default:""`
	Row_created_date       *string  `json:"row_created_date" default:""`
	Row_effective_date     *string  `json:"row_effective_date" default:""`
	Row_expiry_date        *string  `json:"row_expiry_date" default:""`
	Row_quality            *string  `json:"row_quality" default:""`
}
