package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetClassLevelAlias(c *fiber.Ctx) error {
	rows, err := db.Query("select * from class_level_alias")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Class_level_alias

	for rows.Next() {
		var class_level_alias dto.Class_level_alias
		if err := rows.Scan(&class_level_alias.Classification_system_id, &class_level_alias.Class_level_id, &class_level_alias.Class_level_alias_id, &class_level_alias.Abbreviation, &class_level_alias.Active_ind, &class_level_alias.Alias_long_name, &class_level_alias.Alias_reason_type, &class_level_alias.Alias_short_name, &class_level_alias.Alias_type, &class_level_alias.Alias_version, &class_level_alias.Amended_date, &class_level_alias.Column_name, &class_level_alias.Created_date, &class_level_alias.Effective_date, &class_level_alias.Expiry_date, &class_level_alias.Classification_system_id, &class_level_alias.Original_ind, &class_level_alias.Owner_ba_id, &class_level_alias.Ppdm_guid, &class_level_alias.Preferred_ind, &class_level_alias.Reason_desc, &class_level_alias.Referenced_guid, &class_level_alias.Remark, &class_level_alias.Schema_entity_id, &class_level_alias.Source, &class_level_alias.Source_document_id, &class_level_alias.Struckoff_date, &class_level_alias.Sw_application_id, &class_level_alias.System_id, &class_level_alias.Table_name, &class_level_alias.Use_rule, &class_level_alias.Row_changed_by, &class_level_alias.Row_changed_date, &class_level_alias.Row_created_by, &class_level_alias.Row_created_date, &class_level_alias.Row_effective_date, &class_level_alias.Row_expiry_date, &class_level_alias.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append class_level_alias to result
		result = append(result, class_level_alias)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetClassLevelAlias(c *fiber.Ctx) error {
	var class_level_alias dto.Class_level_alias

	setDefaults(&class_level_alias)

	if err := json.Unmarshal(c.Body(), &class_level_alias); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into class_level_alias values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32, :33, :34, :35, :36, :37, :38)")
	if err != nil {
		return err
	}
	class_level_alias.Row_created_date = formatDate(class_level_alias.Row_created_date)
	class_level_alias.Row_changed_date = formatDate(class_level_alias.Row_changed_date)
	class_level_alias.Amended_date = formatDateString(class_level_alias.Amended_date)
	class_level_alias.Created_date = formatDateString(class_level_alias.Created_date)
	class_level_alias.Effective_date = formatDateString(class_level_alias.Effective_date)
	class_level_alias.Expiry_date = formatDateString(class_level_alias.Expiry_date)
	class_level_alias.Struckoff_date = formatDateString(class_level_alias.Struckoff_date)
	class_level_alias.Row_effective_date = formatDateString(class_level_alias.Row_effective_date)
	class_level_alias.Row_expiry_date = formatDateString(class_level_alias.Row_expiry_date)






	rows, err := stmt.Exec(class_level_alias.Classification_system_id, class_level_alias.Class_level_id, class_level_alias.Class_level_alias_id, class_level_alias.Abbreviation, class_level_alias.Active_ind, class_level_alias.Alias_long_name, class_level_alias.Alias_reason_type, class_level_alias.Alias_short_name, class_level_alias.Alias_type, class_level_alias.Alias_version, class_level_alias.Amended_date, class_level_alias.Column_name, class_level_alias.Created_date, class_level_alias.Effective_date, class_level_alias.Expiry_date, class_level_alias.Classification_system_id, class_level_alias.Original_ind, class_level_alias.Owner_ba_id, class_level_alias.Ppdm_guid, class_level_alias.Preferred_ind, class_level_alias.Reason_desc, class_level_alias.Referenced_guid, class_level_alias.Remark, class_level_alias.Schema_entity_id, class_level_alias.Source, class_level_alias.Source_document_id, class_level_alias.Struckoff_date, class_level_alias.Sw_application_id, class_level_alias.System_id, class_level_alias.Table_name, class_level_alias.Use_rule, class_level_alias.Row_changed_by, class_level_alias.Row_changed_date, class_level_alias.Row_created_by, class_level_alias.Row_created_date, class_level_alias.Row_effective_date, class_level_alias.Row_expiry_date, class_level_alias.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateClassLevelAlias(c *fiber.Ctx) error {
	var class_level_alias dto.Class_level_alias

	setDefaults(&class_level_alias)

	if err := json.Unmarshal(c.Body(), &class_level_alias); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	class_level_alias.Classification_system_id = id

    if class_level_alias.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update class_level_alias set class_level_id = :1, class_level_alias_id = :2, abbreviation = :3, active_ind = :4, alias_long_name = :5, alias_reason_type = :6, alias_short_name = :7, alias_type = :8, alias_version = :9, amended_date = :10, column_name = :11, created_date = :12, effective_date = :13, expiry_date = :14, classification_system_id = :15, original_ind = :16, owner_ba_id = :17, ppdm_guid = :18, preferred_ind = :19, reason_desc = :20, referenced_guid = :21, remark = :22, schema_entity_id = :23, source = :24, source_document_id = :25, struckoff_date = :26, sw_application_id = :27, system_id = :28, table_name = :29, use_rule = :30, row_changed_by = :31, row_changed_date = :32, row_created_by = :33, row_effective_date = :34, row_expiry_date = :35, row_quality = :36 where classification_system_id = :38")
	if err != nil {
		return err
	}

	class_level_alias.Row_changed_date = formatDate(class_level_alias.Row_changed_date)
	class_level_alias.Amended_date = formatDateString(class_level_alias.Amended_date)
	class_level_alias.Created_date = formatDateString(class_level_alias.Created_date)
	class_level_alias.Effective_date = formatDateString(class_level_alias.Effective_date)
	class_level_alias.Expiry_date = formatDateString(class_level_alias.Expiry_date)
	class_level_alias.Struckoff_date = formatDateString(class_level_alias.Struckoff_date)
	class_level_alias.Row_effective_date = formatDateString(class_level_alias.Row_effective_date)
	class_level_alias.Row_expiry_date = formatDateString(class_level_alias.Row_expiry_date)






	rows, err := stmt.Exec(class_level_alias.Class_level_id, class_level_alias.Class_level_alias_id, class_level_alias.Abbreviation, class_level_alias.Active_ind, class_level_alias.Alias_long_name, class_level_alias.Alias_reason_type, class_level_alias.Alias_short_name, class_level_alias.Alias_type, class_level_alias.Alias_version, class_level_alias.Amended_date, class_level_alias.Column_name, class_level_alias.Created_date, class_level_alias.Effective_date, class_level_alias.Expiry_date, class_level_alias.Classification_system_id, class_level_alias.Original_ind, class_level_alias.Owner_ba_id, class_level_alias.Ppdm_guid, class_level_alias.Preferred_ind, class_level_alias.Reason_desc, class_level_alias.Referenced_guid, class_level_alias.Remark, class_level_alias.Schema_entity_id, class_level_alias.Source, class_level_alias.Source_document_id, class_level_alias.Struckoff_date, class_level_alias.Sw_application_id, class_level_alias.System_id, class_level_alias.Table_name, class_level_alias.Use_rule, class_level_alias.Row_changed_by, class_level_alias.Row_changed_date, class_level_alias.Row_created_by, class_level_alias.Row_effective_date, class_level_alias.Row_expiry_date, class_level_alias.Row_quality, class_level_alias.Classification_system_id)
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

func PatchClassLevelAlias(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update class_level_alias set "
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

func DeleteClassLevelAlias(c *fiber.Ctx) error {
	id := c.Params("id")
	var class_level_alias dto.Class_level_alias
	class_level_alias.Classification_system_id = id

	stmt, err := db.Prepare("delete from class_level_alias where classification_system_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(class_level_alias.Classification_system_id)
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


