package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetWellDrillStatistic(c *fiber.Ctx) error {
	rows, err := db.Query("select * from well_drill_statistic")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Well_drill_statistic

	for rows.Next() {
		var well_drill_statistic dto.Well_drill_statistic
		if err := rows.Scan(&well_drill_statistic.Uwi, &well_drill_statistic.Period_obs_no, &well_drill_statistic.Statistic_type, &well_drill_statistic.Statistic_obs_no, &well_drill_statistic.Active_ind, &well_drill_statistic.Currency_conversion, &well_drill_statistic.Date_format_desc, &well_drill_statistic.Description, &well_drill_statistic.Effective_date, &well_drill_statistic.Expiry_date, &well_drill_statistic.Max_value, &well_drill_statistic.Max_value_ouom, &well_drill_statistic.Max_value_uom, &well_drill_statistic.Min_value, &well_drill_statistic.Min_value_ouom, &well_drill_statistic.Min_value_uom, &well_drill_statistic.Ppdm_guid, &well_drill_statistic.Remark, &well_drill_statistic.Source, &well_drill_statistic.Statistic_code, &well_drill_statistic.Statistic_date, &well_drill_statistic.Statistic_time, &well_drill_statistic.Statistic_timezone, &well_drill_statistic.Statistic_value, &well_drill_statistic.Statistic_value_ouom, &well_drill_statistic.Statistic_value_uom, &well_drill_statistic.Row_changed_by, &well_drill_statistic.Row_changed_date, &well_drill_statistic.Row_created_by, &well_drill_statistic.Row_created_date, &well_drill_statistic.Row_effective_date, &well_drill_statistic.Row_expiry_date, &well_drill_statistic.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append well_drill_statistic to result
		result = append(result, well_drill_statistic)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetWellDrillStatistic(c *fiber.Ctx) error {
	var well_drill_statistic dto.Well_drill_statistic

	setDefaults(&well_drill_statistic)

	if err := json.Unmarshal(c.Body(), &well_drill_statistic); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into well_drill_statistic values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32, :33)")
	if err != nil {
		return err
	}
	well_drill_statistic.Row_created_date = formatDate(well_drill_statistic.Row_created_date)
	well_drill_statistic.Row_changed_date = formatDate(well_drill_statistic.Row_changed_date)
	well_drill_statistic.Date_format_desc = formatDateString(well_drill_statistic.Date_format_desc)
	well_drill_statistic.Effective_date = formatDateString(well_drill_statistic.Effective_date)
	well_drill_statistic.Expiry_date = formatDateString(well_drill_statistic.Expiry_date)
	well_drill_statistic.Statistic_date = formatDateString(well_drill_statistic.Statistic_date)
	well_drill_statistic.Statistic_time = formatDateString(well_drill_statistic.Statistic_time)
	well_drill_statistic.Row_effective_date = formatDateString(well_drill_statistic.Row_effective_date)
	well_drill_statistic.Row_expiry_date = formatDateString(well_drill_statistic.Row_expiry_date)






	rows, err := stmt.Exec(well_drill_statistic.Uwi, well_drill_statistic.Period_obs_no, well_drill_statistic.Statistic_type, well_drill_statistic.Statistic_obs_no, well_drill_statistic.Active_ind, well_drill_statistic.Currency_conversion, well_drill_statistic.Date_format_desc, well_drill_statistic.Description, well_drill_statistic.Effective_date, well_drill_statistic.Expiry_date, well_drill_statistic.Max_value, well_drill_statistic.Max_value_ouom, well_drill_statistic.Max_value_uom, well_drill_statistic.Min_value, well_drill_statistic.Min_value_ouom, well_drill_statistic.Min_value_uom, well_drill_statistic.Ppdm_guid, well_drill_statistic.Remark, well_drill_statistic.Source, well_drill_statistic.Statistic_code, well_drill_statistic.Statistic_date, well_drill_statistic.Statistic_time, well_drill_statistic.Statistic_timezone, well_drill_statistic.Statistic_value, well_drill_statistic.Statistic_value_ouom, well_drill_statistic.Statistic_value_uom, well_drill_statistic.Row_changed_by, well_drill_statistic.Row_changed_date, well_drill_statistic.Row_created_by, well_drill_statistic.Row_created_date, well_drill_statistic.Row_effective_date, well_drill_statistic.Row_expiry_date, well_drill_statistic.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateWellDrillStatistic(c *fiber.Ctx) error {
	var well_drill_statistic dto.Well_drill_statistic

	setDefaults(&well_drill_statistic)

	if err := json.Unmarshal(c.Body(), &well_drill_statistic); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	well_drill_statistic.Uwi = id

    if well_drill_statistic.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update well_drill_statistic set period_obs_no = :1, statistic_type = :2, statistic_obs_no = :3, active_ind = :4, currency_conversion = :5, date_format_desc = :6, description = :7, effective_date = :8, expiry_date = :9, max_value = :10, max_value_ouom = :11, max_value_uom = :12, min_value = :13, min_value_ouom = :14, min_value_uom = :15, ppdm_guid = :16, remark = :17, source = :18, statistic_code = :19, statistic_date = :20, statistic_time = :21, statistic_timezone = :22, statistic_value = :23, statistic_value_ouom = :24, statistic_value_uom = :25, row_changed_by = :26, row_changed_date = :27, row_created_by = :28, row_effective_date = :29, row_expiry_date = :30, row_quality = :31 where uwi = :33")
	if err != nil {
		return err
	}

	well_drill_statistic.Row_changed_date = formatDate(well_drill_statistic.Row_changed_date)
	well_drill_statistic.Date_format_desc = formatDateString(well_drill_statistic.Date_format_desc)
	well_drill_statistic.Effective_date = formatDateString(well_drill_statistic.Effective_date)
	well_drill_statistic.Expiry_date = formatDateString(well_drill_statistic.Expiry_date)
	well_drill_statistic.Statistic_date = formatDateString(well_drill_statistic.Statistic_date)
	well_drill_statistic.Statistic_time = formatDateString(well_drill_statistic.Statistic_time)
	well_drill_statistic.Row_effective_date = formatDateString(well_drill_statistic.Row_effective_date)
	well_drill_statistic.Row_expiry_date = formatDateString(well_drill_statistic.Row_expiry_date)






	rows, err := stmt.Exec(well_drill_statistic.Period_obs_no, well_drill_statistic.Statistic_type, well_drill_statistic.Statistic_obs_no, well_drill_statistic.Active_ind, well_drill_statistic.Currency_conversion, well_drill_statistic.Date_format_desc, well_drill_statistic.Description, well_drill_statistic.Effective_date, well_drill_statistic.Expiry_date, well_drill_statistic.Max_value, well_drill_statistic.Max_value_ouom, well_drill_statistic.Max_value_uom, well_drill_statistic.Min_value, well_drill_statistic.Min_value_ouom, well_drill_statistic.Min_value_uom, well_drill_statistic.Ppdm_guid, well_drill_statistic.Remark, well_drill_statistic.Source, well_drill_statistic.Statistic_code, well_drill_statistic.Statistic_date, well_drill_statistic.Statistic_time, well_drill_statistic.Statistic_timezone, well_drill_statistic.Statistic_value, well_drill_statistic.Statistic_value_ouom, well_drill_statistic.Statistic_value_uom, well_drill_statistic.Row_changed_by, well_drill_statistic.Row_changed_date, well_drill_statistic.Row_created_by, well_drill_statistic.Row_effective_date, well_drill_statistic.Row_expiry_date, well_drill_statistic.Row_quality, well_drill_statistic.Uwi)
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

func PatchWellDrillStatistic(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update well_drill_statistic set "
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
		} else if key == "date_format_desc" || key == "effective_date" || key == "expiry_date" || key == "statistic_date" || key == "statistic_time" || key == "row_effective_date" || key == "row_expiry_date"      {
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

func DeleteWellDrillStatistic(c *fiber.Ctx) error {
	id := c.Params("id")
	var well_drill_statistic dto.Well_drill_statistic
	well_drill_statistic.Uwi = id

	stmt, err := db.Prepare("delete from well_drill_statistic where uwi = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(well_drill_statistic.Uwi)
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


