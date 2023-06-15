package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetWellTest(c *fiber.Ctx) error {
	rows, err := db.Query("select * from well_test")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Well_test

	for rows.Next() {
		var well_test dto.Well_test
		if err := rows.Scan(&well_test.Uwi, &well_test.Source, &well_test.Test_type, &well_test.Run_num, &well_test.Test_num, &well_test.Active_ind, &well_test.Base_depth, &well_test.Base_depth_ouom, &well_test.Base_strat_unit_id, &well_test.Bhp_z, &well_test.Bottom_choke_desc, &well_test.Bsw_percent, &well_test.Caliper_hole_diameter, &well_test.Caliper_hole_diameter_ouom, &well_test.Cased_hole_ind, &well_test.Casing_pressure, &well_test.Casing_pressure_ouom, &well_test.Choke_size_desc, &well_test.Completion_obs_no, &well_test.Completion_source, &well_test.Condensate_amount_percent, &well_test.Condensate_flow_amount, &well_test.Condensate_flow_amount_ouom, &well_test.Condensate_flow_amount_uom, &well_test.Condensate_gravity, &well_test.Condensate_ratio, &well_test.Condensate_ratio_ouom, &well_test.Damage_quality, &well_test.Effective_date, &well_test.Event_obs_no, &well_test.Event_source, &well_test.Expiry_date, &well_test.Flow_pressure, &well_test.Flow_pressure_ouom, &well_test.Flow_temperature, &well_test.Flow_temperature_ouom, &well_test.Gas_flow_amount, &well_test.Gas_flow_amount_ouom, &well_test.Gas_flow_amount_uom, &well_test.Gas_gravity, &well_test.Gor, &well_test.Gor_ouom, &well_test.H2s_percent, &well_test.Hole_condition, &well_test.Max_condens_flow_rate, &well_test.Max_condens_flow_rate_ouom, &well_test.Max_gas_flow_rate, &well_test.Max_gas_flow_rate_ouom, &well_test.Max_hydrostatic_pressure, &well_test.Max_hydrostatic_press_ouom, &well_test.Max_oil_flow_rate, &well_test.Max_oil_flow_rate_ouom, &well_test.Max_water_flow_rate, &well_test.Max_water_flow_rate_ouom, &well_test.Oil_amount_percent, &well_test.Oil_flow_amount, &well_test.Oil_flow_amount_ouom, &well_test.Oil_flow_amount_uom, &well_test.Oil_gravity, &well_test.Permeability_quality, &well_test.Ppdm_guid, &well_test.Primary_fluid_recovered, &well_test.Production_method, &well_test.Prod_string_id, &well_test.Prod_str_form_obs_no, &well_test.Rat_hole_diameter, &well_test.Rat_hole_diameter_ouom, &well_test.Rat_hole_length, &well_test.Rat_hole_length_ouom, &well_test.Remark, &well_test.Reported_strat_unit_id, &well_test.Report_temperature, &well_test.Report_temperature_ouom, &well_test.Report_test_num, &well_test.Show_type, &well_test.Si_flow_ratio, &well_test.Start_time, &well_test.Start_timezone, &well_test.Static_pressure, &well_test.Static_pressure_ouom, &well_test.Strat_age, &well_test.Strat_age_name_set_id, &well_test.Strat_name_set_id, &well_test.String_source, &well_test.Sulphur_ind, &well_test.Td, &well_test.Td_ouom, &well_test.Temperature_correction, &well_test.Temperature_correction_ind, &well_test.Temperature_correction_ouom, &well_test.Test_company, &well_test.Test_date, &well_test.Test_duration, &well_test.Test_duration_ouom, &well_test.Test_hole_diameter, &well_test.Test_hole_diameter_ouom, &well_test.Test_result_code, &well_test.Test_subtype, &well_test.Tool_open_time, &well_test.Tool_open_timezone, &well_test.Top_choke_desc, &well_test.Top_depth, &well_test.Top_depth_ouom, &well_test.Top_strat_unit_id, &well_test.Water_amount_percent, &well_test.Water_cut_percent, &well_test.Water_flow_amount, &well_test.Water_flow_amount_ouom, &well_test.Water_flow_amount_uom, &well_test.Wellbore_completion_type, &well_test.Z_factor, &well_test.Row_changed_by, &well_test.Row_changed_date, &well_test.Row_created_by, &well_test.Row_created_date, &well_test.Row_effective_date, &well_test.Row_expiry_date, &well_test.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append well_test to result
		result = append(result, well_test)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetWellTest(c *fiber.Ctx) error {
	var well_test dto.Well_test

	setDefaults(&well_test)

	if err := json.Unmarshal(c.Body(), &well_test); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into well_test values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32, :33, :34, :35, :36, :37, :38, :39, :40, :41, :42, :43, :44, :45, :46, :47, :48, :49, :50, :51, :52, :53, :54, :55, :56, :57, :58, :59, :60, :61, :62, :63, :64, :65, :66, :67, :68, :69, :70, :71, :72, :73, :74, :75, :76, :77, :78, :79, :80, :81, :82, :83, :84, :85, :86, :87, :88, :89, :90, :91, :92, :93, :94, :95, :96, :97, :98, :99, :100, :101, :102, :103, :104, :105, :106, :107, :108, :109, :110, :111, :112, :113, :114, :115, :116, :117, :118)")
	if err != nil {
		return err
	}
	well_test.Row_created_date = formatDate(well_test.Row_created_date)
	well_test.Row_changed_date = formatDate(well_test.Row_changed_date)
	well_test.Effective_date = formatDateString(well_test.Effective_date)
	well_test.Expiry_date = formatDateString(well_test.Expiry_date)
	well_test.Start_time = formatDateString(well_test.Start_time)
	well_test.Test_date = formatDateString(well_test.Test_date)
	well_test.Tool_open_time = formatDateString(well_test.Tool_open_time)
	well_test.Row_effective_date = formatDateString(well_test.Row_effective_date)
	well_test.Row_expiry_date = formatDateString(well_test.Row_expiry_date)






	rows, err := stmt.Exec(well_test.Uwi, well_test.Source, well_test.Test_type, well_test.Run_num, well_test.Test_num, well_test.Active_ind, well_test.Base_depth, well_test.Base_depth_ouom, well_test.Base_strat_unit_id, well_test.Bhp_z, well_test.Bottom_choke_desc, well_test.Bsw_percent, well_test.Caliper_hole_diameter, well_test.Caliper_hole_diameter_ouom, well_test.Cased_hole_ind, well_test.Casing_pressure, well_test.Casing_pressure_ouom, well_test.Choke_size_desc, well_test.Completion_obs_no, well_test.Completion_source, well_test.Condensate_amount_percent, well_test.Condensate_flow_amount, well_test.Condensate_flow_amount_ouom, well_test.Condensate_flow_amount_uom, well_test.Condensate_gravity, well_test.Condensate_ratio, well_test.Condensate_ratio_ouom, well_test.Damage_quality, well_test.Effective_date, well_test.Event_obs_no, well_test.Event_source, well_test.Expiry_date, well_test.Flow_pressure, well_test.Flow_pressure_ouom, well_test.Flow_temperature, well_test.Flow_temperature_ouom, well_test.Gas_flow_amount, well_test.Gas_flow_amount_ouom, well_test.Gas_flow_amount_uom, well_test.Gas_gravity, well_test.Gor, well_test.Gor_ouom, well_test.H2s_percent, well_test.Hole_condition, well_test.Max_condens_flow_rate, well_test.Max_condens_flow_rate_ouom, well_test.Max_gas_flow_rate, well_test.Max_gas_flow_rate_ouom, well_test.Max_hydrostatic_pressure, well_test.Max_hydrostatic_press_ouom, well_test.Max_oil_flow_rate, well_test.Max_oil_flow_rate_ouom, well_test.Max_water_flow_rate, well_test.Max_water_flow_rate_ouom, well_test.Oil_amount_percent, well_test.Oil_flow_amount, well_test.Oil_flow_amount_ouom, well_test.Oil_flow_amount_uom, well_test.Oil_gravity, well_test.Permeability_quality, well_test.Ppdm_guid, well_test.Primary_fluid_recovered, well_test.Production_method, well_test.Prod_string_id, well_test.Prod_str_form_obs_no, well_test.Rat_hole_diameter, well_test.Rat_hole_diameter_ouom, well_test.Rat_hole_length, well_test.Rat_hole_length_ouom, well_test.Remark, well_test.Reported_strat_unit_id, well_test.Report_temperature, well_test.Report_temperature_ouom, well_test.Report_test_num, well_test.Show_type, well_test.Si_flow_ratio, well_test.Start_time, well_test.Start_timezone, well_test.Static_pressure, well_test.Static_pressure_ouom, well_test.Strat_age, well_test.Strat_age_name_set_id, well_test.Strat_name_set_id, well_test.String_source, well_test.Sulphur_ind, well_test.Td, well_test.Td_ouom, well_test.Temperature_correction, well_test.Temperature_correction_ind, well_test.Temperature_correction_ouom, well_test.Test_company, well_test.Test_date, well_test.Test_duration, well_test.Test_duration_ouom, well_test.Test_hole_diameter, well_test.Test_hole_diameter_ouom, well_test.Test_result_code, well_test.Test_subtype, well_test.Tool_open_time, well_test.Tool_open_timezone, well_test.Top_choke_desc, well_test.Top_depth, well_test.Top_depth_ouom, well_test.Top_strat_unit_id, well_test.Water_amount_percent, well_test.Water_cut_percent, well_test.Water_flow_amount, well_test.Water_flow_amount_ouom, well_test.Water_flow_amount_uom, well_test.Wellbore_completion_type, well_test.Z_factor, well_test.Row_changed_by, well_test.Row_changed_date, well_test.Row_created_by, well_test.Row_created_date, well_test.Row_effective_date, well_test.Row_expiry_date, well_test.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateWellTest(c *fiber.Ctx) error {
	var well_test dto.Well_test

	setDefaults(&well_test)

	if err := json.Unmarshal(c.Body(), &well_test); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	well_test.Uwi = id

    if well_test.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update well_test set source = :1, test_type = :2, run_num = :3, test_num = :4, active_ind = :5, base_depth = :6, base_depth_ouom = :7, base_strat_unit_id = :8, bhp_z = :9, bottom_choke_desc = :10, bsw_percent = :11, caliper_hole_diameter = :12, caliper_hole_diameter_ouom = :13, cased_hole_ind = :14, casing_pressure = :15, casing_pressure_ouom = :16, choke_size_desc = :17, completion_obs_no = :18, completion_source = :19, condensate_amount_percent = :20, condensate_flow_amount = :21, condensate_flow_amount_ouom = :22, condensate_flow_amount_uom = :23, condensate_gravity = :24, condensate_ratio = :25, condensate_ratio_ouom = :26, damage_quality = :27, effective_date = :28, event_obs_no = :29, event_source = :30, expiry_date = :31, flow_pressure = :32, flow_pressure_ouom = :33, flow_temperature = :34, flow_temperature_ouom = :35, gas_flow_amount = :36, gas_flow_amount_ouom = :37, gas_flow_amount_uom = :38, gas_gravity = :39, gor = :40, gor_ouom = :41, h2s_percent = :42, hole_condition = :43, max_condens_flow_rate = :44, max_condens_flow_rate_ouom = :45, max_gas_flow_rate = :46, max_gas_flow_rate_ouom = :47, max_hydrostatic_pressure = :48, max_hydrostatic_press_ouom = :49, max_oil_flow_rate = :50, max_oil_flow_rate_ouom = :51, max_water_flow_rate = :52, max_water_flow_rate_ouom = :53, oil_amount_percent = :54, oil_flow_amount = :55, oil_flow_amount_ouom = :56, oil_flow_amount_uom = :57, oil_gravity = :58, permeability_quality = :59, ppdm_guid = :60, primary_fluid_recovered = :61, production_method = :62, prod_string_id = :63, prod_str_form_obs_no = :64, rat_hole_diameter = :65, rat_hole_diameter_ouom = :66, rat_hole_length = :67, rat_hole_length_ouom = :68, remark = :69, reported_strat_unit_id = :70, report_temperature = :71, report_temperature_ouom = :72, report_test_num = :73, show_type = :74, si_flow_ratio = :75, start_time = :76, start_timezone = :77, static_pressure = :78, static_pressure_ouom = :79, strat_age = :80, strat_age_name_set_id = :81, strat_name_set_id = :82, string_source = :83, sulphur_ind = :84, td = :85, td_ouom = :86, temperature_correction = :87, temperature_correction_ind = :88, temperature_correction_ouom = :89, test_company = :90, test_date = :91, test_duration = :92, test_duration_ouom = :93, test_hole_diameter = :94, test_hole_diameter_ouom = :95, test_result_code = :96, test_subtype = :97, tool_open_time = :98, tool_open_timezone = :99, top_choke_desc = :100, top_depth = :101, top_depth_ouom = :102, top_strat_unit_id = :103, water_amount_percent = :104, water_cut_percent = :105, water_flow_amount = :106, water_flow_amount_ouom = :107, water_flow_amount_uom = :108, wellbore_completion_type = :109, z_factor = :110, row_changed_by = :111, row_changed_date = :112, row_created_by = :113, row_effective_date = :114, row_expiry_date = :115, row_quality = :116 where uwi = :118")
	if err != nil {
		return err
	}

	well_test.Row_changed_date = formatDate(well_test.Row_changed_date)
	well_test.Effective_date = formatDateString(well_test.Effective_date)
	well_test.Expiry_date = formatDateString(well_test.Expiry_date)
	well_test.Start_time = formatDateString(well_test.Start_time)
	well_test.Test_date = formatDateString(well_test.Test_date)
	well_test.Tool_open_time = formatDateString(well_test.Tool_open_time)
	well_test.Row_effective_date = formatDateString(well_test.Row_effective_date)
	well_test.Row_expiry_date = formatDateString(well_test.Row_expiry_date)






	rows, err := stmt.Exec(well_test.Source, well_test.Test_type, well_test.Run_num, well_test.Test_num, well_test.Active_ind, well_test.Base_depth, well_test.Base_depth_ouom, well_test.Base_strat_unit_id, well_test.Bhp_z, well_test.Bottom_choke_desc, well_test.Bsw_percent, well_test.Caliper_hole_diameter, well_test.Caliper_hole_diameter_ouom, well_test.Cased_hole_ind, well_test.Casing_pressure, well_test.Casing_pressure_ouom, well_test.Choke_size_desc, well_test.Completion_obs_no, well_test.Completion_source, well_test.Condensate_amount_percent, well_test.Condensate_flow_amount, well_test.Condensate_flow_amount_ouom, well_test.Condensate_flow_amount_uom, well_test.Condensate_gravity, well_test.Condensate_ratio, well_test.Condensate_ratio_ouom, well_test.Damage_quality, well_test.Effective_date, well_test.Event_obs_no, well_test.Event_source, well_test.Expiry_date, well_test.Flow_pressure, well_test.Flow_pressure_ouom, well_test.Flow_temperature, well_test.Flow_temperature_ouom, well_test.Gas_flow_amount, well_test.Gas_flow_amount_ouom, well_test.Gas_flow_amount_uom, well_test.Gas_gravity, well_test.Gor, well_test.Gor_ouom, well_test.H2s_percent, well_test.Hole_condition, well_test.Max_condens_flow_rate, well_test.Max_condens_flow_rate_ouom, well_test.Max_gas_flow_rate, well_test.Max_gas_flow_rate_ouom, well_test.Max_hydrostatic_pressure, well_test.Max_hydrostatic_press_ouom, well_test.Max_oil_flow_rate, well_test.Max_oil_flow_rate_ouom, well_test.Max_water_flow_rate, well_test.Max_water_flow_rate_ouom, well_test.Oil_amount_percent, well_test.Oil_flow_amount, well_test.Oil_flow_amount_ouom, well_test.Oil_flow_amount_uom, well_test.Oil_gravity, well_test.Permeability_quality, well_test.Ppdm_guid, well_test.Primary_fluid_recovered, well_test.Production_method, well_test.Prod_string_id, well_test.Prod_str_form_obs_no, well_test.Rat_hole_diameter, well_test.Rat_hole_diameter_ouom, well_test.Rat_hole_length, well_test.Rat_hole_length_ouom, well_test.Remark, well_test.Reported_strat_unit_id, well_test.Report_temperature, well_test.Report_temperature_ouom, well_test.Report_test_num, well_test.Show_type, well_test.Si_flow_ratio, well_test.Start_time, well_test.Start_timezone, well_test.Static_pressure, well_test.Static_pressure_ouom, well_test.Strat_age, well_test.Strat_age_name_set_id, well_test.Strat_name_set_id, well_test.String_source, well_test.Sulphur_ind, well_test.Td, well_test.Td_ouom, well_test.Temperature_correction, well_test.Temperature_correction_ind, well_test.Temperature_correction_ouom, well_test.Test_company, well_test.Test_date, well_test.Test_duration, well_test.Test_duration_ouom, well_test.Test_hole_diameter, well_test.Test_hole_diameter_ouom, well_test.Test_result_code, well_test.Test_subtype, well_test.Tool_open_time, well_test.Tool_open_timezone, well_test.Top_choke_desc, well_test.Top_depth, well_test.Top_depth_ouom, well_test.Top_strat_unit_id, well_test.Water_amount_percent, well_test.Water_cut_percent, well_test.Water_flow_amount, well_test.Water_flow_amount_ouom, well_test.Water_flow_amount_uom, well_test.Wellbore_completion_type, well_test.Z_factor, well_test.Row_changed_by, well_test.Row_changed_date, well_test.Row_created_by, well_test.Row_effective_date, well_test.Row_expiry_date, well_test.Row_quality, well_test.Uwi)
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

func PatchWellTest(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update well_test set "
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
		} else if key == "effective_date" || key == "expiry_date" || key == "start_time" || key == "test_date" || key == "tool_open_time" || key == "row_effective_date" || key == "row_expiry_date"      {
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

func DeleteWellTest(c *fiber.Ctx) error {
	id := c.Params("id")
	var well_test dto.Well_test
	well_test.Uwi = id

	stmt, err := db.Prepare("delete from well_test where uwi = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(well_test.Uwi)
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


