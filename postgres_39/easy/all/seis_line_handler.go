package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetSeisLine(c *fiber.Ctx) error {
	rows, err := db.Query("select * from seis_line")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Seis_line

	for rows.Next() {
		var seis_line dto.Seis_line
		if err := rows.Scan(&seis_line.Seis_set_subtype, &seis_line.Seis_line_id, &seis_line.Active_ind, &seis_line.Dimension, &seis_line.Effective_date, &seis_line.Expiry_date, &seis_line.Line_name, &seis_line.Ppdm_guid, &seis_line.Refraction_reflection, &seis_line.Remark, &seis_line.Reshoot_of_set_id, &seis_line.Reshoot_seis_set_subtype, &seis_line.Seis_acqtn_set_subtype, &seis_line.Seis_acqtn_survey_id, &seis_line.Seis_spectrum_type, &seis_line.Source, &seis_line.Test_experimental, &seis_line.Row_changed_by, &seis_line.Row_changed_date, &seis_line.Row_created_by, &seis_line.Row_created_date, &seis_line.Row_effective_date, &seis_line.Row_expiry_date, &seis_line.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append seis_line to result
		result = append(result, seis_line)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetSeisLine(c *fiber.Ctx) error {
	var seis_line dto.Seis_line

	setDefaults(&seis_line)

	if err := json.Unmarshal(c.Body(), &seis_line); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into seis_line values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24)")
	if err != nil {
		return err
	}
	seis_line.Row_created_date = formatDate(seis_line.Row_created_date)
	seis_line.Row_changed_date = formatDate(seis_line.Row_changed_date)
	seis_line.Effective_date = formatDateString(seis_line.Effective_date)
	seis_line.Expiry_date = formatDateString(seis_line.Expiry_date)
	seis_line.Row_effective_date = formatDateString(seis_line.Row_effective_date)
	seis_line.Row_expiry_date = formatDateString(seis_line.Row_expiry_date)






	rows, err := stmt.Exec(seis_line.Seis_set_subtype, seis_line.Seis_line_id, seis_line.Active_ind, seis_line.Dimension, seis_line.Effective_date, seis_line.Expiry_date, seis_line.Line_name, seis_line.Ppdm_guid, seis_line.Refraction_reflection, seis_line.Remark, seis_line.Reshoot_of_set_id, seis_line.Reshoot_seis_set_subtype, seis_line.Seis_acqtn_set_subtype, seis_line.Seis_acqtn_survey_id, seis_line.Seis_spectrum_type, seis_line.Source, seis_line.Test_experimental, seis_line.Row_changed_by, seis_line.Row_changed_date, seis_line.Row_created_by, seis_line.Row_created_date, seis_line.Row_effective_date, seis_line.Row_expiry_date, seis_line.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateSeisLine(c *fiber.Ctx) error {
	var seis_line dto.Seis_line

	setDefaults(&seis_line)

	if err := json.Unmarshal(c.Body(), &seis_line); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	seis_line.Seis_set_subtype = id

    if seis_line.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update seis_line set seis_line_id = :1, active_ind = :2, dimension = :3, effective_date = :4, expiry_date = :5, line_name = :6, ppdm_guid = :7, refraction_reflection = :8, remark = :9, reshoot_of_set_id = :10, reshoot_seis_set_subtype = :11, seis_acqtn_set_subtype = :12, seis_acqtn_survey_id = :13, seis_spectrum_type = :14, source = :15, test_experimental = :16, row_changed_by = :17, row_changed_date = :18, row_created_by = :19, row_effective_date = :20, row_expiry_date = :21, row_quality = :22 where seis_set_subtype = :24")
	if err != nil {
		return err
	}

	seis_line.Row_changed_date = formatDate(seis_line.Row_changed_date)
	seis_line.Effective_date = formatDateString(seis_line.Effective_date)
	seis_line.Expiry_date = formatDateString(seis_line.Expiry_date)
	seis_line.Row_effective_date = formatDateString(seis_line.Row_effective_date)
	seis_line.Row_expiry_date = formatDateString(seis_line.Row_expiry_date)






	rows, err := stmt.Exec(seis_line.Seis_line_id, seis_line.Active_ind, seis_line.Dimension, seis_line.Effective_date, seis_line.Expiry_date, seis_line.Line_name, seis_line.Ppdm_guid, seis_line.Refraction_reflection, seis_line.Remark, seis_line.Reshoot_of_set_id, seis_line.Reshoot_seis_set_subtype, seis_line.Seis_acqtn_set_subtype, seis_line.Seis_acqtn_survey_id, seis_line.Seis_spectrum_type, seis_line.Source, seis_line.Test_experimental, seis_line.Row_changed_by, seis_line.Row_changed_date, seis_line.Row_created_by, seis_line.Row_effective_date, seis_line.Row_expiry_date, seis_line.Row_quality, seis_line.Seis_set_subtype)
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

func PatchSeisLine(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update seis_line set "
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
	query += " where seis_set_subtype = :id"

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

func DeleteSeisLine(c *fiber.Ctx) error {
	id := c.Params("id")
	var seis_line dto.Seis_line
	seis_line.Seis_set_subtype = id

	stmt, err := db.Prepare("delete from seis_line where seis_set_subtype = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(seis_line.Seis_set_subtype)
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


