package dto

import (
	"time"
)

type Strat_interp_corr struct {
	Correlation_id             string     `json:"correlation_id"`
	Active_ind                 *string    `json:"active_ind"`
	Area_id                    *string    `json:"area_id"`
	Area_type                  *string    `json:"area_type"`
	Business_associate_id      *string    `json:"business_associate_id"`
	Correlation_date           *time.Time `json:"correlation_date"`
	Effective_date             *time.Time `json:"effective_date"`
	Expiry_date                *time.Time `json:"expiry_date"`
	Field_interp_id_1          *string    `json:"field_interp_id_1"`
	Field_interp_id_2          *string    `json:"field_interp_id_2"`
	Field_station_id_1         *string    `json:"field_station_id_1"`
	Field_station_id_2         *string    `json:"field_station_id_2"`
	Field_strat_name_set_1     *string    `json:"field_strat_name_set_1"`
	Field_strat_name_set_2     *string    `json:"field_strat_name_set_2"`
	Field_strat_unit_id_1      *string    `json:"field_strat_unit_id_1"`
	Field_strat_unit_id_2      *string    `json:"field_strat_unit_id_2"`
	Ppdm_guid                  string     `json:"ppdm_guid"`
	Project_id                 *string    `json:"project_id"`
	Remark                     *string    `json:"remark"`
	Source                     *string    `json:"source"`
	Source_document_id         *string    `json:"source_document_id"`
	Strat_correlation_criteria *string    `json:"strat_correlation_criteria"`
	Strat_correlation_quality  *string    `json:"strat_correlation_quality"`
	Strat_correlation_type     *string    `json:"strat_correlation_type"`
	Strat_interpret_method     *string    `json:"strat_interpret_method"`
	Uwi_1                      *string    `json:"uwi_1"`
	Uwi_2                      *string    `json:"uwi_2"`
	Uwi_interp_id_1            *string    `json:"uwi_interp_id_1"`
	Uwi_interp_id_2            *string    `json:"uwi_interp_id_2"`
	Uwi_strat_name_set_1       *string    `json:"uwi_strat_name_set_1"`
	Uwi_strat_name_set_2       *string    `json:"uwi_strat_name_set_2"`
	Uwi_strat_unit_id_1        *string    `json:"uwi_strat_unit_id_1"`
	Uwi_strat_unit_id_2        *string    `json:"uwi_strat_unit_id_2"`
	Row_changed_by             *string    `json:"row_changed_by"`
	Row_changed_date           *time.Time `json:"row_changed_date"`
	Row_created_by             *string    `json:"row_created_by"`
	Row_created_date           *time.Time `json:"row_created_date"`
	Row_effective_date         *time.Time `json:"row_effective_date"`
	Row_expiry_date            *time.Time `json:"row_expiry_date"`
	Row_quality                *string    `json:"row_quality"`
}
