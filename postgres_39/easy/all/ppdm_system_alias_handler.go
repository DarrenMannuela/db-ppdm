package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetPpdmSystemAlias(c *fiber.Ctx) error {
	rows, err := db.Query("select * from ppdm_system_alias")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Ppdm_system_alias

	for rows.Next() {
		var ppdm_system_alias dto.Ppdm_system_alias
		if err := rows.Scan(&ppdm_system_alias.System_id, &ppdm_system_alias.Alias_id, &ppdm_system_alias.Abbreviation, &ppdm_system_alias.Active_ind, &ppdm_system_alias.Alias_long_name, &ppdm_system_alias.Alias_reason_type, &ppdm_system_alias.Alias_short_name, &ppdm_system_alias.Alias_type, &ppdm_system_alias.Amended_date, &ppdm_system_alias.Created_date, &ppdm_system_alias.Effective_date, &ppdm_system_alias.Expiry_date, &ppdm_system_alias.System_id, &ppdm_system_alias.Original_ind, &ppdm_system_alias.Owner_ba_id, &ppdm_system_alias.Ppdm_guid, &ppdm_system_alias.Preferred_ind, &ppdm_system_alias.Reason_desc, &ppdm_system_alias.Remark, &ppdm_system_alias.Source, &ppdm_system_alias.Source_document_id, &ppdm_system_alias.Struckoff_date, &ppdm_system_alias.Sw_application_id, &ppdm_system_alias.Use_rule, &ppdm_system_alias.Row_changed_by, &ppdm_system_alias.Row_changed_date, &ppdm_system_alias.Row_created_by, &ppdm_system_alias.Row_created_date, &ppdm_system_alias.Row_effective_date, &ppdm_system_alias.Row_expiry_date, &ppdm_system_alias.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append ppdm_system_alias to result
		result = append(result, ppdm_system_alias)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetPpdmSystemAlias(c *fiber.Ctx) error {
	var ppdm_system_alias dto.Ppdm_system_alias

	setDefaults(&ppdm_system_alias)

	if err := json.Unmarshal(c.Body(), &ppdm_system_alias); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into ppdm_system_alias values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31)")
	if err != nil {
		return err
	}
	ppdm_system_alias.Row_created_date = formatDate(ppdm_system_alias.Row_created_date)
	ppdm_system_alias.Row_changed_date = formatDate(ppdm_system_alias.Row_changed_date)
	ppdm_system_alias.Amended_date = formatDateString(ppdm_system_alias.Amended_date)
	ppdm_system_alias.Created_date = formatDateString(ppdm_system_alias.Created_date)
	ppdm_system_alias.Effective_date = formatDateString(ppdm_system_alias.Effective_date)
	ppdm_system_alias.Expiry_date = formatDateString(ppdm_system_alias.Expiry_date)
	ppdm_system_alias.Struckoff_date = formatDateString(ppdm_system_alias.Struckoff_date)
	ppdm_system_alias.Row_effective_date = formatDateString(ppdm_system_alias.Row_effective_date)
	ppdm_system_alias.Row_expiry_date = formatDateString(ppdm_system_alias.Row_expiry_date)






	rows, err := stmt.Exec(ppdm_system_alias.System_id, ppdm_system_alias.Alias_id, ppdm_system_alias.Abbreviation, ppdm_system_alias.Active_ind, ppdm_system_alias.Alias_long_name, ppdm_system_alias.Alias_reason_type, ppdm_system_alias.Alias_short_name, ppdm_system_alias.Alias_type, ppdm_system_alias.Amended_date, ppdm_system_alias.Created_date, ppdm_system_alias.Effective_date, ppdm_system_alias.Expiry_date, ppdm_system_alias.System_id, ppdm_system_alias.Original_ind, ppdm_system_alias.Owner_ba_id, ppdm_system_alias.Ppdm_guid, ppdm_system_alias.Preferred_ind, ppdm_system_alias.Reason_desc, ppdm_system_alias.Remark, ppdm_system_alias.Source, ppdm_system_alias.Source_document_id, ppdm_system_alias.Struckoff_date, ppdm_system_alias.Sw_application_id, ppdm_system_alias.Use_rule, ppdm_system_alias.Row_changed_by, ppdm_system_alias.Row_changed_date, ppdm_system_alias.Row_created_by, ppdm_system_alias.Row_created_date, ppdm_system_alias.Row_effective_date, ppdm_system_alias.Row_expiry_date, ppdm_system_alias.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdatePpdmSystemAlias(c *fiber.Ctx) error {
	var ppdm_system_alias dto.Ppdm_system_alias

	setDefaults(&ppdm_system_alias)

	if err := json.Unmarshal(c.Body(), &ppdm_system_alias); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	ppdm_system_alias.System_id = id

    if ppdm_system_alias.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update ppdm_system_alias set alias_id = :1, abbreviation = :2, active_ind = :3, alias_long_name = :4, alias_reason_type = :5, alias_short_name = :6, alias_type = :7, amended_date = :8, created_date = :9, effective_date = :10, expiry_date = :11, system_id = :12, original_ind = :13, owner_ba_id = :14, ppdm_guid = :15, preferred_ind = :16, reason_desc = :17, remark = :18, source = :19, source_document_id = :20, struckoff_date = :21, sw_application_id = :22, use_rule = :23, row_changed_by = :24, row_changed_date = :25, row_created_by = :26, row_effective_date = :27, row_expiry_date = :28, row_quality = :29 where system_id = :31")
	if err != nil {
		return err
	}

	ppdm_system_alias.Row_changed_date = formatDate(ppdm_system_alias.Row_changed_date)
	ppdm_system_alias.Amended_date = formatDateString(ppdm_system_alias.Amended_date)
	ppdm_system_alias.Created_date = formatDateString(ppdm_system_alias.Created_date)
	ppdm_system_alias.Effective_date = formatDateString(ppdm_system_alias.Effective_date)
	ppdm_system_alias.Expiry_date = formatDateString(ppdm_system_alias.Expiry_date)
	ppdm_system_alias.Struckoff_date = formatDateString(ppdm_system_alias.Struckoff_date)
	ppdm_system_alias.Row_effective_date = formatDateString(ppdm_system_alias.Row_effective_date)
	ppdm_system_alias.Row_expiry_date = formatDateString(ppdm_system_alias.Row_expiry_date)






	rows, err := stmt.Exec(ppdm_system_alias.Alias_id, ppdm_system_alias.Abbreviation, ppdm_system_alias.Active_ind, ppdm_system_alias.Alias_long_name, ppdm_system_alias.Alias_reason_type, ppdm_system_alias.Alias_short_name, ppdm_system_alias.Alias_type, ppdm_system_alias.Amended_date, ppdm_system_alias.Created_date, ppdm_system_alias.Effective_date, ppdm_system_alias.Expiry_date, ppdm_system_alias.System_id, ppdm_system_alias.Original_ind, ppdm_system_alias.Owner_ba_id, ppdm_system_alias.Ppdm_guid, ppdm_system_alias.Preferred_ind, ppdm_system_alias.Reason_desc, ppdm_system_alias.Remark, ppdm_system_alias.Source, ppdm_system_alias.Source_document_id, ppdm_system_alias.Struckoff_date, ppdm_system_alias.Sw_application_id, ppdm_system_alias.Use_rule, ppdm_system_alias.Row_changed_by, ppdm_system_alias.Row_changed_date, ppdm_system_alias.Row_created_by, ppdm_system_alias.Row_effective_date, ppdm_system_alias.Row_expiry_date, ppdm_system_alias.Row_quality, ppdm_system_alias.System_id)
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

func PatchPpdmSystemAlias(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update ppdm_system_alias set "
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
	query += " where system_id = :id"

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

func DeletePpdmSystemAlias(c *fiber.Ctx) error {
	id := c.Params("id")
	var ppdm_system_alias dto.Ppdm_system_alias
	ppdm_system_alias.System_id = id

	stmt, err := db.Prepare("delete from ppdm_system_alias where system_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(ppdm_system_alias.System_id)
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


