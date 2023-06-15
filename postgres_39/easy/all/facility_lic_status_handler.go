package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetFacilityLicStatus(c *fiber.Ctx) error {
	rows, err := db.Query("select * from facility_lic_status")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Facility_lic_status

	for rows.Next() {
		var facility_lic_status dto.Facility_lic_status
		if err := rows.Scan(&facility_lic_status.Facility_id, &facility_lic_status.Facility_type, &facility_lic_status.License_id, &facility_lic_status.Status_id, &facility_lic_status.Active_ind, &facility_lic_status.Effective_date, &facility_lic_status.Expiry_date, &facility_lic_status.License_status, &facility_lic_status.License_status_type, &facility_lic_status.Ppdm_guid, &facility_lic_status.Remark, &facility_lic_status.Source, &facility_lic_status.Row_changed_by, &facility_lic_status.Row_changed_date, &facility_lic_status.Row_created_by, &facility_lic_status.Row_created_date, &facility_lic_status.Row_effective_date, &facility_lic_status.Row_expiry_date, &facility_lic_status.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append facility_lic_status to result
		result = append(result, facility_lic_status)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetFacilityLicStatus(c *fiber.Ctx) error {
	var facility_lic_status dto.Facility_lic_status

	setDefaults(&facility_lic_status)

	if err := json.Unmarshal(c.Body(), &facility_lic_status); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into facility_lic_status values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19)")
	if err != nil {
		return err
	}
	facility_lic_status.Row_created_date = formatDate(facility_lic_status.Row_created_date)
	facility_lic_status.Row_changed_date = formatDate(facility_lic_status.Row_changed_date)
	facility_lic_status.Effective_date = formatDateString(facility_lic_status.Effective_date)
	facility_lic_status.Expiry_date = formatDateString(facility_lic_status.Expiry_date)
	facility_lic_status.Row_effective_date = formatDateString(facility_lic_status.Row_effective_date)
	facility_lic_status.Row_expiry_date = formatDateString(facility_lic_status.Row_expiry_date)






	rows, err := stmt.Exec(facility_lic_status.Facility_id, facility_lic_status.Facility_type, facility_lic_status.License_id, facility_lic_status.Status_id, facility_lic_status.Active_ind, facility_lic_status.Effective_date, facility_lic_status.Expiry_date, facility_lic_status.License_status, facility_lic_status.License_status_type, facility_lic_status.Ppdm_guid, facility_lic_status.Remark, facility_lic_status.Source, facility_lic_status.Row_changed_by, facility_lic_status.Row_changed_date, facility_lic_status.Row_created_by, facility_lic_status.Row_created_date, facility_lic_status.Row_effective_date, facility_lic_status.Row_expiry_date, facility_lic_status.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateFacilityLicStatus(c *fiber.Ctx) error {
	var facility_lic_status dto.Facility_lic_status

	setDefaults(&facility_lic_status)

	if err := json.Unmarshal(c.Body(), &facility_lic_status); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	facility_lic_status.Facility_id = id

    if facility_lic_status.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update facility_lic_status set facility_type = :1, license_id = :2, status_id = :3, active_ind = :4, effective_date = :5, expiry_date = :6, license_status = :7, license_status_type = :8, ppdm_guid = :9, remark = :10, source = :11, row_changed_by = :12, row_changed_date = :13, row_created_by = :14, row_effective_date = :15, row_expiry_date = :16, row_quality = :17 where facility_id = :19")
	if err != nil {
		return err
	}

	facility_lic_status.Row_changed_date = formatDate(facility_lic_status.Row_changed_date)
	facility_lic_status.Effective_date = formatDateString(facility_lic_status.Effective_date)
	facility_lic_status.Expiry_date = formatDateString(facility_lic_status.Expiry_date)
	facility_lic_status.Row_effective_date = formatDateString(facility_lic_status.Row_effective_date)
	facility_lic_status.Row_expiry_date = formatDateString(facility_lic_status.Row_expiry_date)






	rows, err := stmt.Exec(facility_lic_status.Facility_type, facility_lic_status.License_id, facility_lic_status.Status_id, facility_lic_status.Active_ind, facility_lic_status.Effective_date, facility_lic_status.Expiry_date, facility_lic_status.License_status, facility_lic_status.License_status_type, facility_lic_status.Ppdm_guid, facility_lic_status.Remark, facility_lic_status.Source, facility_lic_status.Row_changed_by, facility_lic_status.Row_changed_date, facility_lic_status.Row_created_by, facility_lic_status.Row_effective_date, facility_lic_status.Row_expiry_date, facility_lic_status.Row_quality, facility_lic_status.Facility_id)
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

func PatchFacilityLicStatus(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update facility_lic_status set "
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

func DeleteFacilityLicStatus(c *fiber.Ctx) error {
	id := c.Params("id")
	var facility_lic_status dto.Facility_lic_status
	facility_lic_status.Facility_id = id

	stmt, err := db.Prepare("delete from facility_lic_status where facility_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(facility_lic_status.Facility_id)
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


