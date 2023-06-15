package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetHseIncidentInteraction(c *fiber.Ctx) error {
	rows, err := db.Query("select * from hse_incident_interaction")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Hse_incident_interaction

	for rows.Next() {
		var hse_incident_interaction dto.Hse_incident_interaction
		if err := rows.Scan(&hse_incident_interaction.Incident_id, &hse_incident_interaction.Interaction_obs_no, &hse_incident_interaction.Active_ind, &hse_incident_interaction.Cause_obs_no, &hse_incident_interaction.Description, &hse_incident_interaction.Detail_obs_no, &hse_incident_interaction.Effective_date, &hse_incident_interaction.Equip_obs_no, &hse_incident_interaction.Equip_role_obs_no, &hse_incident_interaction.Expiry_date, &hse_incident_interaction.Incident_substance, &hse_incident_interaction.Interaction_type, &hse_incident_interaction.Party_obs_no, &hse_incident_interaction.Party_role_obs_no, &hse_incident_interaction.Ppdm_guid, &hse_incident_interaction.Remark, &hse_incident_interaction.Response_obs_no, &hse_incident_interaction.Source, &hse_incident_interaction.Substance_seq_no, &hse_incident_interaction.Row_changed_by, &hse_incident_interaction.Row_changed_date, &hse_incident_interaction.Row_created_by, &hse_incident_interaction.Row_created_date, &hse_incident_interaction.Row_effective_date, &hse_incident_interaction.Row_expiry_date, &hse_incident_interaction.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append hse_incident_interaction to result
		result = append(result, hse_incident_interaction)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetHseIncidentInteraction(c *fiber.Ctx) error {
	var hse_incident_interaction dto.Hse_incident_interaction

	setDefaults(&hse_incident_interaction)

	if err := json.Unmarshal(c.Body(), &hse_incident_interaction); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into hse_incident_interaction values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26)")
	if err != nil {
		return err
	}
	hse_incident_interaction.Row_created_date = formatDate(hse_incident_interaction.Row_created_date)
	hse_incident_interaction.Row_changed_date = formatDate(hse_incident_interaction.Row_changed_date)
	hse_incident_interaction.Effective_date = formatDateString(hse_incident_interaction.Effective_date)
	hse_incident_interaction.Expiry_date = formatDateString(hse_incident_interaction.Expiry_date)
	hse_incident_interaction.Row_effective_date = formatDateString(hse_incident_interaction.Row_effective_date)
	hse_incident_interaction.Row_expiry_date = formatDateString(hse_incident_interaction.Row_expiry_date)






	rows, err := stmt.Exec(hse_incident_interaction.Incident_id, hse_incident_interaction.Interaction_obs_no, hse_incident_interaction.Active_ind, hse_incident_interaction.Cause_obs_no, hse_incident_interaction.Description, hse_incident_interaction.Detail_obs_no, hse_incident_interaction.Effective_date, hse_incident_interaction.Equip_obs_no, hse_incident_interaction.Equip_role_obs_no, hse_incident_interaction.Expiry_date, hse_incident_interaction.Incident_substance, hse_incident_interaction.Interaction_type, hse_incident_interaction.Party_obs_no, hse_incident_interaction.Party_role_obs_no, hse_incident_interaction.Ppdm_guid, hse_incident_interaction.Remark, hse_incident_interaction.Response_obs_no, hse_incident_interaction.Source, hse_incident_interaction.Substance_seq_no, hse_incident_interaction.Row_changed_by, hse_incident_interaction.Row_changed_date, hse_incident_interaction.Row_created_by, hse_incident_interaction.Row_created_date, hse_incident_interaction.Row_effective_date, hse_incident_interaction.Row_expiry_date, hse_incident_interaction.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateHseIncidentInteraction(c *fiber.Ctx) error {
	var hse_incident_interaction dto.Hse_incident_interaction

	setDefaults(&hse_incident_interaction)

	if err := json.Unmarshal(c.Body(), &hse_incident_interaction); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	hse_incident_interaction.Incident_id = id

    if hse_incident_interaction.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update hse_incident_interaction set interaction_obs_no = :1, active_ind = :2, cause_obs_no = :3, description = :4, detail_obs_no = :5, effective_date = :6, equip_obs_no = :7, equip_role_obs_no = :8, expiry_date = :9, incident_substance = :10, interaction_type = :11, party_obs_no = :12, party_role_obs_no = :13, ppdm_guid = :14, remark = :15, response_obs_no = :16, source = :17, substance_seq_no = :18, row_changed_by = :19, row_changed_date = :20, row_created_by = :21, row_effective_date = :22, row_expiry_date = :23, row_quality = :24 where incident_id = :26")
	if err != nil {
		return err
	}

	hse_incident_interaction.Row_changed_date = formatDate(hse_incident_interaction.Row_changed_date)
	hse_incident_interaction.Effective_date = formatDateString(hse_incident_interaction.Effective_date)
	hse_incident_interaction.Expiry_date = formatDateString(hse_incident_interaction.Expiry_date)
	hse_incident_interaction.Row_effective_date = formatDateString(hse_incident_interaction.Row_effective_date)
	hse_incident_interaction.Row_expiry_date = formatDateString(hse_incident_interaction.Row_expiry_date)






	rows, err := stmt.Exec(hse_incident_interaction.Interaction_obs_no, hse_incident_interaction.Active_ind, hse_incident_interaction.Cause_obs_no, hse_incident_interaction.Description, hse_incident_interaction.Detail_obs_no, hse_incident_interaction.Effective_date, hse_incident_interaction.Equip_obs_no, hse_incident_interaction.Equip_role_obs_no, hse_incident_interaction.Expiry_date, hse_incident_interaction.Incident_substance, hse_incident_interaction.Interaction_type, hse_incident_interaction.Party_obs_no, hse_incident_interaction.Party_role_obs_no, hse_incident_interaction.Ppdm_guid, hse_incident_interaction.Remark, hse_incident_interaction.Response_obs_no, hse_incident_interaction.Source, hse_incident_interaction.Substance_seq_no, hse_incident_interaction.Row_changed_by, hse_incident_interaction.Row_changed_date, hse_incident_interaction.Row_created_by, hse_incident_interaction.Row_effective_date, hse_incident_interaction.Row_expiry_date, hse_incident_interaction.Row_quality, hse_incident_interaction.Incident_id)
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

func PatchHseIncidentInteraction(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update hse_incident_interaction set "
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

func DeleteHseIncidentInteraction(c *fiber.Ctx) error {
	id := c.Params("id")
	var hse_incident_interaction dto.Hse_incident_interaction
	hse_incident_interaction.Incident_id = id

	stmt, err := db.Prepare("delete from hse_incident_interaction where incident_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(hse_incident_interaction.Incident_id)
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


