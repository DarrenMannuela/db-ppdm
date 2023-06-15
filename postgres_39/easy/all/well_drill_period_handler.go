package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetWellDrillPeriod(c *fiber.Ctx) error {
	rows, err := db.Query("select * from well_drill_period")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Well_drill_period

	for rows.Next() {
		var well_drill_period dto.Well_drill_period
		if err := rows.Scan(&well_drill_period.Uwi, &well_drill_period.Period_obs_no, &well_drill_period.Active_ind, &well_drill_period.Bha_weight, &well_drill_period.Bha_weight_ouom, &well_drill_period.Boiler_fuel, &well_drill_period.Boiler_fuel_ouom, &well_drill_period.Camp_ind, &well_drill_period.Casing_press_max, &well_drill_period.Casing_press_max_ouom, &well_drill_period.Circ_pressure, &well_drill_period.Circ_pressure_max, &well_drill_period.Circ_pressure_max_ouom, &well_drill_period.Circ_pressure_min, &well_drill_period.Circ_pressure_min_ouom, &well_drill_period.Circ_pressure_ouom, &well_drill_period.Circ_pressure_type, &well_drill_period.Contractor_ba_id, &well_drill_period.Contractor_rep_ba_id, &well_drill_period.Contractor_rep_name, &well_drill_period.Contractor_rep_signed_ind, &well_drill_period.Cum_crew_safety_days, &well_drill_period.Cum_rig_safety_days, &well_drill_period.Drill_activity_set_id, &well_drill_period.Drill_activity_type_id, &well_drill_period.Drill_pipe_single_count, &well_drill_period.Drill_pipe_single_len, &well_drill_period.Drill_pipe_single_len_ouom, &well_drill_period.Drill_pipe_stand_count, &well_drill_period.Drill_pipe_stand_len, &well_drill_period.Drill_pipe_stand_len_ouom, &well_drill_period.Drill_string_torque, &well_drill_period.Drill_string_torque_ouom, &well_drill_period.Effective_date, &well_drill_period.End_date, &well_drill_period.End_time, &well_drill_period.End_timezone, &well_drill_period.Expiry_date, &well_drill_period.Finance_id, &well_drill_period.Hole_drag, &well_drill_period.Hole_drag_ouom, &well_drill_period.Hole_fill_length, &well_drill_period.Hole_fill_length_ouom, &well_drill_period.Hookload_max, &well_drill_period.Hookload_max_ouom, &well_drill_period.Job_num, &well_drill_period.Kb_down_length, &well_drill_period.Kb_down_length_ouom, &well_drill_period.Mud_system_type, &well_drill_period.Mud_type_oil_ind, &well_drill_period.Mud_type_water_ind, &well_drill_period.Operator_ba_id, &well_drill_period.Operator_fuel, &well_drill_period.Operator_fuel_ouom, &well_drill_period.Operator_fuel_uom, &well_drill_period.Operator_rep_ba_id, &well_drill_period.Operator_rep_name, &well_drill_period.Operator_rep_signed_ind, &well_drill_period.Ppdm_guid, &well_drill_period.Primary_activity_ind, &well_drill_period.Rate_schedule_id, &well_drill_period.Reentry_ind, &well_drill_period.Remark, &well_drill_period.Reported_base_location, &well_drill_period.Reported_kb_elev, &well_drill_period.Reported_kb_elev_ouom, &well_drill_period.Reported_license_num, &well_drill_period.Reported_period_num, &well_drill_period.Reported_profile_type, &well_drill_period.Reported_rig_name, &well_drill_period.Reported_rig_rel_date, &well_drill_period.Reported_rig_rel_time, &well_drill_period.Reported_rig_rel_timezone, &well_drill_period.Reported_spud_date, &well_drill_period.Reported_spud_time, &well_drill_period.Reported_spud_timezone, &well_drill_period.Reported_top_location, &well_drill_period.Reported_uwi, &well_drill_period.Reported_well_name, &well_drill_period.Report_time_type, &well_drill_period.Rig_fuel, &well_drill_period.Rig_fuel_ouom, &well_drill_period.Rig_fuel_uom, &well_drill_period.Safety_meeting_topic, &well_drill_period.Source, &well_drill_period.Start_date, &well_drill_period.Start_time, &well_drill_period.Start_timezone, &well_drill_period.Tot_drill_tool_length, &well_drill_period.Tot_drill_tool_length_ouom, &well_drill_period.Tot_drill_tool_weight, &well_drill_period.Tot_drill_tool_weight_ouom, &well_drill_period.Row_changed_by, &well_drill_period.Row_changed_date, &well_drill_period.Row_created_by, &well_drill_period.Row_created_date, &well_drill_period.Row_effective_date, &well_drill_period.Row_expiry_date, &well_drill_period.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append well_drill_period to result
		result = append(result, well_drill_period)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetWellDrillPeriod(c *fiber.Ctx) error {
	var well_drill_period dto.Well_drill_period

	setDefaults(&well_drill_period)

	if err := json.Unmarshal(c.Body(), &well_drill_period); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into well_drill_period values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32, :33, :34, :35, :36, :37, :38, :39, :40, :41, :42, :43, :44, :45, :46, :47, :48, :49, :50, :51, :52, :53, :54, :55, :56, :57, :58, :59, :60, :61, :62, :63, :64, :65, :66, :67, :68, :69, :70, :71, :72, :73, :74, :75, :76, :77, :78, :79, :80, :81, :82, :83, :84, :85, :86, :87, :88, :89, :90, :91, :92, :93, :94, :95, :96, :97, :98, :99)")
	if err != nil {
		return err
	}
	well_drill_period.Row_created_date = formatDate(well_drill_period.Row_created_date)
	well_drill_period.Row_changed_date = formatDate(well_drill_period.Row_changed_date)
	well_drill_period.Effective_date = formatDateString(well_drill_period.Effective_date)
	well_drill_period.End_date = formatDateString(well_drill_period.End_date)
	well_drill_period.End_time = formatDateString(well_drill_period.End_time)
	well_drill_period.Expiry_date = formatDateString(well_drill_period.Expiry_date)
	well_drill_period.Reported_rig_rel_date = formatDateString(well_drill_period.Reported_rig_rel_date)
	well_drill_period.Reported_rig_rel_time = formatDateString(well_drill_period.Reported_rig_rel_time)
	well_drill_period.Reported_spud_date = formatDateString(well_drill_period.Reported_spud_date)
	well_drill_period.Reported_spud_time = formatDateString(well_drill_period.Reported_spud_time)
	well_drill_period.Start_date = formatDateString(well_drill_period.Start_date)
	well_drill_period.Start_time = formatDateString(well_drill_period.Start_time)
	well_drill_period.Row_effective_date = formatDateString(well_drill_period.Row_effective_date)
	well_drill_period.Row_expiry_date = formatDateString(well_drill_period.Row_expiry_date)






	rows, err := stmt.Exec(well_drill_period.Uwi, well_drill_period.Period_obs_no, well_drill_period.Active_ind, well_drill_period.Bha_weight, well_drill_period.Bha_weight_ouom, well_drill_period.Boiler_fuel, well_drill_period.Boiler_fuel_ouom, well_drill_period.Camp_ind, well_drill_period.Casing_press_max, well_drill_period.Casing_press_max_ouom, well_drill_period.Circ_pressure, well_drill_period.Circ_pressure_max, well_drill_period.Circ_pressure_max_ouom, well_drill_period.Circ_pressure_min, well_drill_period.Circ_pressure_min_ouom, well_drill_period.Circ_pressure_ouom, well_drill_period.Circ_pressure_type, well_drill_period.Contractor_ba_id, well_drill_period.Contractor_rep_ba_id, well_drill_period.Contractor_rep_name, well_drill_period.Contractor_rep_signed_ind, well_drill_period.Cum_crew_safety_days, well_drill_period.Cum_rig_safety_days, well_drill_period.Drill_activity_set_id, well_drill_period.Drill_activity_type_id, well_drill_period.Drill_pipe_single_count, well_drill_period.Drill_pipe_single_len, well_drill_period.Drill_pipe_single_len_ouom, well_drill_period.Drill_pipe_stand_count, well_drill_period.Drill_pipe_stand_len, well_drill_period.Drill_pipe_stand_len_ouom, well_drill_period.Drill_string_torque, well_drill_period.Drill_string_torque_ouom, well_drill_period.Effective_date, well_drill_period.End_date, well_drill_period.End_time, well_drill_period.End_timezone, well_drill_period.Expiry_date, well_drill_period.Finance_id, well_drill_period.Hole_drag, well_drill_period.Hole_drag_ouom, well_drill_period.Hole_fill_length, well_drill_period.Hole_fill_length_ouom, well_drill_period.Hookload_max, well_drill_period.Hookload_max_ouom, well_drill_period.Job_num, well_drill_period.Kb_down_length, well_drill_period.Kb_down_length_ouom, well_drill_period.Mud_system_type, well_drill_period.Mud_type_oil_ind, well_drill_period.Mud_type_water_ind, well_drill_period.Operator_ba_id, well_drill_period.Operator_fuel, well_drill_period.Operator_fuel_ouom, well_drill_period.Operator_fuel_uom, well_drill_period.Operator_rep_ba_id, well_drill_period.Operator_rep_name, well_drill_period.Operator_rep_signed_ind, well_drill_period.Ppdm_guid, well_drill_period.Primary_activity_ind, well_drill_period.Rate_schedule_id, well_drill_period.Reentry_ind, well_drill_period.Remark, well_drill_period.Reported_base_location, well_drill_period.Reported_kb_elev, well_drill_period.Reported_kb_elev_ouom, well_drill_period.Reported_license_num, well_drill_period.Reported_period_num, well_drill_period.Reported_profile_type, well_drill_period.Reported_rig_name, well_drill_period.Reported_rig_rel_date, well_drill_period.Reported_rig_rel_time, well_drill_period.Reported_rig_rel_timezone, well_drill_period.Reported_spud_date, well_drill_period.Reported_spud_time, well_drill_period.Reported_spud_timezone, well_drill_period.Reported_top_location, well_drill_period.Reported_uwi, well_drill_period.Reported_well_name, well_drill_period.Report_time_type, well_drill_period.Rig_fuel, well_drill_period.Rig_fuel_ouom, well_drill_period.Rig_fuel_uom, well_drill_period.Safety_meeting_topic, well_drill_period.Source, well_drill_period.Start_date, well_drill_period.Start_time, well_drill_period.Start_timezone, well_drill_period.Tot_drill_tool_length, well_drill_period.Tot_drill_tool_length_ouom, well_drill_period.Tot_drill_tool_weight, well_drill_period.Tot_drill_tool_weight_ouom, well_drill_period.Row_changed_by, well_drill_period.Row_changed_date, well_drill_period.Row_created_by, well_drill_period.Row_created_date, well_drill_period.Row_effective_date, well_drill_period.Row_expiry_date, well_drill_period.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateWellDrillPeriod(c *fiber.Ctx) error {
	var well_drill_period dto.Well_drill_period

	setDefaults(&well_drill_period)

	if err := json.Unmarshal(c.Body(), &well_drill_period); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	well_drill_period.Uwi = id

    if well_drill_period.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update well_drill_period set period_obs_no = :1, active_ind = :2, bha_weight = :3, bha_weight_ouom = :4, boiler_fuel = :5, boiler_fuel_ouom = :6, camp_ind = :7, casing_press_max = :8, casing_press_max_ouom = :9, circ_pressure = :10, circ_pressure_max = :11, circ_pressure_max_ouom = :12, circ_pressure_min = :13, circ_pressure_min_ouom = :14, circ_pressure_ouom = :15, circ_pressure_type = :16, contractor_ba_id = :17, contractor_rep_ba_id = :18, contractor_rep_name = :19, contractor_rep_signed_ind = :20, cum_crew_safety_days = :21, cum_rig_safety_days = :22, drill_activity_set_id = :23, drill_activity_type_id = :24, drill_pipe_single_count = :25, drill_pipe_single_len = :26, drill_pipe_single_len_ouom = :27, drill_pipe_stand_count = :28, drill_pipe_stand_len = :29, drill_pipe_stand_len_ouom = :30, drill_string_torque = :31, drill_string_torque_ouom = :32, effective_date = :33, end_date = :34, end_time = :35, end_timezone = :36, expiry_date = :37, finance_id = :38, hole_drag = :39, hole_drag_ouom = :40, hole_fill_length = :41, hole_fill_length_ouom = :42, hookload_max = :43, hookload_max_ouom = :44, job_num = :45, kb_down_length = :46, kb_down_length_ouom = :47, mud_system_type = :48, mud_type_oil_ind = :49, mud_type_water_ind = :50, operator_ba_id = :51, operator_fuel = :52, operator_fuel_ouom = :53, operator_fuel_uom = :54, operator_rep_ba_id = :55, operator_rep_name = :56, operator_rep_signed_ind = :57, ppdm_guid = :58, primary_activity_ind = :59, rate_schedule_id = :60, reentry_ind = :61, remark = :62, reported_base_location = :63, reported_kb_elev = :64, reported_kb_elev_ouom = :65, reported_license_num = :66, reported_period_num = :67, reported_profile_type = :68, reported_rig_name = :69, reported_rig_rel_date = :70, reported_rig_rel_time = :71, reported_rig_rel_timezone = :72, reported_spud_date = :73, reported_spud_time = :74, reported_spud_timezone = :75, reported_top_location = :76, reported_uwi = :77, reported_well_name = :78, report_time_type = :79, rig_fuel = :80, rig_fuel_ouom = :81, rig_fuel_uom = :82, safety_meeting_topic = :83, source = :84, start_date = :85, start_time = :86, start_timezone = :87, tot_drill_tool_length = :88, tot_drill_tool_length_ouom = :89, tot_drill_tool_weight = :90, tot_drill_tool_weight_ouom = :91, row_changed_by = :92, row_changed_date = :93, row_created_by = :94, row_effective_date = :95, row_expiry_date = :96, row_quality = :97 where uwi = :99")
	if err != nil {
		return err
	}

	well_drill_period.Row_changed_date = formatDate(well_drill_period.Row_changed_date)
	well_drill_period.Effective_date = formatDateString(well_drill_period.Effective_date)
	well_drill_period.End_date = formatDateString(well_drill_period.End_date)
	well_drill_period.End_time = formatDateString(well_drill_period.End_time)
	well_drill_period.Expiry_date = formatDateString(well_drill_period.Expiry_date)
	well_drill_period.Reported_rig_rel_date = formatDateString(well_drill_period.Reported_rig_rel_date)
	well_drill_period.Reported_rig_rel_time = formatDateString(well_drill_period.Reported_rig_rel_time)
	well_drill_period.Reported_spud_date = formatDateString(well_drill_period.Reported_spud_date)
	well_drill_period.Reported_spud_time = formatDateString(well_drill_period.Reported_spud_time)
	well_drill_period.Start_date = formatDateString(well_drill_period.Start_date)
	well_drill_period.Start_time = formatDateString(well_drill_period.Start_time)
	well_drill_period.Row_effective_date = formatDateString(well_drill_period.Row_effective_date)
	well_drill_period.Row_expiry_date = formatDateString(well_drill_period.Row_expiry_date)






	rows, err := stmt.Exec(well_drill_period.Period_obs_no, well_drill_period.Active_ind, well_drill_period.Bha_weight, well_drill_period.Bha_weight_ouom, well_drill_period.Boiler_fuel, well_drill_period.Boiler_fuel_ouom, well_drill_period.Camp_ind, well_drill_period.Casing_press_max, well_drill_period.Casing_press_max_ouom, well_drill_period.Circ_pressure, well_drill_period.Circ_pressure_max, well_drill_period.Circ_pressure_max_ouom, well_drill_period.Circ_pressure_min, well_drill_period.Circ_pressure_min_ouom, well_drill_period.Circ_pressure_ouom, well_drill_period.Circ_pressure_type, well_drill_period.Contractor_ba_id, well_drill_period.Contractor_rep_ba_id, well_drill_period.Contractor_rep_name, well_drill_period.Contractor_rep_signed_ind, well_drill_period.Cum_crew_safety_days, well_drill_period.Cum_rig_safety_days, well_drill_period.Drill_activity_set_id, well_drill_period.Drill_activity_type_id, well_drill_period.Drill_pipe_single_count, well_drill_period.Drill_pipe_single_len, well_drill_period.Drill_pipe_single_len_ouom, well_drill_period.Drill_pipe_stand_count, well_drill_period.Drill_pipe_stand_len, well_drill_period.Drill_pipe_stand_len_ouom, well_drill_period.Drill_string_torque, well_drill_period.Drill_string_torque_ouom, well_drill_period.Effective_date, well_drill_period.End_date, well_drill_period.End_time, well_drill_period.End_timezone, well_drill_period.Expiry_date, well_drill_period.Finance_id, well_drill_period.Hole_drag, well_drill_period.Hole_drag_ouom, well_drill_period.Hole_fill_length, well_drill_period.Hole_fill_length_ouom, well_drill_period.Hookload_max, well_drill_period.Hookload_max_ouom, well_drill_period.Job_num, well_drill_period.Kb_down_length, well_drill_period.Kb_down_length_ouom, well_drill_period.Mud_system_type, well_drill_period.Mud_type_oil_ind, well_drill_period.Mud_type_water_ind, well_drill_period.Operator_ba_id, well_drill_period.Operator_fuel, well_drill_period.Operator_fuel_ouom, well_drill_period.Operator_fuel_uom, well_drill_period.Operator_rep_ba_id, well_drill_period.Operator_rep_name, well_drill_period.Operator_rep_signed_ind, well_drill_period.Ppdm_guid, well_drill_period.Primary_activity_ind, well_drill_period.Rate_schedule_id, well_drill_period.Reentry_ind, well_drill_period.Remark, well_drill_period.Reported_base_location, well_drill_period.Reported_kb_elev, well_drill_period.Reported_kb_elev_ouom, well_drill_period.Reported_license_num, well_drill_period.Reported_period_num, well_drill_period.Reported_profile_type, well_drill_period.Reported_rig_name, well_drill_period.Reported_rig_rel_date, well_drill_period.Reported_rig_rel_time, well_drill_period.Reported_rig_rel_timezone, well_drill_period.Reported_spud_date, well_drill_period.Reported_spud_time, well_drill_period.Reported_spud_timezone, well_drill_period.Reported_top_location, well_drill_period.Reported_uwi, well_drill_period.Reported_well_name, well_drill_period.Report_time_type, well_drill_period.Rig_fuel, well_drill_period.Rig_fuel_ouom, well_drill_period.Rig_fuel_uom, well_drill_period.Safety_meeting_topic, well_drill_period.Source, well_drill_period.Start_date, well_drill_period.Start_time, well_drill_period.Start_timezone, well_drill_period.Tot_drill_tool_length, well_drill_period.Tot_drill_tool_length_ouom, well_drill_period.Tot_drill_tool_weight, well_drill_period.Tot_drill_tool_weight_ouom, well_drill_period.Row_changed_by, well_drill_period.Row_changed_date, well_drill_period.Row_created_by, well_drill_period.Row_effective_date, well_drill_period.Row_expiry_date, well_drill_period.Row_quality, well_drill_period.Uwi)
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

func PatchWellDrillPeriod(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update well_drill_period set "
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
		} else if key == "effective_date" || key == "end_date" || key == "end_time" || key == "expiry_date" || key == "reported_rig_rel_date" || key == "reported_rig_rel_time" || key == "reported_spud_date" || key == "reported_spud_time" || key == "start_date" || key == "start_time" || key == "row_effective_date" || key == "row_expiry_date"      {
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

func DeleteWellDrillPeriod(c *fiber.Ctx) error {
	id := c.Params("id")
	var well_drill_period dto.Well_drill_period
	well_drill_period.Uwi = id

	stmt, err := db.Prepare("delete from well_drill_period where uwi = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(well_drill_period.Uwi)
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


