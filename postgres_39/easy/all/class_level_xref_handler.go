package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetClassLevelXref(c *fiber.Ctx) error {
	rows, err := db.Query("select * from class_level_xref")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Class_level_xref

	for rows.Next() {
		var class_level_xref dto.Class_level_xref
		if err := rows.Scan(&class_level_xref.Classification_system_id, &class_level_xref.Class_level_id, &class_level_xref.Classification_system_id2, &class_level_xref.Class_level_id2, &class_level_xref.Xref_id, &class_level_xref.Active_ind, &class_level_xref.Effective_date, &class_level_xref.Expiry_date, &class_level_xref.Level_xref_type, &class_level_xref.Ppdm_guid, &class_level_xref.Remark, &class_level_xref.Source, &class_level_xref.Source_document_id, &class_level_xref.Row_changed_by, &class_level_xref.Row_changed_date, &class_level_xref.Row_created_by, &class_level_xref.Row_created_date, &class_level_xref.Row_effective_date, &class_level_xref.Row_expiry_date, &class_level_xref.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append class_level_xref to result
		result = append(result, class_level_xref)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetClassLevelXref(c *fiber.Ctx) error {
	var class_level_xref dto.Class_level_xref

	setDefaults(&class_level_xref)

	if err := json.Unmarshal(c.Body(), &class_level_xref); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into class_level_xref values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20)")
	if err != nil {
		return err
	}
	class_level_xref.Row_created_date = formatDate(class_level_xref.Row_created_date)
	class_level_xref.Row_changed_date = formatDate(class_level_xref.Row_changed_date)
	class_level_xref.Effective_date = formatDateString(class_level_xref.Effective_date)
	class_level_xref.Expiry_date = formatDateString(class_level_xref.Expiry_date)
	class_level_xref.Row_effective_date = formatDateString(class_level_xref.Row_effective_date)
	class_level_xref.Row_expiry_date = formatDateString(class_level_xref.Row_expiry_date)






	rows, err := stmt.Exec(class_level_xref.Classification_system_id, class_level_xref.Class_level_id, class_level_xref.Classification_system_id2, class_level_xref.Class_level_id2, class_level_xref.Xref_id, class_level_xref.Active_ind, class_level_xref.Effective_date, class_level_xref.Expiry_date, class_level_xref.Level_xref_type, class_level_xref.Ppdm_guid, class_level_xref.Remark, class_level_xref.Source, class_level_xref.Source_document_id, class_level_xref.Row_changed_by, class_level_xref.Row_changed_date, class_level_xref.Row_created_by, class_level_xref.Row_created_date, class_level_xref.Row_effective_date, class_level_xref.Row_expiry_date, class_level_xref.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateClassLevelXref(c *fiber.Ctx) error {
	var class_level_xref dto.Class_level_xref

	setDefaults(&class_level_xref)

	if err := json.Unmarshal(c.Body(), &class_level_xref); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	class_level_xref.Classification_system_id = id

    if class_level_xref.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update class_level_xref set class_level_id = :1, classification_system_id2 = :2, class_level_id2 = :3, xref_id = :4, active_ind = :5, effective_date = :6, expiry_date = :7, level_xref_type = :8, ppdm_guid = :9, remark = :10, source = :11, source_document_id = :12, row_changed_by = :13, row_changed_date = :14, row_created_by = :15, row_effective_date = :16, row_expiry_date = :17, row_quality = :18 where classification_system_id = :20")
	if err != nil {
		return err
	}

	class_level_xref.Row_changed_date = formatDate(class_level_xref.Row_changed_date)
	class_level_xref.Effective_date = formatDateString(class_level_xref.Effective_date)
	class_level_xref.Expiry_date = formatDateString(class_level_xref.Expiry_date)
	class_level_xref.Row_effective_date = formatDateString(class_level_xref.Row_effective_date)
	class_level_xref.Row_expiry_date = formatDateString(class_level_xref.Row_expiry_date)






	rows, err := stmt.Exec(class_level_xref.Class_level_id, class_level_xref.Classification_system_id2, class_level_xref.Class_level_id2, class_level_xref.Xref_id, class_level_xref.Active_ind, class_level_xref.Effective_date, class_level_xref.Expiry_date, class_level_xref.Level_xref_type, class_level_xref.Ppdm_guid, class_level_xref.Remark, class_level_xref.Source, class_level_xref.Source_document_id, class_level_xref.Row_changed_by, class_level_xref.Row_changed_date, class_level_xref.Row_created_by, class_level_xref.Row_effective_date, class_level_xref.Row_expiry_date, class_level_xref.Row_quality, class_level_xref.Classification_system_id)
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

func PatchClassLevelXref(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update class_level_xref set "
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

func DeleteClassLevelXref(c *fiber.Ctx) error {
	id := c.Params("id")
	var class_level_xref dto.Class_level_xref
	class_level_xref.Classification_system_id = id

	stmt, err := db.Prepare("delete from class_level_xref where classification_system_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(class_level_xref.Classification_system_id)
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


