package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetBaAlias(c *fiber.Ctx) error {
	rows, err := db.Query("select * from ba_alias")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Ba_alias

	for rows.Next() {
		var ba_alias dto.Ba_alias
		if err := rows.Scan(&ba_alias.Business_associate_id, &ba_alias.Source, &ba_alias.Ba_alias_id, &ba_alias.Abbreviation, &ba_alias.Active_ind, &ba_alias.Alias_long_name, &ba_alias.Alias_reason_type, &ba_alias.Alias_short_name, &ba_alias.Alias_type, &ba_alias.Amended_date, &ba_alias.Created_date, &ba_alias.Effective_date, &ba_alias.Expiry_date, &ba_alias.Business_associate_id, &ba_alias.Original_ind, &ba_alias.Owner_ba_id, &ba_alias.Password, &ba_alias.Ppdm_guid, &ba_alias.Preferred_ind, &ba_alias.Reason_desc, &ba_alias.Remark, &ba_alias.Source_document_id, &ba_alias.Struckoff_date, &ba_alias.Sw_application_id, &ba_alias.Use_rule, &ba_alias.Row_changed_by, &ba_alias.Row_changed_date, &ba_alias.Row_created_by, &ba_alias.Row_created_date, &ba_alias.Row_effective_date, &ba_alias.Row_expiry_date, &ba_alias.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append ba_alias to result
		result = append(result, ba_alias)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetBaAlias(c *fiber.Ctx) error {
	var ba_alias dto.Ba_alias

	setDefaults(&ba_alias)

	if err := json.Unmarshal(c.Body(), &ba_alias); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into ba_alias values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32)")
	if err != nil {
		return err
	}
	ba_alias.Row_created_date = formatDate(ba_alias.Row_created_date)
	ba_alias.Row_changed_date = formatDate(ba_alias.Row_changed_date)
	ba_alias.Amended_date = formatDateString(ba_alias.Amended_date)
	ba_alias.Created_date = formatDateString(ba_alias.Created_date)
	ba_alias.Effective_date = formatDateString(ba_alias.Effective_date)
	ba_alias.Expiry_date = formatDateString(ba_alias.Expiry_date)
	ba_alias.Struckoff_date = formatDateString(ba_alias.Struckoff_date)
	ba_alias.Row_effective_date = formatDateString(ba_alias.Row_effective_date)
	ba_alias.Row_expiry_date = formatDateString(ba_alias.Row_expiry_date)






	rows, err := stmt.Exec(ba_alias.Business_associate_id, ba_alias.Source, ba_alias.Ba_alias_id, ba_alias.Abbreviation, ba_alias.Active_ind, ba_alias.Alias_long_name, ba_alias.Alias_reason_type, ba_alias.Alias_short_name, ba_alias.Alias_type, ba_alias.Amended_date, ba_alias.Created_date, ba_alias.Effective_date, ba_alias.Expiry_date, ba_alias.Business_associate_id, ba_alias.Original_ind, ba_alias.Owner_ba_id, ba_alias.Password, ba_alias.Ppdm_guid, ba_alias.Preferred_ind, ba_alias.Reason_desc, ba_alias.Remark, ba_alias.Source_document_id, ba_alias.Struckoff_date, ba_alias.Sw_application_id, ba_alias.Use_rule, ba_alias.Row_changed_by, ba_alias.Row_changed_date, ba_alias.Row_created_by, ba_alias.Row_created_date, ba_alias.Row_effective_date, ba_alias.Row_expiry_date, ba_alias.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateBaAlias(c *fiber.Ctx) error {
	var ba_alias dto.Ba_alias

	setDefaults(&ba_alias)

	if err := json.Unmarshal(c.Body(), &ba_alias); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	ba_alias.Business_associate_id = id

    if ba_alias.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update ba_alias set source = :1, ba_alias_id = :2, abbreviation = :3, active_ind = :4, alias_long_name = :5, alias_reason_type = :6, alias_short_name = :7, alias_type = :8, amended_date = :9, created_date = :10, effective_date = :11, expiry_date = :12, business_associate_id = :13, original_ind = :14, owner_ba_id = :15, password = :16, ppdm_guid = :17, preferred_ind = :18, reason_desc = :19, remark = :20, source_document_id = :21, struckoff_date = :22, sw_application_id = :23, use_rule = :24, row_changed_by = :25, row_changed_date = :26, row_created_by = :27, row_effective_date = :28, row_expiry_date = :29, row_quality = :30 where business_associate_id = :32")
	if err != nil {
		return err
	}

	ba_alias.Row_changed_date = formatDate(ba_alias.Row_changed_date)
	ba_alias.Amended_date = formatDateString(ba_alias.Amended_date)
	ba_alias.Created_date = formatDateString(ba_alias.Created_date)
	ba_alias.Effective_date = formatDateString(ba_alias.Effective_date)
	ba_alias.Expiry_date = formatDateString(ba_alias.Expiry_date)
	ba_alias.Struckoff_date = formatDateString(ba_alias.Struckoff_date)
	ba_alias.Row_effective_date = formatDateString(ba_alias.Row_effective_date)
	ba_alias.Row_expiry_date = formatDateString(ba_alias.Row_expiry_date)






	rows, err := stmt.Exec(ba_alias.Source, ba_alias.Ba_alias_id, ba_alias.Abbreviation, ba_alias.Active_ind, ba_alias.Alias_long_name, ba_alias.Alias_reason_type, ba_alias.Alias_short_name, ba_alias.Alias_type, ba_alias.Amended_date, ba_alias.Created_date, ba_alias.Effective_date, ba_alias.Expiry_date, ba_alias.Business_associate_id, ba_alias.Original_ind, ba_alias.Owner_ba_id, ba_alias.Password, ba_alias.Ppdm_guid, ba_alias.Preferred_ind, ba_alias.Reason_desc, ba_alias.Remark, ba_alias.Source_document_id, ba_alias.Struckoff_date, ba_alias.Sw_application_id, ba_alias.Use_rule, ba_alias.Row_changed_by, ba_alias.Row_changed_date, ba_alias.Row_created_by, ba_alias.Row_effective_date, ba_alias.Row_expiry_date, ba_alias.Row_quality, ba_alias.Business_associate_id)
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

func PatchBaAlias(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update ba_alias set "
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
	query += " where business_associate_id = :id"

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

func DeleteBaAlias(c *fiber.Ctx) error {
	id := c.Params("id")
	var ba_alias dto.Ba_alias
	ba_alias.Business_associate_id = id

	stmt, err := db.Prepare("delete from ba_alias where business_associate_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(ba_alias.Business_associate_id)
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


