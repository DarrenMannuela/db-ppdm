package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetWellDrillBitPeriod(c *fiber.Ctx) error {
	rows, err := db.Query("select * from well_drill_bit_period")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Well_drill_bit_period

	for rows.Next() {
		var well_drill_bit_period dto.Well_drill_bit_period
		if err := rows.Scan(&well_drill_bit_period.Uwi, &well_drill_bit_period.Period_obs_no, &well_drill_bit_period.Bit_source, &well_drill_bit_period.Bit_interval_obs_no, &well_drill_bit_period.Drill_bit_period_obs_no, &well_drill_bit_period.Active_ind, &well_drill_bit_period.Avg_force_on_bit, &well_drill_bit_period.Avg_force_on_bit_ouom, &well_drill_bit_period.Avg_penetration_rate, &well_drill_bit_period.Avg_penetration_rate_ouom, &well_drill_bit_period.Avg_rotary_rpm, &well_drill_bit_period.Base_depth, &well_drill_bit_period.Base_depth_ouom, &well_drill_bit_period.Bit_operating_time, &well_drill_bit_period.Bit_operating_time_ouom, &well_drill_bit_period.Drill_op_type, &well_drill_bit_period.Effective_date, &well_drill_bit_period.Expiry_date, &well_drill_bit_period.Flow_rate, &well_drill_bit_period.Flow_rate_ouom, &well_drill_bit_period.Max_force_on_bit, &well_drill_bit_period.Max_force_on_bit_ouom, &well_drill_bit_period.Max_penetration_rate, &well_drill_bit_period.Max_penetration_rate_ouom, &well_drill_bit_period.Max_rotary_rpm, &well_drill_bit_period.Min_force_on_bit, &well_drill_bit_period.Min_force_on_bit_ouom, &well_drill_bit_period.Min_penetration_rate, &well_drill_bit_period.Min_penetration_rate_ouom, &well_drill_bit_period.Min_rotary_rpm, &well_drill_bit_period.Ppdm_guid, &well_drill_bit_period.Remark, &well_drill_bit_period.Source, &well_drill_bit_period.Top_depth, &well_drill_bit_period.Top_depth_ouom, &well_drill_bit_period.Torque, &well_drill_bit_period.Torque_ouom, &well_drill_bit_period.Total_bit_revolution, &well_drill_bit_period.Total_drilled_dist, &well_drill_bit_period.Total_drilled_dist_ouom, &well_drill_bit_period.Total_period_run, &well_drill_bit_period.Total_period_run_ouom, &well_drill_bit_period.Row_changed_by, &well_drill_bit_period.Row_changed_date, &well_drill_bit_period.Row_created_by, &well_drill_bit_period.Row_created_date, &well_drill_bit_period.Row_effective_date, &well_drill_bit_period.Row_expiry_date, &well_drill_bit_period.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append well_drill_bit_period to result
		result = append(result, well_drill_bit_period)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetWellDrillBitPeriod(c *fiber.Ctx) error {
	var well_drill_bit_period dto.Well_drill_bit_period

	setDefaults(&well_drill_bit_period)

	if err := json.Unmarshal(c.Body(), &well_drill_bit_period); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into well_drill_bit_period values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32, :33, :34, :35, :36, :37, :38, :39, :40, :41, :42, :43, :44, :45, :46, :47, :48, :49)")
	if err != nil {
		return err
	}
	well_drill_bit_period.Row_created_date = formatDate(well_drill_bit_period.Row_created_date)
	well_drill_bit_period.Row_changed_date = formatDate(well_drill_bit_period.Row_changed_date)
	well_drill_bit_period.Effective_date = formatDateString(well_drill_bit_period.Effective_date)
	well_drill_bit_period.Expiry_date = formatDateString(well_drill_bit_period.Expiry_date)
	well_drill_bit_period.Row_effective_date = formatDateString(well_drill_bit_period.Row_effective_date)
	well_drill_bit_period.Row_expiry_date = formatDateString(well_drill_bit_period.Row_expiry_date)






	rows, err := stmt.Exec(well_drill_bit_period.Uwi, well_drill_bit_period.Period_obs_no, well_drill_bit_period.Bit_source, well_drill_bit_period.Bit_interval_obs_no, well_drill_bit_period.Drill_bit_period_obs_no, well_drill_bit_period.Active_ind, well_drill_bit_period.Avg_force_on_bit, well_drill_bit_period.Avg_force_on_bit_ouom, well_drill_bit_period.Avg_penetration_rate, well_drill_bit_period.Avg_penetration_rate_ouom, well_drill_bit_period.Avg_rotary_rpm, well_drill_bit_period.Base_depth, well_drill_bit_period.Base_depth_ouom, well_drill_bit_period.Bit_operating_time, well_drill_bit_period.Bit_operating_time_ouom, well_drill_bit_period.Drill_op_type, well_drill_bit_period.Effective_date, well_drill_bit_period.Expiry_date, well_drill_bit_period.Flow_rate, well_drill_bit_period.Flow_rate_ouom, well_drill_bit_period.Max_force_on_bit, well_drill_bit_period.Max_force_on_bit_ouom, well_drill_bit_period.Max_penetration_rate, well_drill_bit_period.Max_penetration_rate_ouom, well_drill_bit_period.Max_rotary_rpm, well_drill_bit_period.Min_force_on_bit, well_drill_bit_period.Min_force_on_bit_ouom, well_drill_bit_period.Min_penetration_rate, well_drill_bit_period.Min_penetration_rate_ouom, well_drill_bit_period.Min_rotary_rpm, well_drill_bit_period.Ppdm_guid, well_drill_bit_period.Remark, well_drill_bit_period.Source, well_drill_bit_period.Top_depth, well_drill_bit_period.Top_depth_ouom, well_drill_bit_period.Torque, well_drill_bit_period.Torque_ouom, well_drill_bit_period.Total_bit_revolution, well_drill_bit_period.Total_drilled_dist, well_drill_bit_period.Total_drilled_dist_ouom, well_drill_bit_period.Total_period_run, well_drill_bit_period.Total_period_run_ouom, well_drill_bit_period.Row_changed_by, well_drill_bit_period.Row_changed_date, well_drill_bit_period.Row_created_by, well_drill_bit_period.Row_created_date, well_drill_bit_period.Row_effective_date, well_drill_bit_period.Row_expiry_date, well_drill_bit_period.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateWellDrillBitPeriod(c *fiber.Ctx) error {
	var well_drill_bit_period dto.Well_drill_bit_period

	setDefaults(&well_drill_bit_period)

	if err := json.Unmarshal(c.Body(), &well_drill_bit_period); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	well_drill_bit_period.Uwi = id

    if well_drill_bit_period.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update well_drill_bit_period set period_obs_no = :1, bit_source = :2, bit_interval_obs_no = :3, drill_bit_period_obs_no = :4, active_ind = :5, avg_force_on_bit = :6, avg_force_on_bit_ouom = :7, avg_penetration_rate = :8, avg_penetration_rate_ouom = :9, avg_rotary_rpm = :10, base_depth = :11, base_depth_ouom = :12, bit_operating_time = :13, bit_operating_time_ouom = :14, drill_op_type = :15, effective_date = :16, expiry_date = :17, flow_rate = :18, flow_rate_ouom = :19, max_force_on_bit = :20, max_force_on_bit_ouom = :21, max_penetration_rate = :22, max_penetration_rate_ouom = :23, max_rotary_rpm = :24, min_force_on_bit = :25, min_force_on_bit_ouom = :26, min_penetration_rate = :27, min_penetration_rate_ouom = :28, min_rotary_rpm = :29, ppdm_guid = :30, remark = :31, source = :32, top_depth = :33, top_depth_ouom = :34, torque = :35, torque_ouom = :36, total_bit_revolution = :37, total_drilled_dist = :38, total_drilled_dist_ouom = :39, total_period_run = :40, total_period_run_ouom = :41, row_changed_by = :42, row_changed_date = :43, row_created_by = :44, row_effective_date = :45, row_expiry_date = :46, row_quality = :47 where uwi = :49")
	if err != nil {
		return err
	}

	well_drill_bit_period.Row_changed_date = formatDate(well_drill_bit_period.Row_changed_date)
	well_drill_bit_period.Effective_date = formatDateString(well_drill_bit_period.Effective_date)
	well_drill_bit_period.Expiry_date = formatDateString(well_drill_bit_period.Expiry_date)
	well_drill_bit_period.Row_effective_date = formatDateString(well_drill_bit_period.Row_effective_date)
	well_drill_bit_period.Row_expiry_date = formatDateString(well_drill_bit_period.Row_expiry_date)






	rows, err := stmt.Exec(well_drill_bit_period.Period_obs_no, well_drill_bit_period.Bit_source, well_drill_bit_period.Bit_interval_obs_no, well_drill_bit_period.Drill_bit_period_obs_no, well_drill_bit_period.Active_ind, well_drill_bit_period.Avg_force_on_bit, well_drill_bit_period.Avg_force_on_bit_ouom, well_drill_bit_period.Avg_penetration_rate, well_drill_bit_period.Avg_penetration_rate_ouom, well_drill_bit_period.Avg_rotary_rpm, well_drill_bit_period.Base_depth, well_drill_bit_period.Base_depth_ouom, well_drill_bit_period.Bit_operating_time, well_drill_bit_period.Bit_operating_time_ouom, well_drill_bit_period.Drill_op_type, well_drill_bit_period.Effective_date, well_drill_bit_period.Expiry_date, well_drill_bit_period.Flow_rate, well_drill_bit_period.Flow_rate_ouom, well_drill_bit_period.Max_force_on_bit, well_drill_bit_period.Max_force_on_bit_ouom, well_drill_bit_period.Max_penetration_rate, well_drill_bit_period.Max_penetration_rate_ouom, well_drill_bit_period.Max_rotary_rpm, well_drill_bit_period.Min_force_on_bit, well_drill_bit_period.Min_force_on_bit_ouom, well_drill_bit_period.Min_penetration_rate, well_drill_bit_period.Min_penetration_rate_ouom, well_drill_bit_period.Min_rotary_rpm, well_drill_bit_period.Ppdm_guid, well_drill_bit_period.Remark, well_drill_bit_period.Source, well_drill_bit_period.Top_depth, well_drill_bit_period.Top_depth_ouom, well_drill_bit_period.Torque, well_drill_bit_period.Torque_ouom, well_drill_bit_period.Total_bit_revolution, well_drill_bit_period.Total_drilled_dist, well_drill_bit_period.Total_drilled_dist_ouom, well_drill_bit_period.Total_period_run, well_drill_bit_period.Total_period_run_ouom, well_drill_bit_period.Row_changed_by, well_drill_bit_period.Row_changed_date, well_drill_bit_period.Row_created_by, well_drill_bit_period.Row_effective_date, well_drill_bit_period.Row_expiry_date, well_drill_bit_period.Row_quality, well_drill_bit_period.Uwi)
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

func PatchWellDrillBitPeriod(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update well_drill_bit_period set "
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

func DeleteWellDrillBitPeriod(c *fiber.Ctx) error {
	id := c.Params("id")
	var well_drill_bit_period dto.Well_drill_bit_period
	well_drill_bit_period.Uwi = id

	stmt, err := db.Prepare("delete from well_drill_bit_period where uwi = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(well_drill_bit_period.Uwi)
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


