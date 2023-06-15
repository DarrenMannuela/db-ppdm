package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetBaLicenseArea(c *fiber.Ctx) error {
	rows, err := db.Query("select * from ba_license_area")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Ba_license_area

	for rows.Next() {
		var ba_license_area dto.Ba_license_area
		if err := rows.Scan(&ba_license_area.Licensee_ba_id, &ba_license_area.License_id, &ba_license_area.Area_id, &ba_license_area.Area_type, &ba_license_area.Active_ind, &ba_license_area.Description, &ba_license_area.Effective_date, &ba_license_area.Expiry_date, &ba_license_area.Ppdm_guid, &ba_license_area.Remark, &ba_license_area.Source, &ba_license_area.Row_changed_by, &ba_license_area.Row_changed_date, &ba_license_area.Row_created_by, &ba_license_area.Row_created_date, &ba_license_area.Row_effective_date, &ba_license_area.Row_expiry_date, &ba_license_area.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append ba_license_area to result
		result = append(result, ba_license_area)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetBaLicenseArea(c *fiber.Ctx) error {
	var ba_license_area dto.Ba_license_area

	setDefaults(&ba_license_area)

	if err := json.Unmarshal(c.Body(), &ba_license_area); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into ba_license_area values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18)")
	if err != nil {
		return err
	}
	ba_license_area.Row_created_date = formatDate(ba_license_area.Row_created_date)
	ba_license_area.Row_changed_date = formatDate(ba_license_area.Row_changed_date)
	ba_license_area.Effective_date = formatDateString(ba_license_area.Effective_date)
	ba_license_area.Expiry_date = formatDateString(ba_license_area.Expiry_date)
	ba_license_area.Row_effective_date = formatDateString(ba_license_area.Row_effective_date)
	ba_license_area.Row_expiry_date = formatDateString(ba_license_area.Row_expiry_date)






	rows, err := stmt.Exec(ba_license_area.Licensee_ba_id, ba_license_area.License_id, ba_license_area.Area_id, ba_license_area.Area_type, ba_license_area.Active_ind, ba_license_area.Description, ba_license_area.Effective_date, ba_license_area.Expiry_date, ba_license_area.Ppdm_guid, ba_license_area.Remark, ba_license_area.Source, ba_license_area.Row_changed_by, ba_license_area.Row_changed_date, ba_license_area.Row_created_by, ba_license_area.Row_created_date, ba_license_area.Row_effective_date, ba_license_area.Row_expiry_date, ba_license_area.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateBaLicenseArea(c *fiber.Ctx) error {
	var ba_license_area dto.Ba_license_area

	setDefaults(&ba_license_area)

	if err := json.Unmarshal(c.Body(), &ba_license_area); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	ba_license_area.Licensee_ba_id = id

    if ba_license_area.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update ba_license_area set license_id = :1, area_id = :2, area_type = :3, active_ind = :4, description = :5, effective_date = :6, expiry_date = :7, ppdm_guid = :8, remark = :9, source = :10, row_changed_by = :11, row_changed_date = :12, row_created_by = :13, row_effective_date = :14, row_expiry_date = :15, row_quality = :16 where licensee_ba_id = :18")
	if err != nil {
		return err
	}

	ba_license_area.Row_changed_date = formatDate(ba_license_area.Row_changed_date)
	ba_license_area.Effective_date = formatDateString(ba_license_area.Effective_date)
	ba_license_area.Expiry_date = formatDateString(ba_license_area.Expiry_date)
	ba_license_area.Row_effective_date = formatDateString(ba_license_area.Row_effective_date)
	ba_license_area.Row_expiry_date = formatDateString(ba_license_area.Row_expiry_date)






	rows, err := stmt.Exec(ba_license_area.License_id, ba_license_area.Area_id, ba_license_area.Area_type, ba_license_area.Active_ind, ba_license_area.Description, ba_license_area.Effective_date, ba_license_area.Expiry_date, ba_license_area.Ppdm_guid, ba_license_area.Remark, ba_license_area.Source, ba_license_area.Row_changed_by, ba_license_area.Row_changed_date, ba_license_area.Row_created_by, ba_license_area.Row_effective_date, ba_license_area.Row_expiry_date, ba_license_area.Row_quality, ba_license_area.Licensee_ba_id)
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

func PatchBaLicenseArea(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update ba_license_area set "
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
	query += " where licensee_ba_id = :id"

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

func DeleteBaLicenseArea(c *fiber.Ctx) error {
	id := c.Params("id")
	var ba_license_area dto.Ba_license_area
	ba_license_area.Licensee_ba_id = id

	stmt, err := db.Prepare("delete from ba_license_area where licensee_ba_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(ba_license_area.Licensee_ba_id)
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


