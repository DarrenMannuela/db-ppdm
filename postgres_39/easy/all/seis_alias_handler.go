package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetSeisAlias(c *fiber.Ctx) error {
	rows, err := db.Query("select * from seis_alias")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Seis_alias

	for rows.Next() {
		var seis_alias dto.Seis_alias
		if err := rows.Scan(&seis_alias.Seis_set_subtype, &seis_alias.Seis_set_id, &seis_alias.Seis_alias_id, &seis_alias.Abbreviation, &seis_alias.Active_ind, &seis_alias.Alias_long_name, &seis_alias.Alias_reason_type, &seis_alias.Alias_short_name, &seis_alias.Alias_type, &seis_alias.Amended_date, &seis_alias.Created_date, &seis_alias.Effective_date, &seis_alias.Expiry_date, &seis_alias.Seis_set_subtype, &seis_alias.Original_ind, &seis_alias.Owner_ba_id, &seis_alias.Ppdm_guid, &seis_alias.Preferred_ind, &seis_alias.Reason_desc, &seis_alias.Remark, &seis_alias.Source, &seis_alias.Source_document_id, &seis_alias.Struckoff_date, &seis_alias.Sw_application_id, &seis_alias.Use_rule, &seis_alias.Row_changed_by, &seis_alias.Row_changed_date, &seis_alias.Row_created_by, &seis_alias.Row_created_date, &seis_alias.Row_effective_date, &seis_alias.Row_expiry_date, &seis_alias.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append seis_alias to result
		result = append(result, seis_alias)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetSeisAlias(c *fiber.Ctx) error {
	var seis_alias dto.Seis_alias

	setDefaults(&seis_alias)

	if err := json.Unmarshal(c.Body(), &seis_alias); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into seis_alias values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32)")
	if err != nil {
		return err
	}
	seis_alias.Row_created_date = formatDate(seis_alias.Row_created_date)
	seis_alias.Row_changed_date = formatDate(seis_alias.Row_changed_date)
	seis_alias.Amended_date = formatDateString(seis_alias.Amended_date)
	seis_alias.Created_date = formatDateString(seis_alias.Created_date)
	seis_alias.Effective_date = formatDateString(seis_alias.Effective_date)
	seis_alias.Expiry_date = formatDateString(seis_alias.Expiry_date)
	seis_alias.Struckoff_date = formatDateString(seis_alias.Struckoff_date)
	seis_alias.Row_effective_date = formatDateString(seis_alias.Row_effective_date)
	seis_alias.Row_expiry_date = formatDateString(seis_alias.Row_expiry_date)






	rows, err := stmt.Exec(seis_alias.Seis_set_subtype, seis_alias.Seis_set_id, seis_alias.Seis_alias_id, seis_alias.Abbreviation, seis_alias.Active_ind, seis_alias.Alias_long_name, seis_alias.Alias_reason_type, seis_alias.Alias_short_name, seis_alias.Alias_type, seis_alias.Amended_date, seis_alias.Created_date, seis_alias.Effective_date, seis_alias.Expiry_date, seis_alias.Seis_set_subtype, seis_alias.Original_ind, seis_alias.Owner_ba_id, seis_alias.Ppdm_guid, seis_alias.Preferred_ind, seis_alias.Reason_desc, seis_alias.Remark, seis_alias.Source, seis_alias.Source_document_id, seis_alias.Struckoff_date, seis_alias.Sw_application_id, seis_alias.Use_rule, seis_alias.Row_changed_by, seis_alias.Row_changed_date, seis_alias.Row_created_by, seis_alias.Row_created_date, seis_alias.Row_effective_date, seis_alias.Row_expiry_date, seis_alias.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateSeisAlias(c *fiber.Ctx) error {
	var seis_alias dto.Seis_alias

	setDefaults(&seis_alias)

	if err := json.Unmarshal(c.Body(), &seis_alias); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	seis_alias.Seis_set_subtype = id

    if seis_alias.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update seis_alias set seis_set_id = :1, seis_alias_id = :2, abbreviation = :3, active_ind = :4, alias_long_name = :5, alias_reason_type = :6, alias_short_name = :7, alias_type = :8, amended_date = :9, created_date = :10, effective_date = :11, expiry_date = :12, seis_set_subtype = :13, original_ind = :14, owner_ba_id = :15, ppdm_guid = :16, preferred_ind = :17, reason_desc = :18, remark = :19, source = :20, source_document_id = :21, struckoff_date = :22, sw_application_id = :23, use_rule = :24, row_changed_by = :25, row_changed_date = :26, row_created_by = :27, row_effective_date = :28, row_expiry_date = :29, row_quality = :30 where seis_set_subtype = :32")
	if err != nil {
		return err
	}

	seis_alias.Row_changed_date = formatDate(seis_alias.Row_changed_date)
	seis_alias.Amended_date = formatDateString(seis_alias.Amended_date)
	seis_alias.Created_date = formatDateString(seis_alias.Created_date)
	seis_alias.Effective_date = formatDateString(seis_alias.Effective_date)
	seis_alias.Expiry_date = formatDateString(seis_alias.Expiry_date)
	seis_alias.Struckoff_date = formatDateString(seis_alias.Struckoff_date)
	seis_alias.Row_effective_date = formatDateString(seis_alias.Row_effective_date)
	seis_alias.Row_expiry_date = formatDateString(seis_alias.Row_expiry_date)






	rows, err := stmt.Exec(seis_alias.Seis_set_id, seis_alias.Seis_alias_id, seis_alias.Abbreviation, seis_alias.Active_ind, seis_alias.Alias_long_name, seis_alias.Alias_reason_type, seis_alias.Alias_short_name, seis_alias.Alias_type, seis_alias.Amended_date, seis_alias.Created_date, seis_alias.Effective_date, seis_alias.Expiry_date, seis_alias.Seis_set_subtype, seis_alias.Original_ind, seis_alias.Owner_ba_id, seis_alias.Ppdm_guid, seis_alias.Preferred_ind, seis_alias.Reason_desc, seis_alias.Remark, seis_alias.Source, seis_alias.Source_document_id, seis_alias.Struckoff_date, seis_alias.Sw_application_id, seis_alias.Use_rule, seis_alias.Row_changed_by, seis_alias.Row_changed_date, seis_alias.Row_created_by, seis_alias.Row_effective_date, seis_alias.Row_expiry_date, seis_alias.Row_quality, seis_alias.Seis_set_subtype)
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

func PatchSeisAlias(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update seis_alias set "
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

func DeleteSeisAlias(c *fiber.Ctx) error {
	id := c.Params("id")
	var seis_alias dto.Seis_alias
	seis_alias.Seis_set_subtype = id

	stmt, err := db.Prepare("delete from seis_alias where seis_set_subtype = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(seis_alias.Seis_set_subtype)
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


