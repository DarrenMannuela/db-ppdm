package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetWellVersion(c *fiber.Ctx) error {
	rows, err := db.Query("select * from well_version")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Well_version

	for rows.Next() {
		var well_version dto.Well_version
		if err := rows.Scan(&well_version.Uwi, &well_version.Source, &well_version.Abandonment_date, &well_version.Active_ind, &well_version.Assigned_field, &well_version.Base_depth, &well_version.Base_depth_ouom, &well_version.Base_node_id, &well_version.Base_strat_name_set_id, &well_version.Base_strat_unit_id, &well_version.Bottom_hole_latitude, &well_version.Bottom_hole_longitude, &well_version.Casing_flange_elev, &well_version.Casing_flange_elev_ouom, &well_version.Completion_date, &well_version.Confidential_date, &well_version.Confidential_depth, &well_version.Confidential_depth_ouom, &well_version.Confidential_type, &well_version.Confid_strat_name_set_id, &well_version.Confid_strat_unit_id, &well_version.Current_class, &well_version.Current_status, &well_version.Current_status_date, &well_version.Deepest_depth, &well_version.Deepest_depth_ouom, &well_version.Depth_datum, &well_version.Depth_datum_elev, &well_version.Depth_datum_elev_ouom, &well_version.Derrick_floor_elev, &well_version.Derrick_floor_elev_ouom, &well_version.Difference_lat_msl, &well_version.Discovery_ind, &well_version.Drill_td, &well_version.Drill_td_ouom, &well_version.Effective_date, &well_version.Elev_ref_datum, &well_version.Environment_type, &well_version.Expiry_date, &well_version.Faulted_ind, &well_version.Final_drill_date, &well_version.Final_td, &well_version.Final_td_ouom, &well_version.Ground_elev, &well_version.Ground_elev_ouom, &well_version.Ground_elev_type, &well_version.Initial_class, &well_version.Kb_elev, &well_version.Kb_elev_ouom, &well_version.Lease_name, &well_version.Lease_num, &well_version.Legal_survey_type, &well_version.Location_type, &well_version.Log_td, &well_version.Log_td_ouom, &well_version.Max_tvd, &well_version.Max_tvd_ouom, &well_version.Net_pay, &well_version.Net_pay_ouom, &well_version.Oldest_strat_age, &well_version.Oldest_strat_name_set_age, &well_version.Oldest_strat_name_set_id, &well_version.Oldest_strat_unit_id, &well_version.Operator, &well_version.Platform_id, &well_version.Platform_sf_subtype, &well_version.Plot_name, &well_version.Plot_symbol, &well_version.Plugback_depth, &well_version.Plugback_depth_ouom, &well_version.Ppdm_guid, &well_version.Profile_type, &well_version.Regulatory_agency, &well_version.Remark, &well_version.Rig_on_site_date, &well_version.Rig_release_date, &well_version.Rotary_table_elev, &well_version.Rotary_table_elev_ouom, &well_version.Source_document_id, &well_version.Spud_date, &well_version.Status_type, &well_version.Subsea_elev_ref_type, &well_version.Surface_latitude, &well_version.Surface_longitude, &well_version.Surface_node_id, &well_version.Tax_credit_code, &well_version.Td_strat_age, &well_version.Td_strat_name_set_age, &well_version.Td_strat_name_set_id, &well_version.Td_strat_unit_id, &well_version.Top_depth, &well_version.Top_depth_ouom, &well_version.Top_strat_name_set_id, &well_version.Top_strat_unit_id, &well_version.Water_acoustic_vel, &well_version.Water_acoustic_vel_ouom, &well_version.Water_depth, &well_version.Water_depth_datum, &well_version.Water_depth_ouom, &well_version.Well_event_num, &well_version.Well_government_id, &well_version.Well_intersect_md, &well_version.Well_level_type, &well_version.Well_name, &well_version.Well_num, &well_version.Whipstock_depth, &well_version.Whipstock_depth_ouom, &well_version.Row_changed_by, &well_version.Row_changed_date, &well_version.Row_created_by, &well_version.Row_created_date, &well_version.Row_effective_date, &well_version.Row_expiry_date, &well_version.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append well_version to result
		result = append(result, well_version)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetWellVersion(c *fiber.Ctx) error {
	var well_version dto.Well_version

	setDefaults(&well_version)

	if err := json.Unmarshal(c.Body(), &well_version); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into well_version values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32, :33, :34, :35, :36, :37, :38, :39, :40, :41, :42, :43, :44, :45, :46, :47, :48, :49, :50, :51, :52, :53, :54, :55, :56, :57, :58, :59, :60, :61, :62, :63, :64, :65, :66, :67, :68, :69, :70, :71, :72, :73, :74, :75, :76, :77, :78, :79, :80, :81, :82, :83, :84, :85, :86, :87, :88, :89, :90, :91, :92, :93, :94, :95, :96, :97, :98, :99, :100, :101, :102, :103, :104, :105, :106, :107, :108, :109, :110, :111, :112, :113, :114)")
	if err != nil {
		return err
	}
	well_version.Row_created_date = formatDate(well_version.Row_created_date)
	well_version.Row_changed_date = formatDate(well_version.Row_changed_date)
	well_version.Abandonment_date = formatDateString(well_version.Abandonment_date)
	well_version.Completion_date = formatDateString(well_version.Completion_date)
	well_version.Confidential_date = formatDateString(well_version.Confidential_date)
	well_version.Current_status_date = formatDateString(well_version.Current_status_date)
	well_version.Effective_date = formatDateString(well_version.Effective_date)
	well_version.Expiry_date = formatDateString(well_version.Expiry_date)
	well_version.Final_drill_date = formatDateString(well_version.Final_drill_date)
	well_version.Rig_on_site_date = formatDateString(well_version.Rig_on_site_date)
	well_version.Rig_release_date = formatDateString(well_version.Rig_release_date)
	well_version.Spud_date = formatDateString(well_version.Spud_date)
	well_version.Row_effective_date = formatDateString(well_version.Row_effective_date)
	well_version.Row_expiry_date = formatDateString(well_version.Row_expiry_date)






	rows, err := stmt.Exec(well_version.Uwi, well_version.Source, well_version.Abandonment_date, well_version.Active_ind, well_version.Assigned_field, well_version.Base_depth, well_version.Base_depth_ouom, well_version.Base_node_id, well_version.Base_strat_name_set_id, well_version.Base_strat_unit_id, well_version.Bottom_hole_latitude, well_version.Bottom_hole_longitude, well_version.Casing_flange_elev, well_version.Casing_flange_elev_ouom, well_version.Completion_date, well_version.Confidential_date, well_version.Confidential_depth, well_version.Confidential_depth_ouom, well_version.Confidential_type, well_version.Confid_strat_name_set_id, well_version.Confid_strat_unit_id, well_version.Current_class, well_version.Current_status, well_version.Current_status_date, well_version.Deepest_depth, well_version.Deepest_depth_ouom, well_version.Depth_datum, well_version.Depth_datum_elev, well_version.Depth_datum_elev_ouom, well_version.Derrick_floor_elev, well_version.Derrick_floor_elev_ouom, well_version.Difference_lat_msl, well_version.Discovery_ind, well_version.Drill_td, well_version.Drill_td_ouom, well_version.Effective_date, well_version.Elev_ref_datum, well_version.Environment_type, well_version.Expiry_date, well_version.Faulted_ind, well_version.Final_drill_date, well_version.Final_td, well_version.Final_td_ouom, well_version.Ground_elev, well_version.Ground_elev_ouom, well_version.Ground_elev_type, well_version.Initial_class, well_version.Kb_elev, well_version.Kb_elev_ouom, well_version.Lease_name, well_version.Lease_num, well_version.Legal_survey_type, well_version.Location_type, well_version.Log_td, well_version.Log_td_ouom, well_version.Max_tvd, well_version.Max_tvd_ouom, well_version.Net_pay, well_version.Net_pay_ouom, well_version.Oldest_strat_age, well_version.Oldest_strat_name_set_age, well_version.Oldest_strat_name_set_id, well_version.Oldest_strat_unit_id, well_version.Operator, well_version.Platform_id, well_version.Platform_sf_subtype, well_version.Plot_name, well_version.Plot_symbol, well_version.Plugback_depth, well_version.Plugback_depth_ouom, well_version.Ppdm_guid, well_version.Profile_type, well_version.Regulatory_agency, well_version.Remark, well_version.Rig_on_site_date, well_version.Rig_release_date, well_version.Rotary_table_elev, well_version.Rotary_table_elev_ouom, well_version.Source_document_id, well_version.Spud_date, well_version.Status_type, well_version.Subsea_elev_ref_type, well_version.Surface_latitude, well_version.Surface_longitude, well_version.Surface_node_id, well_version.Tax_credit_code, well_version.Td_strat_age, well_version.Td_strat_name_set_age, well_version.Td_strat_name_set_id, well_version.Td_strat_unit_id, well_version.Top_depth, well_version.Top_depth_ouom, well_version.Top_strat_name_set_id, well_version.Top_strat_unit_id, well_version.Water_acoustic_vel, well_version.Water_acoustic_vel_ouom, well_version.Water_depth, well_version.Water_depth_datum, well_version.Water_depth_ouom, well_version.Well_event_num, well_version.Well_government_id, well_version.Well_intersect_md, well_version.Well_level_type, well_version.Well_name, well_version.Well_num, well_version.Whipstock_depth, well_version.Whipstock_depth_ouom, well_version.Row_changed_by, well_version.Row_changed_date, well_version.Row_created_by, well_version.Row_created_date, well_version.Row_effective_date, well_version.Row_expiry_date, well_version.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateWellVersion(c *fiber.Ctx) error {
	var well_version dto.Well_version

	setDefaults(&well_version)

	if err := json.Unmarshal(c.Body(), &well_version); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	well_version.Uwi = id

    if well_version.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update well_version set source = :1, abandonment_date = :2, active_ind = :3, assigned_field = :4, base_depth = :5, base_depth_ouom = :6, base_node_id = :7, base_strat_name_set_id = :8, base_strat_unit_id = :9, bottom_hole_latitude = :10, bottom_hole_longitude = :11, casing_flange_elev = :12, casing_flange_elev_ouom = :13, completion_date = :14, confidential_date = :15, confidential_depth = :16, confidential_depth_ouom = :17, confidential_type = :18, confid_strat_name_set_id = :19, confid_strat_unit_id = :20, current_class = :21, current_status = :22, current_status_date = :23, deepest_depth = :24, deepest_depth_ouom = :25, depth_datum = :26, depth_datum_elev = :27, depth_datum_elev_ouom = :28, derrick_floor_elev = :29, derrick_floor_elev_ouom = :30, difference_lat_msl = :31, discovery_ind = :32, drill_td = :33, drill_td_ouom = :34, effective_date = :35, elev_ref_datum = :36, environment_type = :37, expiry_date = :38, faulted_ind = :39, final_drill_date = :40, final_td = :41, final_td_ouom = :42, ground_elev = :43, ground_elev_ouom = :44, ground_elev_type = :45, initial_class = :46, kb_elev = :47, kb_elev_ouom = :48, lease_name = :49, lease_num = :50, legal_survey_type = :51, location_type = :52, log_td = :53, log_td_ouom = :54, max_tvd = :55, max_tvd_ouom = :56, net_pay = :57, net_pay_ouom = :58, oldest_strat_age = :59, oldest_strat_name_set_age = :60, oldest_strat_name_set_id = :61, oldest_strat_unit_id = :62, operator = :63, platform_id = :64, platform_sf_subtype = :65, plot_name = :66, plot_symbol = :67, plugback_depth = :68, plugback_depth_ouom = :69, ppdm_guid = :70, profile_type = :71, regulatory_agency = :72, remark = :73, rig_on_site_date = :74, rig_release_date = :75, rotary_table_elev = :76, rotary_table_elev_ouom = :77, source_document_id = :78, spud_date = :79, status_type = :80, subsea_elev_ref_type = :81, surface_latitude = :82, surface_longitude = :83, surface_node_id = :84, tax_credit_code = :85, td_strat_age = :86, td_strat_name_set_age = :87, td_strat_name_set_id = :88, td_strat_unit_id = :89, top_depth = :90, top_depth_ouom = :91, top_strat_name_set_id = :92, top_strat_unit_id = :93, water_acoustic_vel = :94, water_acoustic_vel_ouom = :95, water_depth = :96, water_depth_datum = :97, water_depth_ouom = :98, well_event_num = :99, well_government_id = :100, well_intersect_md = :101, well_level_type = :102, well_name = :103, well_num = :104, whipstock_depth = :105, whipstock_depth_ouom = :106, row_changed_by = :107, row_changed_date = :108, row_created_by = :109, row_effective_date = :110, row_expiry_date = :111, row_quality = :112 where uwi = :114")
	if err != nil {
		return err
	}

	well_version.Row_changed_date = formatDate(well_version.Row_changed_date)
	well_version.Abandonment_date = formatDateString(well_version.Abandonment_date)
	well_version.Completion_date = formatDateString(well_version.Completion_date)
	well_version.Confidential_date = formatDateString(well_version.Confidential_date)
	well_version.Current_status_date = formatDateString(well_version.Current_status_date)
	well_version.Effective_date = formatDateString(well_version.Effective_date)
	well_version.Expiry_date = formatDateString(well_version.Expiry_date)
	well_version.Final_drill_date = formatDateString(well_version.Final_drill_date)
	well_version.Rig_on_site_date = formatDateString(well_version.Rig_on_site_date)
	well_version.Rig_release_date = formatDateString(well_version.Rig_release_date)
	well_version.Spud_date = formatDateString(well_version.Spud_date)
	well_version.Row_effective_date = formatDateString(well_version.Row_effective_date)
	well_version.Row_expiry_date = formatDateString(well_version.Row_expiry_date)






	rows, err := stmt.Exec(well_version.Source, well_version.Abandonment_date, well_version.Active_ind, well_version.Assigned_field, well_version.Base_depth, well_version.Base_depth_ouom, well_version.Base_node_id, well_version.Base_strat_name_set_id, well_version.Base_strat_unit_id, well_version.Bottom_hole_latitude, well_version.Bottom_hole_longitude, well_version.Casing_flange_elev, well_version.Casing_flange_elev_ouom, well_version.Completion_date, well_version.Confidential_date, well_version.Confidential_depth, well_version.Confidential_depth_ouom, well_version.Confidential_type, well_version.Confid_strat_name_set_id, well_version.Confid_strat_unit_id, well_version.Current_class, well_version.Current_status, well_version.Current_status_date, well_version.Deepest_depth, well_version.Deepest_depth_ouom, well_version.Depth_datum, well_version.Depth_datum_elev, well_version.Depth_datum_elev_ouom, well_version.Derrick_floor_elev, well_version.Derrick_floor_elev_ouom, well_version.Difference_lat_msl, well_version.Discovery_ind, well_version.Drill_td, well_version.Drill_td_ouom, well_version.Effective_date, well_version.Elev_ref_datum, well_version.Environment_type, well_version.Expiry_date, well_version.Faulted_ind, well_version.Final_drill_date, well_version.Final_td, well_version.Final_td_ouom, well_version.Ground_elev, well_version.Ground_elev_ouom, well_version.Ground_elev_type, well_version.Initial_class, well_version.Kb_elev, well_version.Kb_elev_ouom, well_version.Lease_name, well_version.Lease_num, well_version.Legal_survey_type, well_version.Location_type, well_version.Log_td, well_version.Log_td_ouom, well_version.Max_tvd, well_version.Max_tvd_ouom, well_version.Net_pay, well_version.Net_pay_ouom, well_version.Oldest_strat_age, well_version.Oldest_strat_name_set_age, well_version.Oldest_strat_name_set_id, well_version.Oldest_strat_unit_id, well_version.Operator, well_version.Platform_id, well_version.Platform_sf_subtype, well_version.Plot_name, well_version.Plot_symbol, well_version.Plugback_depth, well_version.Plugback_depth_ouom, well_version.Ppdm_guid, well_version.Profile_type, well_version.Regulatory_agency, well_version.Remark, well_version.Rig_on_site_date, well_version.Rig_release_date, well_version.Rotary_table_elev, well_version.Rotary_table_elev_ouom, well_version.Source_document_id, well_version.Spud_date, well_version.Status_type, well_version.Subsea_elev_ref_type, well_version.Surface_latitude, well_version.Surface_longitude, well_version.Surface_node_id, well_version.Tax_credit_code, well_version.Td_strat_age, well_version.Td_strat_name_set_age, well_version.Td_strat_name_set_id, well_version.Td_strat_unit_id, well_version.Top_depth, well_version.Top_depth_ouom, well_version.Top_strat_name_set_id, well_version.Top_strat_unit_id, well_version.Water_acoustic_vel, well_version.Water_acoustic_vel_ouom, well_version.Water_depth, well_version.Water_depth_datum, well_version.Water_depth_ouom, well_version.Well_event_num, well_version.Well_government_id, well_version.Well_intersect_md, well_version.Well_level_type, well_version.Well_name, well_version.Well_num, well_version.Whipstock_depth, well_version.Whipstock_depth_ouom, well_version.Row_changed_by, well_version.Row_changed_date, well_version.Row_created_by, well_version.Row_effective_date, well_version.Row_expiry_date, well_version.Row_quality, well_version.Uwi)
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

func PatchWellVersion(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update well_version set "
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

func DeleteWellVersion(c *fiber.Ctx) error {
	id := c.Params("id")
	var well_version dto.Well_version
	well_version.Uwi = id

	stmt, err := db.Prepare("delete from well_version where uwi = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(well_version.Uwi)
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


