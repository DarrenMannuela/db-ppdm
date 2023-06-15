package dto

type Sf_monument struct {
	Sf_subtype            string   `json:"sf_subtype" default:"source"`
	Monument_id           string   `json:"monument_id" default:"source"`
	Active_ind            *string  `json:"active_ind" default:""`
	Coord_acquisition_id  *string  `json:"coord_acquisition_id" default:""`
	Effective_date        *string  `json:"effective_date" default:""`
	Expiry_date           *string  `json:"expiry_date" default:""`
	Horiz_coord_system    *string  `json:"horiz_coord_system" default:""`
	Local_coord_system_id *string  `json:"local_coord_system_id" default:""`
	Location_type         *string  `json:"location_type" default:""`
	Monument_elev         *float64 `json:"monument_elev" default:""`
	Monument_elev_ouom    *string  `json:"monument_elev_ouom" default:""`
	Monument_latitude     *float64 `json:"monument_latitude" default:""`
	Monument_longitude    *float64 `json:"monument_longitude" default:""`
	Monument_name         *string  `json:"monument_name" default:""`
	Ppdm_guid             *string  `json:"ppdm_guid" default:""`
	Remark                *string  `json:"remark" default:""`
	Source                *string  `json:"source" default:""`
	Source_document_id    *string  `json:"source_document_id" default:""`
	Vert_coord_system     *string  `json:"vert_coord_system" default:""`
	Row_changed_by        *string  `json:"row_changed_by" default:""`
	Row_changed_date      *string  `json:"row_changed_date" default:""`
	Row_created_by        *string  `json:"row_created_by" default:""`
	Row_created_date      *string  `json:"row_created_date" default:""`
	Row_effective_date    *string  `json:"row_effective_date" default:""`
	Row_expiry_date       *string  `json:"row_expiry_date" default:""`
	Row_quality           *string  `json:"row_quality" default:""`
}
