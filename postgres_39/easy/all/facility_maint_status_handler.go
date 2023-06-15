package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetFacilityMaintStatus(c *fiber.Ctx) error {
	rows, err := db.Query("select * from facility_maint_status")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Facility_maint_status

	for rows.Next() {
		var facility_maint_status dto.Facility_maint_status
		if err := rows.Scan(&facility_maint_status.Facility_id, &facility_maint_status.Facility_type, &facility_maint_status.Maintain_id, &facility_maint_status.Status_id, &facility_maint_status.Active_ind, &facility_maint_status.Effective_date, &facility_maint_status.Expiry_date, &facility_maint_status.Maintain_status, &facility_maint_status.Maintain_status_type, &facility_maint_status.Ppdm_guid, &facility_maint_status.Remark, &facility_maint_status.Source, &facility_maint_status.Row_changed_by, &facility_maint_status.Row_changed_date, &facility_maint_status.Row_created_by, &facility_maint_status.Row_created_date, &facility_maint_status.Row_effective_date, &facility_maint_status.Row_expiry_date, &facility_maint_status.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append facility_maint_status to result
		result = append(result, facility_maint_status)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetFacilityMaintStatus(c *fiber.Ctx) error {
	var facility_maint_status dto.Facility_maint_status

	setDefaults(&facility_maint_status)

	if err := json.Unmarshal(c.Body(), &facility_maint_status); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into facility_maint_status values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19)")
	if err != nil {
		return err
	}
	facility_maint_status.Row_created_date = formatDate(facility_maint_status.Row_created_date)
	facility_maint_status.Row_changed_date = formatDate(facility_maint_status.Row_changed_date)
	facility_maint_status.Effective_date = formatDateString(facility_maint_status.Effective_date)
	facility_maint_status.Expiry_date = formatDateString(facility_maint_status.Expiry_date)
	facility_maint_status.Row_effective_date = formatDateString(facility_maint_status.Row_effective_date)
	facility_maint_status.Row_expiry_date = formatDateString(facility_maint_status.Row_expiry_date)






	rows, err := stmt.Exec(facility_maint_status.Facility_id, facility_maint_status.Facility_type, facility_maint_status.Maintain_id, facility_maint_status.Status_id, facility_maint_status.Active_ind, facility_maint_status.Effective_date, facility_maint_status.Expiry_date, facility_maint_status.Maintain_status, facility_maint_status.Maintain_status_type, facility_maint_status.Ppdm_guid, facility_maint_status.Remark, facility_maint_status.Source, facility_maint_status.Row_changed_by, facility_maint_status.Row_changed_date, facility_maint_status.Row_created_by, facility_maint_status.Row_created_date, facility_maint_status.Row_effective_date, facility_maint_status.Row_expiry_date, facility_maint_status.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateFacilityMaintStatus(c *fiber.Ctx) error {
	var facility_maint_status dto.Facility_maint_status

	setDefaults(&facility_maint_status)

	if err := json.Unmarshal(c.Body(), &facility_maint_status); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	facility_maint_status.Facility_id = id

    if facility_maint_status.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update facility_maint_status set facility_type = :1, maintain_id = :2, status_id = :3, active_ind = :4, effective_date = :5, expiry_date = :6, maintain_status = :7, maintain_status_type = :8, ppdm_guid = :9, remark = :10, source = :11, row_changed_by = :12, row_changed_date = :13, row_created_by = :14, row_effective_date = :15, row_expiry_date = :16, row_quality = :17 where facility_id = :19")
	if err != nil {
		return err
	}

	facility_maint_status.Row_changed_date = formatDate(facility_maint_status.Row_changed_date)
	facility_maint_status.Effective_date = formatDateString(facility_maint_status.Effective_date)
	facility_maint_status.Expiry_date = formatDateString(facility_maint_status.Expiry_date)
	facility_maint_status.Row_effective_date = formatDateString(facility_maint_status.Row_effective_date)
	facility_maint_status.Row_expiry_date = formatDateString(facility_maint_status.Row_expiry_date)






	rows, err := stmt.Exec(facility_maint_status.Facility_type, facility_maint_status.Maintain_id, facility_maint_status.Status_id, facility_maint_status.Active_ind, facility_maint_status.Effective_date, facility_maint_status.Expiry_date, facility_maint_status.Maintain_status, facility_maint_status.Maintain_status_type, facility_maint_status.Ppdm_guid, facility_maint_status.Remark, facility_maint_status.Source, facility_maint_status.Row_changed_by, facility_maint_status.Row_changed_date, facility_maint_status.Row_created_by, facility_maint_status.Row_effective_date, facility_maint_status.Row_expiry_date, facility_maint_status.Row_quality, facility_maint_status.Facility_id)
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

func PatchFacilityMaintStatus(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update facility_maint_status set "
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

func DeleteFacilityMaintStatus(c *fiber.Ctx) error {
	id := c.Params("id")
	var facility_maint_status dto.Facility_maint_status
	facility_maint_status.Facility_id = id

	stmt, err := db.Prepare("delete from facility_maint_status where facility_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(facility_maint_status.Facility_id)
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


