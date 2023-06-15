package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetAnlMethod(c *fiber.Ctx) error {
	rows, err := db.Query("select * from anl_method")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Anl_method

	for rows.Next() {
		var anl_method dto.Anl_method
		if err := rows.Scan(&anl_method.Method_set_id, &anl_method.Method_id, &anl_method.Abbreviation, &anl_method.Active_ind, &anl_method.Effective_date, &anl_method.Expiry_date, &anl_method.Long_name, &anl_method.Method_desc, &anl_method.Method_seq_no, &anl_method.Method_type, &anl_method.Ppdm_guid, &anl_method.Preparation_class, &anl_method.Remark, &anl_method.Short_name, &anl_method.Source, &anl_method.Source_document_id, &anl_method.Row_changed_by, &anl_method.Row_changed_date, &anl_method.Row_created_by, &anl_method.Row_created_date, &anl_method.Row_effective_date, &anl_method.Row_expiry_date, &anl_method.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append anl_method to result
		result = append(result, anl_method)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetAnlMethod(c *fiber.Ctx) error {
	var anl_method dto.Anl_method

	setDefaults(&anl_method)

	if err := json.Unmarshal(c.Body(), &anl_method); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into anl_method values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23)")
	if err != nil {
		return err
	}
	anl_method.Row_created_date = formatDate(anl_method.Row_created_date)
	anl_method.Row_changed_date = formatDate(anl_method.Row_changed_date)
	anl_method.Effective_date = formatDateString(anl_method.Effective_date)
	anl_method.Expiry_date = formatDateString(anl_method.Expiry_date)
	anl_method.Row_effective_date = formatDateString(anl_method.Row_effective_date)
	anl_method.Row_expiry_date = formatDateString(anl_method.Row_expiry_date)






	rows, err := stmt.Exec(anl_method.Method_set_id, anl_method.Method_id, anl_method.Abbreviation, anl_method.Active_ind, anl_method.Effective_date, anl_method.Expiry_date, anl_method.Long_name, anl_method.Method_desc, anl_method.Method_seq_no, anl_method.Method_type, anl_method.Ppdm_guid, anl_method.Preparation_class, anl_method.Remark, anl_method.Short_name, anl_method.Source, anl_method.Source_document_id, anl_method.Row_changed_by, anl_method.Row_changed_date, anl_method.Row_created_by, anl_method.Row_created_date, anl_method.Row_effective_date, anl_method.Row_expiry_date, anl_method.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateAnlMethod(c *fiber.Ctx) error {
	var anl_method dto.Anl_method

	setDefaults(&anl_method)

	if err := json.Unmarshal(c.Body(), &anl_method); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	anl_method.Method_set_id = id

    if anl_method.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update anl_method set method_id = :1, abbreviation = :2, active_ind = :3, effective_date = :4, expiry_date = :5, long_name = :6, method_desc = :7, method_seq_no = :8, method_type = :9, ppdm_guid = :10, preparation_class = :11, remark = :12, short_name = :13, source = :14, source_document_id = :15, row_changed_by = :16, row_changed_date = :17, row_created_by = :18, row_effective_date = :19, row_expiry_date = :20, row_quality = :21 where method_set_id = :23")
	if err != nil {
		return err
	}

	anl_method.Row_changed_date = formatDate(anl_method.Row_changed_date)
	anl_method.Effective_date = formatDateString(anl_method.Effective_date)
	anl_method.Expiry_date = formatDateString(anl_method.Expiry_date)
	anl_method.Row_effective_date = formatDateString(anl_method.Row_effective_date)
	anl_method.Row_expiry_date = formatDateString(anl_method.Row_expiry_date)






	rows, err := stmt.Exec(anl_method.Method_id, anl_method.Abbreviation, anl_method.Active_ind, anl_method.Effective_date, anl_method.Expiry_date, anl_method.Long_name, anl_method.Method_desc, anl_method.Method_seq_no, anl_method.Method_type, anl_method.Ppdm_guid, anl_method.Preparation_class, anl_method.Remark, anl_method.Short_name, anl_method.Source, anl_method.Source_document_id, anl_method.Row_changed_by, anl_method.Row_changed_date, anl_method.Row_created_by, anl_method.Row_effective_date, anl_method.Row_expiry_date, anl_method.Row_quality, anl_method.Method_set_id)
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

func PatchAnlMethod(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update anl_method set "
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

func DeleteAnlMethod(c *fiber.Ctx) error {
	id := c.Params("id")
	var anl_method dto.Anl_method
	anl_method.Method_set_id = id

	stmt, err := db.Prepare("delete from anl_method where method_set_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(anl_method.Method_set_id)
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


