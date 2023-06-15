package dto

type Legal_loc_remark struct {
	Legal_loc_remark_id    string  `json:"legal_loc_remark_id" default:"source"`
	Location_type          string  `json:"location_type" default:"source"`
	Source                 string  `json:"source" default:"source"`
	Remark_seq_no          int     `json:"remark_seq_no" default:"1"`
	Active_ind             *string `json:"active_ind" default:""`
	Effective_date         *string `json:"effective_date" default:""`
	Expiry_date            *string `json:"expiry_date" default:""`
	Land_parcel_type       *string `json:"land_parcel_type" default:""`
	Legal_carter_loc_id    *string `json:"legal_carter_loc_id" default:""`
	Legal_congress_loc_id  *string `json:"legal_congress_loc_id" default:""`
	Legal_dls_loc_id       *string `json:"legal_dls_loc_id" default:""`
	Legal_fps_loc_id       *string `json:"legal_fps_loc_id" default:""`
	Legal_geodetic_loc_id  *string `json:"legal_geodetic_loc_id" default:""`
	Legal_ne_loc_id        *string `json:"legal_ne_loc_id" default:""`
	Legal_north_sea_loc_id *string `json:"legal_north_sea_loc_id" default:""`
	Legal_nts_loc_id       *string `json:"legal_nts_loc_id" default:""`
	Legal_offshore_loc_id  *string `json:"legal_offshore_loc_id" default:""`
	Legal_ohio_loc_id      *string `json:"legal_ohio_loc_id" default:""`
	Legal_texas_loc_id     *string `json:"legal_texas_loc_id" default:""`
	Ppdm_guid              *string `json:"ppdm_guid" default:""`
	Remark                 *string `json:"remark" default:""`
	Remark_type            *string `json:"remark_type" default:""`
	Well_node_id           *string `json:"well_node_id" default:""`
	Row_changed_by         *string `json:"row_changed_by" default:""`
	Row_changed_date       *string `json:"row_changed_date" default:""`
	Row_created_by         *string `json:"row_created_by" default:""`
	Row_created_date       *string `json:"row_created_date" default:""`
	Row_effective_date     *string `json:"row_effective_date" default:""`
	Row_expiry_date        *string `json:"row_expiry_date" default:""`
	Row_quality            *string `json:"row_quality" default:""`
}
