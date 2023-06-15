package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetHseIncidentDetail(c *fiber.Ctx) error {
	rows, err := db.Query("select * from hse_incident_detail")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Hse_incident_detail

	for rows.Next() {
		var hse_incident_detail dto.Hse_incident_detail
		if err := rows.Scan(&hse_incident_detail.Incident_id, &hse_incident_detail.Detail_obs_no, &hse_incident_detail.Active_ind, &hse_incident_detail.Detail_type, &hse_incident_detail.Effective_date, &hse_incident_detail.Expiry_date, &hse_incident_detail.Incident_date, &hse_incident_detail.Incident_duration, &hse_incident_detail.Incident_duration_ouom, &hse_incident_detail.Incident_duration_uom, &hse_incident_detail.Incident_set_id, &hse_incident_detail.Incident_severity_id, &hse_incident_detail.Incident_type_id, &hse_incident_detail.Ppdm_guid, &hse_incident_detail.Remark, &hse_incident_detail.Source, &hse_incident_detail.Row_changed_by, &hse_incident_detail.Row_changed_date, &hse_incident_detail.Row_created_by, &hse_incident_detail.Row_created_date, &hse_incident_detail.Row_effective_date, &hse_incident_detail.Row_expiry_date, &hse_incident_detail.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append hse_incident_detail to result
		result = append(result, hse_incident_detail)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetHseIncidentDetail(c *fiber.Ctx) error {
	var hse_incident_detail dto.Hse_incident_detail

	setDefaults(&hse_incident_detail)

	if err := json.Unmarshal(c.Body(), &hse_incident_detail); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into hse_incident_detail values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23)")
	if err != nil {
		return err
	}
	hse_incident_detail.Row_created_date = formatDate(hse_incident_detail.Row_created_date)
	hse_incident_detail.Row_changed_date = formatDate(hse_incident_detail.Row_changed_date)
	hse_incident_detail.Effective_date = formatDateString(hse_incident_detail.Effective_date)
	hse_incident_detail.Expiry_date = formatDateString(hse_incident_detail.Expiry_date)
	hse_incident_detail.Incident_date = formatDateString(hse_incident_detail.Incident_date)
	hse_incident_detail.Row_effective_date = formatDateString(hse_incident_detail.Row_effective_date)
	hse_incident_detail.Row_expiry_date = formatDateString(hse_incident_detail.Row_expiry_date)






	rows, err := stmt.Exec(hse_incident_detail.Incident_id, hse_incident_detail.Detail_obs_no, hse_incident_detail.Active_ind, hse_incident_detail.Detail_type, hse_incident_detail.Effective_date, hse_incident_detail.Expiry_date, hse_incident_detail.Incident_date, hse_incident_detail.Incident_duration, hse_incident_detail.Incident_duration_ouom, hse_incident_detail.Incident_duration_uom, hse_incident_detail.Incident_set_id, hse_incident_detail.Incident_severity_id, hse_incident_detail.Incident_type_id, hse_incident_detail.Ppdm_guid, hse_incident_detail.Remark, hse_incident_detail.Source, hse_incident_detail.Row_changed_by, hse_incident_detail.Row_changed_date, hse_incident_detail.Row_created_by, hse_incident_detail.Row_created_date, hse_incident_detail.Row_effective_date, hse_incident_detail.Row_expiry_date, hse_incident_detail.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateHseIncidentDetail(c *fiber.Ctx) error {
	var hse_incident_detail dto.Hse_incident_detail

	setDefaults(&hse_incident_detail)

	if err := json.Unmarshal(c.Body(), &hse_incident_detail); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	hse_incident_detail.Incident_id = id

    if hse_incident_detail.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update hse_incident_detail set detail_obs_no = :1, active_ind = :2, detail_type = :3, effective_date = :4, expiry_date = :5, incident_date = :6, incident_duration = :7, incident_duration_ouom = :8, incident_duration_uom = :9, incident_set_id = :10, incident_severity_id = :11, incident_type_id = :12, ppdm_guid = :13, remark = :14, source = :15, row_changed_by = :16, row_changed_date = :17, row_created_by = :18, row_effective_date = :19, row_expiry_date = :20, row_quality = :21 where incident_id = :23")
	if err != nil {
		return err
	}

	hse_incident_detail.Row_changed_date = formatDate(hse_incident_detail.Row_changed_date)
	hse_incident_detail.Effective_date = formatDateString(hse_incident_detail.Effective_date)
	hse_incident_detail.Expiry_date = formatDateString(hse_incident_detail.Expiry_date)
	hse_incident_detail.Incident_date = formatDateString(hse_incident_detail.Incident_date)
	hse_incident_detail.Row_effective_date = formatDateString(hse_incident_detail.Row_effective_date)
	hse_incident_detail.Row_expiry_date = formatDateString(hse_incident_detail.Row_expiry_date)






	rows, err := stmt.Exec(hse_incident_detail.Detail_obs_no, hse_incident_detail.Active_ind, hse_incident_detail.Detail_type, hse_incident_detail.Effective_date, hse_incident_detail.Expiry_date, hse_incident_detail.Incident_date, hse_incident_detail.Incident_duration, hse_incident_detail.Incident_duration_ouom, hse_incident_detail.Incident_duration_uom, hse_incident_detail.Incident_set_id, hse_incident_detail.Incident_severity_id, hse_incident_detail.Incident_type_id, hse_incident_detail.Ppdm_guid, hse_incident_detail.Remark, hse_incident_detail.Source, hse_incident_detail.Row_changed_by, hse_incident_detail.Row_changed_date, hse_incident_detail.Row_created_by, hse_incident_detail.Row_effective_date, hse_incident_detail.Row_expiry_date, hse_incident_detail.Row_quality, hse_incident_detail.Incident_id)
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

func PatchHseIncidentDetail(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update hse_incident_detail set "
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
		} else if key == "effective_date" || key == "expiry_date" || key == "incident_date" || key == "row_effective_date" || key == "row_expiry_date"      {
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

func DeleteHseIncidentDetail(c *fiber.Ctx) error {
	id := c.Params("id")
	var hse_incident_detail dto.Hse_incident_detail
	hse_incident_detail.Incident_id = id

	stmt, err := db.Prepare("delete from hse_incident_detail where incident_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(hse_incident_detail.Incident_id)
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


