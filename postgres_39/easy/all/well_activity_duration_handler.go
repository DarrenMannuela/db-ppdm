package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetWellActivityDuration(c *fiber.Ctx) error {
	rows, err := db.Query("select * from well_activity_duration")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Well_activity_duration

	for rows.Next() {
		var well_activity_duration dto.Well_activity_duration
		if err := rows.Scan(&well_activity_duration.Uwi, &well_activity_duration.Source, &well_activity_duration.Activity_obs_no, &well_activity_duration.Duration_obs_no, &well_activity_duration.Active_ind, &well_activity_duration.Activity_duration, &well_activity_duration.Activity_duration_ouom, &well_activity_duration.Downtime_type, &well_activity_duration.Effective_date, &well_activity_duration.End_time, &well_activity_duration.End_timezone, &well_activity_duration.Event_date, &well_activity_duration.Expiry_date, &well_activity_duration.Period_obs_no, &well_activity_duration.Ppdm_guid, &well_activity_duration.Remark, &well_activity_duration.Start_time, &well_activity_duration.Start_timezone, &well_activity_duration.Row_changed_by, &well_activity_duration.Row_changed_date, &well_activity_duration.Row_created_by, &well_activity_duration.Row_created_date, &well_activity_duration.Row_effective_date, &well_activity_duration.Row_expiry_date, &well_activity_duration.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append well_activity_duration to result
		result = append(result, well_activity_duration)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetWellActivityDuration(c *fiber.Ctx) error {
	var well_activity_duration dto.Well_activity_duration

	setDefaults(&well_activity_duration)

	if err := json.Unmarshal(c.Body(), &well_activity_duration); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into well_activity_duration values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25)")
	if err != nil {
		return err
	}
	well_activity_duration.Row_created_date = formatDate(well_activity_duration.Row_created_date)
	well_activity_duration.Row_changed_date = formatDate(well_activity_duration.Row_changed_date)
	well_activity_duration.Effective_date = formatDateString(well_activity_duration.Effective_date)
	well_activity_duration.End_time = formatDateString(well_activity_duration.End_time)
	well_activity_duration.Event_date = formatDateString(well_activity_duration.Event_date)
	well_activity_duration.Expiry_date = formatDateString(well_activity_duration.Expiry_date)
	well_activity_duration.Start_time = formatDateString(well_activity_duration.Start_time)
	well_activity_duration.Row_effective_date = formatDateString(well_activity_duration.Row_effective_date)
	well_activity_duration.Row_expiry_date = formatDateString(well_activity_duration.Row_expiry_date)






	rows, err := stmt.Exec(well_activity_duration.Uwi, well_activity_duration.Source, well_activity_duration.Activity_obs_no, well_activity_duration.Duration_obs_no, well_activity_duration.Active_ind, well_activity_duration.Activity_duration, well_activity_duration.Activity_duration_ouom, well_activity_duration.Downtime_type, well_activity_duration.Effective_date, well_activity_duration.End_time, well_activity_duration.End_timezone, well_activity_duration.Event_date, well_activity_duration.Expiry_date, well_activity_duration.Period_obs_no, well_activity_duration.Ppdm_guid, well_activity_duration.Remark, well_activity_duration.Start_time, well_activity_duration.Start_timezone, well_activity_duration.Row_changed_by, well_activity_duration.Row_changed_date, well_activity_duration.Row_created_by, well_activity_duration.Row_created_date, well_activity_duration.Row_effective_date, well_activity_duration.Row_expiry_date, well_activity_duration.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateWellActivityDuration(c *fiber.Ctx) error {
	var well_activity_duration dto.Well_activity_duration

	setDefaults(&well_activity_duration)

	if err := json.Unmarshal(c.Body(), &well_activity_duration); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	well_activity_duration.Uwi = id

    if well_activity_duration.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update well_activity_duration set source = :1, activity_obs_no = :2, duration_obs_no = :3, active_ind = :4, activity_duration = :5, activity_duration_ouom = :6, downtime_type = :7, effective_date = :8, end_time = :9, end_timezone = :10, event_date = :11, expiry_date = :12, period_obs_no = :13, ppdm_guid = :14, remark = :15, start_time = :16, start_timezone = :17, row_changed_by = :18, row_changed_date = :19, row_created_by = :20, row_effective_date = :21, row_expiry_date = :22, row_quality = :23 where uwi = :25")
	if err != nil {
		return err
	}

	well_activity_duration.Row_changed_date = formatDate(well_activity_duration.Row_changed_date)
	well_activity_duration.Effective_date = formatDateString(well_activity_duration.Effective_date)
	well_activity_duration.End_time = formatDateString(well_activity_duration.End_time)
	well_activity_duration.Event_date = formatDateString(well_activity_duration.Event_date)
	well_activity_duration.Expiry_date = formatDateString(well_activity_duration.Expiry_date)
	well_activity_duration.Start_time = formatDateString(well_activity_duration.Start_time)
	well_activity_duration.Row_effective_date = formatDateString(well_activity_duration.Row_effective_date)
	well_activity_duration.Row_expiry_date = formatDateString(well_activity_duration.Row_expiry_date)






	rows, err := stmt.Exec(well_activity_duration.Source, well_activity_duration.Activity_obs_no, well_activity_duration.Duration_obs_no, well_activity_duration.Active_ind, well_activity_duration.Activity_duration, well_activity_duration.Activity_duration_ouom, well_activity_duration.Downtime_type, well_activity_duration.Effective_date, well_activity_duration.End_time, well_activity_duration.End_timezone, well_activity_duration.Event_date, well_activity_duration.Expiry_date, well_activity_duration.Period_obs_no, well_activity_duration.Ppdm_guid, well_activity_duration.Remark, well_activity_duration.Start_time, well_activity_duration.Start_timezone, well_activity_duration.Row_changed_by, well_activity_duration.Row_changed_date, well_activity_duration.Row_created_by, well_activity_duration.Row_effective_date, well_activity_duration.Row_expiry_date, well_activity_duration.Row_quality, well_activity_duration.Uwi)
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

func PatchWellActivityDuration(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update well_activity_duration set "
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
		} else if key == "effective_date" || key == "end_time" || key == "event_date" || key == "expiry_date" || key == "start_time" || key == "row_effective_date" || key == "row_expiry_date"      {
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

func DeleteWellActivityDuration(c *fiber.Ctx) error {
	id := c.Params("id")
	var well_activity_duration dto.Well_activity_duration
	well_activity_duration.Uwi = id

	stmt, err := db.Prepare("delete from well_activity_duration where uwi = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(well_activity_duration.Uwi)
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


