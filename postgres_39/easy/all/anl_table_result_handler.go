package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetAnlTableResult(c *fiber.Ctx) error {
	rows, err := db.Query("select * from anl_table_result")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Anl_table_result

	for rows.Next() {
		var anl_table_result dto.Anl_table_result
		if err := rows.Scan(&anl_table_result.Analysis_id, &anl_table_result.Anl_source, &anl_table_result.Ppdm_system_id, &anl_table_result.Ppdm_table_name, &anl_table_result.Active_ind, &anl_table_result.Effective_date, &anl_table_result.Expiry_date, &anl_table_result.Ppdm_guid, &anl_table_result.Remark, &anl_table_result.Source, &anl_table_result.Step_seq_no, &anl_table_result.Row_changed_by, &anl_table_result.Row_changed_date, &anl_table_result.Row_created_by, &anl_table_result.Row_created_date, &anl_table_result.Row_effective_date, &anl_table_result.Row_expiry_date, &anl_table_result.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append anl_table_result to result
		result = append(result, anl_table_result)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetAnlTableResult(c *fiber.Ctx) error {
	var anl_table_result dto.Anl_table_result

	setDefaults(&anl_table_result)

	if err := json.Unmarshal(c.Body(), &anl_table_result); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into anl_table_result values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18)")
	if err != nil {
		return err
	}
	anl_table_result.Row_created_date = formatDate(anl_table_result.Row_created_date)
	anl_table_result.Row_changed_date = formatDate(anl_table_result.Row_changed_date)
	anl_table_result.Effective_date = formatDateString(anl_table_result.Effective_date)
	anl_table_result.Expiry_date = formatDateString(anl_table_result.Expiry_date)
	anl_table_result.Row_effective_date = formatDateString(anl_table_result.Row_effective_date)
	anl_table_result.Row_expiry_date = formatDateString(anl_table_result.Row_expiry_date)






	rows, err := stmt.Exec(anl_table_result.Analysis_id, anl_table_result.Anl_source, anl_table_result.Ppdm_system_id, anl_table_result.Ppdm_table_name, anl_table_result.Active_ind, anl_table_result.Effective_date, anl_table_result.Expiry_date, anl_table_result.Ppdm_guid, anl_table_result.Remark, anl_table_result.Source, anl_table_result.Step_seq_no, anl_table_result.Row_changed_by, anl_table_result.Row_changed_date, anl_table_result.Row_created_by, anl_table_result.Row_created_date, anl_table_result.Row_effective_date, anl_table_result.Row_expiry_date, anl_table_result.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateAnlTableResult(c *fiber.Ctx) error {
	var anl_table_result dto.Anl_table_result

	setDefaults(&anl_table_result)

	if err := json.Unmarshal(c.Body(), &anl_table_result); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	anl_table_result.Analysis_id = id

    if anl_table_result.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update anl_table_result set anl_source = :1, ppdm_system_id = :2, ppdm_table_name = :3, active_ind = :4, effective_date = :5, expiry_date = :6, ppdm_guid = :7, remark = :8, source = :9, step_seq_no = :10, row_changed_by = :11, row_changed_date = :12, row_created_by = :13, row_effective_date = :14, row_expiry_date = :15, row_quality = :16 where analysis_id = :18")
	if err != nil {
		return err
	}

	anl_table_result.Row_changed_date = formatDate(anl_table_result.Row_changed_date)
	anl_table_result.Effective_date = formatDateString(anl_table_result.Effective_date)
	anl_table_result.Expiry_date = formatDateString(anl_table_result.Expiry_date)
	anl_table_result.Row_effective_date = formatDateString(anl_table_result.Row_effective_date)
	anl_table_result.Row_expiry_date = formatDateString(anl_table_result.Row_expiry_date)






	rows, err := stmt.Exec(anl_table_result.Anl_source, anl_table_result.Ppdm_system_id, anl_table_result.Ppdm_table_name, anl_table_result.Active_ind, anl_table_result.Effective_date, anl_table_result.Expiry_date, anl_table_result.Ppdm_guid, anl_table_result.Remark, anl_table_result.Source, anl_table_result.Step_seq_no, anl_table_result.Row_changed_by, anl_table_result.Row_changed_date, anl_table_result.Row_created_by, anl_table_result.Row_effective_date, anl_table_result.Row_expiry_date, anl_table_result.Row_quality, anl_table_result.Analysis_id)
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

func PatchAnlTableResult(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update anl_table_result set "
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
	query += " where analysis_id = :id"

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

func DeleteAnlTableResult(c *fiber.Ctx) error {
	id := c.Params("id")
	var anl_table_result dto.Anl_table_result
	anl_table_result.Analysis_id = id

	stmt, err := db.Prepare("delete from anl_table_result where analysis_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(anl_table_result.Analysis_id)
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


