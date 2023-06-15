package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetRaAnlValidMeasValue(c *fiber.Ctx) error {
	rows, err := db.Query("select * from ra_anl_valid_meas_value")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Ra_anl_valid_meas_value

	for rows.Next() {
		var ra_anl_valid_meas_value dto.Ra_anl_valid_meas_value
		if err := rows.Scan(&ra_anl_valid_meas_value.Measurement_type, &ra_anl_valid_meas_value.Valid_value_text, &ra_anl_valid_meas_value.Alias_id, &ra_anl_valid_meas_value.Abbreviation, &ra_anl_valid_meas_value.Active_ind, &ra_anl_valid_meas_value.Alias_long_name, &ra_anl_valid_meas_value.Alias_reason_type, &ra_anl_valid_meas_value.Alias_short_name, &ra_anl_valid_meas_value.Alias_type, &ra_anl_valid_meas_value.Amended_date, &ra_anl_valid_meas_value.Created_date, &ra_anl_valid_meas_value.Effective_date, &ra_anl_valid_meas_value.Expiry_date, &ra_anl_valid_meas_value.Measurement_type, &ra_anl_valid_meas_value.Original_ind, &ra_anl_valid_meas_value.Owner_ba_id, &ra_anl_valid_meas_value.Ppdm_guid, &ra_anl_valid_meas_value.Preferred_ind, &ra_anl_valid_meas_value.Reason_desc, &ra_anl_valid_meas_value.Remark, &ra_anl_valid_meas_value.Source, &ra_anl_valid_meas_value.Source_document, &ra_anl_valid_meas_value.Struckoff_date, &ra_anl_valid_meas_value.Sw_application_id, &ra_anl_valid_meas_value.Use_rule, &ra_anl_valid_meas_value.Row_changed_by, &ra_anl_valid_meas_value.Row_changed_date, &ra_anl_valid_meas_value.Row_created_by, &ra_anl_valid_meas_value.Row_created_date, &ra_anl_valid_meas_value.Row_effective_date, &ra_anl_valid_meas_value.Row_expiry_date, &ra_anl_valid_meas_value.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append ra_anl_valid_meas_value to result
		result = append(result, ra_anl_valid_meas_value)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetRaAnlValidMeasValue(c *fiber.Ctx) error {
	var ra_anl_valid_meas_value dto.Ra_anl_valid_meas_value

	setDefaults(&ra_anl_valid_meas_value)

	if err := json.Unmarshal(c.Body(), &ra_anl_valid_meas_value); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into ra_anl_valid_meas_value values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32)")
	if err != nil {
		return err
	}
	ra_anl_valid_meas_value.Row_created_date = formatDate(ra_anl_valid_meas_value.Row_created_date)
	ra_anl_valid_meas_value.Row_changed_date = formatDate(ra_anl_valid_meas_value.Row_changed_date)
	ra_anl_valid_meas_value.Amended_date = formatDateString(ra_anl_valid_meas_value.Amended_date)
	ra_anl_valid_meas_value.Created_date = formatDateString(ra_anl_valid_meas_value.Created_date)
	ra_anl_valid_meas_value.Effective_date = formatDateString(ra_anl_valid_meas_value.Effective_date)
	ra_anl_valid_meas_value.Expiry_date = formatDateString(ra_anl_valid_meas_value.Expiry_date)
	ra_anl_valid_meas_value.Struckoff_date = formatDateString(ra_anl_valid_meas_value.Struckoff_date)
	ra_anl_valid_meas_value.Row_effective_date = formatDateString(ra_anl_valid_meas_value.Row_effective_date)
	ra_anl_valid_meas_value.Row_expiry_date = formatDateString(ra_anl_valid_meas_value.Row_expiry_date)






	rows, err := stmt.Exec(ra_anl_valid_meas_value.Measurement_type, ra_anl_valid_meas_value.Valid_value_text, ra_anl_valid_meas_value.Alias_id, ra_anl_valid_meas_value.Abbreviation, ra_anl_valid_meas_value.Active_ind, ra_anl_valid_meas_value.Alias_long_name, ra_anl_valid_meas_value.Alias_reason_type, ra_anl_valid_meas_value.Alias_short_name, ra_anl_valid_meas_value.Alias_type, ra_anl_valid_meas_value.Amended_date, ra_anl_valid_meas_value.Created_date, ra_anl_valid_meas_value.Effective_date, ra_anl_valid_meas_value.Expiry_date, ra_anl_valid_meas_value.Measurement_type, ra_anl_valid_meas_value.Original_ind, ra_anl_valid_meas_value.Owner_ba_id, ra_anl_valid_meas_value.Ppdm_guid, ra_anl_valid_meas_value.Preferred_ind, ra_anl_valid_meas_value.Reason_desc, ra_anl_valid_meas_value.Remark, ra_anl_valid_meas_value.Source, ra_anl_valid_meas_value.Source_document, ra_anl_valid_meas_value.Struckoff_date, ra_anl_valid_meas_value.Sw_application_id, ra_anl_valid_meas_value.Use_rule, ra_anl_valid_meas_value.Row_changed_by, ra_anl_valid_meas_value.Row_changed_date, ra_anl_valid_meas_value.Row_created_by, ra_anl_valid_meas_value.Row_created_date, ra_anl_valid_meas_value.Row_effective_date, ra_anl_valid_meas_value.Row_expiry_date, ra_anl_valid_meas_value.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateRaAnlValidMeasValue(c *fiber.Ctx) error {
	var ra_anl_valid_meas_value dto.Ra_anl_valid_meas_value

	setDefaults(&ra_anl_valid_meas_value)

	if err := json.Unmarshal(c.Body(), &ra_anl_valid_meas_value); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	ra_anl_valid_meas_value.Measurement_type = id

    if ra_anl_valid_meas_value.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update ra_anl_valid_meas_value set valid_value_text = :1, alias_id = :2, abbreviation = :3, active_ind = :4, alias_long_name = :5, alias_reason_type = :6, alias_short_name = :7, alias_type = :8, amended_date = :9, created_date = :10, effective_date = :11, expiry_date = :12, measurement_type = :13, original_ind = :14, owner_ba_id = :15, ppdm_guid = :16, preferred_ind = :17, reason_desc = :18, remark = :19, source = :20, source_document = :21, struckoff_date = :22, sw_application_id = :23, use_rule = :24, row_changed_by = :25, row_changed_date = :26, row_created_by = :27, row_effective_date = :28, row_expiry_date = :29, row_quality = :30 where measurement_type = :32")
	if err != nil {
		return err
	}

	ra_anl_valid_meas_value.Row_changed_date = formatDate(ra_anl_valid_meas_value.Row_changed_date)
	ra_anl_valid_meas_value.Amended_date = formatDateString(ra_anl_valid_meas_value.Amended_date)
	ra_anl_valid_meas_value.Created_date = formatDateString(ra_anl_valid_meas_value.Created_date)
	ra_anl_valid_meas_value.Effective_date = formatDateString(ra_anl_valid_meas_value.Effective_date)
	ra_anl_valid_meas_value.Expiry_date = formatDateString(ra_anl_valid_meas_value.Expiry_date)
	ra_anl_valid_meas_value.Struckoff_date = formatDateString(ra_anl_valid_meas_value.Struckoff_date)
	ra_anl_valid_meas_value.Row_effective_date = formatDateString(ra_anl_valid_meas_value.Row_effective_date)
	ra_anl_valid_meas_value.Row_expiry_date = formatDateString(ra_anl_valid_meas_value.Row_expiry_date)






	rows, err := stmt.Exec(ra_anl_valid_meas_value.Valid_value_text, ra_anl_valid_meas_value.Alias_id, ra_anl_valid_meas_value.Abbreviation, ra_anl_valid_meas_value.Active_ind, ra_anl_valid_meas_value.Alias_long_name, ra_anl_valid_meas_value.Alias_reason_type, ra_anl_valid_meas_value.Alias_short_name, ra_anl_valid_meas_value.Alias_type, ra_anl_valid_meas_value.Amended_date, ra_anl_valid_meas_value.Created_date, ra_anl_valid_meas_value.Effective_date, ra_anl_valid_meas_value.Expiry_date, ra_anl_valid_meas_value.Measurement_type, ra_anl_valid_meas_value.Original_ind, ra_anl_valid_meas_value.Owner_ba_id, ra_anl_valid_meas_value.Ppdm_guid, ra_anl_valid_meas_value.Preferred_ind, ra_anl_valid_meas_value.Reason_desc, ra_anl_valid_meas_value.Remark, ra_anl_valid_meas_value.Source, ra_anl_valid_meas_value.Source_document, ra_anl_valid_meas_value.Struckoff_date, ra_anl_valid_meas_value.Sw_application_id, ra_anl_valid_meas_value.Use_rule, ra_anl_valid_meas_value.Row_changed_by, ra_anl_valid_meas_value.Row_changed_date, ra_anl_valid_meas_value.Row_created_by, ra_anl_valid_meas_value.Row_effective_date, ra_anl_valid_meas_value.Row_expiry_date, ra_anl_valid_meas_value.Row_quality, ra_anl_valid_meas_value.Measurement_type)
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

func PatchRaAnlValidMeasValue(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update ra_anl_valid_meas_value set "
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
		} else if key == "amended_date" || key == "created_date" || key == "effective_date" || key == "expiry_date" || key == "struckoff_date" || key == "row_effective_date" || key == "row_expiry_date"      {
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

func DeleteRaAnlValidMeasValue(c *fiber.Ctx) error {
	id := c.Params("id")
	var ra_anl_valid_meas_value dto.Ra_anl_valid_meas_value
	ra_anl_valid_meas_value.Measurement_type = id

	stmt, err := db.Prepare("delete from ra_anl_valid_meas_value where measurement_type = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(ra_anl_valid_meas_value.Measurement_type)
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


