package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetPoolAlias(c *fiber.Ctx) error {
	rows, err := db.Query("select * from pool_alias")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Pool_alias

	for rows.Next() {
		var pool_alias dto.Pool_alias
		if err := rows.Scan(&pool_alias.Pool_id, &pool_alias.Source, &pool_alias.Pool_alias_id, &pool_alias.Abbreviation, &pool_alias.Active_ind, &pool_alias.Alias_long_name, &pool_alias.Alias_reason_type, &pool_alias.Alias_short_name, &pool_alias.Alias_type, &pool_alias.Amended_date, &pool_alias.Created_date, &pool_alias.Effective_date, &pool_alias.Expiry_date, &pool_alias.Pool_id, &pool_alias.Original_ind, &pool_alias.Owner_ba_id, &pool_alias.Ppdm_guid, &pool_alias.Preferred_ind, &pool_alias.Reason_desc, &pool_alias.Remark, &pool_alias.Source_document_id, &pool_alias.Struckoff_date, &pool_alias.Sw_application_id, &pool_alias.Use_rule, &pool_alias.Row_changed_by, &pool_alias.Row_changed_date, &pool_alias.Row_created_by, &pool_alias.Row_created_date, &pool_alias.Row_effective_date, &pool_alias.Row_expiry_date, &pool_alias.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append pool_alias to result
		result = append(result, pool_alias)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetPoolAlias(c *fiber.Ctx) error {
	var pool_alias dto.Pool_alias

	setDefaults(&pool_alias)

	if err := json.Unmarshal(c.Body(), &pool_alias); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into pool_alias values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31)")
	if err != nil {
		return err
	}
	pool_alias.Row_created_date = formatDate(pool_alias.Row_created_date)
	pool_alias.Row_changed_date = formatDate(pool_alias.Row_changed_date)
	pool_alias.Amended_date = formatDateString(pool_alias.Amended_date)
	pool_alias.Created_date = formatDateString(pool_alias.Created_date)
	pool_alias.Effective_date = formatDateString(pool_alias.Effective_date)
	pool_alias.Expiry_date = formatDateString(pool_alias.Expiry_date)
	pool_alias.Struckoff_date = formatDateString(pool_alias.Struckoff_date)
	pool_alias.Row_effective_date = formatDateString(pool_alias.Row_effective_date)
	pool_alias.Row_expiry_date = formatDateString(pool_alias.Row_expiry_date)






	rows, err := stmt.Exec(pool_alias.Pool_id, pool_alias.Source, pool_alias.Pool_alias_id, pool_alias.Abbreviation, pool_alias.Active_ind, pool_alias.Alias_long_name, pool_alias.Alias_reason_type, pool_alias.Alias_short_name, pool_alias.Alias_type, pool_alias.Amended_date, pool_alias.Created_date, pool_alias.Effective_date, pool_alias.Expiry_date, pool_alias.Pool_id, pool_alias.Original_ind, pool_alias.Owner_ba_id, pool_alias.Ppdm_guid, pool_alias.Preferred_ind, pool_alias.Reason_desc, pool_alias.Remark, pool_alias.Source_document_id, pool_alias.Struckoff_date, pool_alias.Sw_application_id, pool_alias.Use_rule, pool_alias.Row_changed_by, pool_alias.Row_changed_date, pool_alias.Row_created_by, pool_alias.Row_created_date, pool_alias.Row_effective_date, pool_alias.Row_expiry_date, pool_alias.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdatePoolAlias(c *fiber.Ctx) error {
	var pool_alias dto.Pool_alias

	setDefaults(&pool_alias)

	if err := json.Unmarshal(c.Body(), &pool_alias); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	pool_alias.Pool_id = id

    if pool_alias.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update pool_alias set source = :1, pool_alias_id = :2, abbreviation = :3, active_ind = :4, alias_long_name = :5, alias_reason_type = :6, alias_short_name = :7, alias_type = :8, amended_date = :9, created_date = :10, effective_date = :11, expiry_date = :12, pool_id = :13, original_ind = :14, owner_ba_id = :15, ppdm_guid = :16, preferred_ind = :17, reason_desc = :18, remark = :19, source_document_id = :20, struckoff_date = :21, sw_application_id = :22, use_rule = :23, row_changed_by = :24, row_changed_date = :25, row_created_by = :26, row_effective_date = :27, row_expiry_date = :28, row_quality = :29 where pool_id = :31")
	if err != nil {
		return err
	}

	pool_alias.Row_changed_date = formatDate(pool_alias.Row_changed_date)
	pool_alias.Amended_date = formatDateString(pool_alias.Amended_date)
	pool_alias.Created_date = formatDateString(pool_alias.Created_date)
	pool_alias.Effective_date = formatDateString(pool_alias.Effective_date)
	pool_alias.Expiry_date = formatDateString(pool_alias.Expiry_date)
	pool_alias.Struckoff_date = formatDateString(pool_alias.Struckoff_date)
	pool_alias.Row_effective_date = formatDateString(pool_alias.Row_effective_date)
	pool_alias.Row_expiry_date = formatDateString(pool_alias.Row_expiry_date)






	rows, err := stmt.Exec(pool_alias.Source, pool_alias.Pool_alias_id, pool_alias.Abbreviation, pool_alias.Active_ind, pool_alias.Alias_long_name, pool_alias.Alias_reason_type, pool_alias.Alias_short_name, pool_alias.Alias_type, pool_alias.Amended_date, pool_alias.Created_date, pool_alias.Effective_date, pool_alias.Expiry_date, pool_alias.Pool_id, pool_alias.Original_ind, pool_alias.Owner_ba_id, pool_alias.Ppdm_guid, pool_alias.Preferred_ind, pool_alias.Reason_desc, pool_alias.Remark, pool_alias.Source_document_id, pool_alias.Struckoff_date, pool_alias.Sw_application_id, pool_alias.Use_rule, pool_alias.Row_changed_by, pool_alias.Row_changed_date, pool_alias.Row_created_by, pool_alias.Row_effective_date, pool_alias.Row_expiry_date, pool_alias.Row_quality, pool_alias.Pool_id)
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

func PatchPoolAlias(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update pool_alias set "
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
	query += " where pool_id = :id"

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

func DeletePoolAlias(c *fiber.Ctx) error {
	id := c.Params("id")
	var pool_alias dto.Pool_alias
	pool_alias.Pool_id = id

	stmt, err := db.Prepare("delete from pool_alias where pool_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(pool_alias.Pool_id)
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


