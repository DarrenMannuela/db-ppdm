package dto

import (
	"time"
)

type Anl_coal_rank struct {
	Coal_rank_id        string     `json:"coal_rank_id"`
	Abbreviation        *string    `json:"abbreviation"`
	Active_ind          *string    `json:"active_ind"`
	Coal_rank_scheme_id *string    `json:"coal_rank_scheme_id"`
	Effective_date      *time.Time `json:"effective_date"`
	Expiry_date         *time.Time `json:"expiry_date"`
	Long_name           *string    `json:"long_name"`
	Max_value           *float64   `json:"max_value"`
	Max_value_ouom      *string    `json:"max_value_ouom"`
	Max_value_uom       *string    `json:"max_value_uom"`
	Min_value           *float64   `json:"min_value"`
	Min_value_ouom      *string    `json:"min_value_ouom"`
	Min_value_uom       *string    `json:"min_value_uom"`
	Ppdm_guid           string     `json:"ppdm_guid"`
	Remark              *string    `json:"remark"`
	Short_name          *string    `json:"short_name"`
	Source              *string    `json:"source"`
	Row_changed_by      *string    `json:"row_changed_by"`
	Row_changed_date    *time.Time `json:"row_changed_date"`
	Row_created_by      *string    `json:"row_created_by"`
	Row_created_date    *time.Time `json:"row_created_date"`
	Row_effective_date  *time.Time `json:"row_effective_date"`
	Row_expiry_date     *time.Time `json:"row_expiry_date"`
	Row_quality         *string    `json:"row_quality"`
}
