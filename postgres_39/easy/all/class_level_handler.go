package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetClassLevel(c *fiber.Ctx) error {
	rows, err := db.Query("select * from class_level")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Class_level

	for rows.Next() {
		var class_level dto.Class_level
		if err := rows.Scan(&class_level.Classification_system_id, &class_level.Class_level_id, &class_level.Active_ind, &class_level.Effective_date, &class_level.Expiry_date, &class_level.Level_definition, &class_level.Level_name, &class_level.Level_ref_num, &class_level.Level_seq_no, &class_level.Level_type, &class_level.Ppdm_guid, &class_level.Remark, &class_level.Retention_period, &class_level.Source, &class_level.Row_changed_by, &class_level.Row_changed_date, &class_level.Row_created_by, &class_level.Row_created_date, &class_level.Row_effective_date, &class_level.Row_expiry_date, &class_level.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append class_level to result
		result = append(result, class_level)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetClassLevel(c *fiber.Ctx) error {
	var class_level dto.Class_level

	setDefaults(&class_level)

	if err := json.Unmarshal(c.Body(), &class_level); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into class_level values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21)")
	if err != nil {
		return err
	}
	class_level.Row_created_date = formatDate(class_level.Row_created_date)
	class_level.Row_changed_date = formatDate(class_level.Row_changed_date)
	class_level.Effective_date = formatDateString(class_level.Effective_date)
	class_level.Expiry_date = formatDateString(class_level.Expiry_date)
	class_level.Row_effective_date = formatDateString(class_level.Row_effective_date)
	class_level.Row_expiry_date = formatDateString(class_level.Row_expiry_date)






	rows, err := stmt.Exec(class_level.Classification_system_id, class_level.Class_level_id, class_level.Active_ind, class_level.Effective_date, class_level.Expiry_date, class_level.Level_definition, class_level.Level_name, class_level.Level_ref_num, class_level.Level_seq_no, class_level.Level_type, class_level.Ppdm_guid, class_level.Remark, class_level.Retention_period, class_level.Source, class_level.Row_changed_by, class_level.Row_changed_date, class_level.Row_created_by, class_level.Row_created_date, class_level.Row_effective_date, class_level.Row_expiry_date, class_level.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateClassLevel(c *fiber.Ctx) error {
	var class_level dto.Class_level

	setDefaults(&class_level)

	if err := json.Unmarshal(c.Body(), &class_level); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	class_level.Classification_system_id = id

    if class_level.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update class_level set class_level_id = :1, active_ind = :2, effective_date = :3, expiry_date = :4, level_definition = :5, level_name = :6, level_ref_num = :7, level_seq_no = :8, level_type = :9, ppdm_guid = :10, remark = :11, retention_period = :12, source = :13, row_changed_by = :14, row_changed_date = :15, row_created_by = :16, row_effective_date = :17, row_expiry_date = :18, row_quality = :19 where classification_system_id = :21")
	if err != nil {
		return err
	}

	class_level.Row_changed_date = formatDate(class_level.Row_changed_date)
	class_level.Effective_date = formatDateString(class_level.Effective_date)
	class_level.Expiry_date = formatDateString(class_level.Expiry_date)
	class_level.Row_effective_date = formatDateString(class_level.Row_effective_date)
	class_level.Row_expiry_date = formatDateString(class_level.Row_expiry_date)






	rows, err := stmt.Exec(class_level.Class_level_id, class_level.Active_ind, class_level.Effective_date, class_level.Expiry_date, class_level.Level_definition, class_level.Level_name, class_level.Level_ref_num, class_level.Level_seq_no, class_level.Level_type, class_level.Ppdm_guid, class_level.Remark, class_level.Retention_period, class_level.Source, class_level.Row_changed_by, class_level.Row_changed_date, class_level.Row_created_by, class_level.Row_effective_date, class_level.Row_expiry_date, class_level.Row_quality, class_level.Classification_system_id)
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

func PatchClassLevel(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update class_level set "
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

func DeleteClassLevel(c *fiber.Ctx) error {
	id := c.Params("id")
	var class_level dto.Class_level
	class_level.Classification_system_id = id

	stmt, err := db.Prepare("delete from class_level where classification_system_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(class_level.Classification_system_id)
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


