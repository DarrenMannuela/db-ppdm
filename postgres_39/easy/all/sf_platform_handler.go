package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetSfPlatform(c *fiber.Ctx) error {
	rows, err := db.Query("select * from sf_platform")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Sf_platform

	for rows.Next() {
		var sf_platform dto.Sf_platform
		if err := rows.Scan(&sf_platform.Sf_subtype, &sf_platform.Platform_id, &sf_platform.Active_ind, &sf_platform.Drill_slot_count, &sf_platform.Effective_date, &sf_platform.Expiry_date, &sf_platform.Installed_date, &sf_platform.Platform_name, &sf_platform.Platform_type, &sf_platform.Ppdm_guid, &sf_platform.Remark, &sf_platform.Removal_date, &sf_platform.Source, &sf_platform.Water_depth, &sf_platform.Water_depth_datum, &sf_platform.Water_depth_ouom, &sf_platform.Row_changed_by, &sf_platform.Row_changed_date, &sf_platform.Row_created_by, &sf_platform.Row_created_date, &sf_platform.Row_effective_date, &sf_platform.Row_expiry_date, &sf_platform.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append sf_platform to result
		result = append(result, sf_platform)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetSfPlatform(c *fiber.Ctx) error {
	var sf_platform dto.Sf_platform

	setDefaults(&sf_platform)

	if err := json.Unmarshal(c.Body(), &sf_platform); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into sf_platform values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23)")
	if err != nil {
		return err
	}
	sf_platform.Row_created_date = formatDate(sf_platform.Row_created_date)
	sf_platform.Row_changed_date = formatDate(sf_platform.Row_changed_date)
	sf_platform.Effective_date = formatDateString(sf_platform.Effective_date)
	sf_platform.Expiry_date = formatDateString(sf_platform.Expiry_date)
	sf_platform.Installed_date = formatDateString(sf_platform.Installed_date)
	sf_platform.Removal_date = formatDateString(sf_platform.Removal_date)
	sf_platform.Row_effective_date = formatDateString(sf_platform.Row_effective_date)
	sf_platform.Row_expiry_date = formatDateString(sf_platform.Row_expiry_date)






	rows, err := stmt.Exec(sf_platform.Sf_subtype, sf_platform.Platform_id, sf_platform.Active_ind, sf_platform.Drill_slot_count, sf_platform.Effective_date, sf_platform.Expiry_date, sf_platform.Installed_date, sf_platform.Platform_name, sf_platform.Platform_type, sf_platform.Ppdm_guid, sf_platform.Remark, sf_platform.Removal_date, sf_platform.Source, sf_platform.Water_depth, sf_platform.Water_depth_datum, sf_platform.Water_depth_ouom, sf_platform.Row_changed_by, sf_platform.Row_changed_date, sf_platform.Row_created_by, sf_platform.Row_created_date, sf_platform.Row_effective_date, sf_platform.Row_expiry_date, sf_platform.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateSfPlatform(c *fiber.Ctx) error {
	var sf_platform dto.Sf_platform

	setDefaults(&sf_platform)

	if err := json.Unmarshal(c.Body(), &sf_platform); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	sf_platform.Sf_subtype = id

    if sf_platform.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update sf_platform set platform_id = :1, active_ind = :2, drill_slot_count = :3, effective_date = :4, expiry_date = :5, installed_date = :6, platform_name = :7, platform_type = :8, ppdm_guid = :9, remark = :10, removal_date = :11, source = :12, water_depth = :13, water_depth_datum = :14, water_depth_ouom = :15, row_changed_by = :16, row_changed_date = :17, row_created_by = :18, row_effective_date = :19, row_expiry_date = :20, row_quality = :21 where sf_subtype = :23")
	if err != nil {
		return err
	}

	sf_platform.Row_changed_date = formatDate(sf_platform.Row_changed_date)
	sf_platform.Effective_date = formatDateString(sf_platform.Effective_date)
	sf_platform.Expiry_date = formatDateString(sf_platform.Expiry_date)
	sf_platform.Installed_date = formatDateString(sf_platform.Installed_date)
	sf_platform.Removal_date = formatDateString(sf_platform.Removal_date)
	sf_platform.Row_effective_date = formatDateString(sf_platform.Row_effective_date)
	sf_platform.Row_expiry_date = formatDateString(sf_platform.Row_expiry_date)






	rows, err := stmt.Exec(sf_platform.Platform_id, sf_platform.Active_ind, sf_platform.Drill_slot_count, sf_platform.Effective_date, sf_platform.Expiry_date, sf_platform.Installed_date, sf_platform.Platform_name, sf_platform.Platform_type, sf_platform.Ppdm_guid, sf_platform.Remark, sf_platform.Removal_date, sf_platform.Source, sf_platform.Water_depth, sf_platform.Water_depth_datum, sf_platform.Water_depth_ouom, sf_platform.Row_changed_by, sf_platform.Row_changed_date, sf_platform.Row_created_by, sf_platform.Row_effective_date, sf_platform.Row_expiry_date, sf_platform.Row_quality, sf_platform.Sf_subtype)
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

func PatchSfPlatform(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update sf_platform set "
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
		} else if key == "effective_date" || key == "expiry_date" || key == "installed_date" || key == "removal_date" || key == "row_effective_date" || key == "row_expiry_date"      {
			if date, ok := value.(string); ok && len(date) > 0 {
				formattedDate := formatDateString(&date)
				value = formattedDate
			}
		}
		values = append(values, value)
	}
	query += " where sf_subtype = :id"

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

func DeleteSfPlatform(c *fiber.Ctx) error {
	id := c.Params("id")
	var sf_platform dto.Sf_platform
	sf_platform.Sf_subtype = id

	stmt, err := db.Prepare("delete from sf_platform where sf_subtype = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(sf_platform.Sf_subtype)
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


