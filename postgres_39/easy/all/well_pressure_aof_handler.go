package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetWellPressureAof(c *fiber.Ctx) error {
	rows, err := db.Query("select * from well_pressure_aof")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Well_pressure_aof

	for rows.Next() {
		var well_pressure_aof dto.Well_pressure_aof
		if err := rows.Scan(&well_pressure_aof.Uwi, &well_pressure_aof.Source, &well_pressure_aof.Pressure_obs_no, &well_pressure_aof.Aof_obs_no, &well_pressure_aof.Active_ind, &well_pressure_aof.Analysis_type, &well_pressure_aof.Aof_calculate_method, &well_pressure_aof.Aof_potential, &well_pressure_aof.Aof_potential_ouom, &well_pressure_aof.Aof_slope, &well_pressure_aof.Bottom_hole_pressure_method, &well_pressure_aof.Caof_rate, &well_pressure_aof.Caof_rate_ouom, &well_pressure_aof.Choke_size_desc, &well_pressure_aof.Effective_date, &well_pressure_aof.Expiry_date, &well_pressure_aof.Flow_period_duration, &well_pressure_aof.Flow_period_duration_ouom, &well_pressure_aof.Flow_pressure, &well_pressure_aof.Flow_pressure_ouom, &well_pressure_aof.Flow_rate, &well_pressure_aof.Flow_rate_ouom, &well_pressure_aof.Four_point_caof_rate, &well_pressure_aof.Four_point_caof_rate_ouom, &well_pressure_aof.Ppdm_guid, &well_pressure_aof.Remark, &well_pressure_aof.Reservoir_pressure, &well_pressure_aof.Reservoir_pressure_ouom, &well_pressure_aof.Shutin_pressure_type, &well_pressure_aof.Test_date, &well_pressure_aof.Row_changed_by, &well_pressure_aof.Row_changed_date, &well_pressure_aof.Row_created_by, &well_pressure_aof.Row_created_date, &well_pressure_aof.Row_effective_date, &well_pressure_aof.Row_expiry_date, &well_pressure_aof.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append well_pressure_aof to result
		result = append(result, well_pressure_aof)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetWellPressureAof(c *fiber.Ctx) error {
	var well_pressure_aof dto.Well_pressure_aof

	setDefaults(&well_pressure_aof)

	if err := json.Unmarshal(c.Body(), &well_pressure_aof); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into well_pressure_aof values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32, :33, :34, :35, :36, :37)")
	if err != nil {
		return err
	}
	well_pressure_aof.Row_created_date = formatDate(well_pressure_aof.Row_created_date)
	well_pressure_aof.Row_changed_date = formatDate(well_pressure_aof.Row_changed_date)
	well_pressure_aof.Effective_date = formatDateString(well_pressure_aof.Effective_date)
	well_pressure_aof.Expiry_date = formatDateString(well_pressure_aof.Expiry_date)
	well_pressure_aof.Test_date = formatDateString(well_pressure_aof.Test_date)
	well_pressure_aof.Row_effective_date = formatDateString(well_pressure_aof.Row_effective_date)
	well_pressure_aof.Row_expiry_date = formatDateString(well_pressure_aof.Row_expiry_date)






	rows, err := stmt.Exec(well_pressure_aof.Uwi, well_pressure_aof.Source, well_pressure_aof.Pressure_obs_no, well_pressure_aof.Aof_obs_no, well_pressure_aof.Active_ind, well_pressure_aof.Analysis_type, well_pressure_aof.Aof_calculate_method, well_pressure_aof.Aof_potential, well_pressure_aof.Aof_potential_ouom, well_pressure_aof.Aof_slope, well_pressure_aof.Bottom_hole_pressure_method, well_pressure_aof.Caof_rate, well_pressure_aof.Caof_rate_ouom, well_pressure_aof.Choke_size_desc, well_pressure_aof.Effective_date, well_pressure_aof.Expiry_date, well_pressure_aof.Flow_period_duration, well_pressure_aof.Flow_period_duration_ouom, well_pressure_aof.Flow_pressure, well_pressure_aof.Flow_pressure_ouom, well_pressure_aof.Flow_rate, well_pressure_aof.Flow_rate_ouom, well_pressure_aof.Four_point_caof_rate, well_pressure_aof.Four_point_caof_rate_ouom, well_pressure_aof.Ppdm_guid, well_pressure_aof.Remark, well_pressure_aof.Reservoir_pressure, well_pressure_aof.Reservoir_pressure_ouom, well_pressure_aof.Shutin_pressure_type, well_pressure_aof.Test_date, well_pressure_aof.Row_changed_by, well_pressure_aof.Row_changed_date, well_pressure_aof.Row_created_by, well_pressure_aof.Row_created_date, well_pressure_aof.Row_effective_date, well_pressure_aof.Row_expiry_date, well_pressure_aof.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateWellPressureAof(c *fiber.Ctx) error {
	var well_pressure_aof dto.Well_pressure_aof

	setDefaults(&well_pressure_aof)

	if err := json.Unmarshal(c.Body(), &well_pressure_aof); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	well_pressure_aof.Uwi = id

    if well_pressure_aof.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update well_pressure_aof set source = :1, pressure_obs_no = :2, aof_obs_no = :3, active_ind = :4, analysis_type = :5, aof_calculate_method = :6, aof_potential = :7, aof_potential_ouom = :8, aof_slope = :9, bottom_hole_pressure_method = :10, caof_rate = :11, caof_rate_ouom = :12, choke_size_desc = :13, effective_date = :14, expiry_date = :15, flow_period_duration = :16, flow_period_duration_ouom = :17, flow_pressure = :18, flow_pressure_ouom = :19, flow_rate = :20, flow_rate_ouom = :21, four_point_caof_rate = :22, four_point_caof_rate_ouom = :23, ppdm_guid = :24, remark = :25, reservoir_pressure = :26, reservoir_pressure_ouom = :27, shutin_pressure_type = :28, test_date = :29, row_changed_by = :30, row_changed_date = :31, row_created_by = :32, row_effective_date = :33, row_expiry_date = :34, row_quality = :35 where uwi = :37")
	if err != nil {
		return err
	}

	well_pressure_aof.Row_changed_date = formatDate(well_pressure_aof.Row_changed_date)
	well_pressure_aof.Effective_date = formatDateString(well_pressure_aof.Effective_date)
	well_pressure_aof.Expiry_date = formatDateString(well_pressure_aof.Expiry_date)
	well_pressure_aof.Test_date = formatDateString(well_pressure_aof.Test_date)
	well_pressure_aof.Row_effective_date = formatDateString(well_pressure_aof.Row_effective_date)
	well_pressure_aof.Row_expiry_date = formatDateString(well_pressure_aof.Row_expiry_date)






	rows, err := stmt.Exec(well_pressure_aof.Source, well_pressure_aof.Pressure_obs_no, well_pressure_aof.Aof_obs_no, well_pressure_aof.Active_ind, well_pressure_aof.Analysis_type, well_pressure_aof.Aof_calculate_method, well_pressure_aof.Aof_potential, well_pressure_aof.Aof_potential_ouom, well_pressure_aof.Aof_slope, well_pressure_aof.Bottom_hole_pressure_method, well_pressure_aof.Caof_rate, well_pressure_aof.Caof_rate_ouom, well_pressure_aof.Choke_size_desc, well_pressure_aof.Effective_date, well_pressure_aof.Expiry_date, well_pressure_aof.Flow_period_duration, well_pressure_aof.Flow_period_duration_ouom, well_pressure_aof.Flow_pressure, well_pressure_aof.Flow_pressure_ouom, well_pressure_aof.Flow_rate, well_pressure_aof.Flow_rate_ouom, well_pressure_aof.Four_point_caof_rate, well_pressure_aof.Four_point_caof_rate_ouom, well_pressure_aof.Ppdm_guid, well_pressure_aof.Remark, well_pressure_aof.Reservoir_pressure, well_pressure_aof.Reservoir_pressure_ouom, well_pressure_aof.Shutin_pressure_type, well_pressure_aof.Test_date, well_pressure_aof.Row_changed_by, well_pressure_aof.Row_changed_date, well_pressure_aof.Row_created_by, well_pressure_aof.Row_effective_date, well_pressure_aof.Row_expiry_date, well_pressure_aof.Row_quality, well_pressure_aof.Uwi)
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

func PatchWellPressureAof(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update well_pressure_aof set "
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
		} else if key == "effective_date" || key == "expiry_date" || key == "test_date" || key == "row_effective_date" || key == "row_expiry_date"      {
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

func DeleteWellPressureAof(c *fiber.Ctx) error {
	id := c.Params("id")
	var well_pressure_aof dto.Well_pressure_aof
	well_pressure_aof.Uwi = id

	stmt, err := db.Prepare("delete from well_pressure_aof where uwi = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(well_pressure_aof.Uwi)
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


