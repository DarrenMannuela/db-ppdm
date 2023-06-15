package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetFacilityField(c *fiber.Ctx) error {
	rows, err := db.Query("select * from facility_field")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Facility_field

	for rows.Next() {
		var facility_field dto.Facility_field
		if err := rows.Scan(&facility_field.Facility_id, &facility_field.Facility_type, &facility_field.Field_id, &facility_field.Active_ind, &facility_field.Effective_date, &facility_field.Expiry_date, &facility_field.Ppdm_guid, &facility_field.Remark, &facility_field.Source, &facility_field.Row_changed_by, &facility_field.Row_changed_date, &facility_field.Row_created_by, &facility_field.Row_created_date, &facility_field.Row_effective_date, &facility_field.Row_expiry_date, &facility_field.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append facility_field to result
		result = append(result, facility_field)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetFacilityField(c *fiber.Ctx) error {
	var facility_field dto.Facility_field

	setDefaults(&facility_field)

	if err := json.Unmarshal(c.Body(), &facility_field); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into facility_field values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16)")
	if err != nil {
		return err
	}
	facility_field.Row_created_date = formatDate(facility_field.Row_created_date)
	facility_field.Row_changed_date = formatDate(facility_field.Row_changed_date)
	facility_field.Effective_date = formatDateString(facility_field.Effective_date)
	facility_field.Expiry_date = formatDateString(facility_field.Expiry_date)
	facility_field.Row_effective_date = formatDateString(facility_field.Row_effective_date)
	facility_field.Row_expiry_date = formatDateString(facility_field.Row_expiry_date)






	rows, err := stmt.Exec(facility_field.Facility_id, facility_field.Facility_type, facility_field.Field_id, facility_field.Active_ind, facility_field.Effective_date, facility_field.Expiry_date, facility_field.Ppdm_guid, facility_field.Remark, facility_field.Source, facility_field.Row_changed_by, facility_field.Row_changed_date, facility_field.Row_created_by, facility_field.Row_created_date, facility_field.Row_effective_date, facility_field.Row_expiry_date, facility_field.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateFacilityField(c *fiber.Ctx) error {
	var facility_field dto.Facility_field

	setDefaults(&facility_field)

	if err := json.Unmarshal(c.Body(), &facility_field); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	facility_field.Facility_id = id

    if facility_field.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update facility_field set facility_type = :1, field_id = :2, active_ind = :3, effective_date = :4, expiry_date = :5, ppdm_guid = :6, remark = :7, source = :8, row_changed_by = :9, row_changed_date = :10, row_created_by = :11, row_effective_date = :12, row_expiry_date = :13, row_quality = :14 where facility_id = :16")
	if err != nil {
		return err
	}

	facility_field.Row_changed_date = formatDate(facility_field.Row_changed_date)
	facility_field.Effective_date = formatDateString(facility_field.Effective_date)
	facility_field.Expiry_date = formatDateString(facility_field.Expiry_date)
	facility_field.Row_effective_date = formatDateString(facility_field.Row_effective_date)
	facility_field.Row_expiry_date = formatDateString(facility_field.Row_expiry_date)






	rows, err := stmt.Exec(facility_field.Facility_type, facility_field.Field_id, facility_field.Active_ind, facility_field.Effective_date, facility_field.Expiry_date, facility_field.Ppdm_guid, facility_field.Remark, facility_field.Source, facility_field.Row_changed_by, facility_field.Row_changed_date, facility_field.Row_created_by, facility_field.Row_effective_date, facility_field.Row_expiry_date, facility_field.Row_quality, facility_field.Facility_id)
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

func PatchFacilityField(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update facility_field set "
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
	query += " where facility_id = :id"

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

func DeleteFacilityField(c *fiber.Ctx) error {
	id := c.Params("id")
	var facility_field dto.Facility_field
	facility_field.Facility_id = id

	stmt, err := db.Prepare("delete from facility_field where facility_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(facility_field.Facility_id)
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


