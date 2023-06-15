package dto

import (
	"time"
)

type Well_log_parm_class struct {
	Parameter_class_id string     `json:"parameter_class_id"`
	Active_ind         *string    `json:"active_ind"`
	Class_long_name    *string    `json:"class_long_name"`
	Class_short_name   *string    `json:"class_short_name"`
	Effective_date     *time.Time `json:"effective_date"`
	Expiry_date        *time.Time `json:"expiry_date"`
	Ppdm_guid          string     `json:"ppdm_guid"`
	Quantity_id        *string    `json:"quantity_id"`
	Remark             *string    `json:"remark"`
	Source             *string    `json:"source"`
	Row_changed_by     *string    `json:"row_changed_by"`
	Row_changed_date   *time.Time `json:"row_changed_date"`
	Row_created_by     *string    `json:"row_created_by"`
	Row_created_date   *time.Time `json:"row_created_date"`
	Row_effective_date *time.Time `json:"row_effective_date"`
	Row_expiry_date    *time.Time `json:"row_expiry_date"`
	Row_quality        *string    `json:"row_quality"`
}
