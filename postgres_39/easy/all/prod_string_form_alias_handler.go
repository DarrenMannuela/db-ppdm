package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetProdStringFormAlias(c *fiber.Ctx) error {
	rows, err := db.Query("select * from prod_string_form_alias")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Prod_string_form_alias

	for rows.Next() {
		var prod_string_form_alias dto.Prod_string_form_alias
		if err := rows.Scan(&prod_string_form_alias.Uwi, &prod_string_form_alias.Prod_string_source, &prod_string_form_alias.String_id, &prod_string_form_alias.Pr_str_form_obs_no, &prod_string_form_alias.Source, &prod_string_form_alias.Prod_string_form_alias_id, &prod_string_form_alias.Abbreviation, &prod_string_form_alias.Active_ind, &prod_string_form_alias.Alias_long_name, &prod_string_form_alias.Alias_reason_type, &prod_string_form_alias.Alias_short_name, &prod_string_form_alias.Alias_type, &prod_string_form_alias.Amended_date, &prod_string_form_alias.Created_date, &prod_string_form_alias.Effective_date, &prod_string_form_alias.Expiry_date, &prod_string_form_alias.Uwi, &prod_string_form_alias.Location_type, &prod_string_form_alias.Original_ind, &prod_string_form_alias.Owner_ba_id, &prod_string_form_alias.Ppdm_guid, &prod_string_form_alias.Preferred_ind, &prod_string_form_alias.Reason_desc, &prod_string_form_alias.Remark, &prod_string_form_alias.Source_document_id, &prod_string_form_alias.Struckoff_date, &prod_string_form_alias.Sw_application_id, &prod_string_form_alias.Use_rule, &prod_string_form_alias.Row_changed_by, &prod_string_form_alias.Row_changed_date, &prod_string_form_alias.Row_created_by, &prod_string_form_alias.Row_created_date, &prod_string_form_alias.Row_effective_date, &prod_string_form_alias.Row_expiry_date, &prod_string_form_alias.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append prod_string_form_alias to result
		result = append(result, prod_string_form_alias)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetProdStringFormAlias(c *fiber.Ctx) error {
	var prod_string_form_alias dto.Prod_string_form_alias

	setDefaults(&prod_string_form_alias)

	if err := json.Unmarshal(c.Body(), &prod_string_form_alias); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into prod_string_form_alias values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32, :33, :34, :35)")
	if err != nil {
		return err
	}
	prod_string_form_alias.Row_created_date = formatDate(prod_string_form_alias.Row_created_date)
	prod_string_form_alias.Row_changed_date = formatDate(prod_string_form_alias.Row_changed_date)
	prod_string_form_alias.Amended_date = formatDateString(prod_string_form_alias.Amended_date)
	prod_string_form_alias.Created_date = formatDateString(prod_string_form_alias.Created_date)
	prod_string_form_alias.Effective_date = formatDateString(prod_string_form_alias.Effective_date)
	prod_string_form_alias.Expiry_date = formatDateString(prod_string_form_alias.Expiry_date)
	prod_string_form_alias.Struckoff_date = formatDateString(prod_string_form_alias.Struckoff_date)
	prod_string_form_alias.Row_effective_date = formatDateString(prod_string_form_alias.Row_effective_date)
	prod_string_form_alias.Row_expiry_date = formatDateString(prod_string_form_alias.Row_expiry_date)






	rows, err := stmt.Exec(prod_string_form_alias.Uwi, prod_string_form_alias.Prod_string_source, prod_string_form_alias.String_id, prod_string_form_alias.Pr_str_form_obs_no, prod_string_form_alias.Source, prod_string_form_alias.Prod_string_form_alias_id, prod_string_form_alias.Abbreviation, prod_string_form_alias.Active_ind, prod_string_form_alias.Alias_long_name, prod_string_form_alias.Alias_reason_type, prod_string_form_alias.Alias_short_name, prod_string_form_alias.Alias_type, prod_string_form_alias.Amended_date, prod_string_form_alias.Created_date, prod_string_form_alias.Effective_date, prod_string_form_alias.Expiry_date, prod_string_form_alias.Uwi, prod_string_form_alias.Location_type, prod_string_form_alias.Original_ind, prod_string_form_alias.Owner_ba_id, prod_string_form_alias.Ppdm_guid, prod_string_form_alias.Preferred_ind, prod_string_form_alias.Reason_desc, prod_string_form_alias.Remark, prod_string_form_alias.Source_document_id, prod_string_form_alias.Struckoff_date, prod_string_form_alias.Sw_application_id, prod_string_form_alias.Use_rule, prod_string_form_alias.Row_changed_by, prod_string_form_alias.Row_changed_date, prod_string_form_alias.Row_created_by, prod_string_form_alias.Row_created_date, prod_string_form_alias.Row_effective_date, prod_string_form_alias.Row_expiry_date, prod_string_form_alias.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateProdStringFormAlias(c *fiber.Ctx) error {
	var prod_string_form_alias dto.Prod_string_form_alias

	setDefaults(&prod_string_form_alias)

	if err := json.Unmarshal(c.Body(), &prod_string_form_alias); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	prod_string_form_alias.Uwi = id

    if prod_string_form_alias.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update prod_string_form_alias set prod_string_source = :1, string_id = :2, pr_str_form_obs_no = :3, source = :4, prod_string_form_alias_id = :5, abbreviation = :6, active_ind = :7, alias_long_name = :8, alias_reason_type = :9, alias_short_name = :10, alias_type = :11, amended_date = :12, created_date = :13, effective_date = :14, expiry_date = :15, uwi = :16, location_type = :17, original_ind = :18, owner_ba_id = :19, ppdm_guid = :20, preferred_ind = :21, reason_desc = :22, remark = :23, source_document_id = :24, struckoff_date = :25, sw_application_id = :26, use_rule = :27, row_changed_by = :28, row_changed_date = :29, row_created_by = :30, row_effective_date = :31, row_expiry_date = :32, row_quality = :33 where uwi = :35")
	if err != nil {
		return err
	}

	prod_string_form_alias.Row_changed_date = formatDate(prod_string_form_alias.Row_changed_date)
	prod_string_form_alias.Amended_date = formatDateString(prod_string_form_alias.Amended_date)
	prod_string_form_alias.Created_date = formatDateString(prod_string_form_alias.Created_date)
	prod_string_form_alias.Effective_date = formatDateString(prod_string_form_alias.Effective_date)
	prod_string_form_alias.Expiry_date = formatDateString(prod_string_form_alias.Expiry_date)
	prod_string_form_alias.Struckoff_date = formatDateString(prod_string_form_alias.Struckoff_date)
	prod_string_form_alias.Row_effective_date = formatDateString(prod_string_form_alias.Row_effective_date)
	prod_string_form_alias.Row_expiry_date = formatDateString(prod_string_form_alias.Row_expiry_date)






	rows, err := stmt.Exec(prod_string_form_alias.Prod_string_source, prod_string_form_alias.String_id, prod_string_form_alias.Pr_str_form_obs_no, prod_string_form_alias.Source, prod_string_form_alias.Prod_string_form_alias_id, prod_string_form_alias.Abbreviation, prod_string_form_alias.Active_ind, prod_string_form_alias.Alias_long_name, prod_string_form_alias.Alias_reason_type, prod_string_form_alias.Alias_short_name, prod_string_form_alias.Alias_type, prod_string_form_alias.Amended_date, prod_string_form_alias.Created_date, prod_string_form_alias.Effective_date, prod_string_form_alias.Expiry_date, prod_string_form_alias.Uwi, prod_string_form_alias.Location_type, prod_string_form_alias.Original_ind, prod_string_form_alias.Owner_ba_id, prod_string_form_alias.Ppdm_guid, prod_string_form_alias.Preferred_ind, prod_string_form_alias.Reason_desc, prod_string_form_alias.Remark, prod_string_form_alias.Source_document_id, prod_string_form_alias.Struckoff_date, prod_string_form_alias.Sw_application_id, prod_string_form_alias.Use_rule, prod_string_form_alias.Row_changed_by, prod_string_form_alias.Row_changed_date, prod_string_form_alias.Row_created_by, prod_string_form_alias.Row_effective_date, prod_string_form_alias.Row_expiry_date, prod_string_form_alias.Row_quality, prod_string_form_alias.Uwi)
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

func PatchProdStringFormAlias(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update prod_string_form_alias set "
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

func DeleteProdStringFormAlias(c *fiber.Ctx) error {
	id := c.Params("id")
	var prod_string_form_alias dto.Prod_string_form_alias
	prod_string_form_alias.Uwi = id

	stmt, err := db.Prepare("delete from prod_string_form_alias where uwi = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(prod_string_form_alias.Uwi)
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


