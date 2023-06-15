package dto

import (
	"time"
)

type Strat_unit_age struct {
	Strat_name_set_id          string     `json:"strat_name_set_id"`
	Strat_unit_id              string     `json:"strat_unit_id"`
	Age_seq_no                 int        `json:"age_seq_no"`
	Active_ind                 *string    `json:"active_ind"`
	Average_age                *float64   `json:"average_age"`
	Average_age_error_minus    *float64   `json:"average_age_error_minus"`
	Average_age_error_plus     *float64   `json:"average_age_error_plus"`
	Average_rel_strat_name_set *string    `json:"average_rel_strat_name_set"`
	Average_rel_strat_unit_id  *string    `json:"average_rel_strat_unit_id"`
	Effective_date             *time.Time `json:"effective_date"`
	Expiry_date                *time.Time `json:"expiry_date"`
	Lower_max_age              *float64   `json:"lower_max_age"`
	Lower_max_age_error_minus  *float64   `json:"lower_max_age_error_minus"`
	Lower_max_age_error_plus   *float64   `json:"lower_max_age_error_plus"`
	Lower_min_age              *float64   `json:"lower_min_age"`
	Lower_min_age_error_minus  *float64   `json:"lower_min_age_error_minus"`
	Lower_min_age_error_plus   *float64   `json:"lower_min_age_error_plus"`
	Lower_rel_strat_name_set   *string    `json:"lower_rel_strat_name_set"`
	Lower_rel_strat_unit_id    *string    `json:"lower_rel_strat_unit_id"`
	Max_age                    *float64   `json:"max_age"`
	Max_age_error_minus        *float64   `json:"max_age_error_minus"`
	Max_age_error_plus         *float64   `json:"max_age_error_plus"`
	Min_age                    *float64   `json:"min_age"`
	Min_age_error_minus        *float64   `json:"min_age_error_minus"`
	Min_age_error_plus         *float64   `json:"min_age_error_plus"`
	Post_qualifier             *string    `json:"post_qualifier"`
	Ppdm_guid                  string     `json:"ppdm_guid"`
	Preferred_ind              *string    `json:"preferred_ind"`
	Pre_qualifier              *string    `json:"pre_qualifier"`
	Remark                     *string    `json:"remark"`
	Source                     *string    `json:"source"`
	Source_document_id         *string    `json:"source_document_id"`
	Strat_age_method           *string    `json:"strat_age_method"`
	Upper_max_age              *float64   `json:"upper_max_age"`
	Upper_max_age_error_minus  *float64   `json:"upper_max_age_error_minus"`
	Upper_max_age_error_plus   *float64   `json:"upper_max_age_error_plus"`
	Upper_min_age              *float64   `json:"upper_min_age"`
	Upper_min_age_error_minus  *float64   `json:"upper_min_age_error_minus"`
	Upper_min_age_error_plus   *float64   `json:"upper_min_age_error_plus"`
	Upper_rel_strat_name_set   *string    `json:"upper_rel_strat_name_set"`
	Upper_rel_strat_unit_id    *string    `json:"upper_rel_strat_unit_id"`
	Row_changed_by             *string    `json:"row_changed_by"`
	Row_changed_date           *time.Time `json:"row_changed_date"`
	Row_created_by             *string    `json:"row_created_by"`
	Row_created_date           *time.Time `json:"row_created_date"`
	Row_effective_date         *time.Time `json:"row_effective_date"`
	Row_expiry_date            *time.Time `json:"row_expiry_date"`
	Row_quality                *string    `json:"row_quality"`
}
