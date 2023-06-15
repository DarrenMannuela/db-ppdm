package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetPpdmConsColumn(c *fiber.Ctx) error {
	rows, err := db.Query("select * from ppdm_cons_column")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Ppdm_cons_column

	for rows.Next() {
		var ppdm_cons_column dto.Ppdm_cons_column
		if err := rows.Scan(&ppdm_cons_column.System_id, &ppdm_cons_column.Table_name, &ppdm_cons_column.Constraint_name, &ppdm_cons_column.Column_name, &ppdm_cons_column.Active_ind, &ppdm_cons_column.Condition, &ppdm_cons_column.Constraint_seq_no, &ppdm_cons_column.Effective_date, &ppdm_cons_column.Expiry_date, &ppdm_cons_column.Extension_ind, &ppdm_cons_column.Ppdm_guid, &ppdm_cons_column.Referenced_column_name, &ppdm_cons_column.Referenced_constraint_name, &ppdm_cons_column.Referenced_system_id, &ppdm_cons_column.Referenced_table_name, &ppdm_cons_column.Remark, &ppdm_cons_column.Source, &ppdm_cons_column.Row_changed_by, &ppdm_cons_column.Row_changed_date, &ppdm_cons_column.Row_created_by, &ppdm_cons_column.Row_created_date, &ppdm_cons_column.Row_effective_date, &ppdm_cons_column.Row_expiry_date, &ppdm_cons_column.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append ppdm_cons_column to result
		result = append(result, ppdm_cons_column)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetPpdmConsColumn(c *fiber.Ctx) error {
	var ppdm_cons_column dto.Ppdm_cons_column

	setDefaults(&ppdm_cons_column)

	if err := json.Unmarshal(c.Body(), &ppdm_cons_column); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into ppdm_cons_column values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24)")
	if err != nil {
		return err
	}
	ppdm_cons_column.Row_created_date = formatDate(ppdm_cons_column.Row_created_date)
	ppdm_cons_column.Row_changed_date = formatDate(ppdm_cons_column.Row_changed_date)
	ppdm_cons_column.Effective_date = formatDateString(ppdm_cons_column.Effective_date)
	ppdm_cons_column.Expiry_date = formatDateString(ppdm_cons_column.Expiry_date)
	ppdm_cons_column.Row_effective_date = formatDateString(ppdm_cons_column.Row_effective_date)
	ppdm_cons_column.Row_expiry_date = formatDateString(ppdm_cons_column.Row_expiry_date)






	rows, err := stmt.Exec(ppdm_cons_column.System_id, ppdm_cons_column.Table_name, ppdm_cons_column.Constraint_name, ppdm_cons_column.Column_name, ppdm_cons_column.Active_ind, ppdm_cons_column.Condition, ppdm_cons_column.Constraint_seq_no, ppdm_cons_column.Effective_date, ppdm_cons_column.Expiry_date, ppdm_cons_column.Extension_ind, ppdm_cons_column.Ppdm_guid, ppdm_cons_column.Referenced_column_name, ppdm_cons_column.Referenced_constraint_name, ppdm_cons_column.Referenced_system_id, ppdm_cons_column.Referenced_table_name, ppdm_cons_column.Remark, ppdm_cons_column.Source, ppdm_cons_column.Row_changed_by, ppdm_cons_column.Row_changed_date, ppdm_cons_column.Row_created_by, ppdm_cons_column.Row_created_date, ppdm_cons_column.Row_effective_date, ppdm_cons_column.Row_expiry_date, ppdm_cons_column.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdatePpdmConsColumn(c *fiber.Ctx) error {
	var ppdm_cons_column dto.Ppdm_cons_column

	setDefaults(&ppdm_cons_column)

	if err := json.Unmarshal(c.Body(), &ppdm_cons_column); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	ppdm_cons_column.System_id = id

    if ppdm_cons_column.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update ppdm_cons_column set table_name = :1, constraint_name = :2, column_name = :3, active_ind = :4, condition = :5, constraint_seq_no = :6, effective_date = :7, expiry_date = :8, extension_ind = :9, ppdm_guid = :10, referenced_column_name = :11, referenced_constraint_name = :12, referenced_system_id = :13, referenced_table_name = :14, remark = :15, source = :16, row_changed_by = :17, row_changed_date = :18, row_created_by = :19, row_effective_date = :20, row_expiry_date = :21, row_quality = :22 where system_id = :24")
	if err != nil {
		return err
	}

	ppdm_cons_column.Row_changed_date = formatDate(ppdm_cons_column.Row_changed_date)
	ppdm_cons_column.Effective_date = formatDateString(ppdm_cons_column.Effective_date)
	ppdm_cons_column.Expiry_date = formatDateString(ppdm_cons_column.Expiry_date)
	ppdm_cons_column.Row_effective_date = formatDateString(ppdm_cons_column.Row_effective_date)
	ppdm_cons_column.Row_expiry_date = formatDateString(ppdm_cons_column.Row_expiry_date)






	rows, err := stmt.Exec(ppdm_cons_column.Table_name, ppdm_cons_column.Constraint_name, ppdm_cons_column.Column_name, ppdm_cons_column.Active_ind, ppdm_cons_column.Condition, ppdm_cons_column.Constraint_seq_no, ppdm_cons_column.Effective_date, ppdm_cons_column.Expiry_date, ppdm_cons_column.Extension_ind, ppdm_cons_column.Ppdm_guid, ppdm_cons_column.Referenced_column_name, ppdm_cons_column.Referenced_constraint_name, ppdm_cons_column.Referenced_system_id, ppdm_cons_column.Referenced_table_name, ppdm_cons_column.Remark, ppdm_cons_column.Source, ppdm_cons_column.Row_changed_by, ppdm_cons_column.Row_changed_date, ppdm_cons_column.Row_created_by, ppdm_cons_column.Row_effective_date, ppdm_cons_column.Row_expiry_date, ppdm_cons_column.Row_quality, ppdm_cons_column.System_id)
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

func PatchPpdmConsColumn(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update ppdm_cons_column set "
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

func DeletePpdmConsColumn(c *fiber.Ctx) error {
	id := c.Params("id")
	var ppdm_cons_column dto.Ppdm_cons_column
	ppdm_cons_column.System_id = id

	stmt, err := db.Prepare("delete from ppdm_cons_column where system_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(ppdm_cons_column.System_id)
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


