package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetProdStringAlias(c *fiber.Ctx) error {
	rows, err := db.Query("select * from prod_string_alias")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Prod_string_alias

	for rows.Next() {
		var prod_string_alias dto.Prod_string_alias
		if err := rows.Scan(&prod_string_alias.Uwi, &prod_string_alias.Prod_string_source, &prod_string_alias.String_id, &prod_string_alias.Source, &prod_string_alias.Alias_id, &prod_string_alias.Abbreviation, &prod_string_alias.Active_ind, &prod_string_alias.Alias_long_name, &prod_string_alias.Alias_reason_type, &prod_string_alias.Alias_short_name, &prod_string_alias.Alias_type, &prod_string_alias.Amended_date, &prod_string_alias.Created_date, &prod_string_alias.Effective_date, &prod_string_alias.Expiry_date, &prod_string_alias.Uwi, &prod_string_alias.Location_type, &prod_string_alias.Original_ind, &prod_string_alias.Owner_ba_id, &prod_string_alias.Ppdm_guid, &prod_string_alias.Preferred_ind, &prod_string_alias.Reason_desc, &prod_string_alias.Remark, &prod_string_alias.Source_document_id, &prod_string_alias.Struckoff_date, &prod_string_alias.Sw_application_id, &prod_string_alias.Use_rule, &prod_string_alias.Row_changed_by, &prod_string_alias.Row_changed_date, &prod_string_alias.Row_created_by, &prod_string_alias.Row_created_date, &prod_string_alias.Row_effective_date, &prod_string_alias.Row_expiry_date, &prod_string_alias.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append prod_string_alias to result
		result = append(result, prod_string_alias)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetProdStringAlias(c *fiber.Ctx) error {
	var prod_string_alias dto.Prod_string_alias

	setDefaults(&prod_string_alias)

	if err := json.Unmarshal(c.Body(), &prod_string_alias); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into prod_string_alias values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32, :33, :34)")
	if err != nil {
		return err
	}
	prod_string_alias.Row_created_date = formatDate(prod_string_alias.Row_created_date)
	prod_string_alias.Row_changed_date = formatDate(prod_string_alias.Row_changed_date)
	prod_string_alias.Amended_date = formatDateString(prod_string_alias.Amended_date)
	prod_string_alias.Created_date = formatDateString(prod_string_alias.Created_date)
	prod_string_alias.Effective_date = formatDateString(prod_string_alias.Effective_date)
	prod_string_alias.Expiry_date = formatDateString(prod_string_alias.Expiry_date)
	prod_string_alias.Struckoff_date = formatDateString(prod_string_alias.Struckoff_date)
	prod_string_alias.Row_effective_date = formatDateString(prod_string_alias.Row_effective_date)
	prod_string_alias.Row_expiry_date = formatDateString(prod_string_alias.Row_expiry_date)






	rows, err := stmt.Exec(prod_string_alias.Uwi, prod_string_alias.Prod_string_source, prod_string_alias.String_id, prod_string_alias.Source, prod_string_alias.Alias_id, prod_string_alias.Abbreviation, prod_string_alias.Active_ind, prod_string_alias.Alias_long_name, prod_string_alias.Alias_reason_type, prod_string_alias.Alias_short_name, prod_string_alias.Alias_type, prod_string_alias.Amended_date, prod_string_alias.Created_date, prod_string_alias.Effective_date, prod_string_alias.Expiry_date, prod_string_alias.Uwi, prod_string_alias.Location_type, prod_string_alias.Original_ind, prod_string_alias.Owner_ba_id, prod_string_alias.Ppdm_guid, prod_string_alias.Preferred_ind, prod_string_alias.Reason_desc, prod_string_alias.Remark, prod_string_alias.Source_document_id, prod_string_alias.Struckoff_date, prod_string_alias.Sw_application_id, prod_string_alias.Use_rule, prod_string_alias.Row_changed_by, prod_string_alias.Row_changed_date, prod_string_alias.Row_created_by, prod_string_alias.Row_created_date, prod_string_alias.Row_effective_date, prod_string_alias.Row_expiry_date, prod_string_alias.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateProdStringAlias(c *fiber.Ctx) error {
	var prod_string_alias dto.Prod_string_alias

	setDefaults(&prod_string_alias)

	if err := json.Unmarshal(c.Body(), &prod_string_alias); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	prod_string_alias.Uwi = id

    if prod_string_alias.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update prod_string_alias set prod_string_source = :1, string_id = :2, source = :3, alias_id = :4, abbreviation = :5, active_ind = :6, alias_long_name = :7, alias_reason_type = :8, alias_short_name = :9, alias_type = :10, amended_date = :11, created_date = :12, effective_date = :13, expiry_date = :14, uwi = :15, location_type = :16, original_ind = :17, owner_ba_id = :18, ppdm_guid = :19, preferred_ind = :20, reason_desc = :21, remark = :22, source_document_id = :23, struckoff_date = :24, sw_application_id = :25, use_rule = :26, row_changed_by = :27, row_changed_date = :28, row_created_by = :29, row_effective_date = :30, row_expiry_date = :31, row_quality = :32 where uwi = :34")
	if err != nil {
		return err
	}

	prod_string_alias.Row_changed_date = formatDate(prod_string_alias.Row_changed_date)
	prod_string_alias.Amended_date = formatDateString(prod_string_alias.Amended_date)
	prod_string_alias.Created_date = formatDateString(prod_string_alias.Created_date)
	prod_string_alias.Effective_date = formatDateString(prod_string_alias.Effective_date)
	prod_string_alias.Expiry_date = formatDateString(prod_string_alias.Expiry_date)
	prod_string_alias.Struckoff_date = formatDateString(prod_string_alias.Struckoff_date)
	prod_string_alias.Row_effective_date = formatDateString(prod_string_alias.Row_effective_date)
	prod_string_alias.Row_expiry_date = formatDateString(prod_string_alias.Row_expiry_date)






	rows, err := stmt.Exec(prod_string_alias.Prod_string_source, prod_string_alias.String_id, prod_string_alias.Source, prod_string_alias.Alias_id, prod_string_alias.Abbreviation, prod_string_alias.Active_ind, prod_string_alias.Alias_long_name, prod_string_alias.Alias_reason_type, prod_string_alias.Alias_short_name, prod_string_alias.Alias_type, prod_string_alias.Amended_date, prod_string_alias.Created_date, prod_string_alias.Effective_date, prod_string_alias.Expiry_date, prod_string_alias.Uwi, prod_string_alias.Location_type, prod_string_alias.Original_ind, prod_string_alias.Owner_ba_id, prod_string_alias.Ppdm_guid, prod_string_alias.Preferred_ind, prod_string_alias.Reason_desc, prod_string_alias.Remark, prod_string_alias.Source_document_id, prod_string_alias.Struckoff_date, prod_string_alias.Sw_application_id, prod_string_alias.Use_rule, prod_string_alias.Row_changed_by, prod_string_alias.Row_changed_date, prod_string_alias.Row_created_by, prod_string_alias.Row_effective_date, prod_string_alias.Row_expiry_date, prod_string_alias.Row_quality, prod_string_alias.Uwi)
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

func PatchProdStringAlias(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update prod_string_alias set "
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

func DeleteProdStringAlias(c *fiber.Ctx) error {
	id := c.Params("id")
	var prod_string_alias dto.Prod_string_alias
	prod_string_alias.Uwi = id

	stmt, err := db.Prepare("delete from prod_string_alias where uwi = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(prod_string_alias.Uwi)
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


