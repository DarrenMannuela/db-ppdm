package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetRaLanguage(c *fiber.Ctx) error {
	rows, err := db.Query("select * from ra_language")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Ra_language

	for rows.Next() {
		var ra_language dto.Ra_language
		if err := rows.Scan(&ra_language.Language_id, &ra_language.Alias_id, &ra_language.Abbreviation, &ra_language.Active_ind, &ra_language.Alias_long_name, &ra_language.Alias_reason_type, &ra_language.Alias_short_name, &ra_language.Alias_type, &ra_language.Amended_date, &ra_language.Created_date, &ra_language.Effective_date, &ra_language.Expiry_date, &ra_language.Original_ind, &ra_language.Owner_ba_id, &ra_language.Ppdm_guid, &ra_language.Preferred_ind, &ra_language.Reason_desc, &ra_language.Remark, &ra_language.Source, &ra_language.Source_document, &ra_language.Struckoff_date, &ra_language.Sw_application_id, &ra_language.Use_rule, &ra_language.Row_changed_by, &ra_language.Row_changed_date, &ra_language.Row_created_by, &ra_language.Row_created_date, &ra_language.Row_effective_date, &ra_language.Row_expiry_date, &ra_language.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append ra_language to result
		result = append(result, ra_language)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetRaLanguage(c *fiber.Ctx) error {
	var ra_language dto.Ra_language

	setDefaults(&ra_language)

	if err := json.Unmarshal(c.Body(), &ra_language); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into ra_language values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30)")
	if err != nil {
		return err
	}
	ra_language.Row_created_date = formatDate(ra_language.Row_created_date)
	ra_language.Row_changed_date = formatDate(ra_language.Row_changed_date)
	ra_language.Amended_date = formatDateString(ra_language.Amended_date)
	ra_language.Created_date = formatDateString(ra_language.Created_date)
	ra_language.Effective_date = formatDateString(ra_language.Effective_date)
	ra_language.Expiry_date = formatDateString(ra_language.Expiry_date)
	ra_language.Struckoff_date = formatDateString(ra_language.Struckoff_date)
	ra_language.Row_effective_date = formatDateString(ra_language.Row_effective_date)
	ra_language.Row_expiry_date = formatDateString(ra_language.Row_expiry_date)






	rows, err := stmt.Exec(ra_language.Language_id, ra_language.Alias_id, ra_language.Abbreviation, ra_language.Active_ind, ra_language.Alias_long_name, ra_language.Alias_reason_type, ra_language.Alias_short_name, ra_language.Alias_type, ra_language.Amended_date, ra_language.Created_date, ra_language.Effective_date, ra_language.Expiry_date, ra_language.Original_ind, ra_language.Owner_ba_id, ra_language.Ppdm_guid, ra_language.Preferred_ind, ra_language.Reason_desc, ra_language.Remark, ra_language.Source, ra_language.Source_document, ra_language.Struckoff_date, ra_language.Sw_application_id, ra_language.Use_rule, ra_language.Row_changed_by, ra_language.Row_changed_date, ra_language.Row_created_by, ra_language.Row_created_date, ra_language.Row_effective_date, ra_language.Row_expiry_date, ra_language.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateRaLanguage(c *fiber.Ctx) error {
	var ra_language dto.Ra_language

	setDefaults(&ra_language)

	if err := json.Unmarshal(c.Body(), &ra_language); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	ra_language.Language_id = id

    if ra_language.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update ra_language set alias_id = :1, abbreviation = :2, active_ind = :3, alias_long_name = :4, alias_reason_type = :5, alias_short_name = :6, alias_type = :7, amended_date = :8, created_date = :9, effective_date = :10, expiry_date = :11, original_ind = :12, owner_ba_id = :13, ppdm_guid = :14, preferred_ind = :15, reason_desc = :16, remark = :17, source = :18, source_document = :19, struckoff_date = :20, sw_application_id = :21, use_rule = :22, row_changed_by = :23, row_changed_date = :24, row_created_by = :25, row_effective_date = :26, row_expiry_date = :27, row_quality = :28 where language_id = :30")
	if err != nil {
		return err
	}

	ra_language.Row_changed_date = formatDate(ra_language.Row_changed_date)
	ra_language.Amended_date = formatDateString(ra_language.Amended_date)
	ra_language.Created_date = formatDateString(ra_language.Created_date)
	ra_language.Effective_date = formatDateString(ra_language.Effective_date)
	ra_language.Expiry_date = formatDateString(ra_language.Expiry_date)
	ra_language.Struckoff_date = formatDateString(ra_language.Struckoff_date)
	ra_language.Row_effective_date = formatDateString(ra_language.Row_effective_date)
	ra_language.Row_expiry_date = formatDateString(ra_language.Row_expiry_date)






	rows, err := stmt.Exec(ra_language.Alias_id, ra_language.Abbreviation, ra_language.Active_ind, ra_language.Alias_long_name, ra_language.Alias_reason_type, ra_language.Alias_short_name, ra_language.Alias_type, ra_language.Amended_date, ra_language.Created_date, ra_language.Effective_date, ra_language.Expiry_date, ra_language.Original_ind, ra_language.Owner_ba_id, ra_language.Ppdm_guid, ra_language.Preferred_ind, ra_language.Reason_desc, ra_language.Remark, ra_language.Source, ra_language.Source_document, ra_language.Struckoff_date, ra_language.Sw_application_id, ra_language.Use_rule, ra_language.Row_changed_by, ra_language.Row_changed_date, ra_language.Row_created_by, ra_language.Row_effective_date, ra_language.Row_expiry_date, ra_language.Row_quality, ra_language.Language_id)
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

func PatchRaLanguage(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update ra_language set "
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
	query += " where language_id = :id"

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

func DeleteRaLanguage(c *fiber.Ctx) error {
	id := c.Params("id")
	var ra_language dto.Ra_language
	ra_language.Language_id = id

	stmt, err := db.Prepare("delete from ra_language where language_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(ra_language.Language_id)
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


