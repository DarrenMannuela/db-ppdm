package dto

import (
	"time"
)

type Seis_proc_set struct {
	Seis_set_subtype   string     `json:"seis_set_subtype"`
	Seis_proc_set_id   string     `json:"seis_proc_set_id"`
	Active_ind         *string    `json:"active_ind"`
	Area_id            *string    `json:"area_id"`
	Area_type          *string    `json:"area_type"`
	Description        *string    `json:"description"`
	Effective_date     *time.Time `json:"effective_date"`
	End_date           *time.Time `json:"end_date"`
	Expiry_date        *time.Time `json:"expiry_date"`
	Max_latitude       *float64   `json:"max_latitude"`
	Max_longitude      *float64   `json:"max_longitude"`
	Min_latitude       *float64   `json:"min_latitude"`
	Min_longitude      *float64   `json:"min_longitude"`
	Objective          *string    `json:"objective"`
	Original_proc_ind  *string    `json:"original_proc_ind"`
	Ppdm_guid          string     `json:"ppdm_guid"`
	Processed_for      *string    `json:"processed_for"`
	Processing_company *string    `json:"processing_company"`
	Processing_name    *string    `json:"processing_name"`
	Process_status     *string    `json:"process_status"`
	Proc_set_type      *string    `json:"proc_set_type"`
	Remark             *string    `json:"remark"`
	Reprocessed_ind    *string    `json:"reprocessed_ind"`
	Source             *string    `json:"source"`
	Start_date         *time.Time `json:"start_date"`
	Row_changed_by     *string    `json:"row_changed_by"`
	Row_changed_date   *time.Time `json:"row_changed_date"`
	Row_created_by     *string    `json:"row_created_by"`
	Row_created_date   *time.Time `json:"row_created_date"`
	Row_effective_date *time.Time `json:"row_effective_date"`
	Row_expiry_date    *time.Time `json:"row_expiry_date"`
	Row_quality        *string    `json:"row_quality"`
}
