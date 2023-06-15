package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetWellActivity(c *fiber.Ctx) error {
	rows, err := db.Query("select * from well_activity")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Well_activity

	for rows.Next() {
		var well_activity dto.Well_activity
		if err := rows.Scan(&well_activity.Uwi, &well_activity.Source, &well_activity.Activity_obs_no, &well_activity.Active_ind, &well_activity.Activity_duration, &well_activity.Activity_duration_ouom, &well_activity.Activity_product, &well_activity.Activity_set_id, &well_activity.Activity_type_id, &well_activity.Base_depth, &well_activity.Base_depth_ouom, &well_activity.Base_strat_unit_id, &well_activity.Blowout_fluid, &well_activity.Control_date, &well_activity.Control_depth, &well_activity.Control_depth_ouom, &well_activity.Downtime_type, &well_activity.Effective_date, &well_activity.End_date, &well_activity.End_time, &well_activity.End_timezone, &well_activity.Event_date, &well_activity.Event_depth, &well_activity.Event_depth_ouom, &well_activity.Expiry_date, &well_activity.Final_mud_density, &well_activity.Final_mud_density_ouom, &well_activity.Final_mud_viscosity, &well_activity.Final_mud_viscosity_ouom, &well_activity.Lithology, &well_activity.Lost_circ_severity, &well_activity.Lost_material_amount, &well_activity.Lost_material_amount_ouom, &well_activity.Lost_material_type, &well_activity.Lost_volume, &well_activity.Lost_volume_ouom, &well_activity.Lost_volume_product, &well_activity.Lost_volume_uom, &well_activity.Period_obs_no, &well_activity.Ppdm_guid, &well_activity.Prod_string_id, &well_activity.Prod_string_source, &well_activity.Pr_str_form_obs_no, &well_activity.Remark, &well_activity.Reported_code, &well_activity.Reported_time_elapsed, &well_activity.Reported_time_elapsed_ouom, &well_activity.Reported_tvd, &well_activity.Reported_tvd_ouom, &well_activity.Start_date, &well_activity.Start_mud_density, &well_activity.Start_mud_density_ouom, &well_activity.Start_mud_viscosity, &well_activity.Start_mud_viscosity_ouom, &well_activity.Start_time, &well_activity.Start_timezone, &well_activity.Strat_name_set_id, &well_activity.Top_depth, &well_activity.Top_depth_ouom, &well_activity.Top_strat_unit_id, &well_activity.Total_time_elapsed, &well_activity.Total_time_elapsed_ouom, &well_activity.Row_changed_by, &well_activity.Row_changed_date, &well_activity.Row_created_by, &well_activity.Row_created_date, &well_activity.Row_effective_date, &well_activity.Row_expiry_date, &well_activity.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append well_activity to result
		result = append(result, well_activity)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetWellActivity(c *fiber.Ctx) error {
	var well_activity dto.Well_activity

	setDefaults(&well_activity)

	if err := json.Unmarshal(c.Body(), &well_activity); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into well_activity values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32, :33, :34, :35, :36, :37, :38, :39, :40, :41, :42, :43, :44, :45, :46, :47, :48, :49, :50, :51, :52, :53, :54, :55, :56, :57, :58, :59, :60, :61, :62, :63, :64, :65, :66, :67, :68, :69)")
	if err != nil {
		return err
	}
	well_activity.Row_created_date = formatDate(well_activity.Row_created_date)
	well_activity.Row_changed_date = formatDate(well_activity.Row_changed_date)
	well_activity.Control_date = formatDateString(well_activity.Control_date)
	well_activity.Effective_date = formatDateString(well_activity.Effective_date)
	well_activity.End_date = formatDateString(well_activity.End_date)
	well_activity.End_time = formatDateString(well_activity.End_time)
	well_activity.Event_date = formatDateString(well_activity.Event_date)
	well_activity.Expiry_date = formatDateString(well_activity.Expiry_date)
	well_activity.Start_date = formatDateString(well_activity.Start_date)
	well_activity.Start_time = formatDateString(well_activity.Start_time)
	well_activity.Row_effective_date = formatDateString(well_activity.Row_effective_date)
	well_activity.Row_expiry_date = formatDateString(well_activity.Row_expiry_date)






	rows, err := stmt.Exec(well_activity.Uwi, well_activity.Source, well_activity.Activity_obs_no, well_activity.Active_ind, well_activity.Activity_duration, well_activity.Activity_duration_ouom, well_activity.Activity_product, well_activity.Activity_set_id, well_activity.Activity_type_id, well_activity.Base_depth, well_activity.Base_depth_ouom, well_activity.Base_strat_unit_id, well_activity.Blowout_fluid, well_activity.Control_date, well_activity.Control_depth, well_activity.Control_depth_ouom, well_activity.Downtime_type, well_activity.Effective_date, well_activity.End_date, well_activity.End_time, well_activity.End_timezone, well_activity.Event_date, well_activity.Event_depth, well_activity.Event_depth_ouom, well_activity.Expiry_date, well_activity.Final_mud_density, well_activity.Final_mud_density_ouom, well_activity.Final_mud_viscosity, well_activity.Final_mud_viscosity_ouom, well_activity.Lithology, well_activity.Lost_circ_severity, well_activity.Lost_material_amount, well_activity.Lost_material_amount_ouom, well_activity.Lost_material_type, well_activity.Lost_volume, well_activity.Lost_volume_ouom, well_activity.Lost_volume_product, well_activity.Lost_volume_uom, well_activity.Period_obs_no, well_activity.Ppdm_guid, well_activity.Prod_string_id, well_activity.Prod_string_source, well_activity.Pr_str_form_obs_no, well_activity.Remark, well_activity.Reported_code, well_activity.Reported_time_elapsed, well_activity.Reported_time_elapsed_ouom, well_activity.Reported_tvd, well_activity.Reported_tvd_ouom, well_activity.Start_date, well_activity.Start_mud_density, well_activity.Start_mud_density_ouom, well_activity.Start_mud_viscosity, well_activity.Start_mud_viscosity_ouom, well_activity.Start_time, well_activity.Start_timezone, well_activity.Strat_name_set_id, well_activity.Top_depth, well_activity.Top_depth_ouom, well_activity.Top_strat_unit_id, well_activity.Total_time_elapsed, well_activity.Total_time_elapsed_ouom, well_activity.Row_changed_by, well_activity.Row_changed_date, well_activity.Row_created_by, well_activity.Row_created_date, well_activity.Row_effective_date, well_activity.Row_expiry_date, well_activity.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateWellActivity(c *fiber.Ctx) error {
	var well_activity dto.Well_activity

	setDefaults(&well_activity)

	if err := json.Unmarshal(c.Body(), &well_activity); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	well_activity.Uwi = id

    if well_activity.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update well_activity set source = :1, activity_obs_no = :2, active_ind = :3, activity_duration = :4, activity_duration_ouom = :5, activity_product = :6, activity_set_id = :7, activity_type_id = :8, base_depth = :9, base_depth_ouom = :10, base_strat_unit_id = :11, blowout_fluid = :12, control_date = :13, control_depth = :14, control_depth_ouom = :15, downtime_type = :16, effective_date = :17, end_date = :18, end_time = :19, end_timezone = :20, event_date = :21, event_depth = :22, event_depth_ouom = :23, expiry_date = :24, final_mud_density = :25, final_mud_density_ouom = :26, final_mud_viscosity = :27, final_mud_viscosity_ouom = :28, lithology = :29, lost_circ_severity = :30, lost_material_amount = :31, lost_material_amount_ouom = :32, lost_material_type = :33, lost_volume = :34, lost_volume_ouom = :35, lost_volume_product = :36, lost_volume_uom = :37, period_obs_no = :38, ppdm_guid = :39, prod_string_id = :40, prod_string_source = :41, pr_str_form_obs_no = :42, remark = :43, reported_code = :44, reported_time_elapsed = :45, reported_time_elapsed_ouom = :46, reported_tvd = :47, reported_tvd_ouom = :48, start_date = :49, start_mud_density = :50, start_mud_density_ouom = :51, start_mud_viscosity = :52, start_mud_viscosity_ouom = :53, start_time = :54, start_timezone = :55, strat_name_set_id = :56, top_depth = :57, top_depth_ouom = :58, top_strat_unit_id = :59, total_time_elapsed = :60, total_time_elapsed_ouom = :61, row_changed_by = :62, row_changed_date = :63, row_created_by = :64, row_effective_date = :65, row_expiry_date = :66, row_quality = :67 where uwi = :69")
	if err != nil {
		return err
	}

	well_activity.Row_changed_date = formatDate(well_activity.Row_changed_date)
	well_activity.Control_date = formatDateString(well_activity.Control_date)
	well_activity.Effective_date = formatDateString(well_activity.Effective_date)
	well_activity.End_date = formatDateString(well_activity.End_date)
	well_activity.End_time = formatDateString(well_activity.End_time)
	well_activity.Event_date = formatDateString(well_activity.Event_date)
	well_activity.Expiry_date = formatDateString(well_activity.Expiry_date)
	well_activity.Start_date = formatDateString(well_activity.Start_date)
	well_activity.Start_time = formatDateString(well_activity.Start_time)
	well_activity.Row_effective_date = formatDateString(well_activity.Row_effective_date)
	well_activity.Row_expiry_date = formatDateString(well_activity.Row_expiry_date)






	rows, err := stmt.Exec(well_activity.Source, well_activity.Activity_obs_no, well_activity.Active_ind, well_activity.Activity_duration, well_activity.Activity_duration_ouom, well_activity.Activity_product, well_activity.Activity_set_id, well_activity.Activity_type_id, well_activity.Base_depth, well_activity.Base_depth_ouom, well_activity.Base_strat_unit_id, well_activity.Blowout_fluid, well_activity.Control_date, well_activity.Control_depth, well_activity.Control_depth_ouom, well_activity.Downtime_type, well_activity.Effective_date, well_activity.End_date, well_activity.End_time, well_activity.End_timezone, well_activity.Event_date, well_activity.Event_depth, well_activity.Event_depth_ouom, well_activity.Expiry_date, well_activity.Final_mud_density, well_activity.Final_mud_density_ouom, well_activity.Final_mud_viscosity, well_activity.Final_mud_viscosity_ouom, well_activity.Lithology, well_activity.Lost_circ_severity, well_activity.Lost_material_amount, well_activity.Lost_material_amount_ouom, well_activity.Lost_material_type, well_activity.Lost_volume, well_activity.Lost_volume_ouom, well_activity.Lost_volume_product, well_activity.Lost_volume_uom, well_activity.Period_obs_no, well_activity.Ppdm_guid, well_activity.Prod_string_id, well_activity.Prod_string_source, well_activity.Pr_str_form_obs_no, well_activity.Remark, well_activity.Reported_code, well_activity.Reported_time_elapsed, well_activity.Reported_time_elapsed_ouom, well_activity.Reported_tvd, well_activity.Reported_tvd_ouom, well_activity.Start_date, well_activity.Start_mud_density, well_activity.Start_mud_density_ouom, well_activity.Start_mud_viscosity, well_activity.Start_mud_viscosity_ouom, well_activity.Start_time, well_activity.Start_timezone, well_activity.Strat_name_set_id, well_activity.Top_depth, well_activity.Top_depth_ouom, well_activity.Top_strat_unit_id, well_activity.Total_time_elapsed, well_activity.Total_time_elapsed_ouom, well_activity.Row_changed_by, well_activity.Row_changed_date, well_activity.Row_created_by, well_activity.Row_effective_date, well_activity.Row_expiry_date, well_activity.Row_quality, well_activity.Uwi)
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

func PatchWellActivity(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update well_activity set "
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
		} else if key == "control_date" || key == "effective_date" || key == "end_date" || key == "end_time" || key == "event_date" || key == "expiry_date" || key == "start_date" || key == "start_time" || key == "row_effective_date" || key == "row_expiry_date"      {
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

func DeleteWellActivity(c *fiber.Ctx) error {
	id := c.Params("id")
	var well_activity dto.Well_activity
	well_activity.Uwi = id

	stmt, err := db.Prepare("delete from well_activity where uwi = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(well_activity.Uwi)
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


