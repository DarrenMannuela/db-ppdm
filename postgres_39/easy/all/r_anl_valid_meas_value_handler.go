package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetRAnlValidMeasValue(c *fiber.Ctx) error {
	rows, err := db.Query("select * from r_anl_valid_meas_value")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.R_anl_valid_meas_value

	for rows.Next() {
		var r_anl_valid_meas_value dto.R_anl_valid_meas_value
		if err := rows.Scan(&r_anl_valid_meas_value.Measurement_type, &r_anl_valid_meas_value.Valid_value_text, &r_anl_valid_meas_value.Abbreviation, &r_anl_valid_meas_value.Active_ind, &r_anl_valid_meas_value.Effective_date, &r_anl_valid_meas_value.Expiry_date, &r_anl_valid_meas_value.Long_name, &r_anl_valid_meas_value.Ppdm_guid, &r_anl_valid_meas_value.Remark, &r_anl_valid_meas_value.Short_name, &r_anl_valid_meas_value.Source, &r_anl_valid_meas_value.Row_changed_by, &r_anl_valid_meas_value.Row_changed_date, &r_anl_valid_meas_value.Row_created_by, &r_anl_valid_meas_value.Row_created_date, &r_anl_valid_meas_value.Row_effective_date, &r_anl_valid_meas_value.Row_expiry_date, &r_anl_valid_meas_value.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append r_anl_valid_meas_value to result
		result = append(result, r_anl_valid_meas_value)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetRAnlValidMeasValue(c *fiber.Ctx) error {
	var r_anl_valid_meas_value dto.R_anl_valid_meas_value

	setDefaults(&r_anl_valid_meas_value)

	if err := json.Unmarshal(c.Body(), &r_anl_valid_meas_value); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into r_anl_valid_meas_value values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18)")
	if err != nil {
		return err
	}
	r_anl_valid_meas_value.Row_created_date = formatDate(r_anl_valid_meas_value.Row_created_date)
	r_anl_valid_meas_value.Row_changed_date = formatDate(r_anl_valid_meas_value.Row_changed_date)
	r_anl_valid_meas_value.Effective_date = formatDateString(r_anl_valid_meas_value.Effective_date)
	r_anl_valid_meas_value.Expiry_date = formatDateString(r_anl_valid_meas_value.Expiry_date)
	r_anl_valid_meas_value.Row_effective_date = formatDateString(r_anl_valid_meas_value.Row_effective_date)
	r_anl_valid_meas_value.Row_expiry_date = formatDateString(r_anl_valid_meas_value.Row_expiry_date)






	rows, err := stmt.Exec(r_anl_valid_meas_value.Measurement_type, r_anl_valid_meas_value.Valid_value_text, r_anl_valid_meas_value.Abbreviation, r_anl_valid_meas_value.Active_ind, r_anl_valid_meas_value.Effective_date, r_anl_valid_meas_value.Expiry_date, r_anl_valid_meas_value.Long_name, r_anl_valid_meas_value.Ppdm_guid, r_anl_valid_meas_value.Remark, r_anl_valid_meas_value.Short_name, r_anl_valid_meas_value.Source, r_anl_valid_meas_value.Row_changed_by, r_anl_valid_meas_value.Row_changed_date, r_anl_valid_meas_value.Row_created_by, r_anl_valid_meas_value.Row_created_date, r_anl_valid_meas_value.Row_effective_date, r_anl_valid_meas_value.Row_expiry_date, r_anl_valid_meas_value.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateRAnlValidMeasValue(c *fiber.Ctx) error {
	var r_anl_valid_meas_value dto.R_anl_valid_meas_value

	setDefaults(&r_anl_valid_meas_value)

	if err := json.Unmarshal(c.Body(), &r_anl_valid_meas_value); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	r_anl_valid_meas_value.Measurement_type = id

    if r_anl_valid_meas_value.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update r_anl_valid_meas_value set valid_value_text = :1, abbreviation = :2, active_ind = :3, effective_date = :4, expiry_date = :5, long_name = :6, ppdm_guid = :7, remark = :8, short_name = :9, source = :10, row_changed_by = :11, row_changed_date = :12, row_created_by = :13, row_effective_date = :14, row_expiry_date = :15, row_quality = :16 where measurement_type = :18")
	if err != nil {
		return err
	}

	r_anl_valid_meas_value.Row_changed_date = formatDate(r_anl_valid_meas_value.Row_changed_date)
	r_anl_valid_meas_value.Effective_date = formatDateString(r_anl_valid_meas_value.Effective_date)
	r_anl_valid_meas_value.Expiry_date = formatDateString(r_anl_valid_meas_value.Expiry_date)
	r_anl_valid_meas_value.Row_effective_date = formatDateString(r_anl_valid_meas_value.Row_effective_date)
	r_anl_valid_meas_value.Row_expiry_date = formatDateString(r_anl_valid_meas_value.Row_expiry_date)






	rows, err := stmt.Exec(r_anl_valid_meas_value.Valid_value_text, r_anl_valid_meas_value.Abbreviation, r_anl_valid_meas_value.Active_ind, r_anl_valid_meas_value.Effective_date, r_anl_valid_meas_value.Expiry_date, r_anl_valid_meas_value.Long_name, r_anl_valid_meas_value.Ppdm_guid, r_anl_valid_meas_value.Remark, r_anl_valid_meas_value.Short_name, r_anl_valid_meas_value.Source, r_anl_valid_meas_value.Row_changed_by, r_anl_valid_meas_value.Row_changed_date, r_anl_valid_meas_value.Row_created_by, r_anl_valid_meas_value.Row_effective_date, r_anl_valid_meas_value.Row_expiry_date, r_anl_valid_meas_value.Row_quality, r_anl_valid_meas_value.Measurement_type)
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

func PatchRAnlValidMeasValue(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update r_anl_valid_meas_value set "
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
	query += " where measurement_type = :id"

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

func DeleteRAnlValidMeasValue(c *fiber.Ctx) error {
	id := c.Params("id")
	var r_anl_valid_meas_value dto.R_anl_valid_meas_value
	r_anl_valid_meas_value.Measurement_type = id

	stmt, err := db.Prepare("delete from r_anl_valid_meas_value where measurement_type = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(r_anl_valid_meas_value.Measurement_type)
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


