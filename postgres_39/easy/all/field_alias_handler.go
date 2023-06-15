package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetFieldAlias(c *fiber.Ctx) error {
	rows, err := db.Query("select * from field_alias")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Field_alias

	for rows.Next() {
		var field_alias dto.Field_alias
		if err := rows.Scan(&field_alias.Field_id, &field_alias.Field_alias_id, &field_alias.Source, &field_alias.Abbreviation, &field_alias.Active_ind, &field_alias.Alias_long_name, &field_alias.Alias_reason_type, &field_alias.Alias_short_name, &field_alias.Alias_type, &field_alias.Amended_date, &field_alias.Area_id, &field_alias.Area_type, &field_alias.Created_date, &field_alias.Effective_date, &field_alias.Expiry_date, &field_alias.Field_id, &field_alias.Original_ind, &field_alias.Owner_ba_id, &field_alias.Ppdm_guid, &field_alias.Preferred_ind, &field_alias.Reason_desc, &field_alias.Remark, &field_alias.Source_document_id, &field_alias.Struckoff_date, &field_alias.Sw_application_id, &field_alias.Use_rule, &field_alias.Row_changed_by, &field_alias.Row_changed_date, &field_alias.Row_created_by, &field_alias.Row_created_date, &field_alias.Row_effective_date, &field_alias.Row_expiry_date, &field_alias.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append field_alias to result
		result = append(result, field_alias)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetFieldAlias(c *fiber.Ctx) error {
	var field_alias dto.Field_alias

	setDefaults(&field_alias)

	if err := json.Unmarshal(c.Body(), &field_alias); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into field_alias values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32, :33)")
	if err != nil {
		return err
	}
	field_alias.Row_created_date = formatDate(field_alias.Row_created_date)
	field_alias.Row_changed_date = formatDate(field_alias.Row_changed_date)
	field_alias.Amended_date = formatDateString(field_alias.Amended_date)
	field_alias.Created_date = formatDateString(field_alias.Created_date)
	field_alias.Effective_date = formatDateString(field_alias.Effective_date)
	field_alias.Expiry_date = formatDateString(field_alias.Expiry_date)
	field_alias.Struckoff_date = formatDateString(field_alias.Struckoff_date)
	field_alias.Row_effective_date = formatDateString(field_alias.Row_effective_date)
	field_alias.Row_expiry_date = formatDateString(field_alias.Row_expiry_date)






	rows, err := stmt.Exec(field_alias.Field_id, field_alias.Field_alias_id, field_alias.Source, field_alias.Abbreviation, field_alias.Active_ind, field_alias.Alias_long_name, field_alias.Alias_reason_type, field_alias.Alias_short_name, field_alias.Alias_type, field_alias.Amended_date, field_alias.Area_id, field_alias.Area_type, field_alias.Created_date, field_alias.Effective_date, field_alias.Expiry_date, field_alias.Field_id, field_alias.Original_ind, field_alias.Owner_ba_id, field_alias.Ppdm_guid, field_alias.Preferred_ind, field_alias.Reason_desc, field_alias.Remark, field_alias.Source_document_id, field_alias.Struckoff_date, field_alias.Sw_application_id, field_alias.Use_rule, field_alias.Row_changed_by, field_alias.Row_changed_date, field_alias.Row_created_by, field_alias.Row_created_date, field_alias.Row_effective_date, field_alias.Row_expiry_date, field_alias.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateFieldAlias(c *fiber.Ctx) error {
	var field_alias dto.Field_alias

	setDefaults(&field_alias)

	if err := json.Unmarshal(c.Body(), &field_alias); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	field_alias.Field_id = id

    if field_alias.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update field_alias set field_alias_id = :1, source = :2, abbreviation = :3, active_ind = :4, alias_long_name = :5, alias_reason_type = :6, alias_short_name = :7, alias_type = :8, amended_date = :9, area_id = :10, area_type = :11, created_date = :12, effective_date = :13, expiry_date = :14, field_id = :15, original_ind = :16, owner_ba_id = :17, ppdm_guid = :18, preferred_ind = :19, reason_desc = :20, remark = :21, source_document_id = :22, struckoff_date = :23, sw_application_id = :24, use_rule = :25, row_changed_by = :26, row_changed_date = :27, row_created_by = :28, row_effective_date = :29, row_expiry_date = :30, row_quality = :31 where field_id = :33")
	if err != nil {
		return err
	}

	field_alias.Row_changed_date = formatDate(field_alias.Row_changed_date)
	field_alias.Amended_date = formatDateString(field_alias.Amended_date)
	field_alias.Created_date = formatDateString(field_alias.Created_date)
	field_alias.Effective_date = formatDateString(field_alias.Effective_date)
	field_alias.Expiry_date = formatDateString(field_alias.Expiry_date)
	field_alias.Struckoff_date = formatDateString(field_alias.Struckoff_date)
	field_alias.Row_effective_date = formatDateString(field_alias.Row_effective_date)
	field_alias.Row_expiry_date = formatDateString(field_alias.Row_expiry_date)






	rows, err := stmt.Exec(field_alias.Field_alias_id, field_alias.Source, field_alias.Abbreviation, field_alias.Active_ind, field_alias.Alias_long_name, field_alias.Alias_reason_type, field_alias.Alias_short_name, field_alias.Alias_type, field_alias.Amended_date, field_alias.Area_id, field_alias.Area_type, field_alias.Created_date, field_alias.Effective_date, field_alias.Expiry_date, field_alias.Field_id, field_alias.Original_ind, field_alias.Owner_ba_id, field_alias.Ppdm_guid, field_alias.Preferred_ind, field_alias.Reason_desc, field_alias.Remark, field_alias.Source_document_id, field_alias.Struckoff_date, field_alias.Sw_application_id, field_alias.Use_rule, field_alias.Row_changed_by, field_alias.Row_changed_date, field_alias.Row_created_by, field_alias.Row_effective_date, field_alias.Row_expiry_date, field_alias.Row_quality, field_alias.Field_id)
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

func PatchFieldAlias(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update field_alias set "
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
	query += " where field_id = :id"

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

func DeleteFieldAlias(c *fiber.Ctx) error {
	id := c.Params("id")
	var field_alias dto.Field_alias
	field_alias.Field_id = id

	stmt, err := db.Prepare("delete from field_alias where field_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(field_alias.Field_id)
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


