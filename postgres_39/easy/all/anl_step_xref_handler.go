package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetAnlStepXref(c *fiber.Ctx) error {
	rows, err := db.Query("select * from anl_step_xref")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Anl_step_xref

	for rows.Next() {
		var anl_step_xref dto.Anl_step_xref
		if err := rows.Scan(&anl_step_xref.Analysis_id, &anl_step_xref.Anl_source, &anl_step_xref.Step_seq_no1, &anl_step_xref.Step_seq_no2, &anl_step_xref.Active_ind, &anl_step_xref.Effective_date, &anl_step_xref.Expiry_date, &anl_step_xref.Ppdm_guid, &anl_step_xref.Remark, &anl_step_xref.Source, &anl_step_xref.Step_xref_reason, &anl_step_xref.Row_changed_by, &anl_step_xref.Row_changed_date, &anl_step_xref.Row_created_by, &anl_step_xref.Row_created_date, &anl_step_xref.Row_effective_date, &anl_step_xref.Row_expiry_date, &anl_step_xref.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append anl_step_xref to result
		result = append(result, anl_step_xref)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetAnlStepXref(c *fiber.Ctx) error {
	var anl_step_xref dto.Anl_step_xref

	setDefaults(&anl_step_xref)

	if err := json.Unmarshal(c.Body(), &anl_step_xref); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into anl_step_xref values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18)")
	if err != nil {
		return err
	}
	anl_step_xref.Row_created_date = formatDate(anl_step_xref.Row_created_date)
	anl_step_xref.Row_changed_date = formatDate(anl_step_xref.Row_changed_date)
	anl_step_xref.Effective_date = formatDateString(anl_step_xref.Effective_date)
	anl_step_xref.Expiry_date = formatDateString(anl_step_xref.Expiry_date)
	anl_step_xref.Row_effective_date = formatDateString(anl_step_xref.Row_effective_date)
	anl_step_xref.Row_expiry_date = formatDateString(anl_step_xref.Row_expiry_date)






	rows, err := stmt.Exec(anl_step_xref.Analysis_id, anl_step_xref.Anl_source, anl_step_xref.Step_seq_no1, anl_step_xref.Step_seq_no2, anl_step_xref.Active_ind, anl_step_xref.Effective_date, anl_step_xref.Expiry_date, anl_step_xref.Ppdm_guid, anl_step_xref.Remark, anl_step_xref.Source, anl_step_xref.Step_xref_reason, anl_step_xref.Row_changed_by, anl_step_xref.Row_changed_date, anl_step_xref.Row_created_by, anl_step_xref.Row_created_date, anl_step_xref.Row_effective_date, anl_step_xref.Row_expiry_date, anl_step_xref.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateAnlStepXref(c *fiber.Ctx) error {
	var anl_step_xref dto.Anl_step_xref

	setDefaults(&anl_step_xref)

	if err := json.Unmarshal(c.Body(), &anl_step_xref); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	anl_step_xref.Analysis_id = id

    if anl_step_xref.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update anl_step_xref set anl_source = :1, step_seq_no1 = :2, step_seq_no2 = :3, active_ind = :4, effective_date = :5, expiry_date = :6, ppdm_guid = :7, remark = :8, source = :9, step_xref_reason = :10, row_changed_by = :11, row_changed_date = :12, row_created_by = :13, row_effective_date = :14, row_expiry_date = :15, row_quality = :16 where analysis_id = :18")
	if err != nil {
		return err
	}

	anl_step_xref.Row_changed_date = formatDate(anl_step_xref.Row_changed_date)
	anl_step_xref.Effective_date = formatDateString(anl_step_xref.Effective_date)
	anl_step_xref.Expiry_date = formatDateString(anl_step_xref.Expiry_date)
	anl_step_xref.Row_effective_date = formatDateString(anl_step_xref.Row_effective_date)
	anl_step_xref.Row_expiry_date = formatDateString(anl_step_xref.Row_expiry_date)






	rows, err := stmt.Exec(anl_step_xref.Anl_source, anl_step_xref.Step_seq_no1, anl_step_xref.Step_seq_no2, anl_step_xref.Active_ind, anl_step_xref.Effective_date, anl_step_xref.Expiry_date, anl_step_xref.Ppdm_guid, anl_step_xref.Remark, anl_step_xref.Source, anl_step_xref.Step_xref_reason, anl_step_xref.Row_changed_by, anl_step_xref.Row_changed_date, anl_step_xref.Row_created_by, anl_step_xref.Row_effective_date, anl_step_xref.Row_expiry_date, anl_step_xref.Row_quality, anl_step_xref.Analysis_id)
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

func PatchAnlStepXref(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update anl_step_xref set "
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

func DeleteAnlStepXref(c *fiber.Ctx) error {
	id := c.Params("id")
	var anl_step_xref dto.Anl_step_xref
	anl_step_xref.Analysis_id = id

	stmt, err := db.Prepare("delete from anl_step_xref where analysis_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(anl_step_xref.Analysis_id)
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


