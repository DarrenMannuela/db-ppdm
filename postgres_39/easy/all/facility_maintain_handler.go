package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetFacilityMaintain(c *fiber.Ctx) error {
	rows, err := db.Query("select * from facility_maintain")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Facility_maintain

	for rows.Next() {
		var facility_maintain dto.Facility_maintain
		if err := rows.Scan(&facility_maintain.Facility_id, &facility_maintain.Facility_type, &facility_maintain.Maintain_id, &facility_maintain.Active_ind, &facility_maintain.Actual_end_date, &facility_maintain.Actual_start_date, &facility_maintain.Effective_date, &facility_maintain.Expiry_date, &facility_maintain.Maintain_ba_id, &facility_maintain.Maintain_type, &facility_maintain.Ppdm_guid, &facility_maintain.Remark, &facility_maintain.Schedule_end_date, &facility_maintain.Schedule_start_date, &facility_maintain.Source, &facility_maintain.Row_changed_by, &facility_maintain.Row_changed_date, &facility_maintain.Row_created_by, &facility_maintain.Row_created_date, &facility_maintain.Row_effective_date, &facility_maintain.Row_expiry_date, &facility_maintain.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append facility_maintain to result
		result = append(result, facility_maintain)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetFacilityMaintain(c *fiber.Ctx) error {
	var facility_maintain dto.Facility_maintain

	setDefaults(&facility_maintain)

	if err := json.Unmarshal(c.Body(), &facility_maintain); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into facility_maintain values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22)")
	if err != nil {
		return err
	}
	facility_maintain.Row_created_date = formatDate(facility_maintain.Row_created_date)
	facility_maintain.Row_changed_date = formatDate(facility_maintain.Row_changed_date)
	facility_maintain.Actual_end_date = formatDateString(facility_maintain.Actual_end_date)
	facility_maintain.Actual_start_date = formatDateString(facility_maintain.Actual_start_date)
	facility_maintain.Effective_date = formatDateString(facility_maintain.Effective_date)
	facility_maintain.Expiry_date = formatDateString(facility_maintain.Expiry_date)
	facility_maintain.Schedule_end_date = formatDateString(facility_maintain.Schedule_end_date)
	facility_maintain.Schedule_start_date = formatDateString(facility_maintain.Schedule_start_date)
	facility_maintain.Row_effective_date = formatDateString(facility_maintain.Row_effective_date)
	facility_maintain.Row_expiry_date = formatDateString(facility_maintain.Row_expiry_date)






	rows, err := stmt.Exec(facility_maintain.Facility_id, facility_maintain.Facility_type, facility_maintain.Maintain_id, facility_maintain.Active_ind, facility_maintain.Actual_end_date, facility_maintain.Actual_start_date, facility_maintain.Effective_date, facility_maintain.Expiry_date, facility_maintain.Maintain_ba_id, facility_maintain.Maintain_type, facility_maintain.Ppdm_guid, facility_maintain.Remark, facility_maintain.Schedule_end_date, facility_maintain.Schedule_start_date, facility_maintain.Source, facility_maintain.Row_changed_by, facility_maintain.Row_changed_date, facility_maintain.Row_created_by, facility_maintain.Row_created_date, facility_maintain.Row_effective_date, facility_maintain.Row_expiry_date, facility_maintain.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateFacilityMaintain(c *fiber.Ctx) error {
	var facility_maintain dto.Facility_maintain

	setDefaults(&facility_maintain)

	if err := json.Unmarshal(c.Body(), &facility_maintain); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	facility_maintain.Facility_id = id

    if facility_maintain.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update facility_maintain set facility_type = :1, maintain_id = :2, active_ind = :3, actual_end_date = :4, actual_start_date = :5, effective_date = :6, expiry_date = :7, maintain_ba_id = :8, maintain_type = :9, ppdm_guid = :10, remark = :11, schedule_end_date = :12, schedule_start_date = :13, source = :14, row_changed_by = :15, row_changed_date = :16, row_created_by = :17, row_effective_date = :18, row_expiry_date = :19, row_quality = :20 where facility_id = :22")
	if err != nil {
		return err
	}

	facility_maintain.Row_changed_date = formatDate(facility_maintain.Row_changed_date)
	facility_maintain.Actual_end_date = formatDateString(facility_maintain.Actual_end_date)
	facility_maintain.Actual_start_date = formatDateString(facility_maintain.Actual_start_date)
	facility_maintain.Effective_date = formatDateString(facility_maintain.Effective_date)
	facility_maintain.Expiry_date = formatDateString(facility_maintain.Expiry_date)
	facility_maintain.Schedule_end_date = formatDateString(facility_maintain.Schedule_end_date)
	facility_maintain.Schedule_start_date = formatDateString(facility_maintain.Schedule_start_date)
	facility_maintain.Row_effective_date = formatDateString(facility_maintain.Row_effective_date)
	facility_maintain.Row_expiry_date = formatDateString(facility_maintain.Row_expiry_date)






	rows, err := stmt.Exec(facility_maintain.Facility_type, facility_maintain.Maintain_id, facility_maintain.Active_ind, facility_maintain.Actual_end_date, facility_maintain.Actual_start_date, facility_maintain.Effective_date, facility_maintain.Expiry_date, facility_maintain.Maintain_ba_id, facility_maintain.Maintain_type, facility_maintain.Ppdm_guid, facility_maintain.Remark, facility_maintain.Schedule_end_date, facility_maintain.Schedule_start_date, facility_maintain.Source, facility_maintain.Row_changed_by, facility_maintain.Row_changed_date, facility_maintain.Row_created_by, facility_maintain.Row_effective_date, facility_maintain.Row_expiry_date, facility_maintain.Row_quality, facility_maintain.Facility_id)
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

func PatchFacilityMaintain(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update facility_maintain set "
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
		} else if key == "actual_end_date" || key == "actual_start_date" || key == "effective_date" || key == "expiry_date" || key == "schedule_end_date" || key == "schedule_start_date" || key == "row_effective_date" || key == "row_expiry_date"      {
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

func DeleteFacilityMaintain(c *fiber.Ctx) error {
	id := c.Params("id")
	var facility_maintain dto.Facility_maintain
	facility_maintain.Facility_id = id

	stmt, err := db.Prepare("delete from facility_maintain where facility_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(facility_maintain.Facility_id)
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


