package dto

import (
	"time"
)

type Land_sale_offering_area struct {
	Jurisdiction          string     `json:"jurisdiction"`
	Land_sale_number      string     `json:"land_sale_number"`
	Land_sale_offering_id string     `json:"land_sale_offering_id"`
	Area_id               string     `json:"area_id"`
	Area_type             string     `json:"area_type"`
	Active_ind            *string    `json:"active_ind"`
	Effective_date        *time.Time `json:"effective_date"`
	Expiry_date           *time.Time `json:"expiry_date"`
	Gross_size            *float64   `json:"gross_size"`
	Gross_size_ouom       *string    `json:"gross_size_ouom"`
	Net_size              *float64   `json:"net_size"`
	Net_size_ouom         *string    `json:"net_size_ouom"`
	Ppdm_guid             string     `json:"ppdm_guid"`
	Remark                *string    `json:"remark"`
	Source                *string    `json:"source"`
	Row_changed_by        *string    `json:"row_changed_by"`
	Row_changed_date      *time.Time `json:"row_changed_date"`
	Row_created_by        *string    `json:"row_created_by"`
	Row_created_date      *time.Time `json:"row_created_date"`
	Row_effective_date    *time.Time `json:"row_effective_date"`
	Row_expiry_date       *time.Time `json:"row_expiry_date"`
	Row_quality           *string    `json:"row_quality"`
}
