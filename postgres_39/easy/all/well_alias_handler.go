package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetWellAlias(c *fiber.Ctx) error {
	rows, err := db.Query("select * from well_alias")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Well_alias

	for rows.Next() {
		var well_alias dto.Well_alias
		if err := rows.Scan(&well_alias.Uwi, &well_alias.Source, &well_alias.Well_alias_id, &well_alias.Abbreviation, &well_alias.Active_ind, &well_alias.Alias_long_name, &well_alias.Alias_reason_type, &well_alias.Alias_short_name, &well_alias.Alias_type, &well_alias.Amended_date, &well_alias.Created_date, &well_alias.Effective_date, &well_alias.Expiry_date, &well_alias.Uwi, &well_alias.Location_type, &well_alias.Original_ind, &well_alias.Owner_ba_id, &well_alias.Ppdm_guid, &well_alias.Preferred_ind, &well_alias.Reason_desc, &well_alias.Remark, &well_alias.Source_document_id, &well_alias.Struckoff_date, &well_alias.Sw_application_id, &well_alias.Use_rule, &well_alias.Row_changed_by, &well_alias.Row_changed_date, &well_alias.Row_created_by, &well_alias.Row_created_date, &well_alias.Row_effective_date, &well_alias.Row_expiry_date, &well_alias.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append well_alias to result
		result = append(result, well_alias)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetWellAlias(c *fiber.Ctx) error {
	var well_alias dto.Well_alias

	setDefaults(&well_alias)

	if err := json.Unmarshal(c.Body(), &well_alias); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into well_alias values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32)")
	if err != nil {
		return err
	}
	well_alias.Row_created_date = formatDate(well_alias.Row_created_date)
	well_alias.Row_changed_date = formatDate(well_alias.Row_changed_date)
	well_alias.Amended_date = formatDateString(well_alias.Amended_date)
	well_alias.Created_date = formatDateString(well_alias.Created_date)
	well_alias.Effective_date = formatDateString(well_alias.Effective_date)
	well_alias.Expiry_date = formatDateString(well_alias.Expiry_date)
	well_alias.Struckoff_date = formatDateString(well_alias.Struckoff_date)
	well_alias.Row_effective_date = formatDateString(well_alias.Row_effective_date)
	well_alias.Row_expiry_date = formatDateString(well_alias.Row_expiry_date)






	rows, err := stmt.Exec(well_alias.Uwi, well_alias.Source, well_alias.Well_alias_id, well_alias.Abbreviation, well_alias.Active_ind, well_alias.Alias_long_name, well_alias.Alias_reason_type, well_alias.Alias_short_name, well_alias.Alias_type, well_alias.Amended_date, well_alias.Created_date, well_alias.Effective_date, well_alias.Expiry_date, well_alias.Uwi, well_alias.Location_type, well_alias.Original_ind, well_alias.Owner_ba_id, well_alias.Ppdm_guid, well_alias.Preferred_ind, well_alias.Reason_desc, well_alias.Remark, well_alias.Source_document_id, well_alias.Struckoff_date, well_alias.Sw_application_id, well_alias.Use_rule, well_alias.Row_changed_by, well_alias.Row_changed_date, well_alias.Row_created_by, well_alias.Row_created_date, well_alias.Row_effective_date, well_alias.Row_expiry_date, well_alias.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateWellAlias(c *fiber.Ctx) error {
	var well_alias dto.Well_alias

	setDefaults(&well_alias)

	if err := json.Unmarshal(c.Body(), &well_alias); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	well_alias.Uwi = id

    if well_alias.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update well_alias set source = :1, well_alias_id = :2, abbreviation = :3, active_ind = :4, alias_long_name = :5, alias_reason_type = :6, alias_short_name = :7, alias_type = :8, amended_date = :9, created_date = :10, effective_date = :11, expiry_date = :12, uwi = :13, location_type = :14, original_ind = :15, owner_ba_id = :16, ppdm_guid = :17, preferred_ind = :18, reason_desc = :19, remark = :20, source_document_id = :21, struckoff_date = :22, sw_application_id = :23, use_rule = :24, row_changed_by = :25, row_changed_date = :26, row_created_by = :27, row_effective_date = :28, row_expiry_date = :29, row_quality = :30 where uwi = :32")
	if err != nil {
		return err
	}

	well_alias.Row_changed_date = formatDate(well_alias.Row_changed_date)
	well_alias.Amended_date = formatDateString(well_alias.Amended_date)
	well_alias.Created_date = formatDateString(well_alias.Created_date)
	well_alias.Effective_date = formatDateString(well_alias.Effective_date)
	well_alias.Expiry_date = formatDateString(well_alias.Expiry_date)
	well_alias.Struckoff_date = formatDateString(well_alias.Struckoff_date)
	well_alias.Row_effective_date = formatDateString(well_alias.Row_effective_date)
	well_alias.Row_expiry_date = formatDateString(well_alias.Row_expiry_date)






	rows, err := stmt.Exec(well_alias.Source, well_alias.Well_alias_id, well_alias.Abbreviation, well_alias.Active_ind, well_alias.Alias_long_name, well_alias.Alias_reason_type, well_alias.Alias_short_name, well_alias.Alias_type, well_alias.Amended_date, well_alias.Created_date, well_alias.Effective_date, well_alias.Expiry_date, well_alias.Uwi, well_alias.Location_type, well_alias.Original_ind, well_alias.Owner_ba_id, well_alias.Ppdm_guid, well_alias.Preferred_ind, well_alias.Reason_desc, well_alias.Remark, well_alias.Source_document_id, well_alias.Struckoff_date, well_alias.Sw_application_id, well_alias.Use_rule, well_alias.Row_changed_by, well_alias.Row_changed_date, well_alias.Row_created_by, well_alias.Row_effective_date, well_alias.Row_expiry_date, well_alias.Row_quality, well_alias.Uwi)
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

func PatchWellAlias(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update well_alias set "
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
	query += " where uwi = :id"

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

func DeleteWellAlias(c *fiber.Ctx) error {
	id := c.Params("id")
	var well_alias dto.Well_alias
	well_alias.Uwi = id

	stmt, err := db.Prepare("delete from well_alias where uwi = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(well_alias.Uwi)
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


