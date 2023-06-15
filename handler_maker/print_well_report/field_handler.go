package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_print_well_report/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetField(c *fiber.Ctx) error {
	rows, err := db.Query("select * from field")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Field

	for rows.Next() {
		var field dto.Field
		if err := rows.Scan(&field.Field_id, &field.Active_ind, &field.Area_id, &field.Area_type, &field.Discovery_date, &field.Effective_date, &field.Expiry_date, &field.Field_abbreviation, &field.Field_area_obs_no, &field.Field_name, &field.Field_type, &field.Group_field_id, &field.Ppdm_guid, &field.Remark, &field.Source, &field.Row_changed_by, &field.Row_changed_date, &field.Row_created_by, &field.Row_created_date, &field.Row_effective_date, &field.Row_expiry_date, &field.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append field to result
		result = append(result, field)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetField(c *fiber.Ctx) error {
	var field dto.Field

	setDefaults(&field)

	if err := json.Unmarshal(c.Body(), &field); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into field values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22)")
	if err != nil {
		return err
	}
	field.Row_created_date = formatDate(field.Row_created_date)
	field.Row_changed_date = formatDate(field.Row_changed_date)
	field.Discovery_date = formatDateString(field.Discovery_date)
	field.Effective_date = formatDateString(field.Effective_date)
	field.Expiry_date = formatDateString(field.Expiry_date)
	field.Row_effective_date = formatDateString(field.Row_effective_date)
	field.Row_expiry_date = formatDateString(field.Row_expiry_date)






	rows, err := stmt.Exec(field.Field_id, field.Active_ind, field.Area_id, field.Area_type, field.Discovery_date, field.Effective_date, field.Expiry_date, field.Field_abbreviation, field.Field_area_obs_no, field.Field_name, field.Field_type, field.Group_field_id, field.Ppdm_guid, field.Remark, field.Source, field.Row_changed_by, field.Row_changed_date, field.Row_created_by, field.Row_created_date, field.Row_effective_date, field.Row_expiry_date, field.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateField(c *fiber.Ctx) error {
	var field dto.Field

	setDefaults(&field)

	if err := json.Unmarshal(c.Body(), &field); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	field.Field_id = id

    if field.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update field set active_ind = :1, area_id = :2, area_type = :3, discovery_date = :4, effective_date = :5, expiry_date = :6, field_abbreviation = :7, field_area_obs_no = :8, field_name = :9, field_type = :10, group_field_id = :11, ppdm_guid = :12, remark = :13, source = :14, row_changed_by = :15, row_changed_date = :16, row_created_by = :17, row_effective_date = :18, row_expiry_date = :19, row_quality = :20 where field_id = :22")
	if err != nil {
		return err
	}

	field.Row_changed_date = formatDate(field.Row_changed_date)
	field.Discovery_date = formatDateString(field.Discovery_date)
	field.Effective_date = formatDateString(field.Effective_date)
	field.Expiry_date = formatDateString(field.Expiry_date)
	field.Row_effective_date = formatDateString(field.Row_effective_date)
	field.Row_expiry_date = formatDateString(field.Row_expiry_date)






	rows, err := stmt.Exec(field.Active_ind, field.Area_id, field.Area_type, field.Discovery_date, field.Effective_date, field.Expiry_date, field.Field_abbreviation, field.Field_area_obs_no, field.Field_name, field.Field_type, field.Group_field_id, field.Ppdm_guid, field.Remark, field.Source, field.Row_changed_by, field.Row_changed_date, field.Row_created_by, field.Row_effective_date, field.Row_expiry_date, field.Row_quality, field.Field_id)
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

func PatchField(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update field set "
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
				formattedDate := formatDate(&date)
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

func DeleteField(c *fiber.Ctx) error {
	id := c.Params("id")
	var field dto.Field
	field.Field_id = id

	stmt, err := db.Prepare("delete from field where field_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(field.Field_id)
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


