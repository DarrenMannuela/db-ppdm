package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetHseIncident(c *fiber.Ctx) error {
	rows, err := db.Query("select * from hse_incident")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Hse_incident

	for rows.Next() {
		var hse_incident dto.Hse_incident
		if err := rows.Scan(&hse_incident.Incident_id, &hse_incident.Active_ind, &hse_incident.Effective_date, &hse_incident.Expiry_date, &hse_incident.Incident_class_id, &hse_incident.Incident_date, &hse_incident.Incident_duration, &hse_incident.Incident_duration_ouom, &hse_incident.Incident_duration_uom, &hse_incident.Lost_time_ind, &hse_incident.Ppdm_guid, &hse_incident.Recorded_time, &hse_incident.Recorded_timezone, &hse_incident.Remark, &hse_incident.Reported_by_ba_id, &hse_incident.Reported_by_name, &hse_incident.Reported_ind, &hse_incident.Source, &hse_incident.Work_related_ind, &hse_incident.Row_changed_by, &hse_incident.Row_changed_date, &hse_incident.Row_created_by, &hse_incident.Row_created_date, &hse_incident.Row_effective_date, &hse_incident.Row_expiry_date, &hse_incident.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append hse_incident to result
		result = append(result, hse_incident)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetHseIncident(c *fiber.Ctx) error {
	var hse_incident dto.Hse_incident

	setDefaults(&hse_incident)

	if err := json.Unmarshal(c.Body(), &hse_incident); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into hse_incident values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26)")
	if err != nil {
		return err
	}
	hse_incident.Row_created_date = formatDate(hse_incident.Row_created_date)
	hse_incident.Row_changed_date = formatDate(hse_incident.Row_changed_date)
	hse_incident.Effective_date = formatDateString(hse_incident.Effective_date)
	hse_incident.Expiry_date = formatDateString(hse_incident.Expiry_date)
	hse_incident.Incident_date = formatDateString(hse_incident.Incident_date)
	hse_incident.Recorded_time = formatDateString(hse_incident.Recorded_time)
	hse_incident.Row_effective_date = formatDateString(hse_incident.Row_effective_date)
	hse_incident.Row_expiry_date = formatDateString(hse_incident.Row_expiry_date)






	rows, err := stmt.Exec(hse_incident.Incident_id, hse_incident.Active_ind, hse_incident.Effective_date, hse_incident.Expiry_date, hse_incident.Incident_class_id, hse_incident.Incident_date, hse_incident.Incident_duration, hse_incident.Incident_duration_ouom, hse_incident.Incident_duration_uom, hse_incident.Lost_time_ind, hse_incident.Ppdm_guid, hse_incident.Recorded_time, hse_incident.Recorded_timezone, hse_incident.Remark, hse_incident.Reported_by_ba_id, hse_incident.Reported_by_name, hse_incident.Reported_ind, hse_incident.Source, hse_incident.Work_related_ind, hse_incident.Row_changed_by, hse_incident.Row_changed_date, hse_incident.Row_created_by, hse_incident.Row_created_date, hse_incident.Row_effective_date, hse_incident.Row_expiry_date, hse_incident.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateHseIncident(c *fiber.Ctx) error {
	var hse_incident dto.Hse_incident

	setDefaults(&hse_incident)

	if err := json.Unmarshal(c.Body(), &hse_incident); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	hse_incident.Incident_id = id

    if hse_incident.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update hse_incident set active_ind = :1, effective_date = :2, expiry_date = :3, incident_class_id = :4, incident_date = :5, incident_duration = :6, incident_duration_ouom = :7, incident_duration_uom = :8, lost_time_ind = :9, ppdm_guid = :10, recorded_time = :11, recorded_timezone = :12, remark = :13, reported_by_ba_id = :14, reported_by_name = :15, reported_ind = :16, source = :17, work_related_ind = :18, row_changed_by = :19, row_changed_date = :20, row_created_by = :21, row_effective_date = :22, row_expiry_date = :23, row_quality = :24 where incident_id = :26")
	if err != nil {
		return err
	}

	hse_incident.Row_changed_date = formatDate(hse_incident.Row_changed_date)
	hse_incident.Effective_date = formatDateString(hse_incident.Effective_date)
	hse_incident.Expiry_date = formatDateString(hse_incident.Expiry_date)
	hse_incident.Incident_date = formatDateString(hse_incident.Incident_date)
	hse_incident.Recorded_time = formatDateString(hse_incident.Recorded_time)
	hse_incident.Row_effective_date = formatDateString(hse_incident.Row_effective_date)
	hse_incident.Row_expiry_date = formatDateString(hse_incident.Row_expiry_date)






	rows, err := stmt.Exec(hse_incident.Active_ind, hse_incident.Effective_date, hse_incident.Expiry_date, hse_incident.Incident_class_id, hse_incident.Incident_date, hse_incident.Incident_duration, hse_incident.Incident_duration_ouom, hse_incident.Incident_duration_uom, hse_incident.Lost_time_ind, hse_incident.Ppdm_guid, hse_incident.Recorded_time, hse_incident.Recorded_timezone, hse_incident.Remark, hse_incident.Reported_by_ba_id, hse_incident.Reported_by_name, hse_incident.Reported_ind, hse_incident.Source, hse_incident.Work_related_ind, hse_incident.Row_changed_by, hse_incident.Row_changed_date, hse_incident.Row_created_by, hse_incident.Row_effective_date, hse_incident.Row_expiry_date, hse_incident.Row_quality, hse_incident.Incident_id)
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

func PatchHseIncident(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update hse_incident set "
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
		} else if key == "effective_date" || key == "expiry_date" || key == "incident_date" || key == "recorded_time" || key == "row_effective_date" || key == "row_expiry_date"      {
			if date, ok := value.(string); ok && len(date) > 0 {
				formattedDate := formatDateString(&date)
				value = formattedDate
			}
		}
		values = append(values, value)
	}
	query += " where incident_id = :id"

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

func DeleteHseIncident(c *fiber.Ctx) error {
	id := c.Params("id")
	var hse_incident dto.Hse_incident
	hse_incident.Incident_id = id

	stmt, err := db.Prepare("delete from hse_incident where incident_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(hse_incident.Incident_id)
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


