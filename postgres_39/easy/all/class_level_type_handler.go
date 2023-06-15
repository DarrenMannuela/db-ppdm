package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetClassLevelType(c *fiber.Ctx) error {
	rows, err := db.Query("select * from class_level_type")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Class_level_type

	for rows.Next() {
		var class_level_type dto.Class_level_type
		if err := rows.Scan(&class_level_type.Classification_system_id, &class_level_type.Level_type, &class_level_type.Abbreviation, &class_level_type.Active_ind, &class_level_type.Effective_date, &class_level_type.Expiry_date, &class_level_type.Level_order_seq_no, &class_level_type.Long_name, &class_level_type.Ppdm_guid, &class_level_type.Ppdm_system_id, &class_level_type.Ppdm_table_name, &class_level_type.Remark, &class_level_type.Selection_criteria, &class_level_type.Short_name, &class_level_type.Source, &class_level_type.Row_changed_by, &class_level_type.Row_changed_date, &class_level_type.Row_created_by, &class_level_type.Row_created_date, &class_level_type.Row_effective_date, &class_level_type.Row_expiry_date, &class_level_type.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append class_level_type to result
		result = append(result, class_level_type)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetClassLevelType(c *fiber.Ctx) error {
	var class_level_type dto.Class_level_type

	setDefaults(&class_level_type)

	if err := json.Unmarshal(c.Body(), &class_level_type); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into class_level_type values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22)")
	if err != nil {
		return err
	}
	class_level_type.Row_created_date = formatDate(class_level_type.Row_created_date)
	class_level_type.Row_changed_date = formatDate(class_level_type.Row_changed_date)
	class_level_type.Effective_date = formatDateString(class_level_type.Effective_date)
	class_level_type.Expiry_date = formatDateString(class_level_type.Expiry_date)
	class_level_type.Row_effective_date = formatDateString(class_level_type.Row_effective_date)
	class_level_type.Row_expiry_date = formatDateString(class_level_type.Row_expiry_date)






	rows, err := stmt.Exec(class_level_type.Classification_system_id, class_level_type.Level_type, class_level_type.Abbreviation, class_level_type.Active_ind, class_level_type.Effective_date, class_level_type.Expiry_date, class_level_type.Level_order_seq_no, class_level_type.Long_name, class_level_type.Ppdm_guid, class_level_type.Ppdm_system_id, class_level_type.Ppdm_table_name, class_level_type.Remark, class_level_type.Selection_criteria, class_level_type.Short_name, class_level_type.Source, class_level_type.Row_changed_by, class_level_type.Row_changed_date, class_level_type.Row_created_by, class_level_type.Row_created_date, class_level_type.Row_effective_date, class_level_type.Row_expiry_date, class_level_type.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateClassLevelType(c *fiber.Ctx) error {
	var class_level_type dto.Class_level_type

	setDefaults(&class_level_type)

	if err := json.Unmarshal(c.Body(), &class_level_type); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	class_level_type.Classification_system_id = id

    if class_level_type.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update class_level_type set level_type = :1, abbreviation = :2, active_ind = :3, effective_date = :4, expiry_date = :5, level_order_seq_no = :6, long_name = :7, ppdm_guid = :8, ppdm_system_id = :9, ppdm_table_name = :10, remark = :11, selection_criteria = :12, short_name = :13, source = :14, row_changed_by = :15, row_changed_date = :16, row_created_by = :17, row_effective_date = :18, row_expiry_date = :19, row_quality = :20 where classification_system_id = :22")
	if err != nil {
		return err
	}

	class_level_type.Row_changed_date = formatDate(class_level_type.Row_changed_date)
	class_level_type.Effective_date = formatDateString(class_level_type.Effective_date)
	class_level_type.Expiry_date = formatDateString(class_level_type.Expiry_date)
	class_level_type.Row_effective_date = formatDateString(class_level_type.Row_effective_date)
	class_level_type.Row_expiry_date = formatDateString(class_level_type.Row_expiry_date)






	rows, err := stmt.Exec(class_level_type.Level_type, class_level_type.Abbreviation, class_level_type.Active_ind, class_level_type.Effective_date, class_level_type.Expiry_date, class_level_type.Level_order_seq_no, class_level_type.Long_name, class_level_type.Ppdm_guid, class_level_type.Ppdm_system_id, class_level_type.Ppdm_table_name, class_level_type.Remark, class_level_type.Selection_criteria, class_level_type.Short_name, class_level_type.Source, class_level_type.Row_changed_by, class_level_type.Row_changed_date, class_level_type.Row_created_by, class_level_type.Row_effective_date, class_level_type.Row_expiry_date, class_level_type.Row_quality, class_level_type.Classification_system_id)
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

func PatchClassLevelType(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update class_level_type set "
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
		} else if key == "effective_date" || key == "expiry_date" || key == "row_effective_date" || key == "row_expiry_date"      {
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

func DeleteClassLevelType(c *fiber.Ctx) error {
	id := c.Params("id")
	var class_level_type dto.Class_level_type
	class_level_type.Classification_system_id = id

	stmt, err := db.Prepare("delete from class_level_type where classification_system_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(class_level_type.Classification_system_id)
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


