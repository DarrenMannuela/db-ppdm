package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetSfMaintain(c *fiber.Ctx) error {
	rows, err := db.Query("select * from sf_maintain")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Sf_maintain

	for rows.Next() {
		var sf_maintain dto.Sf_maintain
		if err := rows.Scan(&sf_maintain.Sf_subtype, &sf_maintain.Support_facility_id, &sf_maintain.Maintain_id, &sf_maintain.Active_ind, &sf_maintain.Actual_end_date, &sf_maintain.Actual_start_date, &sf_maintain.Effective_date, &sf_maintain.Expiry_date, &sf_maintain.Maintain_ba_id, &sf_maintain.Maintain_type, &sf_maintain.Ppdm_guid, &sf_maintain.Remark, &sf_maintain.Schedule_end_date, &sf_maintain.Schedule_start_date, &sf_maintain.Source, &sf_maintain.Row_changed_by, &sf_maintain.Row_changed_date, &sf_maintain.Row_created_by, &sf_maintain.Row_created_date, &sf_maintain.Row_effective_date, &sf_maintain.Row_expiry_date, &sf_maintain.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append sf_maintain to result
		result = append(result, sf_maintain)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetSfMaintain(c *fiber.Ctx) error {
	var sf_maintain dto.Sf_maintain

	setDefaults(&sf_maintain)

	if err := json.Unmarshal(c.Body(), &sf_maintain); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into sf_maintain values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22)")
	if err != nil {
		return err
	}
	sf_maintain.Row_created_date = formatDate(sf_maintain.Row_created_date)
	sf_maintain.Row_changed_date = formatDate(sf_maintain.Row_changed_date)
	sf_maintain.Actual_end_date = formatDateString(sf_maintain.Actual_end_date)
	sf_maintain.Actual_start_date = formatDateString(sf_maintain.Actual_start_date)
	sf_maintain.Effective_date = formatDateString(sf_maintain.Effective_date)
	sf_maintain.Expiry_date = formatDateString(sf_maintain.Expiry_date)
	sf_maintain.Schedule_end_date = formatDateString(sf_maintain.Schedule_end_date)
	sf_maintain.Schedule_start_date = formatDateString(sf_maintain.Schedule_start_date)
	sf_maintain.Row_effective_date = formatDateString(sf_maintain.Row_effective_date)
	sf_maintain.Row_expiry_date = formatDateString(sf_maintain.Row_expiry_date)






	rows, err := stmt.Exec(sf_maintain.Sf_subtype, sf_maintain.Support_facility_id, sf_maintain.Maintain_id, sf_maintain.Active_ind, sf_maintain.Actual_end_date, sf_maintain.Actual_start_date, sf_maintain.Effective_date, sf_maintain.Expiry_date, sf_maintain.Maintain_ba_id, sf_maintain.Maintain_type, sf_maintain.Ppdm_guid, sf_maintain.Remark, sf_maintain.Schedule_end_date, sf_maintain.Schedule_start_date, sf_maintain.Source, sf_maintain.Row_changed_by, sf_maintain.Row_changed_date, sf_maintain.Row_created_by, sf_maintain.Row_created_date, sf_maintain.Row_effective_date, sf_maintain.Row_expiry_date, sf_maintain.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateSfMaintain(c *fiber.Ctx) error {
	var sf_maintain dto.Sf_maintain

	setDefaults(&sf_maintain)

	if err := json.Unmarshal(c.Body(), &sf_maintain); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	sf_maintain.Sf_subtype = id

    if sf_maintain.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update sf_maintain set support_facility_id = :1, maintain_id = :2, active_ind = :3, actual_end_date = :4, actual_start_date = :5, effective_date = :6, expiry_date = :7, maintain_ba_id = :8, maintain_type = :9, ppdm_guid = :10, remark = :11, schedule_end_date = :12, schedule_start_date = :13, source = :14, row_changed_by = :15, row_changed_date = :16, row_created_by = :17, row_effective_date = :18, row_expiry_date = :19, row_quality = :20 where sf_subtype = :22")
	if err != nil {
		return err
	}

	sf_maintain.Row_changed_date = formatDate(sf_maintain.Row_changed_date)
	sf_maintain.Actual_end_date = formatDateString(sf_maintain.Actual_end_date)
	sf_maintain.Actual_start_date = formatDateString(sf_maintain.Actual_start_date)
	sf_maintain.Effective_date = formatDateString(sf_maintain.Effective_date)
	sf_maintain.Expiry_date = formatDateString(sf_maintain.Expiry_date)
	sf_maintain.Schedule_end_date = formatDateString(sf_maintain.Schedule_end_date)
	sf_maintain.Schedule_start_date = formatDateString(sf_maintain.Schedule_start_date)
	sf_maintain.Row_effective_date = formatDateString(sf_maintain.Row_effective_date)
	sf_maintain.Row_expiry_date = formatDateString(sf_maintain.Row_expiry_date)






	rows, err := stmt.Exec(sf_maintain.Support_facility_id, sf_maintain.Maintain_id, sf_maintain.Active_ind, sf_maintain.Actual_end_date, sf_maintain.Actual_start_date, sf_maintain.Effective_date, sf_maintain.Expiry_date, sf_maintain.Maintain_ba_id, sf_maintain.Maintain_type, sf_maintain.Ppdm_guid, sf_maintain.Remark, sf_maintain.Schedule_end_date, sf_maintain.Schedule_start_date, sf_maintain.Source, sf_maintain.Row_changed_by, sf_maintain.Row_changed_date, sf_maintain.Row_created_by, sf_maintain.Row_effective_date, sf_maintain.Row_expiry_date, sf_maintain.Row_quality, sf_maintain.Sf_subtype)
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

func PatchSfMaintain(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update sf_maintain set "
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

func DeleteSfMaintain(c *fiber.Ctx) error {
	id := c.Params("id")
	var sf_maintain dto.Sf_maintain
	sf_maintain.Sf_subtype = id

	stmt, err := db.Prepare("delete from sf_maintain where sf_subtype = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(sf_maintain.Sf_subtype)
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


