package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetWellLogDictAlias(c *fiber.Ctx) error {
	rows, err := db.Query("select * from well_log_dict_alias")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Well_log_dict_alias

	for rows.Next() {
		var well_log_dict_alias dto.Well_log_dict_alias
		if err := rows.Scan(&well_log_dict_alias.Dictionary_id, &well_log_dict_alias.Alias_id, &well_log_dict_alias.Abbreviation, &well_log_dict_alias.Active_ind, &well_log_dict_alias.Alias_long_name, &well_log_dict_alias.Alias_reason_type, &well_log_dict_alias.Alias_short_name, &well_log_dict_alias.Alias_type, &well_log_dict_alias.Amended_date, &well_log_dict_alias.Created_date, &well_log_dict_alias.Effective_date, &well_log_dict_alias.Expiry_date, &well_log_dict_alias.Dictionary_id, &well_log_dict_alias.Original_ind, &well_log_dict_alias.Owner_ba_id, &well_log_dict_alias.Ppdm_guid, &well_log_dict_alias.Preferred_ind, &well_log_dict_alias.Reason_desc, &well_log_dict_alias.Remark, &well_log_dict_alias.Source, &well_log_dict_alias.Source_document_id, &well_log_dict_alias.Struckoff_date, &well_log_dict_alias.Sw_application_id, &well_log_dict_alias.Use_rule, &well_log_dict_alias.Row_changed_by, &well_log_dict_alias.Row_changed_date, &well_log_dict_alias.Row_created_by, &well_log_dict_alias.Row_created_date, &well_log_dict_alias.Row_effective_date, &well_log_dict_alias.Row_expiry_date, &well_log_dict_alias.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append well_log_dict_alias to result
		result = append(result, well_log_dict_alias)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetWellLogDictAlias(c *fiber.Ctx) error {
	var well_log_dict_alias dto.Well_log_dict_alias

	setDefaults(&well_log_dict_alias)

	if err := json.Unmarshal(c.Body(), &well_log_dict_alias); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into well_log_dict_alias values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31)")
	if err != nil {
		return err
	}
	well_log_dict_alias.Row_created_date = formatDate(well_log_dict_alias.Row_created_date)
	well_log_dict_alias.Row_changed_date = formatDate(well_log_dict_alias.Row_changed_date)
	well_log_dict_alias.Amended_date = formatDateString(well_log_dict_alias.Amended_date)
	well_log_dict_alias.Created_date = formatDateString(well_log_dict_alias.Created_date)
	well_log_dict_alias.Effective_date = formatDateString(well_log_dict_alias.Effective_date)
	well_log_dict_alias.Expiry_date = formatDateString(well_log_dict_alias.Expiry_date)
	well_log_dict_alias.Struckoff_date = formatDateString(well_log_dict_alias.Struckoff_date)
	well_log_dict_alias.Row_effective_date = formatDateString(well_log_dict_alias.Row_effective_date)
	well_log_dict_alias.Row_expiry_date = formatDateString(well_log_dict_alias.Row_expiry_date)






	rows, err := stmt.Exec(well_log_dict_alias.Dictionary_id, well_log_dict_alias.Alias_id, well_log_dict_alias.Abbreviation, well_log_dict_alias.Active_ind, well_log_dict_alias.Alias_long_name, well_log_dict_alias.Alias_reason_type, well_log_dict_alias.Alias_short_name, well_log_dict_alias.Alias_type, well_log_dict_alias.Amended_date, well_log_dict_alias.Created_date, well_log_dict_alias.Effective_date, well_log_dict_alias.Expiry_date, well_log_dict_alias.Dictionary_id, well_log_dict_alias.Original_ind, well_log_dict_alias.Owner_ba_id, well_log_dict_alias.Ppdm_guid, well_log_dict_alias.Preferred_ind, well_log_dict_alias.Reason_desc, well_log_dict_alias.Remark, well_log_dict_alias.Source, well_log_dict_alias.Source_document_id, well_log_dict_alias.Struckoff_date, well_log_dict_alias.Sw_application_id, well_log_dict_alias.Use_rule, well_log_dict_alias.Row_changed_by, well_log_dict_alias.Row_changed_date, well_log_dict_alias.Row_created_by, well_log_dict_alias.Row_created_date, well_log_dict_alias.Row_effective_date, well_log_dict_alias.Row_expiry_date, well_log_dict_alias.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateWellLogDictAlias(c *fiber.Ctx) error {
	var well_log_dict_alias dto.Well_log_dict_alias

	setDefaults(&well_log_dict_alias)

	if err := json.Unmarshal(c.Body(), &well_log_dict_alias); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	well_log_dict_alias.Dictionary_id = id

    if well_log_dict_alias.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update well_log_dict_alias set alias_id = :1, abbreviation = :2, active_ind = :3, alias_long_name = :4, alias_reason_type = :5, alias_short_name = :6, alias_type = :7, amended_date = :8, created_date = :9, effective_date = :10, expiry_date = :11, dictionary_id = :12, original_ind = :13, owner_ba_id = :14, ppdm_guid = :15, preferred_ind = :16, reason_desc = :17, remark = :18, source = :19, source_document_id = :20, struckoff_date = :21, sw_application_id = :22, use_rule = :23, row_changed_by = :24, row_changed_date = :25, row_created_by = :26, row_effective_date = :27, row_expiry_date = :28, row_quality = :29 where dictionary_id = :31")
	if err != nil {
		return err
	}

	well_log_dict_alias.Row_changed_date = formatDate(well_log_dict_alias.Row_changed_date)
	well_log_dict_alias.Amended_date = formatDateString(well_log_dict_alias.Amended_date)
	well_log_dict_alias.Created_date = formatDateString(well_log_dict_alias.Created_date)
	well_log_dict_alias.Effective_date = formatDateString(well_log_dict_alias.Effective_date)
	well_log_dict_alias.Expiry_date = formatDateString(well_log_dict_alias.Expiry_date)
	well_log_dict_alias.Struckoff_date = formatDateString(well_log_dict_alias.Struckoff_date)
	well_log_dict_alias.Row_effective_date = formatDateString(well_log_dict_alias.Row_effective_date)
	well_log_dict_alias.Row_expiry_date = formatDateString(well_log_dict_alias.Row_expiry_date)






	rows, err := stmt.Exec(well_log_dict_alias.Alias_id, well_log_dict_alias.Abbreviation, well_log_dict_alias.Active_ind, well_log_dict_alias.Alias_long_name, well_log_dict_alias.Alias_reason_type, well_log_dict_alias.Alias_short_name, well_log_dict_alias.Alias_type, well_log_dict_alias.Amended_date, well_log_dict_alias.Created_date, well_log_dict_alias.Effective_date, well_log_dict_alias.Expiry_date, well_log_dict_alias.Dictionary_id, well_log_dict_alias.Original_ind, well_log_dict_alias.Owner_ba_id, well_log_dict_alias.Ppdm_guid, well_log_dict_alias.Preferred_ind, well_log_dict_alias.Reason_desc, well_log_dict_alias.Remark, well_log_dict_alias.Source, well_log_dict_alias.Source_document_id, well_log_dict_alias.Struckoff_date, well_log_dict_alias.Sw_application_id, well_log_dict_alias.Use_rule, well_log_dict_alias.Row_changed_by, well_log_dict_alias.Row_changed_date, well_log_dict_alias.Row_created_by, well_log_dict_alias.Row_effective_date, well_log_dict_alias.Row_expiry_date, well_log_dict_alias.Row_quality, well_log_dict_alias.Dictionary_id)
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

func PatchWellLogDictAlias(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update well_log_dict_alias set "
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
	query += " where dictionary_id = :id"

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

func DeleteWellLogDictAlias(c *fiber.Ctx) error {
	id := c.Params("id")
	var well_log_dict_alias dto.Well_log_dict_alias
	well_log_dict_alias.Dictionary_id = id

	stmt, err := db.Prepare("delete from well_log_dict_alias where dictionary_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(well_log_dict_alias.Dictionary_id)
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


