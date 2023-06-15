package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetClassSystemXref(c *fiber.Ctx) error {
	rows, err := db.Query("select * from class_system_xref")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Class_system_xref

	for rows.Next() {
		var class_system_xref dto.Class_system_xref
		if err := rows.Scan(&class_system_xref.Classification_system_id, &class_system_xref.Classification_system_id2, &class_system_xref.Equiv_id, &class_system_xref.Active_ind, &class_system_xref.Effective_date, &class_system_xref.Expiry_date, &class_system_xref.Ppdm_guid, &class_system_xref.Remark, &class_system_xref.Source, &class_system_xref.Source_document, &class_system_xref.System_xref_type, &class_system_xref.Row_changed_by, &class_system_xref.Row_changed_date, &class_system_xref.Row_created_by, &class_system_xref.Row_created_date, &class_system_xref.Row_effective_date, &class_system_xref.Row_expiry_date, &class_system_xref.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append class_system_xref to result
		result = append(result, class_system_xref)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetClassSystemXref(c *fiber.Ctx) error {
	var class_system_xref dto.Class_system_xref

	setDefaults(&class_system_xref)

	if err := json.Unmarshal(c.Body(), &class_system_xref); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into class_system_xref values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18)")
	if err != nil {
		return err
	}
	class_system_xref.Row_created_date = formatDate(class_system_xref.Row_created_date)
	class_system_xref.Row_changed_date = formatDate(class_system_xref.Row_changed_date)
	class_system_xref.Effective_date = formatDateString(class_system_xref.Effective_date)
	class_system_xref.Expiry_date = formatDateString(class_system_xref.Expiry_date)
	class_system_xref.Row_effective_date = formatDateString(class_system_xref.Row_effective_date)
	class_system_xref.Row_expiry_date = formatDateString(class_system_xref.Row_expiry_date)






	rows, err := stmt.Exec(class_system_xref.Classification_system_id, class_system_xref.Classification_system_id2, class_system_xref.Equiv_id, class_system_xref.Active_ind, class_system_xref.Effective_date, class_system_xref.Expiry_date, class_system_xref.Ppdm_guid, class_system_xref.Remark, class_system_xref.Source, class_system_xref.Source_document, class_system_xref.System_xref_type, class_system_xref.Row_changed_by, class_system_xref.Row_changed_date, class_system_xref.Row_created_by, class_system_xref.Row_created_date, class_system_xref.Row_effective_date, class_system_xref.Row_expiry_date, class_system_xref.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateClassSystemXref(c *fiber.Ctx) error {
	var class_system_xref dto.Class_system_xref

	setDefaults(&class_system_xref)

	if err := json.Unmarshal(c.Body(), &class_system_xref); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	class_system_xref.Classification_system_id = id

    if class_system_xref.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update class_system_xref set classification_system_id2 = :1, equiv_id = :2, active_ind = :3, effective_date = :4, expiry_date = :5, ppdm_guid = :6, remark = :7, source = :8, source_document = :9, system_xref_type = :10, row_changed_by = :11, row_changed_date = :12, row_created_by = :13, row_effective_date = :14, row_expiry_date = :15, row_quality = :16 where classification_system_id = :18")
	if err != nil {
		return err
	}

	class_system_xref.Row_changed_date = formatDate(class_system_xref.Row_changed_date)
	class_system_xref.Effective_date = formatDateString(class_system_xref.Effective_date)
	class_system_xref.Expiry_date = formatDateString(class_system_xref.Expiry_date)
	class_system_xref.Row_effective_date = formatDateString(class_system_xref.Row_effective_date)
	class_system_xref.Row_expiry_date = formatDateString(class_system_xref.Row_expiry_date)






	rows, err := stmt.Exec(class_system_xref.Classification_system_id2, class_system_xref.Equiv_id, class_system_xref.Active_ind, class_system_xref.Effective_date, class_system_xref.Expiry_date, class_system_xref.Ppdm_guid, class_system_xref.Remark, class_system_xref.Source, class_system_xref.Source_document, class_system_xref.System_xref_type, class_system_xref.Row_changed_by, class_system_xref.Row_changed_date, class_system_xref.Row_created_by, class_system_xref.Row_effective_date, class_system_xref.Row_expiry_date, class_system_xref.Row_quality, class_system_xref.Classification_system_id)
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

func PatchClassSystemXref(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update class_system_xref set "
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

func DeleteClassSystemXref(c *fiber.Ctx) error {
	id := c.Params("id")
	var class_system_xref dto.Class_system_xref
	class_system_xref.Classification_system_id = id

	stmt, err := db.Prepare("delete from class_system_xref where classification_system_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(class_system_xref.Classification_system_id)
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


