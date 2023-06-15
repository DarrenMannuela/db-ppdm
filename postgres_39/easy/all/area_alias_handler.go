package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetAreaAlias(c *fiber.Ctx) error {
	rows, err := db.Query("select * from area_alias")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Area_alias

	for rows.Next() {
		var area_alias dto.Area_alias
		if err := rows.Scan(&area_alias.Area_id, &area_alias.Area_type, &area_alias.Area_alias_id, &area_alias.Abbreviation, &area_alias.Active_ind, &area_alias.Alias_long_name, &area_alias.Alias_reason_type, &area_alias.Alias_short_name, &area_alias.Alias_type, &area_alias.Amended_date, &area_alias.Created_date, &area_alias.Effective_date, &area_alias.Expiry_date, &area_alias.Area_id, &area_alias.Original_ind, &area_alias.Owner_ba_id, &area_alias.Ppdm_guid, &area_alias.Preferred_ind, &area_alias.Reason_desc, &area_alias.Remark, &area_alias.Source, &area_alias.Source_document_id, &area_alias.Struckoff_date, &area_alias.Sw_application_id, &area_alias.Use_rule, &area_alias.Row_changed_by, &area_alias.Row_changed_date, &area_alias.Row_created_by, &area_alias.Row_created_date, &area_alias.Row_effective_date, &area_alias.Row_expiry_date, &area_alias.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append area_alias to result
		result = append(result, area_alias)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetAreaAlias(c *fiber.Ctx) error {
	var area_alias dto.Area_alias

	setDefaults(&area_alias)

	if err := json.Unmarshal(c.Body(), &area_alias); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into area_alias values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32)")
	if err != nil {
		return err
	}
	area_alias.Row_created_date = formatDate(area_alias.Row_created_date)
	area_alias.Row_changed_date = formatDate(area_alias.Row_changed_date)
	area_alias.Amended_date = formatDateString(area_alias.Amended_date)
	area_alias.Created_date = formatDateString(area_alias.Created_date)
	area_alias.Effective_date = formatDateString(area_alias.Effective_date)
	area_alias.Expiry_date = formatDateString(area_alias.Expiry_date)
	area_alias.Struckoff_date = formatDateString(area_alias.Struckoff_date)
	area_alias.Row_effective_date = formatDateString(area_alias.Row_effective_date)
	area_alias.Row_expiry_date = formatDateString(area_alias.Row_expiry_date)






	rows, err := stmt.Exec(area_alias.Area_id, area_alias.Area_type, area_alias.Area_alias_id, area_alias.Abbreviation, area_alias.Active_ind, area_alias.Alias_long_name, area_alias.Alias_reason_type, area_alias.Alias_short_name, area_alias.Alias_type, area_alias.Amended_date, area_alias.Created_date, area_alias.Effective_date, area_alias.Expiry_date, area_alias.Area_id, area_alias.Original_ind, area_alias.Owner_ba_id, area_alias.Ppdm_guid, area_alias.Preferred_ind, area_alias.Reason_desc, area_alias.Remark, area_alias.Source, area_alias.Source_document_id, area_alias.Struckoff_date, area_alias.Sw_application_id, area_alias.Use_rule, area_alias.Row_changed_by, area_alias.Row_changed_date, area_alias.Row_created_by, area_alias.Row_created_date, area_alias.Row_effective_date, area_alias.Row_expiry_date, area_alias.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateAreaAlias(c *fiber.Ctx) error {
	var area_alias dto.Area_alias

	setDefaults(&area_alias)

	if err := json.Unmarshal(c.Body(), &area_alias); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	area_alias.Area_id = id

    if area_alias.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update area_alias set area_type = :1, area_alias_id = :2, abbreviation = :3, active_ind = :4, alias_long_name = :5, alias_reason_type = :6, alias_short_name = :7, alias_type = :8, amended_date = :9, created_date = :10, effective_date = :11, expiry_date = :12, area_id = :13, original_ind = :14, owner_ba_id = :15, ppdm_guid = :16, preferred_ind = :17, reason_desc = :18, remark = :19, source = :20, source_document_id = :21, struckoff_date = :22, sw_application_id = :23, use_rule = :24, row_changed_by = :25, row_changed_date = :26, row_created_by = :27, row_effective_date = :28, row_expiry_date = :29, row_quality = :30 where area_id = :32")
	if err != nil {
		return err
	}

	area_alias.Row_changed_date = formatDate(area_alias.Row_changed_date)
	area_alias.Amended_date = formatDateString(area_alias.Amended_date)
	area_alias.Created_date = formatDateString(area_alias.Created_date)
	area_alias.Effective_date = formatDateString(area_alias.Effective_date)
	area_alias.Expiry_date = formatDateString(area_alias.Expiry_date)
	area_alias.Struckoff_date = formatDateString(area_alias.Struckoff_date)
	area_alias.Row_effective_date = formatDateString(area_alias.Row_effective_date)
	area_alias.Row_expiry_date = formatDateString(area_alias.Row_expiry_date)






	rows, err := stmt.Exec(area_alias.Area_type, area_alias.Area_alias_id, area_alias.Abbreviation, area_alias.Active_ind, area_alias.Alias_long_name, area_alias.Alias_reason_type, area_alias.Alias_short_name, area_alias.Alias_type, area_alias.Amended_date, area_alias.Created_date, area_alias.Effective_date, area_alias.Expiry_date, area_alias.Area_id, area_alias.Original_ind, area_alias.Owner_ba_id, area_alias.Ppdm_guid, area_alias.Preferred_ind, area_alias.Reason_desc, area_alias.Remark, area_alias.Source, area_alias.Source_document_id, area_alias.Struckoff_date, area_alias.Sw_application_id, area_alias.Use_rule, area_alias.Row_changed_by, area_alias.Row_changed_date, area_alias.Row_created_by, area_alias.Row_effective_date, area_alias.Row_expiry_date, area_alias.Row_quality, area_alias.Area_id)
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

func PatchAreaAlias(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update area_alias set "
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
	query += " where area_id = :id"

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

func DeleteAreaAlias(c *fiber.Ctx) error {
	id := c.Params("id")
	var area_alias dto.Area_alias
	area_alias.Area_id = id

	stmt, err := db.Prepare("delete from area_alias where area_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(area_alias.Area_id)
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


