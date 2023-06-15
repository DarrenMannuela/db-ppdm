package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetAnlAccuracy(c *fiber.Ctx) error {
	rows, err := db.Query("select * from anl_accuracy")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Anl_accuracy

	for rows.Next() {
		var anl_accuracy dto.Anl_accuracy
		if err := rows.Scan(&anl_accuracy.Analysis_id, &anl_accuracy.Anl_source, &anl_accuracy.Accuracy_obs_no, &anl_accuracy.Accuracy_desc, &anl_accuracy.Accuracy_type, &anl_accuracy.Active_ind, &anl_accuracy.Calculated_ind, &anl_accuracy.Calculate_method_id, &anl_accuracy.Confidence_type, &anl_accuracy.Effective_date, &anl_accuracy.Expiry_date, &anl_accuracy.Measured_value_ind, &anl_accuracy.Ppdm_guid, &anl_accuracy.Raw_value_ind, &anl_accuracy.Referenced_column_name, &anl_accuracy.Referenced_ppdm_guid, &anl_accuracy.Referenced_system_id, &anl_accuracy.Referenced_table_name, &anl_accuracy.Remark, &anl_accuracy.Repeatability_type, &anl_accuracy.Reported_ind, &anl_accuracy.Source, &anl_accuracy.Step_seq_no, &anl_accuracy.Row_changed_by, &anl_accuracy.Row_changed_date, &anl_accuracy.Row_created_by, &anl_accuracy.Row_created_date, &anl_accuracy.Row_effective_date, &anl_accuracy.Row_expiry_date, &anl_accuracy.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append anl_accuracy to result
		result = append(result, anl_accuracy)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetAnlAccuracy(c *fiber.Ctx) error {
	var anl_accuracy dto.Anl_accuracy

	setDefaults(&anl_accuracy)

	if err := json.Unmarshal(c.Body(), &anl_accuracy); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into anl_accuracy values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30)")
	if err != nil {
		return err
	}
	anl_accuracy.Row_created_date = formatDate(anl_accuracy.Row_created_date)
	anl_accuracy.Row_changed_date = formatDate(anl_accuracy.Row_changed_date)
	anl_accuracy.Effective_date = formatDateString(anl_accuracy.Effective_date)
	anl_accuracy.Expiry_date = formatDateString(anl_accuracy.Expiry_date)
	anl_accuracy.Row_effective_date = formatDateString(anl_accuracy.Row_effective_date)
	anl_accuracy.Row_expiry_date = formatDateString(anl_accuracy.Row_expiry_date)






	rows, err := stmt.Exec(anl_accuracy.Analysis_id, anl_accuracy.Anl_source, anl_accuracy.Accuracy_obs_no, anl_accuracy.Accuracy_desc, anl_accuracy.Accuracy_type, anl_accuracy.Active_ind, anl_accuracy.Calculated_ind, anl_accuracy.Calculate_method_id, anl_accuracy.Confidence_type, anl_accuracy.Effective_date, anl_accuracy.Expiry_date, anl_accuracy.Measured_value_ind, anl_accuracy.Ppdm_guid, anl_accuracy.Raw_value_ind, anl_accuracy.Referenced_column_name, anl_accuracy.Referenced_ppdm_guid, anl_accuracy.Referenced_system_id, anl_accuracy.Referenced_table_name, anl_accuracy.Remark, anl_accuracy.Repeatability_type, anl_accuracy.Reported_ind, anl_accuracy.Source, anl_accuracy.Step_seq_no, anl_accuracy.Row_changed_by, anl_accuracy.Row_changed_date, anl_accuracy.Row_created_by, anl_accuracy.Row_created_date, anl_accuracy.Row_effective_date, anl_accuracy.Row_expiry_date, anl_accuracy.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateAnlAccuracy(c *fiber.Ctx) error {
	var anl_accuracy dto.Anl_accuracy

	setDefaults(&anl_accuracy)

	if err := json.Unmarshal(c.Body(), &anl_accuracy); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	anl_accuracy.Analysis_id = id

    if anl_accuracy.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update anl_accuracy set anl_source = :1, accuracy_obs_no = :2, accuracy_desc = :3, accuracy_type = :4, active_ind = :5, calculated_ind = :6, calculate_method_id = :7, confidence_type = :8, effective_date = :9, expiry_date = :10, measured_value_ind = :11, ppdm_guid = :12, raw_value_ind = :13, referenced_column_name = :14, referenced_ppdm_guid = :15, referenced_system_id = :16, referenced_table_name = :17, remark = :18, repeatability_type = :19, reported_ind = :20, source = :21, step_seq_no = :22, row_changed_by = :23, row_changed_date = :24, row_created_by = :25, row_effective_date = :26, row_expiry_date = :27, row_quality = :28 where analysis_id = :30")
	if err != nil {
		return err
	}

	anl_accuracy.Row_changed_date = formatDate(anl_accuracy.Row_changed_date)
	anl_accuracy.Effective_date = formatDateString(anl_accuracy.Effective_date)
	anl_accuracy.Expiry_date = formatDateString(anl_accuracy.Expiry_date)
	anl_accuracy.Row_effective_date = formatDateString(anl_accuracy.Row_effective_date)
	anl_accuracy.Row_expiry_date = formatDateString(anl_accuracy.Row_expiry_date)






	rows, err := stmt.Exec(anl_accuracy.Anl_source, anl_accuracy.Accuracy_obs_no, anl_accuracy.Accuracy_desc, anl_accuracy.Accuracy_type, anl_accuracy.Active_ind, anl_accuracy.Calculated_ind, anl_accuracy.Calculate_method_id, anl_accuracy.Confidence_type, anl_accuracy.Effective_date, anl_accuracy.Expiry_date, anl_accuracy.Measured_value_ind, anl_accuracy.Ppdm_guid, anl_accuracy.Raw_value_ind, anl_accuracy.Referenced_column_name, anl_accuracy.Referenced_ppdm_guid, anl_accuracy.Referenced_system_id, anl_accuracy.Referenced_table_name, anl_accuracy.Remark, anl_accuracy.Repeatability_type, anl_accuracy.Reported_ind, anl_accuracy.Source, anl_accuracy.Step_seq_no, anl_accuracy.Row_changed_by, anl_accuracy.Row_changed_date, anl_accuracy.Row_created_by, anl_accuracy.Row_effective_date, anl_accuracy.Row_expiry_date, anl_accuracy.Row_quality, anl_accuracy.Analysis_id)
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

func PatchAnlAccuracy(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update anl_accuracy set "
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

func DeleteAnlAccuracy(c *fiber.Ctx) error {
	id := c.Params("id")
	var anl_accuracy dto.Anl_accuracy
	anl_accuracy.Analysis_id = id

	stmt, err := db.Prepare("delete from anl_accuracy where analysis_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(anl_accuracy.Analysis_id)
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


