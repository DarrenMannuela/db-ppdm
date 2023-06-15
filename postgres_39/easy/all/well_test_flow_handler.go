package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetWellTestFlow(c *fiber.Ctx) error {
	rows, err := db.Query("select * from well_test_flow")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Well_test_flow

	for rows.Next() {
		var well_test_flow dto.Well_test_flow
		if err := rows.Scan(&well_test_flow.Uwi, &well_test_flow.Source, &well_test_flow.Test_type, &well_test_flow.Run_num, &well_test_flow.Test_num, &well_test_flow.Period_type, &well_test_flow.Period_obs_no, &well_test_flow.Fluid_type, &well_test_flow.Active_ind, &well_test_flow.Basic_sediment_ind, &well_test_flow.Blow_desc, &well_test_flow.Effective_date, &well_test_flow.Expiry_date, &well_test_flow.Gas_riser_diameter, &well_test_flow.Gas_riser_diameter_ouom, &well_test_flow.Max_fluid_rate, &well_test_flow.Max_fluid_rate_ouom, &well_test_flow.Max_fluid_rate_uom, &well_test_flow.Max_surface_pressure, &well_test_flow.Max_surface_pressure_ouom, &well_test_flow.Measurement_technique, &well_test_flow.Ppdm_guid, &well_test_flow.Remark, &well_test_flow.Shakeout_percent, &well_test_flow.Tts_time_elapsed, &well_test_flow.Tts_time_elapsed_ouom, &well_test_flow.Row_changed_by, &well_test_flow.Row_changed_date, &well_test_flow.Row_created_by, &well_test_flow.Row_created_date, &well_test_flow.Row_effective_date, &well_test_flow.Row_expiry_date, &well_test_flow.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append well_test_flow to result
		result = append(result, well_test_flow)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetWellTestFlow(c *fiber.Ctx) error {
	var well_test_flow dto.Well_test_flow

	setDefaults(&well_test_flow)

	if err := json.Unmarshal(c.Body(), &well_test_flow); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into well_test_flow values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32, :33)")
	if err != nil {
		return err
	}
	well_test_flow.Row_created_date = formatDate(well_test_flow.Row_created_date)
	well_test_flow.Row_changed_date = formatDate(well_test_flow.Row_changed_date)
	well_test_flow.Effective_date = formatDateString(well_test_flow.Effective_date)
	well_test_flow.Expiry_date = formatDateString(well_test_flow.Expiry_date)
	well_test_flow.Row_effective_date = formatDateString(well_test_flow.Row_effective_date)
	well_test_flow.Row_expiry_date = formatDateString(well_test_flow.Row_expiry_date)






	rows, err := stmt.Exec(well_test_flow.Uwi, well_test_flow.Source, well_test_flow.Test_type, well_test_flow.Run_num, well_test_flow.Test_num, well_test_flow.Period_type, well_test_flow.Period_obs_no, well_test_flow.Fluid_type, well_test_flow.Active_ind, well_test_flow.Basic_sediment_ind, well_test_flow.Blow_desc, well_test_flow.Effective_date, well_test_flow.Expiry_date, well_test_flow.Gas_riser_diameter, well_test_flow.Gas_riser_diameter_ouom, well_test_flow.Max_fluid_rate, well_test_flow.Max_fluid_rate_ouom, well_test_flow.Max_fluid_rate_uom, well_test_flow.Max_surface_pressure, well_test_flow.Max_surface_pressure_ouom, well_test_flow.Measurement_technique, well_test_flow.Ppdm_guid, well_test_flow.Remark, well_test_flow.Shakeout_percent, well_test_flow.Tts_time_elapsed, well_test_flow.Tts_time_elapsed_ouom, well_test_flow.Row_changed_by, well_test_flow.Row_changed_date, well_test_flow.Row_created_by, well_test_flow.Row_created_date, well_test_flow.Row_effective_date, well_test_flow.Row_expiry_date, well_test_flow.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateWellTestFlow(c *fiber.Ctx) error {
	var well_test_flow dto.Well_test_flow

	setDefaults(&well_test_flow)

	if err := json.Unmarshal(c.Body(), &well_test_flow); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	well_test_flow.Uwi = id

    if well_test_flow.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update well_test_flow set source = :1, test_type = :2, run_num = :3, test_num = :4, period_type = :5, period_obs_no = :6, fluid_type = :7, active_ind = :8, basic_sediment_ind = :9, blow_desc = :10, effective_date = :11, expiry_date = :12, gas_riser_diameter = :13, gas_riser_diameter_ouom = :14, max_fluid_rate = :15, max_fluid_rate_ouom = :16, max_fluid_rate_uom = :17, max_surface_pressure = :18, max_surface_pressure_ouom = :19, measurement_technique = :20, ppdm_guid = :21, remark = :22, shakeout_percent = :23, tts_time_elapsed = :24, tts_time_elapsed_ouom = :25, row_changed_by = :26, row_changed_date = :27, row_created_by = :28, row_effective_date = :29, row_expiry_date = :30, row_quality = :31 where uwi = :33")
	if err != nil {
		return err
	}

	well_test_flow.Row_changed_date = formatDate(well_test_flow.Row_changed_date)
	well_test_flow.Effective_date = formatDateString(well_test_flow.Effective_date)
	well_test_flow.Expiry_date = formatDateString(well_test_flow.Expiry_date)
	well_test_flow.Row_effective_date = formatDateString(well_test_flow.Row_effective_date)
	well_test_flow.Row_expiry_date = formatDateString(well_test_flow.Row_expiry_date)






	rows, err := stmt.Exec(well_test_flow.Source, well_test_flow.Test_type, well_test_flow.Run_num, well_test_flow.Test_num, well_test_flow.Period_type, well_test_flow.Period_obs_no, well_test_flow.Fluid_type, well_test_flow.Active_ind, well_test_flow.Basic_sediment_ind, well_test_flow.Blow_desc, well_test_flow.Effective_date, well_test_flow.Expiry_date, well_test_flow.Gas_riser_diameter, well_test_flow.Gas_riser_diameter_ouom, well_test_flow.Max_fluid_rate, well_test_flow.Max_fluid_rate_ouom, well_test_flow.Max_fluid_rate_uom, well_test_flow.Max_surface_pressure, well_test_flow.Max_surface_pressure_ouom, well_test_flow.Measurement_technique, well_test_flow.Ppdm_guid, well_test_flow.Remark, well_test_flow.Shakeout_percent, well_test_flow.Tts_time_elapsed, well_test_flow.Tts_time_elapsed_ouom, well_test_flow.Row_changed_by, well_test_flow.Row_changed_date, well_test_flow.Row_created_by, well_test_flow.Row_effective_date, well_test_flow.Row_expiry_date, well_test_flow.Row_quality, well_test_flow.Uwi)
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

func PatchWellTestFlow(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update well_test_flow set "
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

func DeleteWellTestFlow(c *fiber.Ctx) error {
	id := c.Params("id")
	var well_test_flow dto.Well_test_flow
	well_test_flow.Uwi = id

	stmt, err := db.Prepare("delete from well_test_flow where uwi = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(well_test_flow.Uwi)
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


