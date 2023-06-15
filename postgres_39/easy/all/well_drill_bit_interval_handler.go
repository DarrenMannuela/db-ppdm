package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetWellDrillBitInterval(c *fiber.Ctx) error {
	rows, err := db.Query("select * from well_drill_bit_interval")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Well_drill_bit_interval

	for rows.Next() {
		var well_drill_bit_interval dto.Well_drill_bit_interval
		if err := rows.Scan(&well_drill_bit_interval.Uwi, &well_drill_bit_interval.Source, &well_drill_bit_interval.Bit_interval_obs_no, &well_drill_bit_interval.Active_ind, &well_drill_bit_interval.Avg_force_on_bit, &well_drill_bit_interval.Avg_force_on_bit_ouom, &well_drill_bit_interval.Avg_penetration_rate, &well_drill_bit_interval.Avg_penetration_rate_ouom, &well_drill_bit_interval.Bit_company, &well_drill_bit_interval.Bit_drilled_rate, &well_drill_bit_interval.Bit_drilled_rate_ouom, &well_drill_bit_interval.Bit_jet_count, &well_drill_bit_interval.Bit_jet_remark, &well_drill_bit_interval.Bit_number, &well_drill_bit_interval.Bit_operating_time, &well_drill_bit_interval.Bit_operating_time_ouom, &well_drill_bit_interval.Bit_remark, &well_drill_bit_interval.Bit_size, &well_drill_bit_interval.Bit_size_desc, &well_drill_bit_interval.Bit_size_ouom, &well_drill_bit_interval.Cutting_structure_loc, &well_drill_bit_interval.Cutting_structure_mdc, &well_drill_bit_interval.Cutting_structure_ti, &well_drill_bit_interval.Cutting_structure_to, &well_drill_bit_interval.Distance_drilled, &well_drill_bit_interval.Distance_drilled_ouom, &well_drill_bit_interval.Drill_bit_type, &well_drill_bit_interval.Effective_date, &well_drill_bit_interval.Equipment_id, &well_drill_bit_interval.Expiry_date, &well_drill_bit_interval.Flow_rate, &well_drill_bit_interval.Flow_rate_ouom, &well_drill_bit_interval.Gage_out_distance, &well_drill_bit_interval.Gage_out_distance_ouom, &well_drill_bit_interval.Max_force_on_bit, &well_drill_bit_interval.Max_force_on_bit_ouom, &well_drill_bit_interval.Max_penetration_rate, &well_drill_bit_interval.Max_penetration_rate_ouom, &well_drill_bit_interval.Min_force_on_bit, &well_drill_bit_interval.Min_force_on_bit_ouom, &well_drill_bit_interval.Min_penetration_rate, &well_drill_bit_interval.Min_penetration_rate_ouom, &well_drill_bit_interval.Ppdm_guid, &well_drill_bit_interval.Reason_pulled, &well_drill_bit_interval.Remark, &well_drill_bit_interval.Reported_tfa, &well_drill_bit_interval.Reported_tfa_ouom, &well_drill_bit_interval.Run_in_depth, &well_drill_bit_interval.Run_in_depth_ouom, &well_drill_bit_interval.Run_out_depth, &well_drill_bit_interval.Run_out_depth_ouom, &well_drill_bit_interval.Torque, &well_drill_bit_interval.Torque_ouom, &well_drill_bit_interval.Total_bit_revolutions, &well_drill_bit_interval.Row_changed_by, &well_drill_bit_interval.Row_changed_date, &well_drill_bit_interval.Row_created_by, &well_drill_bit_interval.Row_created_date, &well_drill_bit_interval.Row_effective_date, &well_drill_bit_interval.Row_expiry_date, &well_drill_bit_interval.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append well_drill_bit_interval to result
		result = append(result, well_drill_bit_interval)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetWellDrillBitInterval(c *fiber.Ctx) error {
	var well_drill_bit_interval dto.Well_drill_bit_interval

	setDefaults(&well_drill_bit_interval)

	if err := json.Unmarshal(c.Body(), &well_drill_bit_interval); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into well_drill_bit_interval values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32, :33, :34, :35, :36, :37, :38, :39, :40, :41, :42, :43, :44, :45, :46, :47, :48, :49, :50, :51, :52, :53, :54, :55, :56, :57, :58, :59, :60, :61)")
	if err != nil {
		return err
	}
	well_drill_bit_interval.Row_created_date = formatDate(well_drill_bit_interval.Row_created_date)
	well_drill_bit_interval.Row_changed_date = formatDate(well_drill_bit_interval.Row_changed_date)
	well_drill_bit_interval.Effective_date = formatDateString(well_drill_bit_interval.Effective_date)
	well_drill_bit_interval.Expiry_date = formatDateString(well_drill_bit_interval.Expiry_date)
	well_drill_bit_interval.Row_effective_date = formatDateString(well_drill_bit_interval.Row_effective_date)
	well_drill_bit_interval.Row_expiry_date = formatDateString(well_drill_bit_interval.Row_expiry_date)






	rows, err := stmt.Exec(well_drill_bit_interval.Uwi, well_drill_bit_interval.Source, well_drill_bit_interval.Bit_interval_obs_no, well_drill_bit_interval.Active_ind, well_drill_bit_interval.Avg_force_on_bit, well_drill_bit_interval.Avg_force_on_bit_ouom, well_drill_bit_interval.Avg_penetration_rate, well_drill_bit_interval.Avg_penetration_rate_ouom, well_drill_bit_interval.Bit_company, well_drill_bit_interval.Bit_drilled_rate, well_drill_bit_interval.Bit_drilled_rate_ouom, well_drill_bit_interval.Bit_jet_count, well_drill_bit_interval.Bit_jet_remark, well_drill_bit_interval.Bit_number, well_drill_bit_interval.Bit_operating_time, well_drill_bit_interval.Bit_operating_time_ouom, well_drill_bit_interval.Bit_remark, well_drill_bit_interval.Bit_size, well_drill_bit_interval.Bit_size_desc, well_drill_bit_interval.Bit_size_ouom, well_drill_bit_interval.Cutting_structure_loc, well_drill_bit_interval.Cutting_structure_mdc, well_drill_bit_interval.Cutting_structure_ti, well_drill_bit_interval.Cutting_structure_to, well_drill_bit_interval.Distance_drilled, well_drill_bit_interval.Distance_drilled_ouom, well_drill_bit_interval.Drill_bit_type, well_drill_bit_interval.Effective_date, well_drill_bit_interval.Equipment_id, well_drill_bit_interval.Expiry_date, well_drill_bit_interval.Flow_rate, well_drill_bit_interval.Flow_rate_ouom, well_drill_bit_interval.Gage_out_distance, well_drill_bit_interval.Gage_out_distance_ouom, well_drill_bit_interval.Max_force_on_bit, well_drill_bit_interval.Max_force_on_bit_ouom, well_drill_bit_interval.Max_penetration_rate, well_drill_bit_interval.Max_penetration_rate_ouom, well_drill_bit_interval.Min_force_on_bit, well_drill_bit_interval.Min_force_on_bit_ouom, well_drill_bit_interval.Min_penetration_rate, well_drill_bit_interval.Min_penetration_rate_ouom, well_drill_bit_interval.Ppdm_guid, well_drill_bit_interval.Reason_pulled, well_drill_bit_interval.Remark, well_drill_bit_interval.Reported_tfa, well_drill_bit_interval.Reported_tfa_ouom, well_drill_bit_interval.Run_in_depth, well_drill_bit_interval.Run_in_depth_ouom, well_drill_bit_interval.Run_out_depth, well_drill_bit_interval.Run_out_depth_ouom, well_drill_bit_interval.Torque, well_drill_bit_interval.Torque_ouom, well_drill_bit_interval.Total_bit_revolutions, well_drill_bit_interval.Row_changed_by, well_drill_bit_interval.Row_changed_date, well_drill_bit_interval.Row_created_by, well_drill_bit_interval.Row_created_date, well_drill_bit_interval.Row_effective_date, well_drill_bit_interval.Row_expiry_date, well_drill_bit_interval.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateWellDrillBitInterval(c *fiber.Ctx) error {
	var well_drill_bit_interval dto.Well_drill_bit_interval

	setDefaults(&well_drill_bit_interval)

	if err := json.Unmarshal(c.Body(), &well_drill_bit_interval); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	well_drill_bit_interval.Uwi = id

    if well_drill_bit_interval.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update well_drill_bit_interval set source = :1, bit_interval_obs_no = :2, active_ind = :3, avg_force_on_bit = :4, avg_force_on_bit_ouom = :5, avg_penetration_rate = :6, avg_penetration_rate_ouom = :7, bit_company = :8, bit_drilled_rate = :9, bit_drilled_rate_ouom = :10, bit_jet_count = :11, bit_jet_remark = :12, bit_number = :13, bit_operating_time = :14, bit_operating_time_ouom = :15, bit_remark = :16, bit_size = :17, bit_size_desc = :18, bit_size_ouom = :19, cutting_structure_loc = :20, cutting_structure_mdc = :21, cutting_structure_ti = :22, cutting_structure_to = :23, distance_drilled = :24, distance_drilled_ouom = :25, drill_bit_type = :26, effective_date = :27, equipment_id = :28, expiry_date = :29, flow_rate = :30, flow_rate_ouom = :31, gage_out_distance = :32, gage_out_distance_ouom = :33, max_force_on_bit = :34, max_force_on_bit_ouom = :35, max_penetration_rate = :36, max_penetration_rate_ouom = :37, min_force_on_bit = :38, min_force_on_bit_ouom = :39, min_penetration_rate = :40, min_penetration_rate_ouom = :41, ppdm_guid = :42, reason_pulled = :43, remark = :44, reported_tfa = :45, reported_tfa_ouom = :46, run_in_depth = :47, run_in_depth_ouom = :48, run_out_depth = :49, run_out_depth_ouom = :50, torque = :51, torque_ouom = :52, total_bit_revolutions = :53, row_changed_by = :54, row_changed_date = :55, row_created_by = :56, row_effective_date = :57, row_expiry_date = :58, row_quality = :59 where uwi = :61")
	if err != nil {
		return err
	}

	well_drill_bit_interval.Row_changed_date = formatDate(well_drill_bit_interval.Row_changed_date)
	well_drill_bit_interval.Effective_date = formatDateString(well_drill_bit_interval.Effective_date)
	well_drill_bit_interval.Expiry_date = formatDateString(well_drill_bit_interval.Expiry_date)
	well_drill_bit_interval.Row_effective_date = formatDateString(well_drill_bit_interval.Row_effective_date)
	well_drill_bit_interval.Row_expiry_date = formatDateString(well_drill_bit_interval.Row_expiry_date)






	rows, err := stmt.Exec(well_drill_bit_interval.Source, well_drill_bit_interval.Bit_interval_obs_no, well_drill_bit_interval.Active_ind, well_drill_bit_interval.Avg_force_on_bit, well_drill_bit_interval.Avg_force_on_bit_ouom, well_drill_bit_interval.Avg_penetration_rate, well_drill_bit_interval.Avg_penetration_rate_ouom, well_drill_bit_interval.Bit_company, well_drill_bit_interval.Bit_drilled_rate, well_drill_bit_interval.Bit_drilled_rate_ouom, well_drill_bit_interval.Bit_jet_count, well_drill_bit_interval.Bit_jet_remark, well_drill_bit_interval.Bit_number, well_drill_bit_interval.Bit_operating_time, well_drill_bit_interval.Bit_operating_time_ouom, well_drill_bit_interval.Bit_remark, well_drill_bit_interval.Bit_size, well_drill_bit_interval.Bit_size_desc, well_drill_bit_interval.Bit_size_ouom, well_drill_bit_interval.Cutting_structure_loc, well_drill_bit_interval.Cutting_structure_mdc, well_drill_bit_interval.Cutting_structure_ti, well_drill_bit_interval.Cutting_structure_to, well_drill_bit_interval.Distance_drilled, well_drill_bit_interval.Distance_drilled_ouom, well_drill_bit_interval.Drill_bit_type, well_drill_bit_interval.Effective_date, well_drill_bit_interval.Equipment_id, well_drill_bit_interval.Expiry_date, well_drill_bit_interval.Flow_rate, well_drill_bit_interval.Flow_rate_ouom, well_drill_bit_interval.Gage_out_distance, well_drill_bit_interval.Gage_out_distance_ouom, well_drill_bit_interval.Max_force_on_bit, well_drill_bit_interval.Max_force_on_bit_ouom, well_drill_bit_interval.Max_penetration_rate, well_drill_bit_interval.Max_penetration_rate_ouom, well_drill_bit_interval.Min_force_on_bit, well_drill_bit_interval.Min_force_on_bit_ouom, well_drill_bit_interval.Min_penetration_rate, well_drill_bit_interval.Min_penetration_rate_ouom, well_drill_bit_interval.Ppdm_guid, well_drill_bit_interval.Reason_pulled, well_drill_bit_interval.Remark, well_drill_bit_interval.Reported_tfa, well_drill_bit_interval.Reported_tfa_ouom, well_drill_bit_interval.Run_in_depth, well_drill_bit_interval.Run_in_depth_ouom, well_drill_bit_interval.Run_out_depth, well_drill_bit_interval.Run_out_depth_ouom, well_drill_bit_interval.Torque, well_drill_bit_interval.Torque_ouom, well_drill_bit_interval.Total_bit_revolutions, well_drill_bit_interval.Row_changed_by, well_drill_bit_interval.Row_changed_date, well_drill_bit_interval.Row_created_by, well_drill_bit_interval.Row_effective_date, well_drill_bit_interval.Row_expiry_date, well_drill_bit_interval.Row_quality, well_drill_bit_interval.Uwi)
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

func PatchWellDrillBitInterval(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update well_drill_bit_interval set "
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

func DeleteWellDrillBitInterval(c *fiber.Ctx) error {
	id := c.Params("id")
	var well_drill_bit_interval dto.Well_drill_bit_interval
	well_drill_bit_interval.Uwi = id

	stmt, err := db.Prepare("delete from well_drill_bit_interval where uwi = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(well_drill_bit_interval.Uwi)
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


