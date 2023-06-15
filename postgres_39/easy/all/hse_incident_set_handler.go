package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetHseIncidentSet(c *fiber.Ctx) error {
	rows, err := db.Query("select * from hse_incident_set")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Hse_incident_set

	for rows.Next() {
		var hse_incident_set dto.Hse_incident_set
		if err := rows.Scan(&hse_incident_set.Incident_set_id, &hse_incident_set.Active_ind, &hse_incident_set.Effective_date, &hse_incident_set.Expiry_date, &hse_incident_set.Full_name, &hse_incident_set.Owner_ba_id, &hse_incident_set.Ppdm_guid, &hse_incident_set.Reference_num, &hse_incident_set.Remark, &hse_incident_set.Source, &hse_incident_set.Row_changed_by, &hse_incident_set.Row_changed_date, &hse_incident_set.Row_created_by, &hse_incident_set.Row_created_date, &hse_incident_set.Row_effective_date, &hse_incident_set.Row_expiry_date, &hse_incident_set.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append hse_incident_set to result
		result = append(result, hse_incident_set)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetHseIncidentSet(c *fiber.Ctx) error {
	var hse_incident_set dto.Hse_incident_set

	setDefaults(&hse_incident_set)

	if err := json.Unmarshal(c.Body(), &hse_incident_set); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into hse_incident_set values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17)")
	if err != nil {
		return err
	}
	hse_incident_set.Row_created_date = formatDate(hse_incident_set.Row_created_date)
	hse_incident_set.Row_changed_date = formatDate(hse_incident_set.Row_changed_date)
	hse_incident_set.Effective_date = formatDateString(hse_incident_set.Effective_date)
	hse_incident_set.Expiry_date = formatDateString(hse_incident_set.Expiry_date)
	hse_incident_set.Row_effective_date = formatDateString(hse_incident_set.Row_effective_date)
	hse_incident_set.Row_expiry_date = formatDateString(hse_incident_set.Row_expiry_date)






	rows, err := stmt.Exec(hse_incident_set.Incident_set_id, hse_incident_set.Active_ind, hse_incident_set.Effective_date, hse_incident_set.Expiry_date, hse_incident_set.Full_name, hse_incident_set.Owner_ba_id, hse_incident_set.Ppdm_guid, hse_incident_set.Reference_num, hse_incident_set.Remark, hse_incident_set.Source, hse_incident_set.Row_changed_by, hse_incident_set.Row_changed_date, hse_incident_set.Row_created_by, hse_incident_set.Row_created_date, hse_incident_set.Row_effective_date, hse_incident_set.Row_expiry_date, hse_incident_set.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateHseIncidentSet(c *fiber.Ctx) error {
	var hse_incident_set dto.Hse_incident_set

	setDefaults(&hse_incident_set)

	if err := json.Unmarshal(c.Body(), &hse_incident_set); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	hse_incident_set.Incident_set_id = id

    if hse_incident_set.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update hse_incident_set set active_ind = :1, effective_date = :2, expiry_date = :3, full_name = :4, owner_ba_id = :5, ppdm_guid = :6, reference_num = :7, remark = :8, source = :9, row_changed_by = :10, row_changed_date = :11, row_created_by = :12, row_effective_date = :13, row_expiry_date = :14, row_quality = :15 where incident_set_id = :17")
	if err != nil {
		return err
	}

	hse_incident_set.Row_changed_date = formatDate(hse_incident_set.Row_changed_date)
	hse_incident_set.Effective_date = formatDateString(hse_incident_set.Effective_date)
	hse_incident_set.Expiry_date = formatDateString(hse_incident_set.Expiry_date)
	hse_incident_set.Row_effective_date = formatDateString(hse_incident_set.Row_effective_date)
	hse_incident_set.Row_expiry_date = formatDateString(hse_incident_set.Row_expiry_date)






	rows, err := stmt.Exec(hse_incident_set.Active_ind, hse_incident_set.Effective_date, hse_incident_set.Expiry_date, hse_incident_set.Full_name, hse_incident_set.Owner_ba_id, hse_incident_set.Ppdm_guid, hse_incident_set.Reference_num, hse_incident_set.Remark, hse_incident_set.Source, hse_incident_set.Row_changed_by, hse_incident_set.Row_changed_date, hse_incident_set.Row_created_by, hse_incident_set.Row_effective_date, hse_incident_set.Row_expiry_date, hse_incident_set.Row_quality, hse_incident_set.Incident_set_id)
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

func PatchHseIncidentSet(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update hse_incident_set set "
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

func DeleteHseIncidentSet(c *fiber.Ctx) error {
	id := c.Params("id")
	var hse_incident_set dto.Hse_incident_set
	hse_incident_set.Incident_set_id = id

	stmt, err := db.Prepare("delete from hse_incident_set where incident_set_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(hse_incident_set.Incident_set_id)
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


