package dto

import (
	"time"
)

type Anl_elemental struct {
	Analysis_id              string     `json:"analysis_id"`
	Anl_source               string     `json:"anl_source"`
	Elemental_anl_obs_no     int        `json:"elemental_anl_obs_no"`
	Active_ind               *string    `json:"active_ind"`
	Ash_element              *float64   `json:"ash_element"`
	Ash_element_ouom         *string    `json:"ash_element_ouom"`
	C_element                *float64   `json:"c_element"`
	C_element_ouom           *string    `json:"c_element_ouom"`
	Effective_date           *time.Time `json:"effective_date"`
	Expiry_date              *time.Time `json:"expiry_date"`
	Fe_element               *float64   `json:"fe_element"`
	Fe_element_ouom          *string    `json:"fe_element_ouom"`
	H_element                *float64   `json:"h_element"`
	H_element_ouom           *string    `json:"h_element_ouom"`
	Measurement_type         *string    `json:"measurement_type"`
	Ni_element               *float64   `json:"ni_element"`
	Ni_element_ouom          *string    `json:"ni_element_ouom"`
	N_element                *float64   `json:"n_element"`
	N_element_ouom           *string    `json:"n_element_ouom"`
	O_element                *float64   `json:"o_element"`
	O_element_ouom           *string    `json:"o_element_ouom"`
	Ppdm_guid                string     `json:"ppdm_guid"`
	Problem_ind              *string    `json:"problem_ind"`
	Remark                   *string    `json:"remark"`
	Reported_h_c_ratio       *float64   `json:"reported_h_c_ratio"`
	Reported_h_c_ratio_ouom  *string    `json:"reported_h_c_ratio_ouom"`
	Reported_ni_c_ratio      *float64   `json:"reported_ni_c_ratio"`
	Reported_ni_c_ratio_ouom *string    `json:"reported_ni_c_ratio_ouom"`
	Reported_n_c_ratio       *float64   `json:"reported_n_c_ratio"`
	Reported_n_c_ratio_ouom  *string    `json:"reported_n_c_ratio_ouom"`
	Reported_o_c_ratio       *float64   `json:"reported_o_c_ratio"`
	Reported_o_c_ratio_ouom  *string    `json:"reported_o_c_ratio_ouom"`
	Source                   *string    `json:"source"`
	Step_seq_no              *int       `json:"step_seq_no"`
	S_element                *float64   `json:"s_element"`
	S_element_ouom           *string    `json:"s_element_ouom"`
	V_element                *float64   `json:"v_element"`
	V_element_ouom           *string    `json:"v_element_ouom"`
	Row_changed_by           *string    `json:"row_changed_by"`
	Row_changed_date         *time.Time `json:"row_changed_date"`
	Row_created_by           *string    `json:"row_created_by"`
	Row_created_date         *time.Time `json:"row_created_date"`
	Row_effective_date       *time.Time `json:"row_effective_date"`
	Row_expiry_date          *time.Time `json:"row_expiry_date"`
	Row_quality              *string    `json:"row_quality"`
}
