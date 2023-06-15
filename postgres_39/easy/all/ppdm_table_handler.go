package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetPpdmTable(c *fiber.Ctx) error {
	rows, err := db.Query("select * from ppdm_table")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Ppdm_table

	for rows.Next() {
		var ppdm_table dto.Ppdm_table
		if err := rows.Scan(&ppdm_table.System_id, &ppdm_table.Table_name, &ppdm_table.Active_ind, &ppdm_table.Effective_date, &ppdm_table.Expiry_date, &ppdm_table.Extension_ind, &ppdm_table.Ppdm_guid, &ppdm_table.Primary_key_name, &ppdm_table.Remark, &ppdm_table.Row_count, &ppdm_table.Row_count_date, &ppdm_table.Source, &ppdm_table.Table_comment, &ppdm_table.Table_type, &ppdm_table.Row_changed_by, &ppdm_table.Row_changed_date, &ppdm_table.Row_created_by, &ppdm_table.Row_created_date, &ppdm_table.Row_effective_date, &ppdm_table.Row_expiry_date, &ppdm_table.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append ppdm_table to result
		result = append(result, ppdm_table)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetPpdmTable(c *fiber.Ctx) error {
	var ppdm_table dto.Ppdm_table

	setDefaults(&ppdm_table)

	if err := json.Unmarshal(c.Body(), &ppdm_table); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into ppdm_table values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21)")
	if err != nil {
		return err
	}
	ppdm_table.Row_created_date = formatDate(ppdm_table.Row_created_date)
	ppdm_table.Row_changed_date = formatDate(ppdm_table.Row_changed_date)
	ppdm_table.Effective_date = formatDateString(ppdm_table.Effective_date)
	ppdm_table.Expiry_date = formatDateString(ppdm_table.Expiry_date)
	ppdm_table.Row_count_date = formatDateString(ppdm_table.Row_count_date)
	ppdm_table.Row_effective_date = formatDateString(ppdm_table.Row_effective_date)
	ppdm_table.Row_expiry_date = formatDateString(ppdm_table.Row_expiry_date)






	rows, err := stmt.Exec(ppdm_table.System_id, ppdm_table.Table_name, ppdm_table.Active_ind, ppdm_table.Effective_date, ppdm_table.Expiry_date, ppdm_table.Extension_ind, ppdm_table.Ppdm_guid, ppdm_table.Primary_key_name, ppdm_table.Remark, ppdm_table.Row_count, ppdm_table.Row_count_date, ppdm_table.Source, ppdm_table.Table_comment, ppdm_table.Table_type, ppdm_table.Row_changed_by, ppdm_table.Row_changed_date, ppdm_table.Row_created_by, ppdm_table.Row_created_date, ppdm_table.Row_effective_date, ppdm_table.Row_expiry_date, ppdm_table.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdatePpdmTable(c *fiber.Ctx) error {
	var ppdm_table dto.Ppdm_table

	setDefaults(&ppdm_table)

	if err := json.Unmarshal(c.Body(), &ppdm_table); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	ppdm_table.System_id = id

    if ppdm_table.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update ppdm_table set table_name = :1, active_ind = :2, effective_date = :3, expiry_date = :4, extension_ind = :5, ppdm_guid = :6, primary_key_name = :7, remark = :8, row_count = :9, row_count_date = :10, source = :11, table_comment = :12, table_type = :13, row_changed_by = :14, row_changed_date = :15, row_created_by = :16, row_effective_date = :17, row_expiry_date = :18, row_quality = :19 where system_id = :21")
	if err != nil {
		return err
	}

	ppdm_table.Row_changed_date = formatDate(ppdm_table.Row_changed_date)
	ppdm_table.Effective_date = formatDateString(ppdm_table.Effective_date)
	ppdm_table.Expiry_date = formatDateString(ppdm_table.Expiry_date)
	ppdm_table.Row_count_date = formatDateString(ppdm_table.Row_count_date)
	ppdm_table.Row_effective_date = formatDateString(ppdm_table.Row_effective_date)
	ppdm_table.Row_expiry_date = formatDateString(ppdm_table.Row_expiry_date)






	rows, err := stmt.Exec(ppdm_table.Table_name, ppdm_table.Active_ind, ppdm_table.Effective_date, ppdm_table.Expiry_date, ppdm_table.Extension_ind, ppdm_table.Ppdm_guid, ppdm_table.Primary_key_name, ppdm_table.Remark, ppdm_table.Row_count, ppdm_table.Row_count_date, ppdm_table.Source, ppdm_table.Table_comment, ppdm_table.Table_type, ppdm_table.Row_changed_by, ppdm_table.Row_changed_date, ppdm_table.Row_created_by, ppdm_table.Row_effective_date, ppdm_table.Row_expiry_date, ppdm_table.Row_quality, ppdm_table.System_id)
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

func PatchPpdmTable(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update ppdm_table set "
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
		} else if key == "effective_date" || key == "expiry_date" || key == "row_count_date" || key == "row_effective_date" || key == "row_expiry_date"      {
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

func DeletePpdmTable(c *fiber.Ctx) error {
	id := c.Params("id")
	var ppdm_table dto.Ppdm_table
	ppdm_table.System_id = id

	stmt, err := db.Prepare("delete from ppdm_table where system_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(ppdm_table.System_id)
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


