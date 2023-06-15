package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetStratAlias(c *fiber.Ctx) error {
	rows, err := db.Query("select * from strat_alias")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Strat_alias

	for rows.Next() {
		var strat_alias dto.Strat_alias
		if err := rows.Scan(&strat_alias.Strat_name_set_id, &strat_alias.Strat_unit_id, &strat_alias.Alias_strat_name_set_id, &strat_alias.Alias_strat_unit_id, &strat_alias.Source, &strat_alias.Abbreviation, &strat_alias.Active_ind, &strat_alias.Alias_long_name, &strat_alias.Alias_reason_type, &strat_alias.Alias_short_name, &strat_alias.Alias_type, &strat_alias.Amended_date, &strat_alias.Created_date, &strat_alias.Effective_date, &strat_alias.Expiry_date, &strat_alias.Strat_name_set_id, &strat_alias.Original_ind, &strat_alias.Owner_ba_id, &strat_alias.Ppdm_guid, &strat_alias.Preferred_ind, &strat_alias.Reason_desc, &strat_alias.Remark, &strat_alias.Source_document_id, &strat_alias.Struckoff_date, &strat_alias.Sw_application_id, &strat_alias.Use_rule, &strat_alias.Row_changed_by, &strat_alias.Row_changed_date, &strat_alias.Row_created_by, &strat_alias.Row_created_date, &strat_alias.Row_effective_date, &strat_alias.Row_expiry_date, &strat_alias.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append strat_alias to result
		result = append(result, strat_alias)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetStratAlias(c *fiber.Ctx) error {
	var strat_alias dto.Strat_alias

	setDefaults(&strat_alias)

	if err := json.Unmarshal(c.Body(), &strat_alias); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into strat_alias values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32, :33)")
	if err != nil {
		return err
	}
	strat_alias.Row_created_date = formatDate(strat_alias.Row_created_date)
	strat_alias.Row_changed_date = formatDate(strat_alias.Row_changed_date)
	strat_alias.Amended_date = formatDateString(strat_alias.Amended_date)
	strat_alias.Created_date = formatDateString(strat_alias.Created_date)
	strat_alias.Effective_date = formatDateString(strat_alias.Effective_date)
	strat_alias.Expiry_date = formatDateString(strat_alias.Expiry_date)
	strat_alias.Struckoff_date = formatDateString(strat_alias.Struckoff_date)
	strat_alias.Row_effective_date = formatDateString(strat_alias.Row_effective_date)
	strat_alias.Row_expiry_date = formatDateString(strat_alias.Row_expiry_date)






	rows, err := stmt.Exec(strat_alias.Strat_name_set_id, strat_alias.Strat_unit_id, strat_alias.Alias_strat_name_set_id, strat_alias.Alias_strat_unit_id, strat_alias.Source, strat_alias.Abbreviation, strat_alias.Active_ind, strat_alias.Alias_long_name, strat_alias.Alias_reason_type, strat_alias.Alias_short_name, strat_alias.Alias_type, strat_alias.Amended_date, strat_alias.Created_date, strat_alias.Effective_date, strat_alias.Expiry_date, strat_alias.Strat_name_set_id, strat_alias.Original_ind, strat_alias.Owner_ba_id, strat_alias.Ppdm_guid, strat_alias.Preferred_ind, strat_alias.Reason_desc, strat_alias.Remark, strat_alias.Source_document_id, strat_alias.Struckoff_date, strat_alias.Sw_application_id, strat_alias.Use_rule, strat_alias.Row_changed_by, strat_alias.Row_changed_date, strat_alias.Row_created_by, strat_alias.Row_created_date, strat_alias.Row_effective_date, strat_alias.Row_expiry_date, strat_alias.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateStratAlias(c *fiber.Ctx) error {
	var strat_alias dto.Strat_alias

	setDefaults(&strat_alias)

	if err := json.Unmarshal(c.Body(), &strat_alias); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	strat_alias.Strat_name_set_id = id

    if strat_alias.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update strat_alias set strat_unit_id = :1, alias_strat_name_set_id = :2, alias_strat_unit_id = :3, source = :4, abbreviation = :5, active_ind = :6, alias_long_name = :7, alias_reason_type = :8, alias_short_name = :9, alias_type = :10, amended_date = :11, created_date = :12, effective_date = :13, expiry_date = :14, strat_name_set_id = :15, original_ind = :16, owner_ba_id = :17, ppdm_guid = :18, preferred_ind = :19, reason_desc = :20, remark = :21, source_document_id = :22, struckoff_date = :23, sw_application_id = :24, use_rule = :25, row_changed_by = :26, row_changed_date = :27, row_created_by = :28, row_effective_date = :29, row_expiry_date = :30, row_quality = :31 where strat_name_set_id = :33")
	if err != nil {
		return err
	}

	strat_alias.Row_changed_date = formatDate(strat_alias.Row_changed_date)
	strat_alias.Amended_date = formatDateString(strat_alias.Amended_date)
	strat_alias.Created_date = formatDateString(strat_alias.Created_date)
	strat_alias.Effective_date = formatDateString(strat_alias.Effective_date)
	strat_alias.Expiry_date = formatDateString(strat_alias.Expiry_date)
	strat_alias.Struckoff_date = formatDateString(strat_alias.Struckoff_date)
	strat_alias.Row_effective_date = formatDateString(strat_alias.Row_effective_date)
	strat_alias.Row_expiry_date = formatDateString(strat_alias.Row_expiry_date)






	rows, err := stmt.Exec(strat_alias.Strat_unit_id, strat_alias.Alias_strat_name_set_id, strat_alias.Alias_strat_unit_id, strat_alias.Source, strat_alias.Abbreviation, strat_alias.Active_ind, strat_alias.Alias_long_name, strat_alias.Alias_reason_type, strat_alias.Alias_short_name, strat_alias.Alias_type, strat_alias.Amended_date, strat_alias.Created_date, strat_alias.Effective_date, strat_alias.Expiry_date, strat_alias.Strat_name_set_id, strat_alias.Original_ind, strat_alias.Owner_ba_id, strat_alias.Ppdm_guid, strat_alias.Preferred_ind, strat_alias.Reason_desc, strat_alias.Remark, strat_alias.Source_document_id, strat_alias.Struckoff_date, strat_alias.Sw_application_id, strat_alias.Use_rule, strat_alias.Row_changed_by, strat_alias.Row_changed_date, strat_alias.Row_created_by, strat_alias.Row_effective_date, strat_alias.Row_expiry_date, strat_alias.Row_quality, strat_alias.Strat_name_set_id)
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

func PatchStratAlias(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update strat_alias set "
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
	query += " where strat_name_set_id = :id"

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

func DeleteStratAlias(c *fiber.Ctx) error {
	id := c.Params("id")
	var strat_alias dto.Strat_alias
	strat_alias.Strat_name_set_id = id

	stmt, err := db.Prepare("delete from strat_alias where strat_name_set_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(strat_alias.Strat_name_set_id)
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


