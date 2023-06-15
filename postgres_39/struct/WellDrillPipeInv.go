package dto

import (
	"time"
)

type Well_drill_pipe_inv struct {
	Uwi                         string     `json:"uwi"`
	Period_obs_no               int        `json:"period_obs_no"`
	Inventory_obs_no            int        `json:"inventory_obs_no"`
	Active_ind                  *string    `json:"active_ind"`
	Cat_equip_id                *string    `json:"cat_equip_id"`
	Coupling_type               *string    `json:"coupling_type"`
	Damaged_joint_count         *int       `json:"damaged_joint_count"`
	Effective_date              *time.Time `json:"effective_date"`
	Expiry_date                 *time.Time `json:"expiry_date"`
	Joint_count                 *int       `json:"joint_count"`
	Linear_mass                 *float64   `json:"linear_mass"`
	Linear_mass_ouom            *string    `json:"linear_mass_ouom"`
	Loan_authorized_by_ba_id    *string    `json:"loan_authorized_by_ba_id"`
	Loan_count                  *int       `json:"loan_count"`
	Loan_to_ba_id               *string    `json:"loan_to_ba_id"`
	Lost_count                  *int       `json:"lost_count"`
	Max_joint_outside_diam      *float64   `json:"max_joint_outside_diam"`
	Max_joint_outside_diam_ouom *string    `json:"max_joint_outside_diam_ouom"`
	Max_outside_diameter        *float64   `json:"max_outside_diameter"`
	Max_outside_diameter_ouom   *string    `json:"max_outside_diameter_ouom"`
	Min_inside_diameter         *float64   `json:"min_inside_diameter"`
	Min_inside_diameter_ouom    *string    `json:"min_inside_diameter_ouom"`
	Ppdm_guid                   string     `json:"ppdm_guid"`
	Remark                      *string    `json:"remark"`
	Report_time                 *time.Time `json:"report_time"`
	Report_timezone             *string    `json:"report_timezone"`
	Report_time_type            *string    `json:"report_time_type"`
	Source                      *string    `json:"source"`
	Tubing_grade                *string    `json:"tubing_grade"`
	Tubing_type                 *string    `json:"tubing_type"`
	Row_changed_by              *string    `json:"row_changed_by"`
	Row_changed_date            *time.Time `json:"row_changed_date"`
	Row_created_by              *string    `json:"row_created_by"`
	Row_created_date            *time.Time `json:"row_created_date"`
	Row_effective_date          *time.Time `json:"row_effective_date"`
	Row_expiry_date             *time.Time `json:"row_expiry_date"`
	Row_quality                 *string    `json:"row_quality"`
}
