package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetWellMiscData(c *fiber.Ctx) error {
	rows, err := db.Query("select * from well_misc_data")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Well_misc_data

	for rows.Next() {
		var well_misc_data dto.Well_misc_data
		if err := rows.Scan(&well_misc_data.Uwi, &well_misc_data.Source, &well_misc_data.Misc_data_type, &well_misc_data.Misc_data_obs_no, &well_misc_data.Active_ind, &well_misc_data.Cost, &well_misc_data.Currency_conversion, &well_misc_data.Currency_ouom, &well_misc_data.Currency_uom, &well_misc_data.Data_value, &well_misc_data.Data_value_ouom, &well_misc_data.Data_value_uom, &well_misc_data.Date_format_desc, &well_misc_data.Effective_date, &well_misc_data.Expiry_date, &well_misc_data.Max_date, &well_misc_data.Max_value, &well_misc_data.Max_value_ouom, &well_misc_data.Max_value_uom, &well_misc_data.Min_date, &well_misc_data.Min_value, &well_misc_data.Min_value_ouom, &well_misc_data.Min_value_uom, &well_misc_data.Misc_data_code, &well_misc_data.Misc_data_desc, &well_misc_data.Ppdm_guid, &well_misc_data.Reference_value, &well_misc_data.Reference_value_ouom, &well_misc_data.Reference_value_type, &well_misc_data.Reference_value_uom, &well_misc_data.Remark, &well_misc_data.Row_changed_by, &well_misc_data.Row_changed_date, &well_misc_data.Row_created_by, &well_misc_data.Row_created_date, &well_misc_data.Row_effective_date, &well_misc_data.Row_expiry_date, &well_misc_data.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append well_misc_data to result
		result = append(result, well_misc_data)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetWellMiscData(c *fiber.Ctx) error {
	var well_misc_data dto.Well_misc_data

	setDefaults(&well_misc_data)

	if err := json.Unmarshal(c.Body(), &well_misc_data); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into well_misc_data values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32, :33, :34, :35, :36, :37, :38)")
	if err != nil {
		return err
	}
	well_misc_data.Row_created_date = formatDate(well_misc_data.Row_created_date)
	well_misc_data.Row_changed_date = formatDate(well_misc_data.Row_changed_date)
	well_misc_data.Date_format_desc = formatDateString(well_misc_data.Date_format_desc)
	well_misc_data.Effective_date = formatDateString(well_misc_data.Effective_date)
	well_misc_data.Expiry_date = formatDateString(well_misc_data.Expiry_date)
	well_misc_data.Max_date = formatDateString(well_misc_data.Max_date)
	well_misc_data.Min_date = formatDateString(well_misc_data.Min_date)
	well_misc_data.Row_effective_date = formatDateString(well_misc_data.Row_effective_date)
	well_misc_data.Row_expiry_date = formatDateString(well_misc_data.Row_expiry_date)






	rows, err := stmt.Exec(well_misc_data.Uwi, well_misc_data.Source, well_misc_data.Misc_data_type, well_misc_data.Misc_data_obs_no, well_misc_data.Active_ind, well_misc_data.Cost, well_misc_data.Currency_conversion, well_misc_data.Currency_ouom, well_misc_data.Currency_uom, well_misc_data.Data_value, well_misc_data.Data_value_ouom, well_misc_data.Data_value_uom, well_misc_data.Date_format_desc, well_misc_data.Effective_date, well_misc_data.Expiry_date, well_misc_data.Max_date, well_misc_data.Max_value, well_misc_data.Max_value_ouom, well_misc_data.Max_value_uom, well_misc_data.Min_date, well_misc_data.Min_value, well_misc_data.Min_value_ouom, well_misc_data.Min_value_uom, well_misc_data.Misc_data_code, well_misc_data.Misc_data_desc, well_misc_data.Ppdm_guid, well_misc_data.Reference_value, well_misc_data.Reference_value_ouom, well_misc_data.Reference_value_type, well_misc_data.Reference_value_uom, well_misc_data.Remark, well_misc_data.Row_changed_by, well_misc_data.Row_changed_date, well_misc_data.Row_created_by, well_misc_data.Row_created_date, well_misc_data.Row_effective_date, well_misc_data.Row_expiry_date, well_misc_data.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateWellMiscData(c *fiber.Ctx) error {
	var well_misc_data dto.Well_misc_data

	setDefaults(&well_misc_data)

	if err := json.Unmarshal(c.Body(), &well_misc_data); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	well_misc_data.Uwi = id

    if well_misc_data.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update well_misc_data set source = :1, misc_data_type = :2, misc_data_obs_no = :3, active_ind = :4, cost = :5, currency_conversion = :6, currency_ouom = :7, currency_uom = :8, data_value = :9, data_value_ouom = :10, data_value_uom = :11, date_format_desc = :12, effective_date = :13, expiry_date = :14, max_date = :15, max_value = :16, max_value_ouom = :17, max_value_uom = :18, min_date = :19, min_value = :20, min_value_ouom = :21, min_value_uom = :22, misc_data_code = :23, misc_data_desc = :24, ppdm_guid = :25, reference_value = :26, reference_value_ouom = :27, reference_value_type = :28, reference_value_uom = :29, remark = :30, row_changed_by = :31, row_changed_date = :32, row_created_by = :33, row_effective_date = :34, row_expiry_date = :35, row_quality = :36 where uwi = :38")
	if err != nil {
		return err
	}

	well_misc_data.Row_changed_date = formatDate(well_misc_data.Row_changed_date)
	well_misc_data.Date_format_desc = formatDateString(well_misc_data.Date_format_desc)
	well_misc_data.Effective_date = formatDateString(well_misc_data.Effective_date)
	well_misc_data.Expiry_date = formatDateString(well_misc_data.Expiry_date)
	well_misc_data.Max_date = formatDateString(well_misc_data.Max_date)
	well_misc_data.Min_date = formatDateString(well_misc_data.Min_date)
	well_misc_data.Row_effective_date = formatDateString(well_misc_data.Row_effective_date)
	well_misc_data.Row_expiry_date = formatDateString(well_misc_data.Row_expiry_date)






	rows, err := stmt.Exec(well_misc_data.Source, well_misc_data.Misc_data_type, well_misc_data.Misc_data_obs_no, well_misc_data.Active_ind, well_misc_data.Cost, well_misc_data.Currency_conversion, well_misc_data.Currency_ouom, well_misc_data.Currency_uom, well_misc_data.Data_value, well_misc_data.Data_value_ouom, well_misc_data.Data_value_uom, well_misc_data.Date_format_desc, well_misc_data.Effective_date, well_misc_data.Expiry_date, well_misc_data.Max_date, well_misc_data.Max_value, well_misc_data.Max_value_ouom, well_misc_data.Max_value_uom, well_misc_data.Min_date, well_misc_data.Min_value, well_misc_data.Min_value_ouom, well_misc_data.Min_value_uom, well_misc_data.Misc_data_code, well_misc_data.Misc_data_desc, well_misc_data.Ppdm_guid, well_misc_data.Reference_value, well_misc_data.Reference_value_ouom, well_misc_data.Reference_value_type, well_misc_data.Reference_value_uom, well_misc_data.Remark, well_misc_data.Row_changed_by, well_misc_data.Row_changed_date, well_misc_data.Row_created_by, well_misc_data.Row_effective_date, well_misc_data.Row_expiry_date, well_misc_data.Row_quality, well_misc_data.Uwi)
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

func PatchWellMiscData(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update well_misc_data set "
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
		} else if key == "date_format_desc" || key == "effective_date" || key == "expiry_date" || key == "max_date" || key == "min_date" || key == "row_effective_date" || key == "row_expiry_date"      {
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

func DeleteWellMiscData(c *fiber.Ctx) error {
	id := c.Params("id")
	var well_misc_data dto.Well_misc_data
	well_misc_data.Uwi = id

	stmt, err := db.Prepare("delete from well_misc_data where uwi = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(well_misc_data.Uwi)
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


