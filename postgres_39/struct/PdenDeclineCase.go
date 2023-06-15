package dto

import (
	"time"
)

type Pden_decline_case struct {
	Pden_subtype       string     `json:"pden_subtype"`
	Pden_id            string     `json:"pden_id"`
	Pden_source        string     `json:"pden_source"`
	Product_type       string     `json:"product_type"`
	Case_id            string     `json:"case_id"`
	Active_ind         *string    `json:"active_ind"`
	Case_name          *string    `json:"case_name"`
	Decline_curve_type *string    `json:"decline_curve_type"`
	Decline_type       *string    `json:"decline_type"`
	Effective_date     *time.Time `json:"effective_date"`
	End_date           *time.Time `json:"end_date"`
	Expiry_date        *time.Time `json:"expiry_date"`
	Final_decline      *float64   `json:"final_decline"`
	Final_rate         *float64   `json:"final_rate"`
	Initial_decline    *float64   `json:"initial_decline"`
	Initial_rate       *float64   `json:"initial_rate"`
	Ppdm_guid          string     `json:"ppdm_guid"`
	Project_id         *string    `json:"project_id"`
	Remark             *string    `json:"remark"`
	Source             *string    `json:"source"`
	Start_date         *time.Time `json:"start_date"`
	Volume             *float64   `json:"volume"`
	Volume_ouom        *string    `json:"volume_ouom"`
	Row_changed_by     *string    `json:"row_changed_by"`
	Row_changed_date   *time.Time `json:"row_changed_date"`
	Row_created_by     *string    `json:"row_created_by"`
	Row_created_date   *time.Time `json:"row_created_date"`
	Row_effective_date *time.Time `json:"row_effective_date"`
	Row_expiry_date    *time.Time `json:"row_expiry_date"`
	Row_quality        *string    `json:"row_quality"`
}
