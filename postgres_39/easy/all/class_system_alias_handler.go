package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetClassSystemAlias(c *fiber.Ctx) error {
	rows, err := db.Query("select * from class_system_alias")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Class_system_alias

	for rows.Next() {
		var class_system_alias dto.Class_system_alias
		if err := rows.Scan(&class_system_alias.Classification_system_id, &class_system_alias.Class_system_alias_id, &class_system_alias.Abbreviation, &class_system_alias.Active_ind, &class_system_alias.Alias_long_name, &class_system_alias.Alias_reason_type, &class_system_alias.Alias_short_name, &class_system_alias.Alias_type, &class_system_alias.Alias_version, &class_system_alias.Amended_date, &class_system_alias.Created_date, &class_system_alias.Effective_date, &class_system_alias.Expiry_date, &class_system_alias.Classification_system_id, &class_system_alias.Original_ind, &class_system_alias.Owner_ba_id, &class_system_alias.Ppdm_guid, &class_system_alias.Preferred_ind, &class_system_alias.Reason_desc, &class_system_alias.Remark, &class_system_alias.Source, &class_system_alias.Source_document_id, &class_system_alias.Struckoff_date, &class_system_alias.Sw_application_id, &class_system_alias.Use_rule, &class_system_alias.Row_changed_by, &class_system_alias.Row_changed_date, &class_system_alias.Row_created_by, &class_system_alias.Row_created_date, &class_system_alias.Row_effective_date, &class_system_alias.Row_expiry_date, &class_system_alias.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append class_system_alias to result
		result = append(result, class_system_alias)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetClassSystemAlias(c *fiber.Ctx) error {
	var class_system_alias dto.Class_system_alias

	setDefaults(&class_system_alias)

	if err := json.Unmarshal(c.Body(), &class_system_alias); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into class_system_alias values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32)")
	if err != nil {
		return err
	}
	class_system_alias.Row_created_date = formatDate(class_system_alias.Row_created_date)
	class_system_alias.Row_changed_date = formatDate(class_system_alias.Row_changed_date)
	class_system_alias.Amended_date = formatDateString(class_system_alias.Amended_date)
	class_system_alias.Created_date = formatDateString(class_system_alias.Created_date)
	class_system_alias.Effective_date = formatDateString(class_system_alias.Effective_date)
	class_system_alias.Expiry_date = formatDateString(class_system_alias.Expiry_date)
	class_system_alias.Struckoff_date = formatDateString(class_system_alias.Struckoff_date)
	class_system_alias.Row_effective_date = formatDateString(class_system_alias.Row_effective_date)
	class_system_alias.Row_expiry_date = formatDateString(class_system_alias.Row_expiry_date)






	rows, err := stmt.Exec(class_system_alias.Classification_system_id, class_system_alias.Class_system_alias_id, class_system_alias.Abbreviation, class_system_alias.Active_ind, class_system_alias.Alias_long_name, class_system_alias.Alias_reason_type, class_system_alias.Alias_short_name, class_system_alias.Alias_type, class_system_alias.Alias_version, class_system_alias.Amended_date, class_system_alias.Created_date, class_system_alias.Effective_date, class_system_alias.Expiry_date, class_system_alias.Classification_system_id, class_system_alias.Original_ind, class_system_alias.Owner_ba_id, class_system_alias.Ppdm_guid, class_system_alias.Preferred_ind, class_system_alias.Reason_desc, class_system_alias.Remark, class_system_alias.Source, class_system_alias.Source_document_id, class_system_alias.Struckoff_date, class_system_alias.Sw_application_id, class_system_alias.Use_rule, class_system_alias.Row_changed_by, class_system_alias.Row_changed_date, class_system_alias.Row_created_by, class_system_alias.Row_created_date, class_system_alias.Row_effective_date, class_system_alias.Row_expiry_date, class_system_alias.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateClassSystemAlias(c *fiber.Ctx) error {
	var class_system_alias dto.Class_system_alias

	setDefaults(&class_system_alias)

	if err := json.Unmarshal(c.Body(), &class_system_alias); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	class_system_alias.Classification_system_id = id

    if class_system_alias.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update class_system_alias set class_system_alias_id = :1, abbreviation = :2, active_ind = :3, alias_long_name = :4, alias_reason_type = :5, alias_short_name = :6, alias_type = :7, alias_version = :8, amended_date = :9, created_date = :10, effective_date = :11, expiry_date = :12, classification_system_id = :13, original_ind = :14, owner_ba_id = :15, ppdm_guid = :16, preferred_ind = :17, reason_desc = :18, remark = :19, source = :20, source_document_id = :21, struckoff_date = :22, sw_application_id = :23, use_rule = :24, row_changed_by = :25, row_changed_date = :26, row_created_by = :27, row_effective_date = :28, row_expiry_date = :29, row_quality = :30 where classification_system_id = :32")
	if err != nil {
		return err
	}

	class_system_alias.Row_changed_date = formatDate(class_system_alias.Row_changed_date)
	class_system_alias.Amended_date = formatDateString(class_system_alias.Amended_date)
	class_system_alias.Created_date = formatDateString(class_system_alias.Created_date)
	class_system_alias.Effective_date = formatDateString(class_system_alias.Effective_date)
	class_system_alias.Expiry_date = formatDateString(class_system_alias.Expiry_date)
	class_system_alias.Struckoff_date = formatDateString(class_system_alias.Struckoff_date)
	class_system_alias.Row_effective_date = formatDateString(class_system_alias.Row_effective_date)
	class_system_alias.Row_expiry_date = formatDateString(class_system_alias.Row_expiry_date)






	rows, err := stmt.Exec(class_system_alias.Class_system_alias_id, class_system_alias.Abbreviation, class_system_alias.Active_ind, class_system_alias.Alias_long_name, class_system_alias.Alias_reason_type, class_system_alias.Alias_short_name, class_system_alias.Alias_type, class_system_alias.Alias_version, class_system_alias.Amended_date, class_system_alias.Created_date, class_system_alias.Effective_date, class_system_alias.Expiry_date, class_system_alias.Classification_system_id, class_system_alias.Original_ind, class_system_alias.Owner_ba_id, class_system_alias.Ppdm_guid, class_system_alias.Preferred_ind, class_system_alias.Reason_desc, class_system_alias.Remark, class_system_alias.Source, class_system_alias.Source_document_id, class_system_alias.Struckoff_date, class_system_alias.Sw_application_id, class_system_alias.Use_rule, class_system_alias.Row_changed_by, class_system_alias.Row_changed_date, class_system_alias.Row_created_by, class_system_alias.Row_effective_date, class_system_alias.Row_expiry_date, class_system_alias.Row_quality, class_system_alias.Classification_system_id)
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

func PatchClassSystemAlias(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update class_system_alias set "
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
	query += " where classification_system_id = :id"

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

func DeleteClassSystemAlias(c *fiber.Ctx) error {
	id := c.Params("id")
	var class_system_alias dto.Class_system_alias
	class_system_alias.Classification_system_id = id

	stmt, err := db.Prepare("delete from class_system_alias where classification_system_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(class_system_alias.Classification_system_id)
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


