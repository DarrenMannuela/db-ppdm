package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetPpdmTableAlias(c *fiber.Ctx) error {
	rows, err := db.Query("select * from ppdm_table_alias")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Ppdm_table_alias

	for rows.Next() {
		var ppdm_table_alias dto.Ppdm_table_alias
		if err := rows.Scan(&ppdm_table_alias.System_id, &ppdm_table_alias.Table_name, &ppdm_table_alias.Table_alias, &ppdm_table_alias.Abbreviation, &ppdm_table_alias.Active_ind, &ppdm_table_alias.Alias_long_name, &ppdm_table_alias.Alias_reason_type, &ppdm_table_alias.Alias_short_name, &ppdm_table_alias.Alias_type, &ppdm_table_alias.Amended_date, &ppdm_table_alias.Created_date, &ppdm_table_alias.Effective_date, &ppdm_table_alias.Expiry_date, &ppdm_table_alias.System_id, &ppdm_table_alias.Original_ind, &ppdm_table_alias.Owner_ba_id, &ppdm_table_alias.Ppdm_guid, &ppdm_table_alias.Preferred_ind, &ppdm_table_alias.Reason_desc, &ppdm_table_alias.Remark, &ppdm_table_alias.Source, &ppdm_table_alias.Source_document_id, &ppdm_table_alias.Struckoff_date, &ppdm_table_alias.Sw_application_id, &ppdm_table_alias.Use_rule, &ppdm_table_alias.Row_changed_by, &ppdm_table_alias.Row_changed_date, &ppdm_table_alias.Row_created_by, &ppdm_table_alias.Row_created_date, &ppdm_table_alias.Row_effective_date, &ppdm_table_alias.Row_expiry_date, &ppdm_table_alias.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append ppdm_table_alias to result
		result = append(result, ppdm_table_alias)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetPpdmTableAlias(c *fiber.Ctx) error {
	var ppdm_table_alias dto.Ppdm_table_alias

	setDefaults(&ppdm_table_alias)

	if err := json.Unmarshal(c.Body(), &ppdm_table_alias); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into ppdm_table_alias values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32)")
	if err != nil {
		return err
	}
	ppdm_table_alias.Row_created_date = formatDate(ppdm_table_alias.Row_created_date)
	ppdm_table_alias.Row_changed_date = formatDate(ppdm_table_alias.Row_changed_date)
	ppdm_table_alias.Amended_date = formatDateString(ppdm_table_alias.Amended_date)
	ppdm_table_alias.Created_date = formatDateString(ppdm_table_alias.Created_date)
	ppdm_table_alias.Effective_date = formatDateString(ppdm_table_alias.Effective_date)
	ppdm_table_alias.Expiry_date = formatDateString(ppdm_table_alias.Expiry_date)
	ppdm_table_alias.Struckoff_date = formatDateString(ppdm_table_alias.Struckoff_date)
	ppdm_table_alias.Row_effective_date = formatDateString(ppdm_table_alias.Row_effective_date)
	ppdm_table_alias.Row_expiry_date = formatDateString(ppdm_table_alias.Row_expiry_date)






	rows, err := stmt.Exec(ppdm_table_alias.System_id, ppdm_table_alias.Table_name, ppdm_table_alias.Table_alias, ppdm_table_alias.Abbreviation, ppdm_table_alias.Active_ind, ppdm_table_alias.Alias_long_name, ppdm_table_alias.Alias_reason_type, ppdm_table_alias.Alias_short_name, ppdm_table_alias.Alias_type, ppdm_table_alias.Amended_date, ppdm_table_alias.Created_date, ppdm_table_alias.Effective_date, ppdm_table_alias.Expiry_date, ppdm_table_alias.System_id, ppdm_table_alias.Original_ind, ppdm_table_alias.Owner_ba_id, ppdm_table_alias.Ppdm_guid, ppdm_table_alias.Preferred_ind, ppdm_table_alias.Reason_desc, ppdm_table_alias.Remark, ppdm_table_alias.Source, ppdm_table_alias.Source_document_id, ppdm_table_alias.Struckoff_date, ppdm_table_alias.Sw_application_id, ppdm_table_alias.Use_rule, ppdm_table_alias.Row_changed_by, ppdm_table_alias.Row_changed_date, ppdm_table_alias.Row_created_by, ppdm_table_alias.Row_created_date, ppdm_table_alias.Row_effective_date, ppdm_table_alias.Row_expiry_date, ppdm_table_alias.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdatePpdmTableAlias(c *fiber.Ctx) error {
	var ppdm_table_alias dto.Ppdm_table_alias

	setDefaults(&ppdm_table_alias)

	if err := json.Unmarshal(c.Body(), &ppdm_table_alias); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	ppdm_table_alias.System_id = id

    if ppdm_table_alias.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update ppdm_table_alias set table_name = :1, table_alias = :2, abbreviation = :3, active_ind = :4, alias_long_name = :5, alias_reason_type = :6, alias_short_name = :7, alias_type = :8, amended_date = :9, created_date = :10, effective_date = :11, expiry_date = :12, system_id = :13, original_ind = :14, owner_ba_id = :15, ppdm_guid = :16, preferred_ind = :17, reason_desc = :18, remark = :19, source = :20, source_document_id = :21, struckoff_date = :22, sw_application_id = :23, use_rule = :24, row_changed_by = :25, row_changed_date = :26, row_created_by = :27, row_effective_date = :28, row_expiry_date = :29, row_quality = :30 where system_id = :32")
	if err != nil {
		return err
	}

	ppdm_table_alias.Row_changed_date = formatDate(ppdm_table_alias.Row_changed_date)
	ppdm_table_alias.Amended_date = formatDateString(ppdm_table_alias.Amended_date)
	ppdm_table_alias.Created_date = formatDateString(ppdm_table_alias.Created_date)
	ppdm_table_alias.Effective_date = formatDateString(ppdm_table_alias.Effective_date)
	ppdm_table_alias.Expiry_date = formatDateString(ppdm_table_alias.Expiry_date)
	ppdm_table_alias.Struckoff_date = formatDateString(ppdm_table_alias.Struckoff_date)
	ppdm_table_alias.Row_effective_date = formatDateString(ppdm_table_alias.Row_effective_date)
	ppdm_table_alias.Row_expiry_date = formatDateString(ppdm_table_alias.Row_expiry_date)






	rows, err := stmt.Exec(ppdm_table_alias.Table_name, ppdm_table_alias.Table_alias, ppdm_table_alias.Abbreviation, ppdm_table_alias.Active_ind, ppdm_table_alias.Alias_long_name, ppdm_table_alias.Alias_reason_type, ppdm_table_alias.Alias_short_name, ppdm_table_alias.Alias_type, ppdm_table_alias.Amended_date, ppdm_table_alias.Created_date, ppdm_table_alias.Effective_date, ppdm_table_alias.Expiry_date, ppdm_table_alias.System_id, ppdm_table_alias.Original_ind, ppdm_table_alias.Owner_ba_id, ppdm_table_alias.Ppdm_guid, ppdm_table_alias.Preferred_ind, ppdm_table_alias.Reason_desc, ppdm_table_alias.Remark, ppdm_table_alias.Source, ppdm_table_alias.Source_document_id, ppdm_table_alias.Struckoff_date, ppdm_table_alias.Sw_application_id, ppdm_table_alias.Use_rule, ppdm_table_alias.Row_changed_by, ppdm_table_alias.Row_changed_date, ppdm_table_alias.Row_created_by, ppdm_table_alias.Row_effective_date, ppdm_table_alias.Row_expiry_date, ppdm_table_alias.Row_quality, ppdm_table_alias.System_id)
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

func PatchPpdmTableAlias(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update ppdm_table_alias set "
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

func DeletePpdmTableAlias(c *fiber.Ctx) error {
	id := c.Params("id")
	var ppdm_table_alias dto.Ppdm_table_alias
	ppdm_table_alias.System_id = id

	stmt, err := db.Prepare("delete from ppdm_table_alias where system_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(ppdm_table_alias.System_id)
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


