package dto

type Sf_road struct {
	Sf_subtype              string   `json:"sf_subtype" default:"source"`
	Road_id                 string   `json:"road_id" default:"source"`
	Active_ind              *string  `json:"active_ind" default:""`
	Capacity                *float64 `json:"capacity" default:""`
	Capacity_ouom           *string  `json:"capacity_ouom" default:""`
	Communication_freq      *float64 `json:"communication_freq" default:""`
	Communication_freq_desc *string  `json:"communication_freq_desc" default:""`
	Control_type            *string  `json:"control_type" default:""`
	Current_operator_ba_id  *string  `json:"current_operator_ba_id" default:""`
	Direction               *string  `json:"direction" default:""`
	Driving_side            *string  `json:"driving_side" default:""`
	Effective_date          *string  `json:"effective_date" default:""`
	Expiry_date             *string  `json:"expiry_date" default:""`
	Ppdm_guid               *string  `json:"ppdm_guid" default:""`
	Remark                  *string  `json:"remark" default:""`
	Road_length             *float64 `json:"road_length" default:""`
	Road_length_ouom        *string  `json:"road_length_ouom" default:""`
	Road_type               *string  `json:"road_type" default:""`
	Road_width              *float64 `json:"road_width" default:""`
	Road_width_ouom         *string  `json:"road_width_ouom" default:""`
	Source                  *string  `json:"source" default:""`
	Surface_type            *string  `json:"surface_type" default:""`
	Traffic_flow_type       *string  `json:"traffic_flow_type" default:""`
	Row_changed_by          *string  `json:"row_changed_by" default:""`
	Row_changed_date        *string  `json:"row_changed_date" default:""`
	Row_created_by          *string  `json:"row_created_by" default:""`
	Row_created_date        *string  `json:"row_created_date" default:""`
	Row_effective_date      *string  `json:"row_effective_date" default:""`
	Row_expiry_date         *string  `json:"row_expiry_date" default:""`
	Row_quality             *string  `json:"row_quality" default:""`
}
