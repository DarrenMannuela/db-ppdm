package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetRMediaType(c *fiber.Ctx) error {
	rows, err := db.Query("select * from r_media_type")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.R_media_type

	for rows.Next() {
		var r_media_type dto.R_media_type
		if err := rows.Scan(&r_media_type.Media_type, &r_media_type.Abbreviation, &r_media_type.Active_ind, &r_media_type.Digital_capacity, &r_media_type.Digital_capacity_uom, &r_media_type.Digital_density, &r_media_type.Digital_density_uom, &r_media_type.Effective_date, &r_media_type.Expiry_date, &r_media_type.Long_name, &r_media_type.Ppdm_guid, &r_media_type.Remark, &r_media_type.Short_name, &r_media_type.Source, &r_media_type.Row_changed_by, &r_media_type.Row_changed_date, &r_media_type.Row_created_by, &r_media_type.Row_created_date, &r_media_type.Row_effective_date, &r_media_type.Row_expiry_date, &r_media_type.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append r_media_type to result
		result = append(result, r_media_type)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetRMediaType(c *fiber.Ctx) error {
	var r_media_type dto.R_media_type

	setDefaults(&r_media_type)

	if err := json.Unmarshal(c.Body(), &r_media_type); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into r_media_type values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21)")
	if err != nil {
		return err
	}
	r_media_type.Row_created_date = formatDate(r_media_type.Row_created_date)
	r_media_type.Row_changed_date = formatDate(r_media_type.Row_changed_date)
	r_media_type.Effective_date = formatDateString(r_media_type.Effective_date)
	r_media_type.Expiry_date = formatDateString(r_media_type.Expiry_date)
	r_media_type.Row_effective_date = formatDateString(r_media_type.Row_effective_date)
	r_media_type.Row_expiry_date = formatDateString(r_media_type.Row_expiry_date)






	rows, err := stmt.Exec(r_media_type.Media_type, r_media_type.Abbreviation, r_media_type.Active_ind, r_media_type.Digital_capacity, r_media_type.Digital_capacity_uom, r_media_type.Digital_density, r_media_type.Digital_density_uom, r_media_type.Effective_date, r_media_type.Expiry_date, r_media_type.Long_name, r_media_type.Ppdm_guid, r_media_type.Remark, r_media_type.Short_name, r_media_type.Source, r_media_type.Row_changed_by, r_media_type.Row_changed_date, r_media_type.Row_created_by, r_media_type.Row_created_date, r_media_type.Row_effective_date, r_media_type.Row_expiry_date, r_media_type.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateRMediaType(c *fiber.Ctx) error {
	var r_media_type dto.R_media_type

	setDefaults(&r_media_type)

	if err := json.Unmarshal(c.Body(), &r_media_type); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	r_media_type.Media_type = id

    if r_media_type.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update r_media_type set abbreviation = :1, active_ind = :2, digital_capacity = :3, digital_capacity_uom = :4, digital_density = :5, digital_density_uom = :6, effective_date = :7, expiry_date = :8, long_name = :9, ppdm_guid = :10, remark = :11, short_name = :12, source = :13, row_changed_by = :14, row_changed_date = :15, row_created_by = :16, row_effective_date = :17, row_expiry_date = :18, row_quality = :19 where media_type = :21")
	if err != nil {
		return err
	}

	r_media_type.Row_changed_date = formatDate(r_media_type.Row_changed_date)
	r_media_type.Effective_date = formatDateString(r_media_type.Effective_date)
	r_media_type.Expiry_date = formatDateString(r_media_type.Expiry_date)
	r_media_type.Row_effective_date = formatDateString(r_media_type.Row_effective_date)
	r_media_type.Row_expiry_date = formatDateString(r_media_type.Row_expiry_date)






	rows, err := stmt.Exec(r_media_type.Abbreviation, r_media_type.Active_ind, r_media_type.Digital_capacity, r_media_type.Digital_capacity_uom, r_media_type.Digital_density, r_media_type.Digital_density_uom, r_media_type.Effective_date, r_media_type.Expiry_date, r_media_type.Long_name, r_media_type.Ppdm_guid, r_media_type.Remark, r_media_type.Short_name, r_media_type.Source, r_media_type.Row_changed_by, r_media_type.Row_changed_date, r_media_type.Row_created_by, r_media_type.Row_effective_date, r_media_type.Row_expiry_date, r_media_type.Row_quality, r_media_type.Media_type)
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

func PatchRMediaType(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update r_media_type set "
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
	query += " where media_type = :id"

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

func DeleteRMediaType(c *fiber.Ctx) error {
	id := c.Params("id")
	var r_media_type dto.R_media_type
	r_media_type.Media_type = id

	stmt, err := db.Prepare("delete from r_media_type where media_type = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(r_media_type.Media_type)
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


