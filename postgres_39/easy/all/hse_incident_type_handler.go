package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetHseIncidentType(c *fiber.Ctx) error {
	rows, err := db.Query("select * from hse_incident_type")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Hse_incident_type

	for rows.Next() {
		var hse_incident_type dto.Hse_incident_type
		if err := rows.Scan(&hse_incident_type.Incident_set_id, &hse_incident_type.Incident_type_id, &hse_incident_type.Active_ind, &hse_incident_type.Description, &hse_incident_type.Determination_method, &hse_incident_type.Effective_date, &hse_incident_type.Expiry_date, &hse_incident_type.Ppdm_guid, &hse_incident_type.Project_plan_id, &hse_incident_type.Remark, &hse_incident_type.Report_ind, &hse_incident_type.Source, &hse_incident_type.Source_document_id, &hse_incident_type.Row_changed_by, &hse_incident_type.Row_changed_date, &hse_incident_type.Row_created_by, &hse_incident_type.Row_created_date, &hse_incident_type.Row_effective_date, &hse_incident_type.Row_expiry_date, &hse_incident_type.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append hse_incident_type to result
		result = append(result, hse_incident_type)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetHseIncidentType(c *fiber.Ctx) error {
	var hse_incident_type dto.Hse_incident_type

	setDefaults(&hse_incident_type)

	if err := json.Unmarshal(c.Body(), &hse_incident_type); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into hse_incident_type values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20)")
	if err != nil {
		return err
	}
	hse_incident_type.Row_created_date = formatDate(hse_incident_type.Row_created_date)
	hse_incident_type.Row_changed_date = formatDate(hse_incident_type.Row_changed_date)
	hse_incident_type.Effective_date = formatDateString(hse_incident_type.Effective_date)
	hse_incident_type.Expiry_date = formatDateString(hse_incident_type.Expiry_date)
	hse_incident_type.Row_effective_date = formatDateString(hse_incident_type.Row_effective_date)
	hse_incident_type.Row_expiry_date = formatDateString(hse_incident_type.Row_expiry_date)






	rows, err := stmt.Exec(hse_incident_type.Incident_set_id, hse_incident_type.Incident_type_id, hse_incident_type.Active_ind, hse_incident_type.Description, hse_incident_type.Determination_method, hse_incident_type.Effective_date, hse_incident_type.Expiry_date, hse_incident_type.Ppdm_guid, hse_incident_type.Project_plan_id, hse_incident_type.Remark, hse_incident_type.Report_ind, hse_incident_type.Source, hse_incident_type.Source_document_id, hse_incident_type.Row_changed_by, hse_incident_type.Row_changed_date, hse_incident_type.Row_created_by, hse_incident_type.Row_created_date, hse_incident_type.Row_effective_date, hse_incident_type.Row_expiry_date, hse_incident_type.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateHseIncidentType(c *fiber.Ctx) error {
	var hse_incident_type dto.Hse_incident_type

	setDefaults(&hse_incident_type)

	if err := json.Unmarshal(c.Body(), &hse_incident_type); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	hse_incident_type.Incident_set_id = id

    if hse_incident_type.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update hse_incident_type set incident_type_id = :1, active_ind = :2, description = :3, determination_method = :4, effective_date = :5, expiry_date = :6, ppdm_guid = :7, project_plan_id = :8, remark = :9, report_ind = :10, source = :11, source_document_id = :12, row_changed_by = :13, row_changed_date = :14, row_created_by = :15, row_effective_date = :16, row_expiry_date = :17, row_quality = :18 where incident_set_id = :20")
	if err != nil {
		return err
	}

	hse_incident_type.Row_changed_date = formatDate(hse_incident_type.Row_changed_date)
	hse_incident_type.Effective_date = formatDateString(hse_incident_type.Effective_date)
	hse_incident_type.Expiry_date = formatDateString(hse_incident_type.Expiry_date)
	hse_incident_type.Row_effective_date = formatDateString(hse_incident_type.Row_effective_date)
	hse_incident_type.Row_expiry_date = formatDateString(hse_incident_type.Row_expiry_date)






	rows, err := stmt.Exec(hse_incident_type.Incident_type_id, hse_incident_type.Active_ind, hse_incident_type.Description, hse_incident_type.Determination_method, hse_incident_type.Effective_date, hse_incident_type.Expiry_date, hse_incident_type.Ppdm_guid, hse_incident_type.Project_plan_id, hse_incident_type.Remark, hse_incident_type.Report_ind, hse_incident_type.Source, hse_incident_type.Source_document_id, hse_incident_type.Row_changed_by, hse_incident_type.Row_changed_date, hse_incident_type.Row_created_by, hse_incident_type.Row_effective_date, hse_incident_type.Row_expiry_date, hse_incident_type.Row_quality, hse_incident_type.Incident_set_id)
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

func PatchHseIncidentType(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update hse_incident_type set "
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
	query += " where incident_set_id = :id"

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

func DeleteHseIncidentType(c *fiber.Ctx) error {
	id := c.Params("id")
	var hse_incident_type dto.Hse_incident_type
	hse_incident_type.Incident_set_id = id

	stmt, err := db.Prepare("delete from hse_incident_type where incident_set_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(hse_incident_type.Incident_set_id)
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


