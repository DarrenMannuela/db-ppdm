package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetCatAdditiveAlias(c *fiber.Ctx) error {
	rows, err := db.Query("select * from cat_additive_alias")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Cat_additive_alias

	for rows.Next() {
		var cat_additive_alias dto.Cat_additive_alias
		if err := rows.Scan(&cat_additive_alias.Catalogue_additive_id, &cat_additive_alias.Cat_additive_alias_id, &cat_additive_alias.Abbreviation, &cat_additive_alias.Active_ind, &cat_additive_alias.Alias_long_name, &cat_additive_alias.Alias_reason_type, &cat_additive_alias.Alias_short_name, &cat_additive_alias.Alias_type, &cat_additive_alias.Amended_date, &cat_additive_alias.Created_date, &cat_additive_alias.Effective_date, &cat_additive_alias.Expiry_date, &cat_additive_alias.Catalogue_additive_id, &cat_additive_alias.Original_ind, &cat_additive_alias.Owner_ba_id, &cat_additive_alias.Ppdm_guid, &cat_additive_alias.Preferred_ind, &cat_additive_alias.Reason_desc, &cat_additive_alias.Remark, &cat_additive_alias.Source, &cat_additive_alias.Source_document_id, &cat_additive_alias.Struckoff_date, &cat_additive_alias.Sw_application_id, &cat_additive_alias.Use_rule, &cat_additive_alias.Row_changed_by, &cat_additive_alias.Row_changed_date, &cat_additive_alias.Row_created_by, &cat_additive_alias.Row_created_date, &cat_additive_alias.Row_effective_date, &cat_additive_alias.Row_expiry_date, &cat_additive_alias.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append cat_additive_alias to result
		result = append(result, cat_additive_alias)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetCatAdditiveAlias(c *fiber.Ctx) error {
	var cat_additive_alias dto.Cat_additive_alias

	setDefaults(&cat_additive_alias)

	if err := json.Unmarshal(c.Body(), &cat_additive_alias); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into cat_additive_alias values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31)")
	if err != nil {
		return err
	}
	cat_additive_alias.Row_created_date = formatDate(cat_additive_alias.Row_created_date)
	cat_additive_alias.Row_changed_date = formatDate(cat_additive_alias.Row_changed_date)
	cat_additive_alias.Amended_date = formatDateString(cat_additive_alias.Amended_date)
	cat_additive_alias.Created_date = formatDateString(cat_additive_alias.Created_date)
	cat_additive_alias.Effective_date = formatDateString(cat_additive_alias.Effective_date)
	cat_additive_alias.Expiry_date = formatDateString(cat_additive_alias.Expiry_date)
	cat_additive_alias.Struckoff_date = formatDateString(cat_additive_alias.Struckoff_date)
	cat_additive_alias.Row_effective_date = formatDateString(cat_additive_alias.Row_effective_date)
	cat_additive_alias.Row_expiry_date = formatDateString(cat_additive_alias.Row_expiry_date)






	rows, err := stmt.Exec(cat_additive_alias.Catalogue_additive_id, cat_additive_alias.Cat_additive_alias_id, cat_additive_alias.Abbreviation, cat_additive_alias.Active_ind, cat_additive_alias.Alias_long_name, cat_additive_alias.Alias_reason_type, cat_additive_alias.Alias_short_name, cat_additive_alias.Alias_type, cat_additive_alias.Amended_date, cat_additive_alias.Created_date, cat_additive_alias.Effective_date, cat_additive_alias.Expiry_date, cat_additive_alias.Catalogue_additive_id, cat_additive_alias.Original_ind, cat_additive_alias.Owner_ba_id, cat_additive_alias.Ppdm_guid, cat_additive_alias.Preferred_ind, cat_additive_alias.Reason_desc, cat_additive_alias.Remark, cat_additive_alias.Source, cat_additive_alias.Source_document_id, cat_additive_alias.Struckoff_date, cat_additive_alias.Sw_application_id, cat_additive_alias.Use_rule, cat_additive_alias.Row_changed_by, cat_additive_alias.Row_changed_date, cat_additive_alias.Row_created_by, cat_additive_alias.Row_created_date, cat_additive_alias.Row_effective_date, cat_additive_alias.Row_expiry_date, cat_additive_alias.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateCatAdditiveAlias(c *fiber.Ctx) error {
	var cat_additive_alias dto.Cat_additive_alias

	setDefaults(&cat_additive_alias)

	if err := json.Unmarshal(c.Body(), &cat_additive_alias); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	cat_additive_alias.Catalogue_additive_id = id

    if cat_additive_alias.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update cat_additive_alias set cat_additive_alias_id = :1, abbreviation = :2, active_ind = :3, alias_long_name = :4, alias_reason_type = :5, alias_short_name = :6, alias_type = :7, amended_date = :8, created_date = :9, effective_date = :10, expiry_date = :11, catalogue_additive_id = :12, original_ind = :13, owner_ba_id = :14, ppdm_guid = :15, preferred_ind = :16, reason_desc = :17, remark = :18, source = :19, source_document_id = :20, struckoff_date = :21, sw_application_id = :22, use_rule = :23, row_changed_by = :24, row_changed_date = :25, row_created_by = :26, row_effective_date = :27, row_expiry_date = :28, row_quality = :29 where catalogue_additive_id = :31")
	if err != nil {
		return err
	}

	cat_additive_alias.Row_changed_date = formatDate(cat_additive_alias.Row_changed_date)
	cat_additive_alias.Amended_date = formatDateString(cat_additive_alias.Amended_date)
	cat_additive_alias.Created_date = formatDateString(cat_additive_alias.Created_date)
	cat_additive_alias.Effective_date = formatDateString(cat_additive_alias.Effective_date)
	cat_additive_alias.Expiry_date = formatDateString(cat_additive_alias.Expiry_date)
	cat_additive_alias.Struckoff_date = formatDateString(cat_additive_alias.Struckoff_date)
	cat_additive_alias.Row_effective_date = formatDateString(cat_additive_alias.Row_effective_date)
	cat_additive_alias.Row_expiry_date = formatDateString(cat_additive_alias.Row_expiry_date)






	rows, err := stmt.Exec(cat_additive_alias.Cat_additive_alias_id, cat_additive_alias.Abbreviation, cat_additive_alias.Active_ind, cat_additive_alias.Alias_long_name, cat_additive_alias.Alias_reason_type, cat_additive_alias.Alias_short_name, cat_additive_alias.Alias_type, cat_additive_alias.Amended_date, cat_additive_alias.Created_date, cat_additive_alias.Effective_date, cat_additive_alias.Expiry_date, cat_additive_alias.Catalogue_additive_id, cat_additive_alias.Original_ind, cat_additive_alias.Owner_ba_id, cat_additive_alias.Ppdm_guid, cat_additive_alias.Preferred_ind, cat_additive_alias.Reason_desc, cat_additive_alias.Remark, cat_additive_alias.Source, cat_additive_alias.Source_document_id, cat_additive_alias.Struckoff_date, cat_additive_alias.Sw_application_id, cat_additive_alias.Use_rule, cat_additive_alias.Row_changed_by, cat_additive_alias.Row_changed_date, cat_additive_alias.Row_created_by, cat_additive_alias.Row_effective_date, cat_additive_alias.Row_expiry_date, cat_additive_alias.Row_quality, cat_additive_alias.Catalogue_additive_id)
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

func PatchCatAdditiveAlias(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update cat_additive_alias set "
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
	query += " where catalogue_additive_id = :id"

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

func DeleteCatAdditiveAlias(c *fiber.Ctx) error {
	id := c.Params("id")
	var cat_additive_alias dto.Cat_additive_alias
	cat_additive_alias.Catalogue_additive_id = id

	stmt, err := db.Prepare("delete from cat_additive_alias where catalogue_additive_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(cat_additive_alias.Catalogue_additive_id)
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


