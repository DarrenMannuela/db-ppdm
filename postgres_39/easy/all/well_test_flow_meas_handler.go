package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetWellTestFlowMeas(c *fiber.Ctx) error {
	rows, err := db.Query("select * from well_test_flow_meas")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Well_test_flow_meas

	for rows.Next() {
		var well_test_flow_meas dto.Well_test_flow_meas
		if err := rows.Scan(&well_test_flow_meas.Uwi, &well_test_flow_meas.Source, &well_test_flow_meas.Test_type, &well_test_flow_meas.Run_num, &well_test_flow_meas.Test_num, &well_test_flow_meas.Measurement_obs_no, &well_test_flow_meas.Active_ind, &well_test_flow_meas.Effective_date, &well_test_flow_meas.Expiry_date, &well_test_flow_meas.Flow_duration, &well_test_flow_meas.Flow_duration_ouom, &well_test_flow_meas.Fluid_type, &well_test_flow_meas.Measurement_pressure, &well_test_flow_meas.Measurement_pressure_ouom, &well_test_flow_meas.Measurement_time_elapsed, &well_test_flow_meas.Measurement_time_elapsed_ouom, &well_test_flow_meas.Measurement_volume, &well_test_flow_meas.Measurement_volume_ouom, &well_test_flow_meas.Measurement_volume_uom, &well_test_flow_meas.Period_obs_no, &well_test_flow_meas.Period_type, &well_test_flow_meas.Ppdm_guid, &well_test_flow_meas.Remark, &well_test_flow_meas.Surface_choke_diameter, &well_test_flow_meas.Surface_choke_diameter_ouom, &well_test_flow_meas.Row_changed_by, &well_test_flow_meas.Row_changed_date, &well_test_flow_meas.Row_created_by, &well_test_flow_meas.Row_created_date, &well_test_flow_meas.Row_effective_date, &well_test_flow_meas.Row_expiry_date, &well_test_flow_meas.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append well_test_flow_meas to result
		result = append(result, well_test_flow_meas)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetWellTestFlowMeas(c *fiber.Ctx) error {
	var well_test_flow_meas dto.Well_test_flow_meas

	setDefaults(&well_test_flow_meas)

	if err := json.Unmarshal(c.Body(), &well_test_flow_meas); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into well_test_flow_meas values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32)")
	if err != nil {
		return err
	}
	well_test_flow_meas.Row_created_date = formatDate(well_test_flow_meas.Row_created_date)
	well_test_flow_meas.Row_changed_date = formatDate(well_test_flow_meas.Row_changed_date)
	well_test_flow_meas.Effective_date = formatDateString(well_test_flow_meas.Effective_date)
	well_test_flow_meas.Expiry_date = formatDateString(well_test_flow_meas.Expiry_date)
	well_test_flow_meas.Row_effective_date = formatDateString(well_test_flow_meas.Row_effective_date)
	well_test_flow_meas.Row_expiry_date = formatDateString(well_test_flow_meas.Row_expiry_date)






	rows, err := stmt.Exec(well_test_flow_meas.Uwi, well_test_flow_meas.Source, well_test_flow_meas.Test_type, well_test_flow_meas.Run_num, well_test_flow_meas.Test_num, well_test_flow_meas.Measurement_obs_no, well_test_flow_meas.Active_ind, well_test_flow_meas.Effective_date, well_test_flow_meas.Expiry_date, well_test_flow_meas.Flow_duration, well_test_flow_meas.Flow_duration_ouom, well_test_flow_meas.Fluid_type, well_test_flow_meas.Measurement_pressure, well_test_flow_meas.Measurement_pressure_ouom, well_test_flow_meas.Measurement_time_elapsed, well_test_flow_meas.Measurement_time_elapsed_ouom, well_test_flow_meas.Measurement_volume, well_test_flow_meas.Measurement_volume_ouom, well_test_flow_meas.Measurement_volume_uom, well_test_flow_meas.Period_obs_no, well_test_flow_meas.Period_type, well_test_flow_meas.Ppdm_guid, well_test_flow_meas.Remark, well_test_flow_meas.Surface_choke_diameter, well_test_flow_meas.Surface_choke_diameter_ouom, well_test_flow_meas.Row_changed_by, well_test_flow_meas.Row_changed_date, well_test_flow_meas.Row_created_by, well_test_flow_meas.Row_created_date, well_test_flow_meas.Row_effective_date, well_test_flow_meas.Row_expiry_date, well_test_flow_meas.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateWellTestFlowMeas(c *fiber.Ctx) error {
	var well_test_flow_meas dto.Well_test_flow_meas

	setDefaults(&well_test_flow_meas)

	if err := json.Unmarshal(c.Body(), &well_test_flow_meas); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	well_test_flow_meas.Uwi = id

    if well_test_flow_meas.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update well_test_flow_meas set source = :1, test_type = :2, run_num = :3, test_num = :4, measurement_obs_no = :5, active_ind = :6, effective_date = :7, expiry_date = :8, flow_duration = :9, flow_duration_ouom = :10, fluid_type = :11, measurement_pressure = :12, measurement_pressure_ouom = :13, measurement_time_elapsed = :14, measurement_time_elapsed_ouom = :15, measurement_volume = :16, measurement_volume_ouom = :17, measurement_volume_uom = :18, period_obs_no = :19, period_type = :20, ppdm_guid = :21, remark = :22, surface_choke_diameter = :23, surface_choke_diameter_ouom = :24, row_changed_by = :25, row_changed_date = :26, row_created_by = :27, row_effective_date = :28, row_expiry_date = :29, row_quality = :30 where uwi = :32")
	if err != nil {
		return err
	}

	well_test_flow_meas.Row_changed_date = formatDate(well_test_flow_meas.Row_changed_date)
	well_test_flow_meas.Effective_date = formatDateString(well_test_flow_meas.Effective_date)
	well_test_flow_meas.Expiry_date = formatDateString(well_test_flow_meas.Expiry_date)
	well_test_flow_meas.Row_effective_date = formatDateString(well_test_flow_meas.Row_effective_date)
	well_test_flow_meas.Row_expiry_date = formatDateString(well_test_flow_meas.Row_expiry_date)






	rows, err := stmt.Exec(well_test_flow_meas.Source, well_test_flow_meas.Test_type, well_test_flow_meas.Run_num, well_test_flow_meas.Test_num, well_test_flow_meas.Measurement_obs_no, well_test_flow_meas.Active_ind, well_test_flow_meas.Effective_date, well_test_flow_meas.Expiry_date, well_test_flow_meas.Flow_duration, well_test_flow_meas.Flow_duration_ouom, well_test_flow_meas.Fluid_type, well_test_flow_meas.Measurement_pressure, well_test_flow_meas.Measurement_pressure_ouom, well_test_flow_meas.Measurement_time_elapsed, well_test_flow_meas.Measurement_time_elapsed_ouom, well_test_flow_meas.Measurement_volume, well_test_flow_meas.Measurement_volume_ouom, well_test_flow_meas.Measurement_volume_uom, well_test_flow_meas.Period_obs_no, well_test_flow_meas.Period_type, well_test_flow_meas.Ppdm_guid, well_test_flow_meas.Remark, well_test_flow_meas.Surface_choke_diameter, well_test_flow_meas.Surface_choke_diameter_ouom, well_test_flow_meas.Row_changed_by, well_test_flow_meas.Row_changed_date, well_test_flow_meas.Row_created_by, well_test_flow_meas.Row_effective_date, well_test_flow_meas.Row_expiry_date, well_test_flow_meas.Row_quality, well_test_flow_meas.Uwi)
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

func PatchWellTestFlowMeas(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update well_test_flow_meas set "
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

func DeleteWellTestFlowMeas(c *fiber.Ctx) error {
	id := c.Params("id")
	var well_test_flow_meas dto.Well_test_flow_meas
	well_test_flow_meas.Uwi = id

	stmt, err := db.Prepare("delete from well_test_flow_meas where uwi = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(well_test_flow_meas.Uwi)
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


