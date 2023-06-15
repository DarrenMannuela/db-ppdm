package dto

import (
	"time"
)

type Sf_disposal struct {
	Sf_subtype          string     `json:"sf_subtype"`
	Support_facility_id string     `json:"support_facility_id"`
	Active_ind          *string    `json:"active_ind"`
	Area_size           *float64   `json:"area_size"`
	Area_size_ouom      *string    `json:"area_size_ouom"`
	Area_size_uom       *string    `json:"area_size_uom"`
	Capacity            *float64   `json:"capacity"`
	Capacity_ouom       *string    `json:"capacity_ouom"`
	Capacity_uom        *string    `json:"capacity_uom"`
	Effective_date      *time.Time `json:"effective_date"`
	Expiry_date         *time.Time `json:"expiry_date"`
	Ppdm_guid           string     `json:"ppdm_guid"`
	Remark              *string    `json:"remark"`
	Source              *string    `json:"source"`
	Row_changed_by      *string    `json:"row_changed_by"`
	Row_changed_date    *time.Time `json:"row_changed_date"`
	Row_created_by      *string    `json:"row_created_by"`
	Row_created_date    *time.Time `json:"row_created_date"`
	Row_effective_date  *time.Time `json:"row_effective_date"`
	Row_expiry_date     *time.Time `json:"row_expiry_date"`
	Row_quality         *string    `json:"row_quality"`
}
