package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetRMeasureTechnique(c *fiber.Ctx) error {
	rows, err := db.Query("select * from r_measure_technique")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.R_measure_technique

	for rows.Next() {
		var r_measure_technique dto.R_measure_technique
		if err := rows.Scan(&r_measure_technique.Measurement_technique, &r_measure_technique.Abbreviation, &r_measure_technique.Active_ind, &r_measure_technique.Effective_date, &r_measure_technique.Expiry_date, &r_measure_technique.Long_name, &r_measure_technique.Ppdm_guid, &r_measure_technique.Remark, &r_measure_technique.Short_name, &r_measure_technique.Source, &r_measure_technique.Row_changed_by, &r_measure_technique.Row_changed_date, &r_measure_technique.Row_created_by, &r_measure_technique.Row_created_date, &r_measure_technique.Row_effective_date, &r_measure_technique.Row_expiry_date, &r_measure_technique.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append r_measure_technique to result
		result = append(result, r_measure_technique)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetRMeasureTechnique(c *fiber.Ctx) error {
	var r_measure_technique dto.R_measure_technique

	setDefaults(&r_measure_technique)

	if err := json.Unmarshal(c.Body(), &r_measure_technique); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into r_measure_technique values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17)")
	if err != nil {
		return err
	}
	r_measure_technique.Row_created_date = formatDate(r_measure_technique.Row_created_date)
	r_measure_technique.Row_changed_date = formatDate(r_measure_technique.Row_changed_date)
	r_measure_technique.Effective_date = formatDateString(r_measure_technique.Effective_date)
	r_measure_technique.Expiry_date = formatDateString(r_measure_technique.Expiry_date)
	r_measure_technique.Row_effective_date = formatDateString(r_measure_technique.Row_effective_date)
	r_measure_technique.Row_expiry_date = formatDateString(r_measure_technique.Row_expiry_date)






	rows, err := stmt.Exec(r_measure_technique.Measurement_technique, r_measure_technique.Abbreviation, r_measure_technique.Active_ind, r_measure_technique.Effective_date, r_measure_technique.Expiry_date, r_measure_technique.Long_name, r_measure_technique.Ppdm_guid, r_measure_technique.Remark, r_measure_technique.Short_name, r_measure_technique.Source, r_measure_technique.Row_changed_by, r_measure_technique.Row_changed_date, r_measure_technique.Row_created_by, r_measure_technique.Row_created_date, r_measure_technique.Row_effective_date, r_measure_technique.Row_expiry_date, r_measure_technique.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateRMeasureTechnique(c *fiber.Ctx) error {
	var r_measure_technique dto.R_measure_technique

	setDefaults(&r_measure_technique)

	if err := json.Unmarshal(c.Body(), &r_measure_technique); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	r_measure_technique.Measurement_technique = id

    if r_measure_technique.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update r_measure_technique set abbreviation = :1, active_ind = :2, effective_date = :3, expiry_date = :4, long_name = :5, ppdm_guid = :6, remark = :7, short_name = :8, source = :9, row_changed_by = :10, row_changed_date = :11, row_created_by = :12, row_effective_date = :13, row_expiry_date = :14, row_quality = :15 where measurement_technique = :17")
	if err != nil {
		return err
	}

	r_measure_technique.Row_changed_date = formatDate(r_measure_technique.Row_changed_date)
	r_measure_technique.Effective_date = formatDateString(r_measure_technique.Effective_date)
	r_measure_technique.Expiry_date = formatDateString(r_measure_technique.Expiry_date)
	r_measure_technique.Row_effective_date = formatDateString(r_measure_technique.Row_effective_date)
	r_measure_technique.Row_expiry_date = formatDateString(r_measure_technique.Row_expiry_date)






	rows, err := stmt.Exec(r_measure_technique.Abbreviation, r_measure_technique.Active_ind, r_measure_technique.Effective_date, r_measure_technique.Expiry_date, r_measure_technique.Long_name, r_measure_technique.Ppdm_guid, r_measure_technique.Remark, r_measure_technique.Short_name, r_measure_technique.Source, r_measure_technique.Row_changed_by, r_measure_technique.Row_changed_date, r_measure_technique.Row_created_by, r_measure_technique.Row_effective_date, r_measure_technique.Row_expiry_date, r_measure_technique.Row_quality, r_measure_technique.Measurement_technique)
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

func PatchRMeasureTechnique(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update r_measure_technique set "
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
	query += " where measurement_technique = :id"

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

func DeleteRMeasureTechnique(c *fiber.Ctx) error {
	id := c.Params("id")
	var r_measure_technique dto.R_measure_technique
	r_measure_technique.Measurement_technique = id

	stmt, err := db.Prepare("delete from r_measure_technique where measurement_technique = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(r_measure_technique.Measurement_technique)
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


