package dto

import (
	"time"
)

type Rm_seis_trace struct {
	Info_item_subtype         string     `json:"info_item_subtype"`
	Information_item_id       string     `json:"information_item_id"`
	Active_ind                *string    `json:"active_ind"`
	Datum_vel_correction      *string    `json:"datum_vel_correction"`
	Datum_vel_correction_ouom *string    `json:"datum_vel_correction_ouom"`
	Effective_date            *time.Time `json:"effective_date"`
	Expiry_date               *time.Time `json:"expiry_date"`
	First_sample_time         *float64   `json:"first_sample_time"`
	First_sample_timezone     *string    `json:"first_sample_timezone"`
	Horizontal_scale          *string    `json:"horizontal_scale"`
	Horizontal_scale_ouom     *string    `json:"horizontal_scale_ouom"`
	Phase                     *string    `json:"phase"`
	Polarity                  *string    `json:"polarity"`
	Ppdm_guid                 string     `json:"ppdm_guid"`
	Record_length             *float64   `json:"record_length"`
	Record_length_ouom        *string    `json:"record_length_ouom"`
	Reference_datum           *float64   `json:"reference_datum"`
	Reference_datum_ouom      *string    `json:"reference_datum_ouom"`
	Remark                    *string    `json:"remark"`
	Replacement_vel           *float64   `json:"replacement_vel"`
	Replacement_vel_ouom      *string    `json:"replacement_vel_ouom"`
	Reported_first_ref_num    *string    `json:"reported_first_ref_num"`
	Reported_last_ref_num     *string    `json:"reported_last_ref_num"`
	Reported_ref_num_type     *string    `json:"reported_ref_num_type"`
	Sample_rate               *float64   `json:"sample_rate"`
	Sample_rate_ouom          *string    `json:"sample_rate_ouom"`
	Sample_type               *string    `json:"sample_type"`
	Source                    *string    `json:"source"`
	Static_correction         *float64   `json:"static_correction"`
	Static_correction_ouom    *string    `json:"static_correction_ouom"`
	Variable_area_ind         *string    `json:"variable_area_ind"`
	Vertical_scale            *string    `json:"vertical_scale"`
	Vertical_scale_ouom       *string    `json:"vertical_scale_ouom"`
	Row_changed_by            *string    `json:"row_changed_by"`
	Row_changed_date          *time.Time `json:"row_changed_date"`
	Row_created_by            *string    `json:"row_created_by"`
	Row_created_date          *time.Time `json:"row_created_date"`
	Row_effective_date        *time.Time `json:"row_effective_date"`
	Row_expiry_date           *time.Time `json:"row_expiry_date"`
	Row_quality               *string    `json:"row_quality"`
}
