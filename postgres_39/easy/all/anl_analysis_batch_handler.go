package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetAnlAnalysisBatch(c *fiber.Ctx) error {
	rows, err := db.Query("select * from anl_analysis_batch")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Anl_analysis_batch

	for rows.Next() {
		var anl_analysis_batch dto.Anl_analysis_batch
		if err := rows.Scan(&anl_analysis_batch.Batch_id, &anl_analysis_batch.Active_ind, &anl_analysis_batch.Batch_desc, &anl_analysis_batch.Batch_name, &anl_analysis_batch.Batch_ref_num, &anl_analysis_batch.Create_date, &anl_analysis_batch.Effective_date, &anl_analysis_batch.Expiry_date, &anl_analysis_batch.Ppdm_guid, &anl_analysis_batch.Remark, &anl_analysis_batch.Sample_count, &anl_analysis_batch.Source, &anl_analysis_batch.Standard_sample_count, &anl_analysis_batch.Row_changed_by, &anl_analysis_batch.Row_changed_date, &anl_analysis_batch.Row_created_by, &anl_analysis_batch.Row_created_date, &anl_analysis_batch.Row_effective_date, &anl_analysis_batch.Row_expiry_date, &anl_analysis_batch.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append anl_analysis_batch to result
		result = append(result, anl_analysis_batch)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetAnlAnalysisBatch(c *fiber.Ctx) error {
	var anl_analysis_batch dto.Anl_analysis_batch

	setDefaults(&anl_analysis_batch)

	if err := json.Unmarshal(c.Body(), &anl_analysis_batch); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into anl_analysis_batch values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20)")
	if err != nil {
		return err
	}
	anl_analysis_batch.Row_created_date = formatDate(anl_analysis_batch.Row_created_date)
	anl_analysis_batch.Row_changed_date = formatDate(anl_analysis_batch.Row_changed_date)
	anl_analysis_batch.Create_date = formatDateString(anl_analysis_batch.Create_date)
	anl_analysis_batch.Effective_date = formatDateString(anl_analysis_batch.Effective_date)
	anl_analysis_batch.Expiry_date = formatDateString(anl_analysis_batch.Expiry_date)
	anl_analysis_batch.Row_effective_date = formatDateString(anl_analysis_batch.Row_effective_date)
	anl_analysis_batch.Row_expiry_date = formatDateString(anl_analysis_batch.Row_expiry_date)






	rows, err := stmt.Exec(anl_analysis_batch.Batch_id, anl_analysis_batch.Active_ind, anl_analysis_batch.Batch_desc, anl_analysis_batch.Batch_name, anl_analysis_batch.Batch_ref_num, anl_analysis_batch.Create_date, anl_analysis_batch.Effective_date, anl_analysis_batch.Expiry_date, anl_analysis_batch.Ppdm_guid, anl_analysis_batch.Remark, anl_analysis_batch.Sample_count, anl_analysis_batch.Source, anl_analysis_batch.Standard_sample_count, anl_analysis_batch.Row_changed_by, anl_analysis_batch.Row_changed_date, anl_analysis_batch.Row_created_by, anl_analysis_batch.Row_created_date, anl_analysis_batch.Row_effective_date, anl_analysis_batch.Row_expiry_date, anl_analysis_batch.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateAnlAnalysisBatch(c *fiber.Ctx) error {
	var anl_analysis_batch dto.Anl_analysis_batch

	setDefaults(&anl_analysis_batch)

	if err := json.Unmarshal(c.Body(), &anl_analysis_batch); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	anl_analysis_batch.Batch_id = id

    if anl_analysis_batch.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update anl_analysis_batch set active_ind = :1, batch_desc = :2, batch_name = :3, batch_ref_num = :4, create_date = :5, effective_date = :6, expiry_date = :7, ppdm_guid = :8, remark = :9, sample_count = :10, source = :11, standard_sample_count = :12, row_changed_by = :13, row_changed_date = :14, row_created_by = :15, row_effective_date = :16, row_expiry_date = :17, row_quality = :18 where batch_id = :20")
	if err != nil {
		return err
	}

	anl_analysis_batch.Row_changed_date = formatDate(anl_analysis_batch.Row_changed_date)
	anl_analysis_batch.Create_date = formatDateString(anl_analysis_batch.Create_date)
	anl_analysis_batch.Effective_date = formatDateString(anl_analysis_batch.Effective_date)
	anl_analysis_batch.Expiry_date = formatDateString(anl_analysis_batch.Expiry_date)
	anl_analysis_batch.Row_effective_date = formatDateString(anl_analysis_batch.Row_effective_date)
	anl_analysis_batch.Row_expiry_date = formatDateString(anl_analysis_batch.Row_expiry_date)






	rows, err := stmt.Exec(anl_analysis_batch.Active_ind, anl_analysis_batch.Batch_desc, anl_analysis_batch.Batch_name, anl_analysis_batch.Batch_ref_num, anl_analysis_batch.Create_date, anl_analysis_batch.Effective_date, anl_analysis_batch.Expiry_date, anl_analysis_batch.Ppdm_guid, anl_analysis_batch.Remark, anl_analysis_batch.Sample_count, anl_analysis_batch.Source, anl_analysis_batch.Standard_sample_count, anl_analysis_batch.Row_changed_by, anl_analysis_batch.Row_changed_date, anl_analysis_batch.Row_created_by, anl_analysis_batch.Row_effective_date, anl_analysis_batch.Row_expiry_date, anl_analysis_batch.Row_quality, anl_analysis_batch.Batch_id)
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

func PatchAnlAnalysisBatch(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update anl_analysis_batch set "
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
		} else if key == "create_date" || key == "effective_date" || key == "expiry_date" || key == "row_effective_date" || key == "row_expiry_date"      {
			if date, ok := value.(string); ok && len(date) > 0 {
				formattedDate := formatDateString(&date)
				value = formattedDate
			}
		}
		values = append(values, value)
	}
	query += " where batch_id = :id"

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

func DeleteAnlAnalysisBatch(c *fiber.Ctx) error {
	id := c.Params("id")
	var anl_analysis_batch dto.Anl_analysis_batch
	anl_analysis_batch.Batch_id = id

	stmt, err := db.Prepare("delete from anl_analysis_batch where batch_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(anl_analysis_batch.Batch_id)
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


