package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetAnlCalcAlias(c *fiber.Ctx) error {
	rows, err := db.Query("select * from anl_calc_alias")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Anl_calc_alias

	for rows.Next() {
		var anl_calc_alias dto.Anl_calc_alias
		if err := rows.Scan(&anl_calc_alias.Calculate_method_id, &anl_calc_alias.Alias_id, &anl_calc_alias.Abbreviation, &anl_calc_alias.Active_ind, &anl_calc_alias.Alias_long_name, &anl_calc_alias.Alias_reason_type, &anl_calc_alias.Alias_short_name, &anl_calc_alias.Alias_type, &anl_calc_alias.Amended_date, &anl_calc_alias.Anl_method_id, &anl_calc_alias.Anl_method_set_id, &anl_calc_alias.Created_date, &anl_calc_alias.Effective_date, &anl_calc_alias.Expiry_date, &anl_calc_alias.Calculate_method_id, &anl_calc_alias.Original_ind, &anl_calc_alias.Owner_ba_id, &anl_calc_alias.Ppdm_guid, &anl_calc_alias.Preferred_ind, &anl_calc_alias.Reason_desc, &anl_calc_alias.Remark, &anl_calc_alias.Source, &anl_calc_alias.Source_document_id, &anl_calc_alias.Struckoff_date, &anl_calc_alias.Sw_application_id, &anl_calc_alias.Use_rule, &anl_calc_alias.Row_changed_by, &anl_calc_alias.Row_changed_date, &anl_calc_alias.Row_created_by, &anl_calc_alias.Row_created_date, &anl_calc_alias.Row_effective_date, &anl_calc_alias.Row_expiry_date, &anl_calc_alias.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append anl_calc_alias to result
		result = append(result, anl_calc_alias)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetAnlCalcAlias(c *fiber.Ctx) error {
	var anl_calc_alias dto.Anl_calc_alias

	setDefaults(&anl_calc_alias)

	if err := json.Unmarshal(c.Body(), &anl_calc_alias); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into anl_calc_alias values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32, :33)")
	if err != nil {
		return err
	}
	anl_calc_alias.Row_created_date = formatDate(anl_calc_alias.Row_created_date)
	anl_calc_alias.Row_changed_date = formatDate(anl_calc_alias.Row_changed_date)
	anl_calc_alias.Amended_date = formatDateString(anl_calc_alias.Amended_date)
	anl_calc_alias.Created_date = formatDateString(anl_calc_alias.Created_date)
	anl_calc_alias.Effective_date = formatDateString(anl_calc_alias.Effective_date)
	anl_calc_alias.Expiry_date = formatDateString(anl_calc_alias.Expiry_date)
	anl_calc_alias.Struckoff_date = formatDateString(anl_calc_alias.Struckoff_date)
	anl_calc_alias.Row_effective_date = formatDateString(anl_calc_alias.Row_effective_date)
	anl_calc_alias.Row_expiry_date = formatDateString(anl_calc_alias.Row_expiry_date)






	rows, err := stmt.Exec(anl_calc_alias.Calculate_method_id, anl_calc_alias.Alias_id, anl_calc_alias.Abbreviation, anl_calc_alias.Active_ind, anl_calc_alias.Alias_long_name, anl_calc_alias.Alias_reason_type, anl_calc_alias.Alias_short_name, anl_calc_alias.Alias_type, anl_calc_alias.Amended_date, anl_calc_alias.Anl_method_id, anl_calc_alias.Anl_method_set_id, anl_calc_alias.Created_date, anl_calc_alias.Effective_date, anl_calc_alias.Expiry_date, anl_calc_alias.Calculate_method_id, anl_calc_alias.Original_ind, anl_calc_alias.Owner_ba_id, anl_calc_alias.Ppdm_guid, anl_calc_alias.Preferred_ind, anl_calc_alias.Reason_desc, anl_calc_alias.Remark, anl_calc_alias.Source, anl_calc_alias.Source_document_id, anl_calc_alias.Struckoff_date, anl_calc_alias.Sw_application_id, anl_calc_alias.Use_rule, anl_calc_alias.Row_changed_by, anl_calc_alias.Row_changed_date, anl_calc_alias.Row_created_by, anl_calc_alias.Row_created_date, anl_calc_alias.Row_effective_date, anl_calc_alias.Row_expiry_date, anl_calc_alias.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateAnlCalcAlias(c *fiber.Ctx) error {
	var anl_calc_alias dto.Anl_calc_alias

	setDefaults(&anl_calc_alias)

	if err := json.Unmarshal(c.Body(), &anl_calc_alias); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	anl_calc_alias.Calculate_method_id = id

    if anl_calc_alias.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update anl_calc_alias set alias_id = :1, abbreviation = :2, active_ind = :3, alias_long_name = :4, alias_reason_type = :5, alias_short_name = :6, alias_type = :7, amended_date = :8, anl_method_id = :9, anl_method_set_id = :10, created_date = :11, effective_date = :12, expiry_date = :13, calculate_method_id = :14, original_ind = :15, owner_ba_id = :16, ppdm_guid = :17, preferred_ind = :18, reason_desc = :19, remark = :20, source = :21, source_document_id = :22, struckoff_date = :23, sw_application_id = :24, use_rule = :25, row_changed_by = :26, row_changed_date = :27, row_created_by = :28, row_effective_date = :29, row_expiry_date = :30, row_quality = :31 where calculate_method_id = :33")
	if err != nil {
		return err
	}

	anl_calc_alias.Row_changed_date = formatDate(anl_calc_alias.Row_changed_date)
	anl_calc_alias.Amended_date = formatDateString(anl_calc_alias.Amended_date)
	anl_calc_alias.Created_date = formatDateString(anl_calc_alias.Created_date)
	anl_calc_alias.Effective_date = formatDateString(anl_calc_alias.Effective_date)
	anl_calc_alias.Expiry_date = formatDateString(anl_calc_alias.Expiry_date)
	anl_calc_alias.Struckoff_date = formatDateString(anl_calc_alias.Struckoff_date)
	anl_calc_alias.Row_effective_date = formatDateString(anl_calc_alias.Row_effective_date)
	anl_calc_alias.Row_expiry_date = formatDateString(anl_calc_alias.Row_expiry_date)






	rows, err := stmt.Exec(anl_calc_alias.Alias_id, anl_calc_alias.Abbreviation, anl_calc_alias.Active_ind, anl_calc_alias.Alias_long_name, anl_calc_alias.Alias_reason_type, anl_calc_alias.Alias_short_name, anl_calc_alias.Alias_type, anl_calc_alias.Amended_date, anl_calc_alias.Anl_method_id, anl_calc_alias.Anl_method_set_id, anl_calc_alias.Created_date, anl_calc_alias.Effective_date, anl_calc_alias.Expiry_date, anl_calc_alias.Calculate_method_id, anl_calc_alias.Original_ind, anl_calc_alias.Owner_ba_id, anl_calc_alias.Ppdm_guid, anl_calc_alias.Preferred_ind, anl_calc_alias.Reason_desc, anl_calc_alias.Remark, anl_calc_alias.Source, anl_calc_alias.Source_document_id, anl_calc_alias.Struckoff_date, anl_calc_alias.Sw_application_id, anl_calc_alias.Use_rule, anl_calc_alias.Row_changed_by, anl_calc_alias.Row_changed_date, anl_calc_alias.Row_created_by, anl_calc_alias.Row_effective_date, anl_calc_alias.Row_expiry_date, anl_calc_alias.Row_quality, anl_calc_alias.Calculate_method_id)
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

func PatchAnlCalcAlias(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update anl_calc_alias set "
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
	query += " where calculate_method_id = :id"

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

func DeleteAnlCalcAlias(c *fiber.Ctx) error {
	id := c.Params("id")
	var anl_calc_alias dto.Anl_calc_alias
	anl_calc_alias.Calculate_method_id = id

	stmt, err := db.Prepare("delete from anl_calc_alias where calculate_method_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(anl_calc_alias.Calculate_method_id)
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


