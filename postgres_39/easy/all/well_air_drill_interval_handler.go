package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetWellAirDrillInterval(c *fiber.Ctx) error {
	rows, err := db.Query("select * from well_air_drill_interval")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Well_air_drill_interval

	for rows.Next() {
		var well_air_drill_interval dto.Well_air_drill_interval
		if err := rows.Scan(&well_air_drill_interval.Uwi, &well_air_drill_interval.Air_drill_source, &well_air_drill_interval.Air_drill_obs_no, &well_air_drill_interval.Depth_obs_no, &well_air_drill_interval.Active_ind, &well_air_drill_interval.Air_gas_code, &well_air_drill_interval.Base_depth, &well_air_drill_interval.Base_depth_ouom, &well_air_drill_interval.Compressor_rate_volume, &well_air_drill_interval.Compressor_rate_volume_ouom, &well_air_drill_interval.Effective_date, &well_air_drill_interval.Expiry_date, &well_air_drill_interval.Max_interval_pressure, &well_air_drill_interval.Max_interval_pressure_ouom, &well_air_drill_interval.Ppdm_guid, &well_air_drill_interval.Remark, &well_air_drill_interval.Source, &well_air_drill_interval.Top_depth, &well_air_drill_interval.Top_depth_ouom, &well_air_drill_interval.Water_influx_amount, &well_air_drill_interval.Water_influx_amount_ouom, &well_air_drill_interval.Row_changed_by, &well_air_drill_interval.Row_changed_date, &well_air_drill_interval.Row_created_by, &well_air_drill_interval.Row_created_date, &well_air_drill_interval.Row_effective_date, &well_air_drill_interval.Row_expiry_date, &well_air_drill_interval.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append well_air_drill_interval to result
		result = append(result, well_air_drill_interval)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetWellAirDrillInterval(c *fiber.Ctx) error {
	var well_air_drill_interval dto.Well_air_drill_interval

	setDefaults(&well_air_drill_interval)

	if err := json.Unmarshal(c.Body(), &well_air_drill_interval); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into well_air_drill_interval values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28)")
	if err != nil {
		return err
	}
	well_air_drill_interval.Row_created_date = formatDate(well_air_drill_interval.Row_created_date)
	well_air_drill_interval.Row_changed_date = formatDate(well_air_drill_interval.Row_changed_date)
	well_air_drill_interval.Effective_date = formatDateString(well_air_drill_interval.Effective_date)
	well_air_drill_interval.Expiry_date = formatDateString(well_air_drill_interval.Expiry_date)
	well_air_drill_interval.Row_effective_date = formatDateString(well_air_drill_interval.Row_effective_date)
	well_air_drill_interval.Row_expiry_date = formatDateString(well_air_drill_interval.Row_expiry_date)






	rows, err := stmt.Exec(well_air_drill_interval.Uwi, well_air_drill_interval.Air_drill_source, well_air_drill_interval.Air_drill_obs_no, well_air_drill_interval.Depth_obs_no, well_air_drill_interval.Active_ind, well_air_drill_interval.Air_gas_code, well_air_drill_interval.Base_depth, well_air_drill_interval.Base_depth_ouom, well_air_drill_interval.Compressor_rate_volume, well_air_drill_interval.Compressor_rate_volume_ouom, well_air_drill_interval.Effective_date, well_air_drill_interval.Expiry_date, well_air_drill_interval.Max_interval_pressure, well_air_drill_interval.Max_interval_pressure_ouom, well_air_drill_interval.Ppdm_guid, well_air_drill_interval.Remark, well_air_drill_interval.Source, well_air_drill_interval.Top_depth, well_air_drill_interval.Top_depth_ouom, well_air_drill_interval.Water_influx_amount, well_air_drill_interval.Water_influx_amount_ouom, well_air_drill_interval.Row_changed_by, well_air_drill_interval.Row_changed_date, well_air_drill_interval.Row_created_by, well_air_drill_interval.Row_created_date, well_air_drill_interval.Row_effective_date, well_air_drill_interval.Row_expiry_date, well_air_drill_interval.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateWellAirDrillInterval(c *fiber.Ctx) error {
	var well_air_drill_interval dto.Well_air_drill_interval

	setDefaults(&well_air_drill_interval)

	if err := json.Unmarshal(c.Body(), &well_air_drill_interval); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	well_air_drill_interval.Uwi = id

    if well_air_drill_interval.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update well_air_drill_interval set air_drill_source = :1, air_drill_obs_no = :2, depth_obs_no = :3, active_ind = :4, air_gas_code = :5, base_depth = :6, base_depth_ouom = :7, compressor_rate_volume = :8, compressor_rate_volume_ouom = :9, effective_date = :10, expiry_date = :11, max_interval_pressure = :12, max_interval_pressure_ouom = :13, ppdm_guid = :14, remark = :15, source = :16, top_depth = :17, top_depth_ouom = :18, water_influx_amount = :19, water_influx_amount_ouom = :20, row_changed_by = :21, row_changed_date = :22, row_created_by = :23, row_effective_date = :24, row_expiry_date = :25, row_quality = :26 where uwi = :28")
	if err != nil {
		return err
	}

	well_air_drill_interval.Row_changed_date = formatDate(well_air_drill_interval.Row_changed_date)
	well_air_drill_interval.Effective_date = formatDateString(well_air_drill_interval.Effective_date)
	well_air_drill_interval.Expiry_date = formatDateString(well_air_drill_interval.Expiry_date)
	well_air_drill_interval.Row_effective_date = formatDateString(well_air_drill_interval.Row_effective_date)
	well_air_drill_interval.Row_expiry_date = formatDateString(well_air_drill_interval.Row_expiry_date)






	rows, err := stmt.Exec(well_air_drill_interval.Air_drill_source, well_air_drill_interval.Air_drill_obs_no, well_air_drill_interval.Depth_obs_no, well_air_drill_interval.Active_ind, well_air_drill_interval.Air_gas_code, well_air_drill_interval.Base_depth, well_air_drill_interval.Base_depth_ouom, well_air_drill_interval.Compressor_rate_volume, well_air_drill_interval.Compressor_rate_volume_ouom, well_air_drill_interval.Effective_date, well_air_drill_interval.Expiry_date, well_air_drill_interval.Max_interval_pressure, well_air_drill_interval.Max_interval_pressure_ouom, well_air_drill_interval.Ppdm_guid, well_air_drill_interval.Remark, well_air_drill_interval.Source, well_air_drill_interval.Top_depth, well_air_drill_interval.Top_depth_ouom, well_air_drill_interval.Water_influx_amount, well_air_drill_interval.Water_influx_amount_ouom, well_air_drill_interval.Row_changed_by, well_air_drill_interval.Row_changed_date, well_air_drill_interval.Row_created_by, well_air_drill_interval.Row_effective_date, well_air_drill_interval.Row_expiry_date, well_air_drill_interval.Row_quality, well_air_drill_interval.Uwi)
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

func PatchWellAirDrillInterval(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update well_air_drill_interval set "
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

func DeleteWellAirDrillInterval(c *fiber.Ctx) error {
	id := c.Params("id")
	var well_air_drill_interval dto.Well_air_drill_interval
	well_air_drill_interval.Uwi = id

	stmt, err := db.Prepare("delete from well_air_drill_interval where uwi = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(well_air_drill_interval.Uwi)
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


