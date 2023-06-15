package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetClassLevelDesc(c *fiber.Ctx) error {
	rows, err := db.Query("select * from class_level_desc")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Class_level_desc

	for rows.Next() {
		var class_level_desc dto.Class_level_desc
		if err := rows.Scan(&class_level_desc.Classification_system_id, &class_level_desc.Class_level_id, &class_level_desc.Desc_obs_no, &class_level_desc.Active_ind, &class_level_desc.Average_value, &class_level_desc.Average_value_ouom, &class_level_desc.Average_value_uom, &class_level_desc.Date_format_desc, &class_level_desc.Description, &class_level_desc.Desc_code, &class_level_desc.Desc_type, &class_level_desc.Effective_date, &class_level_desc.Expiry_date, &class_level_desc.Max_date, &class_level_desc.Max_value, &class_level_desc.Max_value_ouom, &class_level_desc.Max_value_uom, &class_level_desc.Min_date, &class_level_desc.Min_value, &class_level_desc.Min_value_ouom, &class_level_desc.Min_value_uom, &class_level_desc.Ppdm_guid, &class_level_desc.Remark, &class_level_desc.Source, &class_level_desc.Row_changed_by, &class_level_desc.Row_changed_date, &class_level_desc.Row_created_by, &class_level_desc.Row_created_date, &class_level_desc.Row_effective_date, &class_level_desc.Row_expiry_date, &class_level_desc.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append class_level_desc to result
		result = append(result, class_level_desc)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetClassLevelDesc(c *fiber.Ctx) error {
	var class_level_desc dto.Class_level_desc

	setDefaults(&class_level_desc)

	if err := json.Unmarshal(c.Body(), &class_level_desc); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into class_level_desc values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31)")
	if err != nil {
		return err
	}
	class_level_desc.Row_created_date = formatDate(class_level_desc.Row_created_date)
	class_level_desc.Row_changed_date = formatDate(class_level_desc.Row_changed_date)
	class_level_desc.Date_format_desc = formatDateString(class_level_desc.Date_format_desc)
	class_level_desc.Effective_date = formatDateString(class_level_desc.Effective_date)
	class_level_desc.Expiry_date = formatDateString(class_level_desc.Expiry_date)
	class_level_desc.Max_date = formatDateString(class_level_desc.Max_date)
	class_level_desc.Min_date = formatDateString(class_level_desc.Min_date)
	class_level_desc.Row_effective_date = formatDateString(class_level_desc.Row_effective_date)
	class_level_desc.Row_expiry_date = formatDateString(class_level_desc.Row_expiry_date)






	rows, err := stmt.Exec(class_level_desc.Classification_system_id, class_level_desc.Class_level_id, class_level_desc.Desc_obs_no, class_level_desc.Active_ind, class_level_desc.Average_value, class_level_desc.Average_value_ouom, class_level_desc.Average_value_uom, class_level_desc.Date_format_desc, class_level_desc.Description, class_level_desc.Desc_code, class_level_desc.Desc_type, class_level_desc.Effective_date, class_level_desc.Expiry_date, class_level_desc.Max_date, class_level_desc.Max_value, class_level_desc.Max_value_ouom, class_level_desc.Max_value_uom, class_level_desc.Min_date, class_level_desc.Min_value, class_level_desc.Min_value_ouom, class_level_desc.Min_value_uom, class_level_desc.Ppdm_guid, class_level_desc.Remark, class_level_desc.Source, class_level_desc.Row_changed_by, class_level_desc.Row_changed_date, class_level_desc.Row_created_by, class_level_desc.Row_created_date, class_level_desc.Row_effective_date, class_level_desc.Row_expiry_date, class_level_desc.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateClassLevelDesc(c *fiber.Ctx) error {
	var class_level_desc dto.Class_level_desc

	setDefaults(&class_level_desc)

	if err := json.Unmarshal(c.Body(), &class_level_desc); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	class_level_desc.Classification_system_id = id

    if class_level_desc.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update class_level_desc set class_level_id = :1, desc_obs_no = :2, active_ind = :3, average_value = :4, average_value_ouom = :5, average_value_uom = :6, date_format_desc = :7, description = :8, desc_code = :9, desc_type = :10, effective_date = :11, expiry_date = :12, max_date = :13, max_value = :14, max_value_ouom = :15, max_value_uom = :16, min_date = :17, min_value = :18, min_value_ouom = :19, min_value_uom = :20, ppdm_guid = :21, remark = :22, source = :23, row_changed_by = :24, row_changed_date = :25, row_created_by = :26, row_effective_date = :27, row_expiry_date = :28, row_quality = :29 where classification_system_id = :31")
	if err != nil {
		return err
	}

	class_level_desc.Row_changed_date = formatDate(class_level_desc.Row_changed_date)
	class_level_desc.Date_format_desc = formatDateString(class_level_desc.Date_format_desc)
	class_level_desc.Effective_date = formatDateString(class_level_desc.Effective_date)
	class_level_desc.Expiry_date = formatDateString(class_level_desc.Expiry_date)
	class_level_desc.Max_date = formatDateString(class_level_desc.Max_date)
	class_level_desc.Min_date = formatDateString(class_level_desc.Min_date)
	class_level_desc.Row_effective_date = formatDateString(class_level_desc.Row_effective_date)
	class_level_desc.Row_expiry_date = formatDateString(class_level_desc.Row_expiry_date)






	rows, err := stmt.Exec(class_level_desc.Class_level_id, class_level_desc.Desc_obs_no, class_level_desc.Active_ind, class_level_desc.Average_value, class_level_desc.Average_value_ouom, class_level_desc.Average_value_uom, class_level_desc.Date_format_desc, class_level_desc.Description, class_level_desc.Desc_code, class_level_desc.Desc_type, class_level_desc.Effective_date, class_level_desc.Expiry_date, class_level_desc.Max_date, class_level_desc.Max_value, class_level_desc.Max_value_ouom, class_level_desc.Max_value_uom, class_level_desc.Min_date, class_level_desc.Min_value, class_level_desc.Min_value_ouom, class_level_desc.Min_value_uom, class_level_desc.Ppdm_guid, class_level_desc.Remark, class_level_desc.Source, class_level_desc.Row_changed_by, class_level_desc.Row_changed_date, class_level_desc.Row_created_by, class_level_desc.Row_effective_date, class_level_desc.Row_expiry_date, class_level_desc.Row_quality, class_level_desc.Classification_system_id)
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

func PatchClassLevelDesc(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update class_level_desc set "
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
		} else if key == "date_format_desc" || key == "effective_date" || key == "expiry_date" || key == "max_date" || key == "min_date" || key == "row_effective_date" || key == "row_expiry_date"      {
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

func DeleteClassLevelDesc(c *fiber.Ctx) error {
	id := c.Params("id")
	var class_level_desc dto.Class_level_desc
	class_level_desc.Classification_system_id = id

	stmt, err := db.Prepare("delete from class_level_desc where classification_system_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(class_level_desc.Classification_system_id)
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


