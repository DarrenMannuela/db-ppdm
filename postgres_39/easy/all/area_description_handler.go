package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetAreaDescription(c *fiber.Ctx) error {
	rows, err := db.Query("select * from area_description")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Area_description

	for rows.Next() {
		var area_description dto.Area_description
		if err := rows.Scan(&area_description.Area_id, &area_description.Area_type, &area_description.Description_obs_no, &area_description.Active_ind, &area_description.Area_description_date, &area_description.Area_desc_code, &area_description.Area_desc_type, &area_description.Average_value, &area_description.Average_value_ouom, &area_description.Average_value_uom, &area_description.Date_format_desc, &area_description.Description, &area_description.Effective_date, &area_description.Expiry_date, &area_description.Max_value, &area_description.Max_value_ouom, &area_description.Max_value_uom, &area_description.Min_value, &area_description.Min_value_ouom, &area_description.Min_value_uom, &area_description.Ppdm_guid, &area_description.Remark, &area_description.Source, &area_description.Source_document_id, &area_description.Row_changed_by, &area_description.Row_changed_date, &area_description.Row_created_by, &area_description.Row_created_date, &area_description.Row_effective_date, &area_description.Row_expiry_date, &area_description.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append area_description to result
		result = append(result, area_description)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetAreaDescription(c *fiber.Ctx) error {
	var area_description dto.Area_description

	setDefaults(&area_description)

	if err := json.Unmarshal(c.Body(), &area_description); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into area_description values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31)")
	if err != nil {
		return err
	}
	area_description.Row_created_date = formatDate(area_description.Row_created_date)
	area_description.Row_changed_date = formatDate(area_description.Row_changed_date)
	area_description.Area_description_date = formatDateString(area_description.Area_description_date)
	area_description.Date_format_desc = formatDateString(area_description.Date_format_desc)
	area_description.Effective_date = formatDateString(area_description.Effective_date)
	area_description.Expiry_date = formatDateString(area_description.Expiry_date)
	area_description.Row_effective_date = formatDateString(area_description.Row_effective_date)
	area_description.Row_expiry_date = formatDateString(area_description.Row_expiry_date)






	rows, err := stmt.Exec(area_description.Area_id, area_description.Area_type, area_description.Description_obs_no, area_description.Active_ind, area_description.Area_description_date, area_description.Area_desc_code, area_description.Area_desc_type, area_description.Average_value, area_description.Average_value_ouom, area_description.Average_value_uom, area_description.Date_format_desc, area_description.Description, area_description.Effective_date, area_description.Expiry_date, area_description.Max_value, area_description.Max_value_ouom, area_description.Max_value_uom, area_description.Min_value, area_description.Min_value_ouom, area_description.Min_value_uom, area_description.Ppdm_guid, area_description.Remark, area_description.Source, area_description.Source_document_id, area_description.Row_changed_by, area_description.Row_changed_date, area_description.Row_created_by, area_description.Row_created_date, area_description.Row_effective_date, area_description.Row_expiry_date, area_description.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateAreaDescription(c *fiber.Ctx) error {
	var area_description dto.Area_description

	setDefaults(&area_description)

	if err := json.Unmarshal(c.Body(), &area_description); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	area_description.Area_id = id

    if area_description.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update area_description set area_type = :1, description_obs_no = :2, active_ind = :3, area_description_date = :4, area_desc_code = :5, area_desc_type = :6, average_value = :7, average_value_ouom = :8, average_value_uom = :9, date_format_desc = :10, description = :11, effective_date = :12, expiry_date = :13, max_value = :14, max_value_ouom = :15, max_value_uom = :16, min_value = :17, min_value_ouom = :18, min_value_uom = :19, ppdm_guid = :20, remark = :21, source = :22, source_document_id = :23, row_changed_by = :24, row_changed_date = :25, row_created_by = :26, row_effective_date = :27, row_expiry_date = :28, row_quality = :29 where area_id = :31")
	if err != nil {
		return err
	}

	area_description.Row_changed_date = formatDate(area_description.Row_changed_date)
	area_description.Area_description_date = formatDateString(area_description.Area_description_date)
	area_description.Date_format_desc = formatDateString(area_description.Date_format_desc)
	area_description.Effective_date = formatDateString(area_description.Effective_date)
	area_description.Expiry_date = formatDateString(area_description.Expiry_date)
	area_description.Row_effective_date = formatDateString(area_description.Row_effective_date)
	area_description.Row_expiry_date = formatDateString(area_description.Row_expiry_date)






	rows, err := stmt.Exec(area_description.Area_type, area_description.Description_obs_no, area_description.Active_ind, area_description.Area_description_date, area_description.Area_desc_code, area_description.Area_desc_type, area_description.Average_value, area_description.Average_value_ouom, area_description.Average_value_uom, area_description.Date_format_desc, area_description.Description, area_description.Effective_date, area_description.Expiry_date, area_description.Max_value, area_description.Max_value_ouom, area_description.Max_value_uom, area_description.Min_value, area_description.Min_value_ouom, area_description.Min_value_uom, area_description.Ppdm_guid, area_description.Remark, area_description.Source, area_description.Source_document_id, area_description.Row_changed_by, area_description.Row_changed_date, area_description.Row_created_by, area_description.Row_effective_date, area_description.Row_expiry_date, area_description.Row_quality, area_description.Area_id)
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

func PatchAreaDescription(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update area_description set "
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
		} else if key == "area_description_date" || key == "date_format_desc" || key == "effective_date" || key == "expiry_date" || key == "row_effective_date" || key == "row_expiry_date"      {
			if date, ok := value.(string); ok && len(date) > 0 {
				formattedDate := formatDateString(&date)
				value = formattedDate
			}
		}
		values = append(values, value)
	}
	query += " where area_id = :id"

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

func DeleteAreaDescription(c *fiber.Ctx) error {
	id := c.Params("id")
	var area_description dto.Area_description
	area_description.Area_id = id

	stmt, err := db.Prepare("delete from area_description where area_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(area_description.Area_id)
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


