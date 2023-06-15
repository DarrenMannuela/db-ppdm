package dto

import (
	"time"
)

type Seis_proc_plan struct {
	Seis_proc_plan_id  string     `json:"seis_proc_plan_id"`
	Active_ind         *string    `json:"active_ind"`
	Created_by         *string    `json:"created_by"`
	Effective_date     *time.Time `json:"effective_date"`
	Expiry_date        *time.Time `json:"expiry_date"`
	Plan_name          *string    `json:"plan_name"`
	Plan_num           *string    `json:"plan_num"`
	Plan_owner         *string    `json:"plan_owner"`
	Ppdm_guid          string     `json:"ppdm_guid"`
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
