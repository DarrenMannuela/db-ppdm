package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetZRWaterType(c *fiber.Ctx) error {
	rows, err := db.Query("select * from z_r_water_type")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Z_r_water_type

	for rows.Next() {
		var z_r_water_type dto.Z_r_water_type
		if err := rows.Scan(&z_r_water_type.Water_type, &z_r_water_type.Abbreviation, &z_r_water_type.Active_ind, &z_r_water_type.Effective_date, &z_r_water_type.Expiry_date, &z_r_water_type.Long_name, &z_r_water_type.Ppdm_guid, &z_r_water_type.Remark, &z_r_water_type.Short_name, &z_r_water_type.Source, &z_r_water_type.Row_changed_by, &z_r_water_type.Row_changed_date, &z_r_water_type.Row_created_by, &z_r_water_type.Row_created_date, &z_r_water_type.Row_effective_date, &z_r_water_type.Row_expiry_date, &z_r_water_type.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append z_r_water_type to result
		result = append(result, z_r_water_type)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetZRWaterType(c *fiber.Ctx) error {
	var z_r_water_type dto.Z_r_water_type

	setDefaults(&z_r_water_type)

	if err := json.Unmarshal(c.Body(), &z_r_water_type); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into z_r_water_type values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17)")
	if err != nil {
		return err
	}
	z_r_water_type.Row_created_date = formatDate(z_r_water_type.Row_created_date)
	z_r_water_type.Row_changed_date = formatDate(z_r_water_type.Row_changed_date)
	z_r_water_type.Effective_date = formatDateString(z_r_water_type.Effective_date)
	z_r_water_type.Expiry_date = formatDateString(z_r_water_type.Expiry_date)
	z_r_water_type.Row_effective_date = formatDateString(z_r_water_type.Row_effective_date)
	z_r_water_type.Row_expiry_date = formatDateString(z_r_water_type.Row_expiry_date)






	rows, err := stmt.Exec(z_r_water_type.Water_type, z_r_water_type.Abbreviation, z_r_water_type.Active_ind, z_r_water_type.Effective_date, z_r_water_type.Expiry_date, z_r_water_type.Long_name, z_r_water_type.Ppdm_guid, z_r_water_type.Remark, z_r_water_type.Short_name, z_r_water_type.Source, z_r_water_type.Row_changed_by, z_r_water_type.Row_changed_date, z_r_water_type.Row_created_by, z_r_water_type.Row_created_date, z_r_water_type.Row_effective_date, z_r_water_type.Row_expiry_date, z_r_water_type.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateZRWaterType(c *fiber.Ctx) error {
	var z_r_water_type dto.Z_r_water_type

	setDefaults(&z_r_water_type)

	if err := json.Unmarshal(c.Body(), &z_r_water_type); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	z_r_water_type.Water_type = id

    if z_r_water_type.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update z_r_water_type set abbreviation = :1, active_ind = :2, effective_date = :3, expiry_date = :4, long_name = :5, ppdm_guid = :6, remark = :7, short_name = :8, source = :9, row_changed_by = :10, row_changed_date = :11, row_created_by = :12, row_effective_date = :13, row_expiry_date = :14, row_quality = :15 where water_type = :17")
	if err != nil {
		return err
	}

	z_r_water_type.Row_changed_date = formatDate(z_r_water_type.Row_changed_date)
	z_r_water_type.Effective_date = formatDateString(z_r_water_type.Effective_date)
	z_r_water_type.Expiry_date = formatDateString(z_r_water_type.Expiry_date)
	z_r_water_type.Row_effective_date = formatDateString(z_r_water_type.Row_effective_date)
	z_r_water_type.Row_expiry_date = formatDateString(z_r_water_type.Row_expiry_date)






	rows, err := stmt.Exec(z_r_water_type.Abbreviation, z_r_water_type.Active_ind, z_r_water_type.Effective_date, z_r_water_type.Expiry_date, z_r_water_type.Long_name, z_r_water_type.Ppdm_guid, z_r_water_type.Remark, z_r_water_type.Short_name, z_r_water_type.Source, z_r_water_type.Row_changed_by, z_r_water_type.Row_changed_date, z_r_water_type.Row_created_by, z_r_water_type.Row_effective_date, z_r_water_type.Row_expiry_date, z_r_water_type.Row_quality, z_r_water_type.Water_type)
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

func PatchZRWaterType(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update z_r_water_type set "
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
	query += " where water_type = :id"

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

func DeleteZRWaterType(c *fiber.Ctx) error {
	id := c.Params("id")
	var z_r_water_type dto.Z_r_water_type
	z_r_water_type.Water_type = id

	stmt, err := db.Prepare("delete from z_r_water_type where water_type = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(z_r_water_type.Water_type)
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


