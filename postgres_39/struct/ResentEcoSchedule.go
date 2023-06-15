package dto

import (
	"time"
)

type Resent_eco_schedule struct {
	Resent_id           string     `json:"resent_id"`
	Reserve_class_id    string     `json:"reserve_class_id"`
	Economics_run_id    string     `json:"economics_run_id"`
	Schedule_id         string     `json:"schedule_id"`
	Active_ind          *string    `json:"active_ind"`
	Economic_schedule   *string    `json:"economic_schedule"`
	Effective_date      *time.Time `json:"effective_date"`
	Expiry_date         *time.Time `json:"expiry_date"`
	Max_value           *float64   `json:"max_value"`
	Max_value_ouom      *string    `json:"max_value_ouom"`
	Max_value_uom       *string    `json:"max_value_uom"`
	Min_value           *float64   `json:"min_value"`
	Min_value_ouom      *string    `json:"min_value_ouom"`
	Min_value_uom       *string    `json:"min_value_uom"`
	Period_type         *string    `json:"period_type"`
	Ppdm_guid           string     `json:"ppdm_guid"`
	Product_ind         *string    `json:"product_ind"`
	Product_type        *string    `json:"product_type"`
	Remark              *string    `json:"remark"`
	Schedule_date       *time.Time `json:"schedule_date"`
	Schedule_date_desc  *string    `json:"schedule_date_desc"`
	Schedule_desc       *string    `json:"schedule_desc"`
	Schedule_value      *float64   `json:"schedule_value"`
	Schedule_value_ouom *string    `json:"schedule_value_ouom"`
	Schedule_value_uom  *string    `json:"schedule_value_uom"`
	Source              *string    `json:"source"`
	Row_changed_by      *string    `json:"row_changed_by"`
	Row_changed_date    *time.Time `json:"row_changed_date"`
	Row_created_by      *string    `json:"row_created_by"`
	Row_created_date    *time.Time `json:"row_created_date"`
	Row_effective_date  *time.Time `json:"row_effective_date"`
	Row_expiry_date     *time.Time `json:"row_expiry_date"`
	Row_quality         *string    `json:"row_quality"`
}
