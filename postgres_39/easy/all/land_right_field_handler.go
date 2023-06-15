package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetLandRightField(c *fiber.Ctx) error {
	rows, err := db.Query("select * from land_right_field")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Land_right_field

	for rows.Next() {
		var land_right_field dto.Land_right_field
		if err := rows.Scan(&land_right_field.Land_right_subtype, &land_right_field.Land_right_id, &land_right_field.Field_id, &land_right_field.Active_ind, &land_right_field.Effective_date, &land_right_field.Expiry_date, &land_right_field.Ppdm_guid, &land_right_field.Remark, &land_right_field.Source, &land_right_field.Unit_operated_ind, &land_right_field.Xref_type, &land_right_field.Row_changed_by, &land_right_field.Row_changed_date, &land_right_field.Row_created_by, &land_right_field.Row_created_date, &land_right_field.Row_effective_date, &land_right_field.Row_expiry_date, &land_right_field.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append land_right_field to result
		result = append(result, land_right_field)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetLandRightField(c *fiber.Ctx) error {
	var land_right_field dto.Land_right_field

	setDefaults(&land_right_field)

	if err := json.Unmarshal(c.Body(), &land_right_field); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into land_right_field values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18)")
	if err != nil {
		return err
	}
	land_right_field.Row_created_date = formatDate(land_right_field.Row_created_date)
	land_right_field.Row_changed_date = formatDate(land_right_field.Row_changed_date)
	land_right_field.Effective_date = formatDateString(land_right_field.Effective_date)
	land_right_field.Expiry_date = formatDateString(land_right_field.Expiry_date)
	land_right_field.Row_effective_date = formatDateString(land_right_field.Row_effective_date)
	land_right_field.Row_expiry_date = formatDateString(land_right_field.Row_expiry_date)






	rows, err := stmt.Exec(land_right_field.Land_right_subtype, land_right_field.Land_right_id, land_right_field.Field_id, land_right_field.Active_ind, land_right_field.Effective_date, land_right_field.Expiry_date, land_right_field.Ppdm_guid, land_right_field.Remark, land_right_field.Source, land_right_field.Unit_operated_ind, land_right_field.Xref_type, land_right_field.Row_changed_by, land_right_field.Row_changed_date, land_right_field.Row_created_by, land_right_field.Row_created_date, land_right_field.Row_effective_date, land_right_field.Row_expiry_date, land_right_field.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateLandRightField(c *fiber.Ctx) error {
	var land_right_field dto.Land_right_field

	setDefaults(&land_right_field)

	if err := json.Unmarshal(c.Body(), &land_right_field); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	land_right_field.Land_right_subtype = id

    if land_right_field.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update land_right_field set land_right_id = :1, field_id = :2, active_ind = :3, effective_date = :4, expiry_date = :5, ppdm_guid = :6, remark = :7, source = :8, unit_operated_ind = :9, xref_type = :10, row_changed_by = :11, row_changed_date = :12, row_created_by = :13, row_effective_date = :14, row_expiry_date = :15, row_quality = :16 where land_right_subtype = :18")
	if err != nil {
		return err
	}

	land_right_field.Row_changed_date = formatDate(land_right_field.Row_changed_date)
	land_right_field.Effective_date = formatDateString(land_right_field.Effective_date)
	land_right_field.Expiry_date = formatDateString(land_right_field.Expiry_date)
	land_right_field.Row_effective_date = formatDateString(land_right_field.Row_effective_date)
	land_right_field.Row_expiry_date = formatDateString(land_right_field.Row_expiry_date)






	rows, err := stmt.Exec(land_right_field.Land_right_id, land_right_field.Field_id, land_right_field.Active_ind, land_right_field.Effective_date, land_right_field.Expiry_date, land_right_field.Ppdm_guid, land_right_field.Remark, land_right_field.Source, land_right_field.Unit_operated_ind, land_right_field.Xref_type, land_right_field.Row_changed_by, land_right_field.Row_changed_date, land_right_field.Row_created_by, land_right_field.Row_effective_date, land_right_field.Row_expiry_date, land_right_field.Row_quality, land_right_field.Land_right_subtype)
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

func PatchLandRightField(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update land_right_field set "
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
	query += " where land_right_subtype = :id"

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

func DeleteLandRightField(c *fiber.Ctx) error {
	id := c.Params("id")
	var land_right_field dto.Land_right_field
	land_right_field.Land_right_subtype = id

	stmt, err := db.Prepare("delete from land_right_field where land_right_subtype = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(land_right_field.Land_right_subtype)
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


