package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetHseIncidentEquiv(c *fiber.Ctx) error {
	rows, err := db.Query("select * from hse_incident_equiv")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Hse_incident_equiv

	for rows.Next() {
		var hse_incident_equiv dto.Hse_incident_equiv
		if err := rows.Scan(&hse_incident_equiv.Incident_set_id, &hse_incident_equiv.Incident_type_id, &hse_incident_equiv.Incident_set_id2, &hse_incident_equiv.Incident_type_id2, &hse_incident_equiv.Active_ind, &hse_incident_equiv.Description, &hse_incident_equiv.Effective_date, &hse_incident_equiv.Expiry_date, &hse_incident_equiv.Ppdm_guid, &hse_incident_equiv.Remark, &hse_incident_equiv.Source, &hse_incident_equiv.Source_document_id, &hse_incident_equiv.Row_changed_by, &hse_incident_equiv.Row_changed_date, &hse_incident_equiv.Row_created_by, &hse_incident_equiv.Row_created_date, &hse_incident_equiv.Row_effective_date, &hse_incident_equiv.Row_expiry_date, &hse_incident_equiv.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append hse_incident_equiv to result
		result = append(result, hse_incident_equiv)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetHseIncidentEquiv(c *fiber.Ctx) error {
	var hse_incident_equiv dto.Hse_incident_equiv

	setDefaults(&hse_incident_equiv)

	if err := json.Unmarshal(c.Body(), &hse_incident_equiv); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into hse_incident_equiv values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19)")
	if err != nil {
		return err
	}
	hse_incident_equiv.Row_created_date = formatDate(hse_incident_equiv.Row_created_date)
	hse_incident_equiv.Row_changed_date = formatDate(hse_incident_equiv.Row_changed_date)
	hse_incident_equiv.Effective_date = formatDateString(hse_incident_equiv.Effective_date)
	hse_incident_equiv.Expiry_date = formatDateString(hse_incident_equiv.Expiry_date)
	hse_incident_equiv.Row_effective_date = formatDateString(hse_incident_equiv.Row_effective_date)
	hse_incident_equiv.Row_expiry_date = formatDateString(hse_incident_equiv.Row_expiry_date)






	rows, err := stmt.Exec(hse_incident_equiv.Incident_set_id, hse_incident_equiv.Incident_type_id, hse_incident_equiv.Incident_set_id2, hse_incident_equiv.Incident_type_id2, hse_incident_equiv.Active_ind, hse_incident_equiv.Description, hse_incident_equiv.Effective_date, hse_incident_equiv.Expiry_date, hse_incident_equiv.Ppdm_guid, hse_incident_equiv.Remark, hse_incident_equiv.Source, hse_incident_equiv.Source_document_id, hse_incident_equiv.Row_changed_by, hse_incident_equiv.Row_changed_date, hse_incident_equiv.Row_created_by, hse_incident_equiv.Row_created_date, hse_incident_equiv.Row_effective_date, hse_incident_equiv.Row_expiry_date, hse_incident_equiv.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateHseIncidentEquiv(c *fiber.Ctx) error {
	var hse_incident_equiv dto.Hse_incident_equiv

	setDefaults(&hse_incident_equiv)

	if err := json.Unmarshal(c.Body(), &hse_incident_equiv); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	hse_incident_equiv.Incident_set_id = id

    if hse_incident_equiv.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update hse_incident_equiv set incident_type_id = :1, incident_set_id2 = :2, incident_type_id2 = :3, active_ind = :4, description = :5, effective_date = :6, expiry_date = :7, ppdm_guid = :8, remark = :9, source = :10, source_document_id = :11, row_changed_by = :12, row_changed_date = :13, row_created_by = :14, row_effective_date = :15, row_expiry_date = :16, row_quality = :17 where incident_set_id = :19")
	if err != nil {
		return err
	}

	hse_incident_equiv.Row_changed_date = formatDate(hse_incident_equiv.Row_changed_date)
	hse_incident_equiv.Effective_date = formatDateString(hse_incident_equiv.Effective_date)
	hse_incident_equiv.Expiry_date = formatDateString(hse_incident_equiv.Expiry_date)
	hse_incident_equiv.Row_effective_date = formatDateString(hse_incident_equiv.Row_effective_date)
	hse_incident_equiv.Row_expiry_date = formatDateString(hse_incident_equiv.Row_expiry_date)






	rows, err := stmt.Exec(hse_incident_equiv.Incident_type_id, hse_incident_equiv.Incident_set_id2, hse_incident_equiv.Incident_type_id2, hse_incident_equiv.Active_ind, hse_incident_equiv.Description, hse_incident_equiv.Effective_date, hse_incident_equiv.Expiry_date, hse_incident_equiv.Ppdm_guid, hse_incident_equiv.Remark, hse_incident_equiv.Source, hse_incident_equiv.Source_document_id, hse_incident_equiv.Row_changed_by, hse_incident_equiv.Row_changed_date, hse_incident_equiv.Row_created_by, hse_incident_equiv.Row_effective_date, hse_incident_equiv.Row_expiry_date, hse_incident_equiv.Row_quality, hse_incident_equiv.Incident_set_id)
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

func PatchHseIncidentEquiv(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update hse_incident_equiv set "
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

func DeleteHseIncidentEquiv(c *fiber.Ctx) error {
	id := c.Params("id")
	var hse_incident_equiv dto.Hse_incident_equiv
	hse_incident_equiv.Incident_set_id = id

	stmt, err := db.Prepare("delete from hse_incident_equiv where incident_set_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(hse_incident_equiv.Incident_set_id)
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


