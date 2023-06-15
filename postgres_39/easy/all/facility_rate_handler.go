package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetFacilityRate(c *fiber.Ctx) error {
	rows, err := db.Query("select * from facility_rate")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Facility_rate

	for rows.Next() {
		var facility_rate dto.Facility_rate
		if err := rows.Scan(&facility_rate.Facility_id, &facility_rate.Facility_type, &facility_rate.Rate_schedule_id, &facility_rate.Active_ind, &facility_rate.Effective_date, &facility_rate.Expiry_date, &facility_rate.Ppdm_guid, &facility_rate.Remark, &facility_rate.Source, &facility_rate.Row_changed_by, &facility_rate.Row_changed_date, &facility_rate.Row_created_by, &facility_rate.Row_created_date, &facility_rate.Row_effective_date, &facility_rate.Row_expiry_date, &facility_rate.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append facility_rate to result
		result = append(result, facility_rate)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetFacilityRate(c *fiber.Ctx) error {
	var facility_rate dto.Facility_rate

	setDefaults(&facility_rate)

	if err := json.Unmarshal(c.Body(), &facility_rate); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into facility_rate values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16)")
	if err != nil {
		return err
	}
	facility_rate.Row_created_date = formatDate(facility_rate.Row_created_date)
	facility_rate.Row_changed_date = formatDate(facility_rate.Row_changed_date)
	facility_rate.Effective_date = formatDateString(facility_rate.Effective_date)
	facility_rate.Expiry_date = formatDateString(facility_rate.Expiry_date)
	facility_rate.Row_effective_date = formatDateString(facility_rate.Row_effective_date)
	facility_rate.Row_expiry_date = formatDateString(facility_rate.Row_expiry_date)






	rows, err := stmt.Exec(facility_rate.Facility_id, facility_rate.Facility_type, facility_rate.Rate_schedule_id, facility_rate.Active_ind, facility_rate.Effective_date, facility_rate.Expiry_date, facility_rate.Ppdm_guid, facility_rate.Remark, facility_rate.Source, facility_rate.Row_changed_by, facility_rate.Row_changed_date, facility_rate.Row_created_by, facility_rate.Row_created_date, facility_rate.Row_effective_date, facility_rate.Row_expiry_date, facility_rate.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateFacilityRate(c *fiber.Ctx) error {
	var facility_rate dto.Facility_rate

	setDefaults(&facility_rate)

	if err := json.Unmarshal(c.Body(), &facility_rate); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	facility_rate.Facility_id = id

    if facility_rate.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update facility_rate set facility_type = :1, rate_schedule_id = :2, active_ind = :3, effective_date = :4, expiry_date = :5, ppdm_guid = :6, remark = :7, source = :8, row_changed_by = :9, row_changed_date = :10, row_created_by = :11, row_effective_date = :12, row_expiry_date = :13, row_quality = :14 where facility_id = :16")
	if err != nil {
		return err
	}

	facility_rate.Row_changed_date = formatDate(facility_rate.Row_changed_date)
	facility_rate.Effective_date = formatDateString(facility_rate.Effective_date)
	facility_rate.Expiry_date = formatDateString(facility_rate.Expiry_date)
	facility_rate.Row_effective_date = formatDateString(facility_rate.Row_effective_date)
	facility_rate.Row_expiry_date = formatDateString(facility_rate.Row_expiry_date)






	rows, err := stmt.Exec(facility_rate.Facility_type, facility_rate.Rate_schedule_id, facility_rate.Active_ind, facility_rate.Effective_date, facility_rate.Expiry_date, facility_rate.Ppdm_guid, facility_rate.Remark, facility_rate.Source, facility_rate.Row_changed_by, facility_rate.Row_changed_date, facility_rate.Row_created_by, facility_rate.Row_effective_date, facility_rate.Row_expiry_date, facility_rate.Row_quality, facility_rate.Facility_id)
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

func PatchFacilityRate(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update facility_rate set "
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

func DeleteFacilityRate(c *fiber.Ctx) error {
	id := c.Params("id")
	var facility_rate dto.Facility_rate
	facility_rate.Facility_id = id

	stmt, err := db.Prepare("delete from facility_rate where facility_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(facility_rate.Facility_id)
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


