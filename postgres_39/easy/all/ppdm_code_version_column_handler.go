package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetPpdmCodeVersionColumn(c *fiber.Ctx) error {
	rows, err := db.Query("select * from ppdm_code_version_column")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Ppdm_code_version_column

	for rows.Next() {
		var ppdm_code_version_column dto.Ppdm_code_version_column
		if err := rows.Scan(&ppdm_code_version_column.System_id, &ppdm_code_version_column.Table_name, &ppdm_code_version_column.Source, &ppdm_code_version_column.Version_obs_no, &ppdm_code_version_column.Column_name, &ppdm_code_version_column.Abbreviation, &ppdm_code_version_column.Active_ind, &ppdm_code_version_column.Definition, &ppdm_code_version_column.Effective_date, &ppdm_code_version_column.Expiry_date, &ppdm_code_version_column.Language, &ppdm_code_version_column.Long_name, &ppdm_code_version_column.Ppdm_guid, &ppdm_code_version_column.Primary_key, &ppdm_code_version_column.Remark, &ppdm_code_version_column.Short_name, &ppdm_code_version_column.Row_changed_by, &ppdm_code_version_column.Row_changed_date, &ppdm_code_version_column.Row_created_by, &ppdm_code_version_column.Row_created_date, &ppdm_code_version_column.Row_effective_date, &ppdm_code_version_column.Row_expiry_date, &ppdm_code_version_column.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append ppdm_code_version_column to result
		result = append(result, ppdm_code_version_column)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetPpdmCodeVersionColumn(c *fiber.Ctx) error {
	var ppdm_code_version_column dto.Ppdm_code_version_column

	setDefaults(&ppdm_code_version_column)

	if err := json.Unmarshal(c.Body(), &ppdm_code_version_column); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into ppdm_code_version_column values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23)")
	if err != nil {
		return err
	}
	ppdm_code_version_column.Row_created_date = formatDate(ppdm_code_version_column.Row_created_date)
	ppdm_code_version_column.Row_changed_date = formatDate(ppdm_code_version_column.Row_changed_date)
	ppdm_code_version_column.Effective_date = formatDateString(ppdm_code_version_column.Effective_date)
	ppdm_code_version_column.Expiry_date = formatDateString(ppdm_code_version_column.Expiry_date)
	ppdm_code_version_column.Row_effective_date = formatDateString(ppdm_code_version_column.Row_effective_date)
	ppdm_code_version_column.Row_expiry_date = formatDateString(ppdm_code_version_column.Row_expiry_date)






	rows, err := stmt.Exec(ppdm_code_version_column.System_id, ppdm_code_version_column.Table_name, ppdm_code_version_column.Source, ppdm_code_version_column.Version_obs_no, ppdm_code_version_column.Column_name, ppdm_code_version_column.Abbreviation, ppdm_code_version_column.Active_ind, ppdm_code_version_column.Definition, ppdm_code_version_column.Effective_date, ppdm_code_version_column.Expiry_date, ppdm_code_version_column.Language, ppdm_code_version_column.Long_name, ppdm_code_version_column.Ppdm_guid, ppdm_code_version_column.Primary_key, ppdm_code_version_column.Remark, ppdm_code_version_column.Short_name, ppdm_code_version_column.Row_changed_by, ppdm_code_version_column.Row_changed_date, ppdm_code_version_column.Row_created_by, ppdm_code_version_column.Row_created_date, ppdm_code_version_column.Row_effective_date, ppdm_code_version_column.Row_expiry_date, ppdm_code_version_column.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdatePpdmCodeVersionColumn(c *fiber.Ctx) error {
	var ppdm_code_version_column dto.Ppdm_code_version_column

	setDefaults(&ppdm_code_version_column)

	if err := json.Unmarshal(c.Body(), &ppdm_code_version_column); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	ppdm_code_version_column.System_id = id

    if ppdm_code_version_column.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update ppdm_code_version_column set table_name = :1, source = :2, version_obs_no = :3, column_name = :4, abbreviation = :5, active_ind = :6, definition = :7, effective_date = :8, expiry_date = :9, language = :10, long_name = :11, ppdm_guid = :12, primary_key = :13, remark = :14, short_name = :15, row_changed_by = :16, row_changed_date = :17, row_created_by = :18, row_effective_date = :19, row_expiry_date = :20, row_quality = :21 where system_id = :23")
	if err != nil {
		return err
	}

	ppdm_code_version_column.Row_changed_date = formatDate(ppdm_code_version_column.Row_changed_date)
	ppdm_code_version_column.Effective_date = formatDateString(ppdm_code_version_column.Effective_date)
	ppdm_code_version_column.Expiry_date = formatDateString(ppdm_code_version_column.Expiry_date)
	ppdm_code_version_column.Row_effective_date = formatDateString(ppdm_code_version_column.Row_effective_date)
	ppdm_code_version_column.Row_expiry_date = formatDateString(ppdm_code_version_column.Row_expiry_date)






	rows, err := stmt.Exec(ppdm_code_version_column.Table_name, ppdm_code_version_column.Source, ppdm_code_version_column.Version_obs_no, ppdm_code_version_column.Column_name, ppdm_code_version_column.Abbreviation, ppdm_code_version_column.Active_ind, ppdm_code_version_column.Definition, ppdm_code_version_column.Effective_date, ppdm_code_version_column.Expiry_date, ppdm_code_version_column.Language, ppdm_code_version_column.Long_name, ppdm_code_version_column.Ppdm_guid, ppdm_code_version_column.Primary_key, ppdm_code_version_column.Remark, ppdm_code_version_column.Short_name, ppdm_code_version_column.Row_changed_by, ppdm_code_version_column.Row_changed_date, ppdm_code_version_column.Row_created_by, ppdm_code_version_column.Row_effective_date, ppdm_code_version_column.Row_expiry_date, ppdm_code_version_column.Row_quality, ppdm_code_version_column.System_id)
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

func PatchPpdmCodeVersionColumn(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update ppdm_code_version_column set "
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

func DeletePpdmCodeVersionColumn(c *fiber.Ctx) error {
	id := c.Params("id")
	var ppdm_code_version_column dto.Ppdm_code_version_column
	ppdm_code_version_column.System_id = id

	stmt, err := db.Prepare("delete from ppdm_code_version_column where system_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(ppdm_code_version_column.System_id)
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


