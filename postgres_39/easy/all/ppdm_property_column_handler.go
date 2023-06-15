package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetPpdmPropertyColumn(c *fiber.Ctx) error {
	rows, err := db.Query("select * from ppdm_property_column")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Ppdm_property_column

	for rows.Next() {
		var ppdm_property_column dto.Ppdm_property_column
		if err := rows.Scan(&ppdm_property_column.Property_set_id, &ppdm_property_column.Property_obs_no, &ppdm_property_column.Active_ind, &ppdm_property_column.Column_precision, &ppdm_property_column.Column_scale, &ppdm_property_column.Column_size, &ppdm_property_column.Data_type, &ppdm_property_column.Domain, &ppdm_property_column.Effective_date, &ppdm_property_column.Expiry_date, &ppdm_property_column.Ppdm_guid, &ppdm_property_column.Preferred_currency_uom, &ppdm_property_column.Preferred_uom, &ppdm_property_column.Quantity_type, &ppdm_property_column.Referenced_system_id, &ppdm_property_column.Referenced_table_name, &ppdm_property_column.Remark, &ppdm_property_column.Source, &ppdm_property_column.Use_column_name, &ppdm_property_column.Use_indicator_ind, &ppdm_property_column.Use_system_id, &ppdm_property_column.Use_table_name, &ppdm_property_column.Row_changed_by, &ppdm_property_column.Row_changed_date, &ppdm_property_column.Row_created_by, &ppdm_property_column.Row_created_date, &ppdm_property_column.Row_effective_date, &ppdm_property_column.Row_expiry_date, &ppdm_property_column.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append ppdm_property_column to result
		result = append(result, ppdm_property_column)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetPpdmPropertyColumn(c *fiber.Ctx) error {
	var ppdm_property_column dto.Ppdm_property_column

	setDefaults(&ppdm_property_column)

	if err := json.Unmarshal(c.Body(), &ppdm_property_column); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into ppdm_property_column values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29)")
	if err != nil {
		return err
	}
	ppdm_property_column.Row_created_date = formatDate(ppdm_property_column.Row_created_date)
	ppdm_property_column.Row_changed_date = formatDate(ppdm_property_column.Row_changed_date)
	ppdm_property_column.Effective_date = formatDateString(ppdm_property_column.Effective_date)
	ppdm_property_column.Expiry_date = formatDateString(ppdm_property_column.Expiry_date)
	ppdm_property_column.Row_effective_date = formatDateString(ppdm_property_column.Row_effective_date)
	ppdm_property_column.Row_expiry_date = formatDateString(ppdm_property_column.Row_expiry_date)






	rows, err := stmt.Exec(ppdm_property_column.Property_set_id, ppdm_property_column.Property_obs_no, ppdm_property_column.Active_ind, ppdm_property_column.Column_precision, ppdm_property_column.Column_scale, ppdm_property_column.Column_size, ppdm_property_column.Data_type, ppdm_property_column.Domain, ppdm_property_column.Effective_date, ppdm_property_column.Expiry_date, ppdm_property_column.Ppdm_guid, ppdm_property_column.Preferred_currency_uom, ppdm_property_column.Preferred_uom, ppdm_property_column.Quantity_type, ppdm_property_column.Referenced_system_id, ppdm_property_column.Referenced_table_name, ppdm_property_column.Remark, ppdm_property_column.Source, ppdm_property_column.Use_column_name, ppdm_property_column.Use_indicator_ind, ppdm_property_column.Use_system_id, ppdm_property_column.Use_table_name, ppdm_property_column.Row_changed_by, ppdm_property_column.Row_changed_date, ppdm_property_column.Row_created_by, ppdm_property_column.Row_created_date, ppdm_property_column.Row_effective_date, ppdm_property_column.Row_expiry_date, ppdm_property_column.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdatePpdmPropertyColumn(c *fiber.Ctx) error {
	var ppdm_property_column dto.Ppdm_property_column

	setDefaults(&ppdm_property_column)

	if err := json.Unmarshal(c.Body(), &ppdm_property_column); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	ppdm_property_column.Property_set_id = id

    if ppdm_property_column.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update ppdm_property_column set property_obs_no = :1, active_ind = :2, column_precision = :3, column_scale = :4, column_size = :5, data_type = :6, domain = :7, effective_date = :8, expiry_date = :9, ppdm_guid = :10, preferred_currency_uom = :11, preferred_uom = :12, quantity_type = :13, referenced_system_id = :14, referenced_table_name = :15, remark = :16, source = :17, use_column_name = :18, use_indicator_ind = :19, use_system_id = :20, use_table_name = :21, row_changed_by = :22, row_changed_date = :23, row_created_by = :24, row_effective_date = :25, row_expiry_date = :26, row_quality = :27 where property_set_id = :29")
	if err != nil {
		return err
	}

	ppdm_property_column.Row_changed_date = formatDate(ppdm_property_column.Row_changed_date)
	ppdm_property_column.Effective_date = formatDateString(ppdm_property_column.Effective_date)
	ppdm_property_column.Expiry_date = formatDateString(ppdm_property_column.Expiry_date)
	ppdm_property_column.Row_effective_date = formatDateString(ppdm_property_column.Row_effective_date)
	ppdm_property_column.Row_expiry_date = formatDateString(ppdm_property_column.Row_expiry_date)






	rows, err := stmt.Exec(ppdm_property_column.Property_obs_no, ppdm_property_column.Active_ind, ppdm_property_column.Column_precision, ppdm_property_column.Column_scale, ppdm_property_column.Column_size, ppdm_property_column.Data_type, ppdm_property_column.Domain, ppdm_property_column.Effective_date, ppdm_property_column.Expiry_date, ppdm_property_column.Ppdm_guid, ppdm_property_column.Preferred_currency_uom, ppdm_property_column.Preferred_uom, ppdm_property_column.Quantity_type, ppdm_property_column.Referenced_system_id, ppdm_property_column.Referenced_table_name, ppdm_property_column.Remark, ppdm_property_column.Source, ppdm_property_column.Use_column_name, ppdm_property_column.Use_indicator_ind, ppdm_property_column.Use_system_id, ppdm_property_column.Use_table_name, ppdm_property_column.Row_changed_by, ppdm_property_column.Row_changed_date, ppdm_property_column.Row_created_by, ppdm_property_column.Row_effective_date, ppdm_property_column.Row_expiry_date, ppdm_property_column.Row_quality, ppdm_property_column.Property_set_id)
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

func PatchPpdmPropertyColumn(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update ppdm_property_column set "
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
	query += " where property_set_id = :id"

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

func DeletePpdmPropertyColumn(c *fiber.Ctx) error {
	id := c.Params("id")
	var ppdm_property_column dto.Ppdm_property_column
	ppdm_property_column.Property_set_id = id

	stmt, err := db.Prepare("delete from ppdm_property_column where property_set_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(ppdm_property_column.Property_set_id)
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


