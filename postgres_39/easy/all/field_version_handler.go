package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetFieldVersion(c *fiber.Ctx) error {
	rows, err := db.Query("select * from field_version")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Field_version

	for rows.Next() {
		var field_version dto.Field_version
		if err := rows.Scan(&field_version.Field_id, &field_version.Source, &field_version.Active_ind, &field_version.Area_id, &field_version.Area_type, &field_version.Discovery_date, &field_version.Effective_date, &field_version.Expiry_date, &field_version.Field_abbreviation, &field_version.Field_area_obs_no, &field_version.Field_name, &field_version.Field_type, &field_version.Ppdm_guid, &field_version.Remark, &field_version.Row_changed_by, &field_version.Row_changed_date, &field_version.Row_created_by, &field_version.Row_created_date, &field_version.Row_effective_date, &field_version.Row_expiry_date, &field_version.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append field_version to result
		result = append(result, field_version)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetFieldVersion(c *fiber.Ctx) error {
	var field_version dto.Field_version

	setDefaults(&field_version)

	if err := json.Unmarshal(c.Body(), &field_version); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into field_version values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21)")
	if err != nil {
		return err
	}
	field_version.Row_created_date = formatDate(field_version.Row_created_date)
	field_version.Row_changed_date = formatDate(field_version.Row_changed_date)
	field_version.Discovery_date = formatDateString(field_version.Discovery_date)
	field_version.Effective_date = formatDateString(field_version.Effective_date)
	field_version.Expiry_date = formatDateString(field_version.Expiry_date)
	field_version.Row_effective_date = formatDateString(field_version.Row_effective_date)
	field_version.Row_expiry_date = formatDateString(field_version.Row_expiry_date)






	rows, err := stmt.Exec(field_version.Field_id, field_version.Source, field_version.Active_ind, field_version.Area_id, field_version.Area_type, field_version.Discovery_date, field_version.Effective_date, field_version.Expiry_date, field_version.Field_abbreviation, field_version.Field_area_obs_no, field_version.Field_name, field_version.Field_type, field_version.Ppdm_guid, field_version.Remark, field_version.Row_changed_by, field_version.Row_changed_date, field_version.Row_created_by, field_version.Row_created_date, field_version.Row_effective_date, field_version.Row_expiry_date, field_version.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateFieldVersion(c *fiber.Ctx) error {
	var field_version dto.Field_version

	setDefaults(&field_version)

	if err := json.Unmarshal(c.Body(), &field_version); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	field_version.Field_id = id

    if field_version.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update field_version set source = :1, active_ind = :2, area_id = :3, area_type = :4, discovery_date = :5, effective_date = :6, expiry_date = :7, field_abbreviation = :8, field_area_obs_no = :9, field_name = :10, field_type = :11, ppdm_guid = :12, remark = :13, row_changed_by = :14, row_changed_date = :15, row_created_by = :16, row_effective_date = :17, row_expiry_date = :18, row_quality = :19 where field_id = :21")
	if err != nil {
		return err
	}

	field_version.Row_changed_date = formatDate(field_version.Row_changed_date)
	field_version.Discovery_date = formatDateString(field_version.Discovery_date)
	field_version.Effective_date = formatDateString(field_version.Effective_date)
	field_version.Expiry_date = formatDateString(field_version.Expiry_date)
	field_version.Row_effective_date = formatDateString(field_version.Row_effective_date)
	field_version.Row_expiry_date = formatDateString(field_version.Row_expiry_date)






	rows, err := stmt.Exec(field_version.Source, field_version.Active_ind, field_version.Area_id, field_version.Area_type, field_version.Discovery_date, field_version.Effective_date, field_version.Expiry_date, field_version.Field_abbreviation, field_version.Field_area_obs_no, field_version.Field_name, field_version.Field_type, field_version.Ppdm_guid, field_version.Remark, field_version.Row_changed_by, field_version.Row_changed_date, field_version.Row_created_by, field_version.Row_effective_date, field_version.Row_expiry_date, field_version.Row_quality, field_version.Field_id)
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

func PatchFieldVersion(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update field_version set "
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
		} else if key == "discovery_date" || key == "effective_date" || key == "expiry_date" || key == "row_effective_date" || key == "row_expiry_date"      {
			if date, ok := value.(string); ok && len(date) > 0 {
				formattedDate := formatDateString(&date)
				value = formattedDate
			}
		}
		values = append(values, value)
	}
	query += " where field_id = :id"

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

func DeleteFieldVersion(c *fiber.Ctx) error {
	id := c.Params("id")
	var field_version dto.Field_version
	field_version.Field_id = id

	stmt, err := db.Prepare("delete from field_version where field_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(field_version.Field_id)
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


