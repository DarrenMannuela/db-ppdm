package dto

type Cs_geodetic_datum struct {
	Geodetic_datum           string   `json:"geodetic_datum" default:"source"`
	Active_ind               *string  `json:"active_ind" default:""`
	Datum_origin             *string  `json:"datum_origin" default:""`
	Effective_date           *string  `json:"effective_date" default:""`
	Ellipsoid_id             *string  `json:"ellipsoid_id" default:""`
	Expiry_date              *string  `json:"expiry_date" default:""`
	Geodetic_datum_area_id   *string  `json:"geodetic_datum_area_id" default:""`
	Geodetic_datum_area_type *string  `json:"geodetic_datum_area_type" default:""`
	Geodetic_datum_name      *string  `json:"geodetic_datum_name" default:""`
	Origin_angular_ouom      *string  `json:"origin_angular_ouom" default:""`
	Origin_latitude          *float64 `json:"origin_latitude" default:""`
	Origin_longitude         *float64 `json:"origin_longitude" default:""`
	Origin_name              *string  `json:"origin_name" default:""`
	Ppdm_guid                *string  `json:"ppdm_guid" default:""`
	Remark                   *string  `json:"remark" default:""`
	Source                   *string  `json:"source" default:""`
	Source_document_id       *string  `json:"source_document_id" default:""`
	Row_changed_by           *string  `json:"row_changed_by" default:""`
	Row_changed_date         *string  `json:"row_changed_date" default:""`
	Row_created_by           *string  `json:"row_created_by" default:""`
	Row_created_date         *string  `json:"row_created_date" default:""`
	Row_effective_date       *string  `json:"row_effective_date" default:""`
	Row_expiry_date          *string  `json:"row_expiry_date" default:""`
	Row_quality              *string  `json:"row_quality" default:""`
}
