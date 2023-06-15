package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetPpdmColumn(c *fiber.Ctx) error {
	rows, err := db.Query("select * from ppdm_column")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Ppdm_column

	for rows.Next() {
		var ppdm_column dto.Ppdm_column
		if err := rows.Scan(&ppdm_column.System_id, &ppdm_column.Table_name, &ppdm_column.Column_name, &ppdm_column.Active_ind, &ppdm_column.Column_comment, &ppdm_column.Column_key_method, &ppdm_column.Column_precision, &ppdm_column.Column_scale, &ppdm_column.Column_sequence, &ppdm_column.Column_size, &ppdm_column.Control_column_ind, &ppdm_column.Control_column_name, &ppdm_column.Data_type, &ppdm_column.Default_ind_value, &ppdm_column.Default_uom_id, &ppdm_column.Default_value_method, &ppdm_column.Distinct_value_count, &ppdm_column.Distinct_value_count_date, &ppdm_column.Domain, &ppdm_column.Effective_date, &ppdm_column.Expiry_date, &ppdm_column.Extension_ind, &ppdm_column.Last_system_key, &ppdm_column.Nullable_ind, &ppdm_column.Ouom_column_name, &ppdm_column.Ppdm_guid, &ppdm_column.Primary_key_sequence, &ppdm_column.Remark, &ppdm_column.Source, &ppdm_column.Surrogate_ind, &ppdm_column.Uom_column_name, &ppdm_column.Row_changed_by, &ppdm_column.Row_changed_date, &ppdm_column.Row_created_by, &ppdm_column.Row_created_date, &ppdm_column.Row_effective_date, &ppdm_column.Row_expiry_date, &ppdm_column.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append ppdm_column to result
		result = append(result, ppdm_column)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetPpdmColumn(c *fiber.Ctx) error {
	var ppdm_column dto.Ppdm_column

	setDefaults(&ppdm_column)

	if err := json.Unmarshal(c.Body(), &ppdm_column); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into ppdm_column values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32, :33, :34, :35, :36, :37, :38)")
	if err != nil {
		return err
	}
	ppdm_column.Row_created_date = formatDate(ppdm_column.Row_created_date)
	ppdm_column.Row_changed_date = formatDate(ppdm_column.Row_changed_date)
	ppdm_column.Distinct_value_count_date = formatDateString(ppdm_column.Distinct_value_count_date)
	ppdm_column.Effective_date = formatDateString(ppdm_column.Effective_date)
	ppdm_column.Expiry_date = formatDateString(ppdm_column.Expiry_date)
	ppdm_column.Row_effective_date = formatDateString(ppdm_column.Row_effective_date)
	ppdm_column.Row_expiry_date = formatDateString(ppdm_column.Row_expiry_date)






	rows, err := stmt.Exec(ppdm_column.System_id, ppdm_column.Table_name, ppdm_column.Column_name, ppdm_column.Active_ind, ppdm_column.Column_comment, ppdm_column.Column_key_method, ppdm_column.Column_precision, ppdm_column.Column_scale, ppdm_column.Column_sequence, ppdm_column.Column_size, ppdm_column.Control_column_ind, ppdm_column.Control_column_name, ppdm_column.Data_type, ppdm_column.Default_ind_value, ppdm_column.Default_uom_id, ppdm_column.Default_value_method, ppdm_column.Distinct_value_count, ppdm_column.Distinct_value_count_date, ppdm_column.Domain, ppdm_column.Effective_date, ppdm_column.Expiry_date, ppdm_column.Extension_ind, ppdm_column.Last_system_key, ppdm_column.Nullable_ind, ppdm_column.Ouom_column_name, ppdm_column.Ppdm_guid, ppdm_column.Primary_key_sequence, ppdm_column.Remark, ppdm_column.Source, ppdm_column.Surrogate_ind, ppdm_column.Uom_column_name, ppdm_column.Row_changed_by, ppdm_column.Row_changed_date, ppdm_column.Row_created_by, ppdm_column.Row_created_date, ppdm_column.Row_effective_date, ppdm_column.Row_expiry_date, ppdm_column.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdatePpdmColumn(c *fiber.Ctx) error {
	var ppdm_column dto.Ppdm_column

	setDefaults(&ppdm_column)

	if err := json.Unmarshal(c.Body(), &ppdm_column); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	ppdm_column.System_id = id

    if ppdm_column.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update ppdm_column set table_name = :1, column_name = :2, active_ind = :3, column_comment = :4, column_key_method = :5, column_precision = :6, column_scale = :7, column_sequence = :8, column_size = :9, control_column_ind = :10, control_column_name = :11, data_type = :12, default_ind_value = :13, default_uom_id = :14, default_value_method = :15, distinct_value_count = :16, distinct_value_count_date = :17, domain = :18, effective_date = :19, expiry_date = :20, extension_ind = :21, last_system_key = :22, nullable_ind = :23, ouom_column_name = :24, ppdm_guid = :25, primary_key_sequence = :26, remark = :27, source = :28, surrogate_ind = :29, uom_column_name = :30, row_changed_by = :31, row_changed_date = :32, row_created_by = :33, row_effective_date = :34, row_expiry_date = :35, row_quality = :36 where system_id = :38")
	if err != nil {
		return err
	}

	ppdm_column.Row_changed_date = formatDate(ppdm_column.Row_changed_date)
	ppdm_column.Distinct_value_count_date = formatDateString(ppdm_column.Distinct_value_count_date)
	ppdm_column.Effective_date = formatDateString(ppdm_column.Effective_date)
	ppdm_column.Expiry_date = formatDateString(ppdm_column.Expiry_date)
	ppdm_column.Row_effective_date = formatDateString(ppdm_column.Row_effective_date)
	ppdm_column.Row_expiry_date = formatDateString(ppdm_column.Row_expiry_date)






	rows, err := stmt.Exec(ppdm_column.Table_name, ppdm_column.Column_name, ppdm_column.Active_ind, ppdm_column.Column_comment, ppdm_column.Column_key_method, ppdm_column.Column_precision, ppdm_column.Column_scale, ppdm_column.Column_sequence, ppdm_column.Column_size, ppdm_column.Control_column_ind, ppdm_column.Control_column_name, ppdm_column.Data_type, ppdm_column.Default_ind_value, ppdm_column.Default_uom_id, ppdm_column.Default_value_method, ppdm_column.Distinct_value_count, ppdm_column.Distinct_value_count_date, ppdm_column.Domain, ppdm_column.Effective_date, ppdm_column.Expiry_date, ppdm_column.Extension_ind, ppdm_column.Last_system_key, ppdm_column.Nullable_ind, ppdm_column.Ouom_column_name, ppdm_column.Ppdm_guid, ppdm_column.Primary_key_sequence, ppdm_column.Remark, ppdm_column.Source, ppdm_column.Surrogate_ind, ppdm_column.Uom_column_name, ppdm_column.Row_changed_by, ppdm_column.Row_changed_date, ppdm_column.Row_created_by, ppdm_column.Row_effective_date, ppdm_column.Row_expiry_date, ppdm_column.Row_quality, ppdm_column.System_id)
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

func PatchPpdmColumn(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update ppdm_column set "
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
		} else if key == "distinct_value_count_date" || key == "effective_date" || key == "expiry_date" || key == "row_effective_date" || key == "row_expiry_date"      {
			if date, ok := value.(string); ok && len(date) > 0 {
				formattedDate := formatDateString(&date)
				value = formattedDate
			}
		}
		values = append(values, value)
	}
	query += " where system_id = :id"

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

func DeletePpdmColumn(c *fiber.Ctx) error {
	id := c.Params("id")
	var ppdm_column dto.Ppdm_column
	ppdm_column.System_id = id

	stmt, err := db.Prepare("delete from ppdm_column where system_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(ppdm_column.System_id)
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


