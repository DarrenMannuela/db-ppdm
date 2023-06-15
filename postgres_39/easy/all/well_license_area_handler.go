package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetWellLicenseArea(c *fiber.Ctx) error {
	rows, err := db.Query("select * from well_license_area")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Well_license_area

	for rows.Next() {
		var well_license_area dto.Well_license_area
		if err := rows.Scan(&well_license_area.Uwi, &well_license_area.License_id, &well_license_area.Well_lic_source, &well_license_area.Area_id, &well_license_area.Area_type, &well_license_area.Active_ind, &well_license_area.Description, &well_license_area.Effective_date, &well_license_area.Expiry_date, &well_license_area.Ppdm_guid, &well_license_area.Remark, &well_license_area.Source, &well_license_area.Row_changed_by, &well_license_area.Row_changed_date, &well_license_area.Row_created_by, &well_license_area.Row_created_date, &well_license_area.Row_effective_date, &well_license_area.Row_expiry_date, &well_license_area.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append well_license_area to result
		result = append(result, well_license_area)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetWellLicenseArea(c *fiber.Ctx) error {
	var well_license_area dto.Well_license_area

	setDefaults(&well_license_area)

	if err := json.Unmarshal(c.Body(), &well_license_area); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into well_license_area values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19)")
	if err != nil {
		return err
	}
	well_license_area.Row_created_date = formatDate(well_license_area.Row_created_date)
	well_license_area.Row_changed_date = formatDate(well_license_area.Row_changed_date)
	well_license_area.Effective_date = formatDateString(well_license_area.Effective_date)
	well_license_area.Expiry_date = formatDateString(well_license_area.Expiry_date)
	well_license_area.Row_effective_date = formatDateString(well_license_area.Row_effective_date)
	well_license_area.Row_expiry_date = formatDateString(well_license_area.Row_expiry_date)






	rows, err := stmt.Exec(well_license_area.Uwi, well_license_area.License_id, well_license_area.Well_lic_source, well_license_area.Area_id, well_license_area.Area_type, well_license_area.Active_ind, well_license_area.Description, well_license_area.Effective_date, well_license_area.Expiry_date, well_license_area.Ppdm_guid, well_license_area.Remark, well_license_area.Source, well_license_area.Row_changed_by, well_license_area.Row_changed_date, well_license_area.Row_created_by, well_license_area.Row_created_date, well_license_area.Row_effective_date, well_license_area.Row_expiry_date, well_license_area.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateWellLicenseArea(c *fiber.Ctx) error {
	var well_license_area dto.Well_license_area

	setDefaults(&well_license_area)

	if err := json.Unmarshal(c.Body(), &well_license_area); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	well_license_area.Uwi = id

    if well_license_area.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update well_license_area set license_id = :1, well_lic_source = :2, area_id = :3, area_type = :4, active_ind = :5, description = :6, effective_date = :7, expiry_date = :8, ppdm_guid = :9, remark = :10, source = :11, row_changed_by = :12, row_changed_date = :13, row_created_by = :14, row_effective_date = :15, row_expiry_date = :16, row_quality = :17 where uwi = :19")
	if err != nil {
		return err
	}

	well_license_area.Row_changed_date = formatDate(well_license_area.Row_changed_date)
	well_license_area.Effective_date = formatDateString(well_license_area.Effective_date)
	well_license_area.Expiry_date = formatDateString(well_license_area.Expiry_date)
	well_license_area.Row_effective_date = formatDateString(well_license_area.Row_effective_date)
	well_license_area.Row_expiry_date = formatDateString(well_license_area.Row_expiry_date)






	rows, err := stmt.Exec(well_license_area.License_id, well_license_area.Well_lic_source, well_license_area.Area_id, well_license_area.Area_type, well_license_area.Active_ind, well_license_area.Description, well_license_area.Effective_date, well_license_area.Expiry_date, well_license_area.Ppdm_guid, well_license_area.Remark, well_license_area.Source, well_license_area.Row_changed_by, well_license_area.Row_changed_date, well_license_area.Row_created_by, well_license_area.Row_effective_date, well_license_area.Row_expiry_date, well_license_area.Row_quality, well_license_area.Uwi)
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

func PatchWellLicenseArea(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update well_license_area set "
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
	query += " where uwi = :id"

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

func DeleteWellLicenseArea(c *fiber.Ctx) error {
	id := c.Params("id")
	var well_license_area dto.Well_license_area
	well_license_area.Uwi = id

	stmt, err := db.Prepare("delete from well_license_area where uwi = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(well_license_area.Uwi)
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


