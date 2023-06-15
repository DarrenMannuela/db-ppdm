package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetWell(c *fiber.Ctx) error {
	rows, err := db.Query("select * from well")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Well

	for rows.Next() {
		var well dto.Well
		if err := rows.Scan(&well.Uwi, &well.Abandonment_date, &well.Active_ind, &well.Area_id, &well.Area_type, &well.Assigned_field, &well.Base_depth, &well.Base_depth_ouom, &well.Base_node_id, &well.Base_strat_name_set_id, &well.Base_strat_unit_id, &well.Bottom_hole_latitude, &well.Bottom_hole_longitude, &well.Casing_flange_elev, &well.Casing_flange_elev_ouom, &well.Completion_date, &well.Confidential_date, &well.Confidential_depth, &well.Confidential_depth_ouom, &well.Confidential_type, &well.Confid_strat_name_set_id, &well.Confid_strat_unit_id, &well.Current_class, &well.Current_status, &well.Current_status_date, &well.Deepest_depth, &well.Deepest_depth_ouom, &well.Depth_datum, &well.Depth_datum_elev, &well.Depth_datum_elev_ouom, &well.Derrick_floor_elev, &well.Derrick_floor_elev_ouom, &well.Difference_lat_msl, &well.Discovery_ind, &well.Drill_td, &well.Drill_td_ouom, &well.Effective_date, &well.Elev_ref_datum, &well.Environment_type, &well.Expiry_date, &well.Faulted_ind, &well.Final_drill_date, &well.Final_td, &well.Final_td_ouom, &well.Ground_elev, &well.Ground_elev_ouom, &well.Ground_elev_type, &well.Initial_class, &well.Kb_elev, &well.Kb_elev_ouom, &well.Lease_name, &well.Lease_num, &well.Legal_survey_type, &well.Location_type, &well.Log_td, &well.Log_td_ouom, &well.Max_tvd, &well.Max_tvd_ouom, &well.Net_pay, &well.Net_pay_ouom, &well.Oldest_strat_age, &well.Oldest_strat_name_set_age, &well.Oldest_strat_name_set_id, &well.Oldest_strat_unit_id, &well.Operator, &well.Platform_id, &well.Platform_sf_subtype, &well.Plot_name, &well.Plot_symbol, &well.Plugback_depth, &well.Plugback_depth_ouom, &well.Ppdm_guid, &well.Primary_source, &well.Profile_type, &well.Regulatory_agency, &well.Remark, &well.Rig_on_site_date, &well.Rig_release_date, &well.Rotary_table_elev, &well.Rotary_table_elev_ouom, &well.Source_document_id, &well.Spud_date, &well.Status_type, &well.Subsea_elev_ref_type, &well.Surface_latitude, &well.Surface_longitude, &well.Surface_node_id, &well.Tax_credit_code, &well.Td_strat_age, &well.Td_strat_name_set_age, &well.Td_strat_name_set_id, &well.Td_strat_unit_id, &well.Top_depth, &well.Top_depth_ouom, &well.Top_strat_name_set_id, &well.Top_strat_unit_id, &well.Water_acoustic_vel, &well.Water_acoustic_vel_ouom, &well.Water_depth, &well.Water_depth_datum, &well.Water_depth_ouom, &well.Well_event_num, &well.Well_government_id, &well.Well_intersect_md, &well.Well_level_type, &well.Well_name, &well.Well_num, &well.Whipstock_depth, &well.Whipstock_depth_ouom, &well.Row_changed_by, &well.Row_changed_date, &well.Row_created_by, &well.Row_created_date, &well.Row_effective_date, &well.Row_expiry_date, &well.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append well to result
		result = append(result, well)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetWell(c *fiber.Ctx) error {
	var well dto.Well

	setDefaults(&well)

	if err := json.Unmarshal(c.Body(), &well); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into well values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32, :33, :34, :35, :36, :37, :38, :39, :40, :41, :42, :43, :44, :45, :46, :47, :48, :49, :50, :51, :52, :53, :54, :55, :56, :57, :58, :59, :60, :61, :62, :63, :64, :65, :66, :67, :68, :69, :70, :71, :72, :73, :74, :75, :76, :77, :78, :79, :80, :81, :82, :83, :84, :85, :86, :87, :88, :89, :90, :91, :92, :93, :94, :95, :96, :97, :98, :99, :100, :101, :102, :103, :104, :105, :106, :107, :108, :109, :110, :111, :112, :113, :114, :115, :116)")
	if err != nil {
		return err
	}
	well.Row_created_date = formatDate(well.Row_created_date)
	well.Row_changed_date = formatDate(well.Row_changed_date)
	well.Abandonment_date = formatDateString(well.Abandonment_date)
	well.Completion_date = formatDateString(well.Completion_date)
	well.Confidential_date = formatDateString(well.Confidential_date)
	well.Current_status_date = formatDateString(well.Current_status_date)
	well.Effective_date = formatDateString(well.Effective_date)
	well.Expiry_date = formatDateString(well.Expiry_date)
	well.Final_drill_date = formatDateString(well.Final_drill_date)
	well.Rig_on_site_date = formatDateString(well.Rig_on_site_date)
	well.Rig_release_date = formatDateString(well.Rig_release_date)
	well.Spud_date = formatDateString(well.Spud_date)
	well.Row_effective_date = formatDateString(well.Row_effective_date)
	well.Row_expiry_date = formatDateString(well.Row_expiry_date)






	rows, err := stmt.Exec(well.Uwi, well.Abandonment_date, well.Active_ind, well.Area_id, well.Area_type, well.Assigned_field, well.Base_depth, well.Base_depth_ouom, well.Base_node_id, well.Base_strat_name_set_id, well.Base_strat_unit_id, well.Bottom_hole_latitude, well.Bottom_hole_longitude, well.Casing_flange_elev, well.Casing_flange_elev_ouom, well.Completion_date, well.Confidential_date, well.Confidential_depth, well.Confidential_depth_ouom, well.Confidential_type, well.Confid_strat_name_set_id, well.Confid_strat_unit_id, well.Current_class, well.Current_status, well.Current_status_date, well.Deepest_depth, well.Deepest_depth_ouom, well.Depth_datum, well.Depth_datum_elev, well.Depth_datum_elev_ouom, well.Derrick_floor_elev, well.Derrick_floor_elev_ouom, well.Difference_lat_msl, well.Discovery_ind, well.Drill_td, well.Drill_td_ouom, well.Effective_date, well.Elev_ref_datum, well.Environment_type, well.Expiry_date, well.Faulted_ind, well.Final_drill_date, well.Final_td, well.Final_td_ouom, well.Ground_elev, well.Ground_elev_ouom, well.Ground_elev_type, well.Initial_class, well.Kb_elev, well.Kb_elev_ouom, well.Lease_name, well.Lease_num, well.Legal_survey_type, well.Location_type, well.Log_td, well.Log_td_ouom, well.Max_tvd, well.Max_tvd_ouom, well.Net_pay, well.Net_pay_ouom, well.Oldest_strat_age, well.Oldest_strat_name_set_age, well.Oldest_strat_name_set_id, well.Oldest_strat_unit_id, well.Operator, well.Platform_id, well.Platform_sf_subtype, well.Plot_name, well.Plot_symbol, well.Plugback_depth, well.Plugback_depth_ouom, well.Ppdm_guid, well.Primary_source, well.Profile_type, well.Regulatory_agency, well.Remark, well.Rig_on_site_date, well.Rig_release_date, well.Rotary_table_elev, well.Rotary_table_elev_ouom, well.Source_document_id, well.Spud_date, well.Status_type, well.Subsea_elev_ref_type, well.Surface_latitude, well.Surface_longitude, well.Surface_node_id, well.Tax_credit_code, well.Td_strat_age, well.Td_strat_name_set_age, well.Td_strat_name_set_id, well.Td_strat_unit_id, well.Top_depth, well.Top_depth_ouom, well.Top_strat_name_set_id, well.Top_strat_unit_id, well.Water_acoustic_vel, well.Water_acoustic_vel_ouom, well.Water_depth, well.Water_depth_datum, well.Water_depth_ouom, well.Well_event_num, well.Well_government_id, well.Well_intersect_md, well.Well_level_type, well.Well_name, well.Well_num, well.Whipstock_depth, well.Whipstock_depth_ouom, well.Row_changed_by, well.Row_changed_date, well.Row_created_by, well.Row_created_date, well.Row_effective_date, well.Row_expiry_date, well.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateWell(c *fiber.Ctx) error {
	var well dto.Well

	setDefaults(&well)

	if err := json.Unmarshal(c.Body(), &well); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	well.Uwi = id

    if well.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update well set abandonment_date = :1, active_ind = :2, area_id = :3, area_type = :4, assigned_field = :5, base_depth = :6, base_depth_ouom = :7, base_node_id = :8, base_strat_name_set_id = :9, base_strat_unit_id = :10, bottom_hole_latitude = :11, bottom_hole_longitude = :12, casing_flange_elev = :13, casing_flange_elev_ouom = :14, completion_date = :15, confidential_date = :16, confidential_depth = :17, confidential_depth_ouom = :18, confidential_type = :19, confid_strat_name_set_id = :20, confid_strat_unit_id = :21, current_class = :22, current_status = :23, current_status_date = :24, deepest_depth = :25, deepest_depth_ouom = :26, depth_datum = :27, depth_datum_elev = :28, depth_datum_elev_ouom = :29, derrick_floor_elev = :30, derrick_floor_elev_ouom = :31, difference_lat_msl = :32, discovery_ind = :33, drill_td = :34, drill_td_ouom = :35, effective_date = :36, elev_ref_datum = :37, environment_type = :38, expiry_date = :39, faulted_ind = :40, final_drill_date = :41, final_td = :42, final_td_ouom = :43, ground_elev = :44, ground_elev_ouom = :45, ground_elev_type = :46, initial_class = :47, kb_elev = :48, kb_elev_ouom = :49, lease_name = :50, lease_num = :51, legal_survey_type = :52, location_type = :53, log_td = :54, log_td_ouom = :55, max_tvd = :56, max_tvd_ouom = :57, net_pay = :58, net_pay_ouom = :59, oldest_strat_age = :60, oldest_strat_name_set_age = :61, oldest_strat_name_set_id = :62, oldest_strat_unit_id = :63, operator = :64, platform_id = :65, platform_sf_subtype = :66, plot_name = :67, plot_symbol = :68, plugback_depth = :69, plugback_depth_ouom = :70, ppdm_guid = :71, primary_source = :72, profile_type = :73, regulatory_agency = :74, remark = :75, rig_on_site_date = :76, rig_release_date = :77, rotary_table_elev = :78, rotary_table_elev_ouom = :79, source_document_id = :80, spud_date = :81, status_type = :82, subsea_elev_ref_type = :83, surface_latitude = :84, surface_longitude = :85, surface_node_id = :86, tax_credit_code = :87, td_strat_age = :88, td_strat_name_set_age = :89, td_strat_name_set_id = :90, td_strat_unit_id = :91, top_depth = :92, top_depth_ouom = :93, top_strat_name_set_id = :94, top_strat_unit_id = :95, water_acoustic_vel = :96, water_acoustic_vel_ouom = :97, water_depth = :98, water_depth_datum = :99, water_depth_ouom = :100, well_event_num = :101, well_government_id = :102, well_intersect_md = :103, well_level_type = :104, well_name = :105, well_num = :106, whipstock_depth = :107, whipstock_depth_ouom = :108, row_changed_by = :109, row_changed_date = :110, row_created_by = :111, row_effective_date = :112, row_expiry_date = :113, row_quality = :114 where uwi = :116")
	if err != nil {
		return err
	}

	well.Row_changed_date = formatDate(well.Row_changed_date)
	well.Abandonment_date = formatDateString(well.Abandonment_date)
	well.Completion_date = formatDateString(well.Completion_date)
	well.Confidential_date = formatDateString(well.Confidential_date)
	well.Current_status_date = formatDateString(well.Current_status_date)
	well.Effective_date = formatDateString(well.Effective_date)
	well.Expiry_date = formatDateString(well.Expiry_date)
	well.Final_drill_date = formatDateString(well.Final_drill_date)
	well.Rig_on_site_date = formatDateString(well.Rig_on_site_date)
	well.Rig_release_date = formatDateString(well.Rig_release_date)
	well.Spud_date = formatDateString(well.Spud_date)
	well.Row_effective_date = formatDateString(well.Row_effective_date)
	well.Row_expiry_date = formatDateString(well.Row_expiry_date)






	rows, err := stmt.Exec(well.Abandonment_date, well.Active_ind, well.Area_id, well.Area_type, well.Assigned_field, well.Base_depth, well.Base_depth_ouom, well.Base_node_id, well.Base_strat_name_set_id, well.Base_strat_unit_id, well.Bottom_hole_latitude, well.Bottom_hole_longitude, well.Casing_flange_elev, well.Casing_flange_elev_ouom, well.Completion_date, well.Confidential_date, well.Confidential_depth, well.Confidential_depth_ouom, well.Confidential_type, well.Confid_strat_name_set_id, well.Confid_strat_unit_id, well.Current_class, well.Current_status, well.Current_status_date, well.Deepest_depth, well.Deepest_depth_ouom, well.Depth_datum, well.Depth_datum_elev, well.Depth_datum_elev_ouom, well.Derrick_floor_elev, well.Derrick_floor_elev_ouom, well.Difference_lat_msl, well.Discovery_ind, well.Drill_td, well.Drill_td_ouom, well.Effective_date, well.Elev_ref_datum, well.Environment_type, well.Expiry_date, well.Faulted_ind, well.Final_drill_date, well.Final_td, well.Final_td_ouom, well.Ground_elev, well.Ground_elev_ouom, well.Ground_elev_type, well.Initial_class, well.Kb_elev, well.Kb_elev_ouom, well.Lease_name, well.Lease_num, well.Legal_survey_type, well.Location_type, well.Log_td, well.Log_td_ouom, well.Max_tvd, well.Max_tvd_ouom, well.Net_pay, well.Net_pay_ouom, well.Oldest_strat_age, well.Oldest_strat_name_set_age, well.Oldest_strat_name_set_id, well.Oldest_strat_unit_id, well.Operator, well.Platform_id, well.Platform_sf_subtype, well.Plot_name, well.Plot_symbol, well.Plugback_depth, well.Plugback_depth_ouom, well.Ppdm_guid, well.Primary_source, well.Profile_type, well.Regulatory_agency, well.Remark, well.Rig_on_site_date, well.Rig_release_date, well.Rotary_table_elev, well.Rotary_table_elev_ouom, well.Source_document_id, well.Spud_date, well.Status_type, well.Subsea_elev_ref_type, well.Surface_latitude, well.Surface_longitude, well.Surface_node_id, well.Tax_credit_code, well.Td_strat_age, well.Td_strat_name_set_age, well.Td_strat_name_set_id, well.Td_strat_unit_id, well.Top_depth, well.Top_depth_ouom, well.Top_strat_name_set_id, well.Top_strat_unit_id, well.Water_acoustic_vel, well.Water_acoustic_vel_ouom, well.Water_depth, well.Water_depth_datum, well.Water_depth_ouom, well.Well_event_num, well.Well_government_id, well.Well_intersect_md, well.Well_level_type, well.Well_name, well.Well_num, well.Whipstock_depth, well.Whipstock_depth_ouom, well.Row_changed_by, well.Row_changed_date, well.Row_created_by, well.Row_effective_date, well.Row_expiry_date, well.Row_quality, well.Uwi)
	if err != nil {
		return err
	}

	if count, err := rows.RowsAffected(); err == nil {
		if count == 0 {
			return c.Status(404).SendString("No matching record found")
		}
	} else {
		return err
	}

	return c.Status(201).JSON(rows)
}

func PatchWell(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update well set "
	values := []interface{}{}
	i := 1
	for key, value := range updatedFields {
		query += key + " = :" + strconv.Itoa(i)
		i++
		if i <= len(updatedFields) {
			query += ", "
		}
		if key == "row_changed_date" {
			if date, ok := value.(string); ok && len(date) > 0 {
				formattedDate := formatDate(&date)
				value = formattedDate
			}
		} else if key == "abandonment_date" || key == "completion_date" || key == "confidential_date" || key == "current_status_date" || key == "effective_date" || key == "expiry_date" || key == "final_drill_date" || key == "rig_on_site_date" || key == "rig_release_date" || key == "spud_date" || key == "row_effective_date" || key == "row_expiry_date"      {
			if date, ok := value.(string); ok && len(date) > 0 {
				formattedDate := formatDateString(&date)
				value = formattedDate
			}
		}
		values = append(values, value)
	}
	query += " where uwi = :id"

	stmt, err := db.Prepare(query)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	values = append(values, id)
	res, err := stmt.Exec(values...)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	if count, err := res.RowsAffected(); err == nil {
		if count == 0 {
			return c.Status(404).SendString("No matching record found")
		}
	} else {
		return err
	}

	return c.JSON(res)
}

func DeleteWell(c *fiber.Ctx) error {
	id := c.Params("id")
	var well dto.Well
	well.Uwi = id

	stmt, err := db.Prepare("delete from well where uwi = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(well.Uwi)
	if err != nil {
		return err
	}

	rowsAffected, err := rows.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return c.Status(404).SendString("No matching record found")
	}

	return c.JSON(rows)
}


