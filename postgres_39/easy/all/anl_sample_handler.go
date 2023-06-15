package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetAnlSample(c *fiber.Ctx) error {
	rows, err := db.Query("select * from anl_sample")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Anl_sample

	for rows.Next() {
		var anl_sample dto.Anl_sample
		if err := rows.Scan(&anl_sample.Analysis_id, &anl_sample.Anl_source, &anl_sample.Sample_id, &anl_sample.Active_ind, &anl_sample.Batch_id, &anl_sample.Created_by_step_ind, &anl_sample.Effective_date, &anl_sample.End_date, &anl_sample.Expiry_date, &anl_sample.Ppdm_guid, &anl_sample.Remark, &anl_sample.Sample_description, &anl_sample.Source, &anl_sample.Standard_sample_ind, &anl_sample.Step_seq_no, &anl_sample.Used_by_step_ind, &anl_sample.Row_changed_by, &anl_sample.Row_changed_date, &anl_sample.Row_created_by, &anl_sample.Row_created_date, &anl_sample.Row_effective_date, &anl_sample.Row_expiry_date, &anl_sample.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append anl_sample to result
		result = append(result, anl_sample)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetAnlSample(c *fiber.Ctx) error {
	var anl_sample dto.Anl_sample

	setDefaults(&anl_sample)

	if err := json.Unmarshal(c.Body(), &anl_sample); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into anl_sample values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23)")
	if err != nil {
		return err
	}
	anl_sample.Row_created_date = formatDate(anl_sample.Row_created_date)
	anl_sample.Row_changed_date = formatDate(anl_sample.Row_changed_date)
	anl_sample.Effective_date = formatDateString(anl_sample.Effective_date)
	anl_sample.End_date = formatDateString(anl_sample.End_date)
	anl_sample.Expiry_date = formatDateString(anl_sample.Expiry_date)
	anl_sample.Row_effective_date = formatDateString(anl_sample.Row_effective_date)
	anl_sample.Row_expiry_date = formatDateString(anl_sample.Row_expiry_date)






	rows, err := stmt.Exec(anl_sample.Analysis_id, anl_sample.Anl_source, anl_sample.Sample_id, anl_sample.Active_ind, anl_sample.Batch_id, anl_sample.Created_by_step_ind, anl_sample.Effective_date, anl_sample.End_date, anl_sample.Expiry_date, anl_sample.Ppdm_guid, anl_sample.Remark, anl_sample.Sample_description, anl_sample.Source, anl_sample.Standard_sample_ind, anl_sample.Step_seq_no, anl_sample.Used_by_step_ind, anl_sample.Row_changed_by, anl_sample.Row_changed_date, anl_sample.Row_created_by, anl_sample.Row_created_date, anl_sample.Row_effective_date, anl_sample.Row_expiry_date, anl_sample.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateAnlSample(c *fiber.Ctx) error {
	var anl_sample dto.Anl_sample

	setDefaults(&anl_sample)

	if err := json.Unmarshal(c.Body(), &anl_sample); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	anl_sample.Analysis_id = id

    if anl_sample.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update anl_sample set anl_source = :1, sample_id = :2, active_ind = :3, batch_id = :4, created_by_step_ind = :5, effective_date = :6, end_date = :7, expiry_date = :8, ppdm_guid = :9, remark = :10, sample_description = :11, source = :12, standard_sample_ind = :13, step_seq_no = :14, used_by_step_ind = :15, row_changed_by = :16, row_changed_date = :17, row_created_by = :18, row_effective_date = :19, row_expiry_date = :20, row_quality = :21 where analysis_id = :23")
	if err != nil {
		return err
	}

	anl_sample.Row_changed_date = formatDate(anl_sample.Row_changed_date)
	anl_sample.Effective_date = formatDateString(anl_sample.Effective_date)
	anl_sample.End_date = formatDateString(anl_sample.End_date)
	anl_sample.Expiry_date = formatDateString(anl_sample.Expiry_date)
	anl_sample.Row_effective_date = formatDateString(anl_sample.Row_effective_date)
	anl_sample.Row_expiry_date = formatDateString(anl_sample.Row_expiry_date)






	rows, err := stmt.Exec(anl_sample.Anl_source, anl_sample.Sample_id, anl_sample.Active_ind, anl_sample.Batch_id, anl_sample.Created_by_step_ind, anl_sample.Effective_date, anl_sample.End_date, anl_sample.Expiry_date, anl_sample.Ppdm_guid, anl_sample.Remark, anl_sample.Sample_description, anl_sample.Source, anl_sample.Standard_sample_ind, anl_sample.Step_seq_no, anl_sample.Used_by_step_ind, anl_sample.Row_changed_by, anl_sample.Row_changed_date, anl_sample.Row_created_by, anl_sample.Row_effective_date, anl_sample.Row_expiry_date, anl_sample.Row_quality, anl_sample.Analysis_id)
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

func PatchAnlSample(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update anl_sample set "
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
		} else if key == "effective_date" || key == "end_date" || key == "expiry_date" || key == "row_effective_date" || key == "row_expiry_date"      {
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

func DeleteAnlSample(c *fiber.Ctx) error {
	id := c.Params("id")
	var anl_sample dto.Anl_sample
	anl_sample.Analysis_id = id

	stmt, err := db.Prepare("delete from anl_sample where analysis_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(anl_sample.Analysis_id)
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


