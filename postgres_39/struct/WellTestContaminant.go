package dto

import (
	"time"
)

type Well_test_contaminant struct {
	Uwi                string     `json:"uwi"`
	Source             string     `json:"source"`
	Test_type          string     `json:"test_type"`
	Run_num            string     `json:"run_num"`
	Test_num           string     `json:"test_num"`
	Recovery_obs_no    int        `json:"recovery_obs_no"`
	Contaminant_seq_no int        `json:"contaminant_seq_no"`
	Active_ind         *string    `json:"active_ind"`
	Contaminant_type   *string    `json:"contaminant_type"`
	Effective_date     *time.Time `json:"effective_date"`
	Expiry_date        *time.Time `json:"expiry_date"`
	Ppdm_guid          string     `json:"ppdm_guid"`
	Remark             *string    `json:"remark"`
	Row_changed_by     *string    `json:"row_changed_by"`
	Row_changed_date   *time.Time `json:"row_changed_date"`
	Row_created_by     *string    `json:"row_created_by"`
	Row_created_date   *time.Time `json:"row_created_date"`
	Row_effective_date *time.Time `json:"row_effective_date"`
	Row_expiry_date    *time.Time `json:"row_expiry_date"`
	Row_quality        *string    `json:"row_quality"`
}
