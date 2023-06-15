package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetApplicAlias(c *fiber.Ctx) error {
	rows, err := db.Query("select * from applic_alias")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Applic_alias

	for rows.Next() {
		var applic_alias dto.Applic_alias
		if err := rows.Scan(&applic_alias.Application_id, &applic_alias.Alias_id, &applic_alias.Abbreviation, &applic_alias.Active_ind, &applic_alias.Alias_long_name, &applic_alias.Alias_reason_type, &applic_alias.Alias_short_name, &applic_alias.Alias_type, &applic_alias.Amended_date, &applic_alias.Created_date, &applic_alias.Effective_date, &applic_alias.Expiry_date, &applic_alias.Application_id, &applic_alias.Original_ind, &applic_alias.Owner_ba_id, &applic_alias.Ppdm_guid, &applic_alias.Preferred_ind, &applic_alias.Reason_desc, &applic_alias.Remark, &applic_alias.Source, &applic_alias.Source_document_id, &applic_alias.Struckoff_date, &applic_alias.Sw_application_id, &applic_alias.Use_rule, &applic_alias.Row_changed_by, &applic_alias.Row_changed_date, &applic_alias.Row_created_by, &applic_alias.Row_created_date, &applic_alias.Row_effective_date, &applic_alias.Row_expiry_date, &applic_alias.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append applic_alias to result
		result = append(result, applic_alias)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetApplicAlias(c *fiber.Ctx) error {
	var applic_alias dto.Applic_alias

	setDefaults(&applic_alias)

	if err := json.Unmarshal(c.Body(), &applic_alias); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into applic_alias values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31)")
	if err != nil {
		return err
	}
	applic_alias.Row_created_date = formatDate(applic_alias.Row_created_date)
	applic_alias.Row_changed_date = formatDate(applic_alias.Row_changed_date)
	applic_alias.Amended_date = formatDateString(applic_alias.Amended_date)
	applic_alias.Created_date = formatDateString(applic_alias.Created_date)
	applic_alias.Effective_date = formatDateString(applic_alias.Effective_date)
	applic_alias.Expiry_date = formatDateString(applic_alias.Expiry_date)
	applic_alias.Struckoff_date = formatDateString(applic_alias.Struckoff_date)
	applic_alias.Row_effective_date = formatDateString(applic_alias.Row_effective_date)
	applic_alias.Row_expiry_date = formatDateString(applic_alias.Row_expiry_date)






	rows, err := stmt.Exec(applic_alias.Application_id, applic_alias.Alias_id, applic_alias.Abbreviation, applic_alias.Active_ind, applic_alias.Alias_long_name, applic_alias.Alias_reason_type, applic_alias.Alias_short_name, applic_alias.Alias_type, applic_alias.Amended_date, applic_alias.Created_date, applic_alias.Effective_date, applic_alias.Expiry_date, applic_alias.Application_id, applic_alias.Original_ind, applic_alias.Owner_ba_id, applic_alias.Ppdm_guid, applic_alias.Preferred_ind, applic_alias.Reason_desc, applic_alias.Remark, applic_alias.Source, applic_alias.Source_document_id, applic_alias.Struckoff_date, applic_alias.Sw_application_id, applic_alias.Use_rule, applic_alias.Row_changed_by, applic_alias.Row_changed_date, applic_alias.Row_created_by, applic_alias.Row_created_date, applic_alias.Row_effective_date, applic_alias.Row_expiry_date, applic_alias.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateApplicAlias(c *fiber.Ctx) error {
	var applic_alias dto.Applic_alias

	setDefaults(&applic_alias)

	if err := json.Unmarshal(c.Body(), &applic_alias); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	applic_alias.Application_id = id

    if applic_alias.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update applic_alias set alias_id = :1, abbreviation = :2, active_ind = :3, alias_long_name = :4, alias_reason_type = :5, alias_short_name = :6, alias_type = :7, amended_date = :8, created_date = :9, effective_date = :10, expiry_date = :11, application_id = :12, original_ind = :13, owner_ba_id = :14, ppdm_guid = :15, preferred_ind = :16, reason_desc = :17, remark = :18, source = :19, source_document_id = :20, struckoff_date = :21, sw_application_id = :22, use_rule = :23, row_changed_by = :24, row_changed_date = :25, row_created_by = :26, row_effective_date = :27, row_expiry_date = :28, row_quality = :29 where application_id = :31")
	if err != nil {
		return err
	}

	applic_alias.Row_changed_date = formatDate(applic_alias.Row_changed_date)
	applic_alias.Amended_date = formatDateString(applic_alias.Amended_date)
	applic_alias.Created_date = formatDateString(applic_alias.Created_date)
	applic_alias.Effective_date = formatDateString(applic_alias.Effective_date)
	applic_alias.Expiry_date = formatDateString(applic_alias.Expiry_date)
	applic_alias.Struckoff_date = formatDateString(applic_alias.Struckoff_date)
	applic_alias.Row_effective_date = formatDateString(applic_alias.Row_effective_date)
	applic_alias.Row_expiry_date = formatDateString(applic_alias.Row_expiry_date)






	rows, err := stmt.Exec(applic_alias.Alias_id, applic_alias.Abbreviation, applic_alias.Active_ind, applic_alias.Alias_long_name, applic_alias.Alias_reason_type, applic_alias.Alias_short_name, applic_alias.Alias_type, applic_alias.Amended_date, applic_alias.Created_date, applic_alias.Effective_date, applic_alias.Expiry_date, applic_alias.Application_id, applic_alias.Original_ind, applic_alias.Owner_ba_id, applic_alias.Ppdm_guid, applic_alias.Preferred_ind, applic_alias.Reason_desc, applic_alias.Remark, applic_alias.Source, applic_alias.Source_document_id, applic_alias.Struckoff_date, applic_alias.Sw_application_id, applic_alias.Use_rule, applic_alias.Row_changed_by, applic_alias.Row_changed_date, applic_alias.Row_created_by, applic_alias.Row_effective_date, applic_alias.Row_expiry_date, applic_alias.Row_quality, applic_alias.Application_id)
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

func PatchApplicAlias(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update applic_alias set "
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
	query += " where application_id = :id"

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

func DeleteApplicAlias(c *fiber.Ctx) error {
	id := c.Params("id")
	var applic_alias dto.Applic_alias
	applic_alias.Application_id = id

	stmt, err := db.Prepare("delete from applic_alias where application_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(applic_alias.Application_id)
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


