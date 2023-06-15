package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetClassSystem(c *fiber.Ctx) error {
	rows, err := db.Query("select * from class_system")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Class_system

	for rows.Next() {
		var class_system dto.Class_system
		if err := rows.Scan(&class_system.Classification_system_id, &class_system.Active_ind, &class_system.Class_dimension, &class_system.Effective_date, &class_system.Expiry_date, &class_system.Ppdm_guid, &class_system.Remark, &class_system.Source, &class_system.Source_document_id, &class_system.System_definition, &class_system.System_name, &class_system.System_owner, &class_system.System_ref_num, &class_system.System_version, &class_system.Row_changed_by, &class_system.Row_changed_date, &class_system.Row_created_by, &class_system.Row_created_date, &class_system.Row_effective_date, &class_system.Row_expiry_date, &class_system.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append class_system to result
		result = append(result, class_system)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetClassSystem(c *fiber.Ctx) error {
	var class_system dto.Class_system

	setDefaults(&class_system)

	if err := json.Unmarshal(c.Body(), &class_system); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into class_system values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21)")
	if err != nil {
		return err
	}
	class_system.Row_created_date = formatDate(class_system.Row_created_date)
	class_system.Row_changed_date = formatDate(class_system.Row_changed_date)
	class_system.Effective_date = formatDateString(class_system.Effective_date)
	class_system.Expiry_date = formatDateString(class_system.Expiry_date)
	class_system.Row_effective_date = formatDateString(class_system.Row_effective_date)
	class_system.Row_expiry_date = formatDateString(class_system.Row_expiry_date)






	rows, err := stmt.Exec(class_system.Classification_system_id, class_system.Active_ind, class_system.Class_dimension, class_system.Effective_date, class_system.Expiry_date, class_system.Ppdm_guid, class_system.Remark, class_system.Source, class_system.Source_document_id, class_system.System_definition, class_system.System_name, class_system.System_owner, class_system.System_ref_num, class_system.System_version, class_system.Row_changed_by, class_system.Row_changed_date, class_system.Row_created_by, class_system.Row_created_date, class_system.Row_effective_date, class_system.Row_expiry_date, class_system.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateClassSystem(c *fiber.Ctx) error {
	var class_system dto.Class_system

	setDefaults(&class_system)

	if err := json.Unmarshal(c.Body(), &class_system); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	class_system.Classification_system_id = id

    if class_system.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update class_system set active_ind = :1, class_dimension = :2, effective_date = :3, expiry_date = :4, ppdm_guid = :5, remark = :6, source = :7, source_document_id = :8, system_definition = :9, system_name = :10, system_owner = :11, system_ref_num = :12, system_version = :13, row_changed_by = :14, row_changed_date = :15, row_created_by = :16, row_effective_date = :17, row_expiry_date = :18, row_quality = :19 where classification_system_id = :21")
	if err != nil {
		return err
	}

	class_system.Row_changed_date = formatDate(class_system.Row_changed_date)
	class_system.Effective_date = formatDateString(class_system.Effective_date)
	class_system.Expiry_date = formatDateString(class_system.Expiry_date)
	class_system.Row_effective_date = formatDateString(class_system.Row_effective_date)
	class_system.Row_expiry_date = formatDateString(class_system.Row_expiry_date)






	rows, err := stmt.Exec(class_system.Active_ind, class_system.Class_dimension, class_system.Effective_date, class_system.Expiry_date, class_system.Ppdm_guid, class_system.Remark, class_system.Source, class_system.Source_document_id, class_system.System_definition, class_system.System_name, class_system.System_owner, class_system.System_ref_num, class_system.System_version, class_system.Row_changed_by, class_system.Row_changed_date, class_system.Row_created_by, class_system.Row_effective_date, class_system.Row_expiry_date, class_system.Row_quality, class_system.Classification_system_id)
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

func PatchClassSystem(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update class_system set "
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

func DeleteClassSystem(c *fiber.Ctx) error {
	id := c.Params("id")
	var class_system dto.Class_system
	class_system.Classification_system_id = id

	stmt, err := db.Prepare("delete from class_system where classification_system_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(class_system.Classification_system_id)
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


