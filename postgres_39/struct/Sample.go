package dto

import (
	"time"
)

type Sample struct {
	Sample_id               string     `json:"sample_id"`
	Active_ind              *string    `json:"active_ind"`
	Base_md                 *float64   `json:"base_md"`
	Base_md_ouom            *string    `json:"base_md_ouom"`
	Base_strat_unit_id      *string    `json:"base_strat_unit_id"`
	Collected_by_ba_id      *string    `json:"collected_by_ba_id"`
	Collected_date          *time.Time `json:"collected_date"`
	Collected_equip_id      *string    `json:"collected_equip_id"`
	Collected_for_ba_id     *string    `json:"collected_for_ba_id"`
	Collected_time          *time.Time `json:"collected_time"`
	Collection_phase_type   *string    `json:"collection_phase_type"`
	Collection_press        *float64   `json:"collection_press"`
	Collection_press_ouom   *string    `json:"collection_press_ouom"`
	Collection_temp         *float64   `json:"collection_temp"`
	Collection_temp_ouom    *string    `json:"collection_temp_ouom"`
	Collection_test_type    *string    `json:"collection_test_type"`
	Collection_type         *string    `json:"collection_type"`
	Composite_sample_count  *int       `json:"composite_sample_count"`
	Composite_sample_ind    *string    `json:"composite_sample_ind"`
	Condensed_section_ind   *string    `json:"condensed_section_ind"`
	Current_length          *float64   `json:"current_length"`
	Current_length_ouom     *string    `json:"current_length_ouom"`
	Current_length_percent  *float64   `json:"current_length_percent"`
	Current_volume          *float64   `json:"current_volume"`
	Current_volume_ouom     *string    `json:"current_volume_ouom"`
	Current_volume_percent  *float64   `json:"current_volume_percent"`
	Current_weight          *float64   `json:"current_weight"`
	Current_weight_ouom     *string    `json:"current_weight_ouom"`
	Current_weight_percent  *float64   `json:"current_weight_percent"`
	Cylinder_press          *float64   `json:"cylinder_press"`
	Cylinder_press_ouom     *string    `json:"cylinder_press_ouom"`
	Cylinder_temp           *float64   `json:"cylinder_temp"`
	Cylinder_temp_ouom      *string    `json:"cylinder_temp_ouom"`
	Destroyed_date          *time.Time `json:"destroyed_date"`
	Effective_date          *time.Time `json:"effective_date"`
	Expiry_date             *time.Time `json:"expiry_date"`
	Initial_length          *float64   `json:"initial_length"`
	Initial_length_ouom     *string    `json:"initial_length_ouom"`
	Initial_length_percent  *float64   `json:"initial_length_percent"`
	Initial_volume          *float64   `json:"initial_volume"`
	Initial_volume_ouom     *string    `json:"initial_volume_ouom"`
	Initial_volume_percent  *float64   `json:"initial_volume_percent"`
	Initial_weight          *float64   `json:"initial_weight"`
	Initial_weight_ouom     *string    `json:"initial_weight_ouom"`
	Initial_weight_percent  *float64   `json:"initial_weight_percent"`
	Interval_percent        *float64   `json:"interval_percent"`
	Oil_stained_ind         *string    `json:"oil_stained_ind"`
	Physical_item_id        *string    `json:"physical_item_id"`
	Ppdm_guid               string     `json:"ppdm_guid"`
	Remark                  *string    `json:"remark"`
	Reported_mud_desc       *string    `json:"reported_mud_desc"`
	Reported_res_press      *float64   `json:"reported_res_press"`
	Reported_res_press_ouom *string    `json:"reported_res_press_ouom"`
	Sample_count            *int       `json:"sample_count"`
	Sample_diameter         *float64   `json:"sample_diameter"`
	Sample_diameter_ouom    *string    `json:"sample_diameter_ouom"`
	Sample_interval         *float64   `json:"sample_interval"`
	Sample_interval_ouom    *string    `json:"sample_interval_ouom"`
	Sample_quality          *string    `json:"sample_quality"`
	Sample_quality_desc     *string    `json:"sample_quality_desc"`
	Sample_ref_num          *string    `json:"sample_ref_num"`
	Sample_shape_type       *string    `json:"sample_shape_type"`
	Sample_type             *string    `json:"sample_type"`
	Sample_width            *float64   `json:"sample_width"`
	Sample_width_ouom       *string    `json:"sample_width_ouom"`
	Separator_press         *float64   `json:"separator_press"`
	Separator_press_ouom    *string    `json:"separator_press_ouom"`
	Separator_temp          *float64   `json:"separator_temp"`
	Separator_temp_ouom     *string    `json:"separator_temp_ouom"`
	Source                  *string    `json:"source"`
	Stored_equip_id         *string    `json:"stored_equip_id"`
	Stored_phase_type       *string    `json:"stored_phase_type"`
	Strat_name_set_id       *string    `json:"strat_name_set_id"`
	Sub_sample_ind          *string    `json:"sub_sample_ind"`
	Top_md                  *float64   `json:"top_md"`
	Top_md_ouom             *string    `json:"top_md_ouom"`
	Top_strat_unit_id       *string    `json:"top_strat_unit_id"`
	Total_length            *float64   `json:"total_length"`
	Total_length_ouom       *string    `json:"total_length_ouom"`
	Total_volume            *float64   `json:"total_volume"`
	Total_volume_ouom       *string    `json:"total_volume_ouom"`
	Total_weight            *float64   `json:"total_weight"`
	Total_weight_ouom       *string    `json:"total_weight_ouom"`
	Well_datum_type         *string    `json:"well_datum_type"`
	Row_changed_by          *string    `json:"row_changed_by"`
	Row_changed_date        *time.Time `json:"row_changed_date"`
	Row_created_by          *string    `json:"row_created_by"`
	Row_created_date        *time.Time `json:"row_created_date"`
	Row_effective_date      *time.Time `json:"row_effective_date"`
	Row_expiry_date         *time.Time `json:"row_expiry_date"`
	Row_quality             *string    `json:"row_quality"`
}
