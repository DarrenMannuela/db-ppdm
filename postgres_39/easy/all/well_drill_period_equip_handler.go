package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetWellDrillPeriodEquip(c *fiber.Ctx) error {
	rows, err := db.Query("select * from well_drill_period_equip")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Well_drill_period_equip

	for rows.Next() {
		var well_drill_period_equip dto.Well_drill_period_equip
		if err := rows.Scan(&well_drill_period_equip.Uwi, &well_drill_period_equip.Period_obs_no, &well_drill_period_equip.Equipment_id, &well_drill_period_equip.Active_ind, &well_drill_period_equip.Avg_boiler_ph, &well_drill_period_equip.Avg_boiler_stack_temp, &well_drill_period_equip.Avg_boiler_stack_temp_ouom, &well_drill_period_equip.Avg_pump_pressure, &well_drill_period_equip.Avg_pump_pressure_ouom, &well_drill_period_equip.Booked_time_elapsed, &well_drill_period_equip.Booked_time_elapsed_ouom, &well_drill_period_equip.Effective_date, &well_drill_period_equip.Equipment_obs_no, &well_drill_period_equip.Equip_ref_num, &well_drill_period_equip.Expiry_date, &well_drill_period_equip.Intake_density, &well_drill_period_equip.Intake_density_ouom, &well_drill_period_equip.Overflow_density, &well_drill_period_equip.Overflow_density_ouom, &well_drill_period_equip.Ppdm_guid, &well_drill_period_equip.Pump_liner_inside_diam, &well_drill_period_equip.Pump_liner_inside_diam_ouom, &well_drill_period_equip.Pump_stroke, &well_drill_period_equip.Pump_stroke_ouom, &well_drill_period_equip.Reduced_pump_depth, &well_drill_period_equip.Reduced_pump_depth_ouom, &well_drill_period_equip.Reduced_pump_press, &well_drill_period_equip.Reduced_pump_press_ouom, &well_drill_period_equip.Reduced_pump_stroke, &well_drill_period_equip.Reduced_pump_stroke_ouom, &well_drill_period_equip.Remark, &well_drill_period_equip.Source, &well_drill_period_equip.Tubing_obs_no, &well_drill_period_equip.Tubing_source, &well_drill_period_equip.Tubing_type, &well_drill_period_equip.Underflow_density, &well_drill_period_equip.Underflow_density_ouom, &well_drill_period_equip.Row_changed_by, &well_drill_period_equip.Row_changed_date, &well_drill_period_equip.Row_created_by, &well_drill_period_equip.Row_created_date, &well_drill_period_equip.Row_effective_date, &well_drill_period_equip.Row_expiry_date, &well_drill_period_equip.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append well_drill_period_equip to result
		result = append(result, well_drill_period_equip)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetWellDrillPeriodEquip(c *fiber.Ctx) error {
	var well_drill_period_equip dto.Well_drill_period_equip

	setDefaults(&well_drill_period_equip)

	if err := json.Unmarshal(c.Body(), &well_drill_period_equip); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into well_drill_period_equip values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32, :33, :34, :35, :36, :37, :38, :39, :40, :41, :42, :43, :44)")
	if err != nil {
		return err
	}
	well_drill_period_equip.Row_created_date = formatDate(well_drill_period_equip.Row_created_date)
	well_drill_period_equip.Row_changed_date = formatDate(well_drill_period_equip.Row_changed_date)
	well_drill_period_equip.Effective_date = formatDateString(well_drill_period_equip.Effective_date)
	well_drill_period_equip.Expiry_date = formatDateString(well_drill_period_equip.Expiry_date)
	well_drill_period_equip.Row_effective_date = formatDateString(well_drill_period_equip.Row_effective_date)
	well_drill_period_equip.Row_expiry_date = formatDateString(well_drill_period_equip.Row_expiry_date)






	rows, err := stmt.Exec(well_drill_period_equip.Uwi, well_drill_period_equip.Period_obs_no, well_drill_period_equip.Equipment_id, well_drill_period_equip.Active_ind, well_drill_period_equip.Avg_boiler_ph, well_drill_period_equip.Avg_boiler_stack_temp, well_drill_period_equip.Avg_boiler_stack_temp_ouom, well_drill_period_equip.Avg_pump_pressure, well_drill_period_equip.Avg_pump_pressure_ouom, well_drill_period_equip.Booked_time_elapsed, well_drill_period_equip.Booked_time_elapsed_ouom, well_drill_period_equip.Effective_date, well_drill_period_equip.Equipment_obs_no, well_drill_period_equip.Equip_ref_num, well_drill_period_equip.Expiry_date, well_drill_period_equip.Intake_density, well_drill_period_equip.Intake_density_ouom, well_drill_period_equip.Overflow_density, well_drill_period_equip.Overflow_density_ouom, well_drill_period_equip.Ppdm_guid, well_drill_period_equip.Pump_liner_inside_diam, well_drill_period_equip.Pump_liner_inside_diam_ouom, well_drill_period_equip.Pump_stroke, well_drill_period_equip.Pump_stroke_ouom, well_drill_period_equip.Reduced_pump_depth, well_drill_period_equip.Reduced_pump_depth_ouom, well_drill_period_equip.Reduced_pump_press, well_drill_period_equip.Reduced_pump_press_ouom, well_drill_period_equip.Reduced_pump_stroke, well_drill_period_equip.Reduced_pump_stroke_ouom, well_drill_period_equip.Remark, well_drill_period_equip.Source, well_drill_period_equip.Tubing_obs_no, well_drill_period_equip.Tubing_source, well_drill_period_equip.Tubing_type, well_drill_period_equip.Underflow_density, well_drill_period_equip.Underflow_density_ouom, well_drill_period_equip.Row_changed_by, well_drill_period_equip.Row_changed_date, well_drill_period_equip.Row_created_by, well_drill_period_equip.Row_created_date, well_drill_period_equip.Row_effective_date, well_drill_period_equip.Row_expiry_date, well_drill_period_equip.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateWellDrillPeriodEquip(c *fiber.Ctx) error {
	var well_drill_period_equip dto.Well_drill_period_equip

	setDefaults(&well_drill_period_equip)

	if err := json.Unmarshal(c.Body(), &well_drill_period_equip); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	well_drill_period_equip.Uwi = id

    if well_drill_period_equip.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update well_drill_period_equip set period_obs_no = :1, equipment_id = :2, active_ind = :3, avg_boiler_ph = :4, avg_boiler_stack_temp = :5, avg_boiler_stack_temp_ouom = :6, avg_pump_pressure = :7, avg_pump_pressure_ouom = :8, booked_time_elapsed = :9, booked_time_elapsed_ouom = :10, effective_date = :11, equipment_obs_no = :12, equip_ref_num = :13, expiry_date = :14, intake_density = :15, intake_density_ouom = :16, overflow_density = :17, overflow_density_ouom = :18, ppdm_guid = :19, pump_liner_inside_diam = :20, pump_liner_inside_diam_ouom = :21, pump_stroke = :22, pump_stroke_ouom = :23, reduced_pump_depth = :24, reduced_pump_depth_ouom = :25, reduced_pump_press = :26, reduced_pump_press_ouom = :27, reduced_pump_stroke = :28, reduced_pump_stroke_ouom = :29, remark = :30, source = :31, tubing_obs_no = :32, tubing_source = :33, tubing_type = :34, underflow_density = :35, underflow_density_ouom = :36, row_changed_by = :37, row_changed_date = :38, row_created_by = :39, row_effective_date = :40, row_expiry_date = :41, row_quality = :42 where uwi = :44")
	if err != nil {
		return err
	}

	well_drill_period_equip.Row_changed_date = formatDate(well_drill_period_equip.Row_changed_date)
	well_drill_period_equip.Effective_date = formatDateString(well_drill_period_equip.Effective_date)
	well_drill_period_equip.Expiry_date = formatDateString(well_drill_period_equip.Expiry_date)
	well_drill_period_equip.Row_effective_date = formatDateString(well_drill_period_equip.Row_effective_date)
	well_drill_period_equip.Row_expiry_date = formatDateString(well_drill_period_equip.Row_expiry_date)






	rows, err := stmt.Exec(well_drill_period_equip.Period_obs_no, well_drill_period_equip.Equipment_id, well_drill_period_equip.Active_ind, well_drill_period_equip.Avg_boiler_ph, well_drill_period_equip.Avg_boiler_stack_temp, well_drill_period_equip.Avg_boiler_stack_temp_ouom, well_drill_period_equip.Avg_pump_pressure, well_drill_period_equip.Avg_pump_pressure_ouom, well_drill_period_equip.Booked_time_elapsed, well_drill_period_equip.Booked_time_elapsed_ouom, well_drill_period_equip.Effective_date, well_drill_period_equip.Equipment_obs_no, well_drill_period_equip.Equip_ref_num, well_drill_period_equip.Expiry_date, well_drill_period_equip.Intake_density, well_drill_period_equip.Intake_density_ouom, well_drill_period_equip.Overflow_density, well_drill_period_equip.Overflow_density_ouom, well_drill_period_equip.Ppdm_guid, well_drill_period_equip.Pump_liner_inside_diam, well_drill_period_equip.Pump_liner_inside_diam_ouom, well_drill_period_equip.Pump_stroke, well_drill_period_equip.Pump_stroke_ouom, well_drill_period_equip.Reduced_pump_depth, well_drill_period_equip.Reduced_pump_depth_ouom, well_drill_period_equip.Reduced_pump_press, well_drill_period_equip.Reduced_pump_press_ouom, well_drill_period_equip.Reduced_pump_stroke, well_drill_period_equip.Reduced_pump_stroke_ouom, well_drill_period_equip.Remark, well_drill_period_equip.Source, well_drill_period_equip.Tubing_obs_no, well_drill_period_equip.Tubing_source, well_drill_period_equip.Tubing_type, well_drill_period_equip.Underflow_density, well_drill_period_equip.Underflow_density_ouom, well_drill_period_equip.Row_changed_by, well_drill_period_equip.Row_changed_date, well_drill_period_equip.Row_created_by, well_drill_period_equip.Row_effective_date, well_drill_period_equip.Row_expiry_date, well_drill_period_equip.Row_quality, well_drill_period_equip.Uwi)
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

func PatchWellDrillPeriodEquip(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update well_drill_period_equip set "
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
		} else if key == "effective_date" || key == "expiry_date" || key == "row_effective_date" || key == "row_expiry_date"      {
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

func DeleteWellDrillPeriodEquip(c *fiber.Ctx) error {
	id := c.Params("id")
	var well_drill_period_equip dto.Well_drill_period_equip
	well_drill_period_equip.Uwi = id

	stmt, err := db.Prepare("delete from well_drill_period_equip where uwi = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(well_drill_period_equip.Uwi)
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


