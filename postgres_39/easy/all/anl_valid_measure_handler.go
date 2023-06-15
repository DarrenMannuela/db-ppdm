package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetAnlValidMeasure(c *fiber.Ctx) error {
	rows, err := db.Query("select * from anl_valid_measure")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Anl_valid_measure

	for rows.Next() {
		var anl_valid_measure dto.Anl_valid_measure
		if err := rows.Scan(&anl_valid_measure.Method_set_id, &anl_valid_measure.Method_id, &anl_valid_measure.Measurement_type, &anl_valid_measure.Valid_value_obs_no, &anl_valid_measure.Accuracy_type, &anl_valid_measure.Active_ind, &anl_valid_measure.Average_ratio_value, &anl_valid_measure.Average_value, &anl_valid_measure.Average_value_ouom, &anl_valid_measure.Average_value_uom, &anl_valid_measure.Calculate_method_id, &anl_valid_measure.Effective_date, &anl_valid_measure.Expiry_date, &anl_valid_measure.Maximum_ratio_value, &anl_valid_measure.Max_value, &anl_valid_measure.Max_value_ouom, &anl_valid_measure.Max_value_uom, &anl_valid_measure.Minimum_ratio_value, &anl_valid_measure.Min_value, &anl_valid_measure.Min_value_ouom, &anl_valid_measure.Min_value_uom, &anl_valid_measure.Missing_representation, &anl_valid_measure.Null_representation, &anl_valid_measure.Ppdm_guid, &anl_valid_measure.Reference_value, &anl_valid_measure.Reference_value_ouom, &anl_valid_measure.Reference_value_uom, &anl_valid_measure.Remark, &anl_valid_measure.Source, &anl_valid_measure.Substance_id, &anl_valid_measure.Valid_desc, &anl_valid_measure.Valid_value_text, &anl_valid_measure.Row_changed_by, &anl_valid_measure.Row_changed_date, &anl_valid_measure.Row_created_by, &anl_valid_measure.Row_created_date, &anl_valid_measure.Row_effective_date, &anl_valid_measure.Row_expiry_date, &anl_valid_measure.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append anl_valid_measure to result
		result = append(result, anl_valid_measure)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetAnlValidMeasure(c *fiber.Ctx) error {
	var anl_valid_measure dto.Anl_valid_measure

	setDefaults(&anl_valid_measure)

	if err := json.Unmarshal(c.Body(), &anl_valid_measure); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into anl_valid_measure values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32, :33, :34, :35, :36, :37, :38, :39)")
	if err != nil {
		return err
	}
	anl_valid_measure.Row_created_date = formatDate(anl_valid_measure.Row_created_date)
	anl_valid_measure.Row_changed_date = formatDate(anl_valid_measure.Row_changed_date)
	anl_valid_measure.Effective_date = formatDateString(anl_valid_measure.Effective_date)
	anl_valid_measure.Expiry_date = formatDateString(anl_valid_measure.Expiry_date)
	anl_valid_measure.Row_effective_date = formatDateString(anl_valid_measure.Row_effective_date)
	anl_valid_measure.Row_expiry_date = formatDateString(anl_valid_measure.Row_expiry_date)






	rows, err := stmt.Exec(anl_valid_measure.Method_set_id, anl_valid_measure.Method_id, anl_valid_measure.Measurement_type, anl_valid_measure.Valid_value_obs_no, anl_valid_measure.Accuracy_type, anl_valid_measure.Active_ind, anl_valid_measure.Average_ratio_value, anl_valid_measure.Average_value, anl_valid_measure.Average_value_ouom, anl_valid_measure.Average_value_uom, anl_valid_measure.Calculate_method_id, anl_valid_measure.Effective_date, anl_valid_measure.Expiry_date, anl_valid_measure.Maximum_ratio_value, anl_valid_measure.Max_value, anl_valid_measure.Max_value_ouom, anl_valid_measure.Max_value_uom, anl_valid_measure.Minimum_ratio_value, anl_valid_measure.Min_value, anl_valid_measure.Min_value_ouom, anl_valid_measure.Min_value_uom, anl_valid_measure.Missing_representation, anl_valid_measure.Null_representation, anl_valid_measure.Ppdm_guid, anl_valid_measure.Reference_value, anl_valid_measure.Reference_value_ouom, anl_valid_measure.Reference_value_uom, anl_valid_measure.Remark, anl_valid_measure.Source, anl_valid_measure.Substance_id, anl_valid_measure.Valid_desc, anl_valid_measure.Valid_value_text, anl_valid_measure.Row_changed_by, anl_valid_measure.Row_changed_date, anl_valid_measure.Row_created_by, anl_valid_measure.Row_created_date, anl_valid_measure.Row_effective_date, anl_valid_measure.Row_expiry_date, anl_valid_measure.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateAnlValidMeasure(c *fiber.Ctx) error {
	var anl_valid_measure dto.Anl_valid_measure

	setDefaults(&anl_valid_measure)

	if err := json.Unmarshal(c.Body(), &anl_valid_measure); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	anl_valid_measure.Method_set_id = id

    if anl_valid_measure.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update anl_valid_measure set method_id = :1, measurement_type = :2, valid_value_obs_no = :3, accuracy_type = :4, active_ind = :5, average_ratio_value = :6, average_value = :7, average_value_ouom = :8, average_value_uom = :9, calculate_method_id = :10, effective_date = :11, expiry_date = :12, maximum_ratio_value = :13, max_value = :14, max_value_ouom = :15, max_value_uom = :16, minimum_ratio_value = :17, min_value = :18, min_value_ouom = :19, min_value_uom = :20, missing_representation = :21, null_representation = :22, ppdm_guid = :23, reference_value = :24, reference_value_ouom = :25, reference_value_uom = :26, remark = :27, source = :28, substance_id = :29, valid_desc = :30, valid_value_text = :31, row_changed_by = :32, row_changed_date = :33, row_created_by = :34, row_effective_date = :35, row_expiry_date = :36, row_quality = :37 where method_set_id = :39")
	if err != nil {
		return err
	}

	anl_valid_measure.Row_changed_date = formatDate(anl_valid_measure.Row_changed_date)
	anl_valid_measure.Effective_date = formatDateString(anl_valid_measure.Effective_date)
	anl_valid_measure.Expiry_date = formatDateString(anl_valid_measure.Expiry_date)
	anl_valid_measure.Row_effective_date = formatDateString(anl_valid_measure.Row_effective_date)
	anl_valid_measure.Row_expiry_date = formatDateString(anl_valid_measure.Row_expiry_date)






	rows, err := stmt.Exec(anl_valid_measure.Method_id, anl_valid_measure.Measurement_type, anl_valid_measure.Valid_value_obs_no, anl_valid_measure.Accuracy_type, anl_valid_measure.Active_ind, anl_valid_measure.Average_ratio_value, anl_valid_measure.Average_value, anl_valid_measure.Average_value_ouom, anl_valid_measure.Average_value_uom, anl_valid_measure.Calculate_method_id, anl_valid_measure.Effective_date, anl_valid_measure.Expiry_date, anl_valid_measure.Maximum_ratio_value, anl_valid_measure.Max_value, anl_valid_measure.Max_value_ouom, anl_valid_measure.Max_value_uom, anl_valid_measure.Minimum_ratio_value, anl_valid_measure.Min_value, anl_valid_measure.Min_value_ouom, anl_valid_measure.Min_value_uom, anl_valid_measure.Missing_representation, anl_valid_measure.Null_representation, anl_valid_measure.Ppdm_guid, anl_valid_measure.Reference_value, anl_valid_measure.Reference_value_ouom, anl_valid_measure.Reference_value_uom, anl_valid_measure.Remark, anl_valid_measure.Source, anl_valid_measure.Substance_id, anl_valid_measure.Valid_desc, anl_valid_measure.Valid_value_text, anl_valid_measure.Row_changed_by, anl_valid_measure.Row_changed_date, anl_valid_measure.Row_created_by, anl_valid_measure.Row_effective_date, anl_valid_measure.Row_expiry_date, anl_valid_measure.Row_quality, anl_valid_measure.Method_set_id)
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

func PatchAnlValidMeasure(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update anl_valid_measure set "
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
	query += " where method_set_id = :id"

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

func DeleteAnlValidMeasure(c *fiber.Ctx) error {
	id := c.Params("id")
	var anl_valid_measure dto.Anl_valid_measure
	anl_valid_measure.Method_set_id = id

	stmt, err := db.Prepare("delete from anl_valid_measure where method_set_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(anl_valid_measure.Method_set_id)
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


