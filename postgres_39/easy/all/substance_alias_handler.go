package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetSubstanceAlias(c *fiber.Ctx) error {
	rows, err := db.Query("select * from substance_alias")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Substance_alias

	for rows.Next() {
		var substance_alias dto.Substance_alias
		if err := rows.Scan(&substance_alias.Substance_id, &substance_alias.Alias_id, &substance_alias.Abbreviation, &substance_alias.Active_ind, &substance_alias.Alias_long_name, &substance_alias.Alias_reason_type, &substance_alias.Alias_short_name, &substance_alias.Alias_type, &substance_alias.Amended_date, &substance_alias.Catalogue_equip_id, &substance_alias.Created_date, &substance_alias.Effective_date, &substance_alias.Expiry_date, &substance_alias.Substance_id, &substance_alias.Original_ind, &substance_alias.Owner_ba_id, &substance_alias.Ppdm_guid, &substance_alias.Preferred_ind, &substance_alias.Reason_desc, &substance_alias.Remark, &substance_alias.Source, &substance_alias.Source_document_id, &substance_alias.Struckoff_date, &substance_alias.Sw_application_id, &substance_alias.Use_rule, &substance_alias.Row_changed_by, &substance_alias.Row_changed_date, &substance_alias.Row_created_by, &substance_alias.Row_created_date, &substance_alias.Row_effective_date, &substance_alias.Row_expiry_date, &substance_alias.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append substance_alias to result
		result = append(result, substance_alias)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetSubstanceAlias(c *fiber.Ctx) error {
	var substance_alias dto.Substance_alias

	setDefaults(&substance_alias)

	if err := json.Unmarshal(c.Body(), &substance_alias); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into substance_alias values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32)")
	if err != nil {
		return err
	}
	substance_alias.Row_created_date = formatDate(substance_alias.Row_created_date)
	substance_alias.Row_changed_date = formatDate(substance_alias.Row_changed_date)
	substance_alias.Amended_date = formatDateString(substance_alias.Amended_date)
	substance_alias.Created_date = formatDateString(substance_alias.Created_date)
	substance_alias.Effective_date = formatDateString(substance_alias.Effective_date)
	substance_alias.Expiry_date = formatDateString(substance_alias.Expiry_date)
	substance_alias.Struckoff_date = formatDateString(substance_alias.Struckoff_date)
	substance_alias.Row_effective_date = formatDateString(substance_alias.Row_effective_date)
	substance_alias.Row_expiry_date = formatDateString(substance_alias.Row_expiry_date)






	rows, err := stmt.Exec(substance_alias.Substance_id, substance_alias.Alias_id, substance_alias.Abbreviation, substance_alias.Active_ind, substance_alias.Alias_long_name, substance_alias.Alias_reason_type, substance_alias.Alias_short_name, substance_alias.Alias_type, substance_alias.Amended_date, substance_alias.Catalogue_equip_id, substance_alias.Created_date, substance_alias.Effective_date, substance_alias.Expiry_date, substance_alias.Substance_id, substance_alias.Original_ind, substance_alias.Owner_ba_id, substance_alias.Ppdm_guid, substance_alias.Preferred_ind, substance_alias.Reason_desc, substance_alias.Remark, substance_alias.Source, substance_alias.Source_document_id, substance_alias.Struckoff_date, substance_alias.Sw_application_id, substance_alias.Use_rule, substance_alias.Row_changed_by, substance_alias.Row_changed_date, substance_alias.Row_created_by, substance_alias.Row_created_date, substance_alias.Row_effective_date, substance_alias.Row_expiry_date, substance_alias.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateSubstanceAlias(c *fiber.Ctx) error {
	var substance_alias dto.Substance_alias

	setDefaults(&substance_alias)

	if err := json.Unmarshal(c.Body(), &substance_alias); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	substance_alias.Substance_id = id

    if substance_alias.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update substance_alias set alias_id = :1, abbreviation = :2, active_ind = :3, alias_long_name = :4, alias_reason_type = :5, alias_short_name = :6, alias_type = :7, amended_date = :8, catalogue_equip_id = :9, created_date = :10, effective_date = :11, expiry_date = :12, substance_id = :13, original_ind = :14, owner_ba_id = :15, ppdm_guid = :16, preferred_ind = :17, reason_desc = :18, remark = :19, source = :20, source_document_id = :21, struckoff_date = :22, sw_application_id = :23, use_rule = :24, row_changed_by = :25, row_changed_date = :26, row_created_by = :27, row_effective_date = :28, row_expiry_date = :29, row_quality = :30 where substance_id = :32")
	if err != nil {
		return err
	}

	substance_alias.Row_changed_date = formatDate(substance_alias.Row_changed_date)
	substance_alias.Amended_date = formatDateString(substance_alias.Amended_date)
	substance_alias.Created_date = formatDateString(substance_alias.Created_date)
	substance_alias.Effective_date = formatDateString(substance_alias.Effective_date)
	substance_alias.Expiry_date = formatDateString(substance_alias.Expiry_date)
	substance_alias.Struckoff_date = formatDateString(substance_alias.Struckoff_date)
	substance_alias.Row_effective_date = formatDateString(substance_alias.Row_effective_date)
	substance_alias.Row_expiry_date = formatDateString(substance_alias.Row_expiry_date)






	rows, err := stmt.Exec(substance_alias.Alias_id, substance_alias.Abbreviation, substance_alias.Active_ind, substance_alias.Alias_long_name, substance_alias.Alias_reason_type, substance_alias.Alias_short_name, substance_alias.Alias_type, substance_alias.Amended_date, substance_alias.Catalogue_equip_id, substance_alias.Created_date, substance_alias.Effective_date, substance_alias.Expiry_date, substance_alias.Substance_id, substance_alias.Original_ind, substance_alias.Owner_ba_id, substance_alias.Ppdm_guid, substance_alias.Preferred_ind, substance_alias.Reason_desc, substance_alias.Remark, substance_alias.Source, substance_alias.Source_document_id, substance_alias.Struckoff_date, substance_alias.Sw_application_id, substance_alias.Use_rule, substance_alias.Row_changed_by, substance_alias.Row_changed_date, substance_alias.Row_created_by, substance_alias.Row_effective_date, substance_alias.Row_expiry_date, substance_alias.Row_quality, substance_alias.Substance_id)
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

func PatchSubstanceAlias(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update substance_alias set "
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
	query += " where substance_id = :id"

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

func DeleteSubstanceAlias(c *fiber.Ctx) error {
	id := c.Params("id")
	var substance_alias dto.Substance_alias
	substance_alias.Substance_id = id

	stmt, err := db.Prepare("delete from substance_alias where substance_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(substance_alias.Substance_id)
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


