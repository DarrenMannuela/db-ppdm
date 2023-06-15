package dto

import (
	"time"
)

type Ppdm_property_column struct {
	Property_set_id        string     `json:"property_set_id"`
	Property_obs_no        int        `json:"property_obs_no"`
	Active_ind             *string    `json:"active_ind"`
	Column_precision       *float64   `json:"column_precision"`
	Column_scale           *float64   `json:"column_scale"`
	Column_size            *float64   `json:"column_size"`
	Data_type              *string    `json:"data_type"`
	Domain                 *string    `json:"domain"`
	Effective_date         *time.Time `json:"effective_date"`
	Expiry_date            *time.Time `json:"expiry_date"`
	Ppdm_guid              string     `json:"ppdm_guid"`
	Preferred_currency_uom *string    `json:"preferred_currency_uom"`
	Preferred_uom          *string    `json:"preferred_uom"`
	Quantity_type          *string    `json:"quantity_type"`
	Referenced_system_id   *string    `json:"referenced_system_id"`
	Referenced_table_name  *string    `json:"referenced_table_name"`
	Remark                 *string    `json:"remark"`
	Source                 *string    `json:"source"`
	Use_column_name        *string    `json:"use_column_name"`
	Use_indicator_ind      *string    `json:"use_indicator_ind"`
	Use_system_id          *string    `json:"use_system_id"`
	Use_table_name         *string    `json:"use_table_name"`
	Row_changed_by         *string    `json:"row_changed_by"`
	Row_changed_date       *time.Time `json:"row_changed_date"`
	Row_created_by         *string    `json:"row_created_by"`
	Row_created_date       *time.Time `json:"row_created_date"`
	Row_effective_date     *time.Time `json:"row_effective_date"`
	Row_expiry_date        *time.Time `json:"row_expiry_date"`
	Row_quality            *string    `json:"row_quality"`
}
