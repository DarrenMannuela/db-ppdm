package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetPaleoFossilObs(c *fiber.Ctx) error {
	rows, err := db.Query("select * from paleo_fossil_obs")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Paleo_fossil_obs

	for rows.Next() {
		var paleo_fossil_obs dto.Paleo_fossil_obs
		if err := rows.Scan(&paleo_fossil_obs.Paleo_summary_id, &paleo_fossil_obs.Fossil_detail_id, &paleo_fossil_obs.Fossil_id, &paleo_fossil_obs.Observation_obs_no, &paleo_fossil_obs.Active_ind, &paleo_fossil_obs.Effective_date, &paleo_fossil_obs.Expiry_date, &paleo_fossil_obs.Fossil_obs_type, &paleo_fossil_obs.Observation, &paleo_fossil_obs.Ppdm_guid, &paleo_fossil_obs.Remark, &paleo_fossil_obs.Source, &paleo_fossil_obs.Row_changed_by, &paleo_fossil_obs.Row_changed_date, &paleo_fossil_obs.Row_created_by, &paleo_fossil_obs.Row_created_date, &paleo_fossil_obs.Row_effective_date, &paleo_fossil_obs.Row_expiry_date, &paleo_fossil_obs.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append paleo_fossil_obs to result
		result = append(result, paleo_fossil_obs)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetPaleoFossilObs(c *fiber.Ctx) error {
	var paleo_fossil_obs dto.Paleo_fossil_obs

	setDefaults(&paleo_fossil_obs)

	if err := json.Unmarshal(c.Body(), &paleo_fossil_obs); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into paleo_fossil_obs values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19)")
	if err != nil {
		return err
	}
	paleo_fossil_obs.Row_created_date = formatDate(paleo_fossil_obs.Row_created_date)
	paleo_fossil_obs.Row_changed_date = formatDate(paleo_fossil_obs.Row_changed_date)
	paleo_fossil_obs.Effective_date = formatDateString(paleo_fossil_obs.Effective_date)
	paleo_fossil_obs.Expiry_date = formatDateString(paleo_fossil_obs.Expiry_date)
	paleo_fossil_obs.Row_effective_date = formatDateString(paleo_fossil_obs.Row_effective_date)
	paleo_fossil_obs.Row_expiry_date = formatDateString(paleo_fossil_obs.Row_expiry_date)






	rows, err := stmt.Exec(paleo_fossil_obs.Paleo_summary_id, paleo_fossil_obs.Fossil_detail_id, paleo_fossil_obs.Fossil_id, paleo_fossil_obs.Observation_obs_no, paleo_fossil_obs.Active_ind, paleo_fossil_obs.Effective_date, paleo_fossil_obs.Expiry_date, paleo_fossil_obs.Fossil_obs_type, paleo_fossil_obs.Observation, paleo_fossil_obs.Ppdm_guid, paleo_fossil_obs.Remark, paleo_fossil_obs.Source, paleo_fossil_obs.Row_changed_by, paleo_fossil_obs.Row_changed_date, paleo_fossil_obs.Row_created_by, paleo_fossil_obs.Row_created_date, paleo_fossil_obs.Row_effective_date, paleo_fossil_obs.Row_expiry_date, paleo_fossil_obs.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdatePaleoFossilObs(c *fiber.Ctx) error {
	var paleo_fossil_obs dto.Paleo_fossil_obs

	setDefaults(&paleo_fossil_obs)

	if err := json.Unmarshal(c.Body(), &paleo_fossil_obs); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	paleo_fossil_obs.Paleo_summary_id = id

    if paleo_fossil_obs.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update paleo_fossil_obs set fossil_detail_id = :1, fossil_id = :2, observation_obs_no = :3, active_ind = :4, effective_date = :5, expiry_date = :6, fossil_obs_type = :7, observation = :8, ppdm_guid = :9, remark = :10, source = :11, row_changed_by = :12, row_changed_date = :13, row_created_by = :14, row_effective_date = :15, row_expiry_date = :16, row_quality = :17 where paleo_summary_id = :19")
	if err != nil {
		return err
	}

	paleo_fossil_obs.Row_changed_date = formatDate(paleo_fossil_obs.Row_changed_date)
	paleo_fossil_obs.Effective_date = formatDateString(paleo_fossil_obs.Effective_date)
	paleo_fossil_obs.Expiry_date = formatDateString(paleo_fossil_obs.Expiry_date)
	paleo_fossil_obs.Row_effective_date = formatDateString(paleo_fossil_obs.Row_effective_date)
	paleo_fossil_obs.Row_expiry_date = formatDateString(paleo_fossil_obs.Row_expiry_date)






	rows, err := stmt.Exec(paleo_fossil_obs.Fossil_detail_id, paleo_fossil_obs.Fossil_id, paleo_fossil_obs.Observation_obs_no, paleo_fossil_obs.Active_ind, paleo_fossil_obs.Effective_date, paleo_fossil_obs.Expiry_date, paleo_fossil_obs.Fossil_obs_type, paleo_fossil_obs.Observation, paleo_fossil_obs.Ppdm_guid, paleo_fossil_obs.Remark, paleo_fossil_obs.Source, paleo_fossil_obs.Row_changed_by, paleo_fossil_obs.Row_changed_date, paleo_fossil_obs.Row_created_by, paleo_fossil_obs.Row_effective_date, paleo_fossil_obs.Row_expiry_date, paleo_fossil_obs.Row_quality, paleo_fossil_obs.Paleo_summary_id)
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

func PatchPaleoFossilObs(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update paleo_fossil_obs set "
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
	query += " where paleo_summary_id = :id"

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

func DeletePaleoFossilObs(c *fiber.Ctx) error {
	id := c.Params("id")
	var paleo_fossil_obs dto.Paleo_fossil_obs
	paleo_fossil_obs.Paleo_summary_id = id

	stmt, err := db.Prepare("delete from paleo_fossil_obs where paleo_summary_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(paleo_fossil_obs.Paleo_summary_id)
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


