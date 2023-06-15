package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetHseIncidentEquip(c *fiber.Ctx) error {
	rows, err := db.Query("select * from hse_incident_equip")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Hse_incident_equip

	for rows.Next() {
		var hse_incident_equip dto.Hse_incident_equip
		if err := rows.Scan(&hse_incident_equip.Incident_id, &hse_incident_equip.Equip_obs_no, &hse_incident_equip.Equip_role_obs_no, &hse_incident_equip.Active_ind, &hse_incident_equip.Crew_role, &hse_incident_equip.Effective_date, &hse_incident_equip.Equipment_group, &hse_incident_equip.Equipment_id, &hse_incident_equip.Equipment_type, &hse_incident_equip.Expiry_date, &hse_incident_equip.Period_obs_no, &hse_incident_equip.Ppdm_guid, &hse_incident_equip.Remark, &hse_incident_equip.Source, &hse_incident_equip.Uwi, &hse_incident_equip.Row_changed_by, &hse_incident_equip.Row_changed_date, &hse_incident_equip.Row_created_by, &hse_incident_equip.Row_created_date, &hse_incident_equip.Row_effective_date, &hse_incident_equip.Row_expiry_date, &hse_incident_equip.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append hse_incident_equip to result
		result = append(result, hse_incident_equip)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetHseIncidentEquip(c *fiber.Ctx) error {
	var hse_incident_equip dto.Hse_incident_equip

	setDefaults(&hse_incident_equip)

	if err := json.Unmarshal(c.Body(), &hse_incident_equip); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into hse_incident_equip values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22)")
	if err != nil {
		return err
	}
	hse_incident_equip.Row_created_date = formatDate(hse_incident_equip.Row_created_date)
	hse_incident_equip.Row_changed_date = formatDate(hse_incident_equip.Row_changed_date)
	hse_incident_equip.Effective_date = formatDateString(hse_incident_equip.Effective_date)
	hse_incident_equip.Expiry_date = formatDateString(hse_incident_equip.Expiry_date)
	hse_incident_equip.Row_effective_date = formatDateString(hse_incident_equip.Row_effective_date)
	hse_incident_equip.Row_expiry_date = formatDateString(hse_incident_equip.Row_expiry_date)






	rows, err := stmt.Exec(hse_incident_equip.Incident_id, hse_incident_equip.Equip_obs_no, hse_incident_equip.Equip_role_obs_no, hse_incident_equip.Active_ind, hse_incident_equip.Crew_role, hse_incident_equip.Effective_date, hse_incident_equip.Equipment_group, hse_incident_equip.Equipment_id, hse_incident_equip.Equipment_type, hse_incident_equip.Expiry_date, hse_incident_equip.Period_obs_no, hse_incident_equip.Ppdm_guid, hse_incident_equip.Remark, hse_incident_equip.Source, hse_incident_equip.Uwi, hse_incident_equip.Row_changed_by, hse_incident_equip.Row_changed_date, hse_incident_equip.Row_created_by, hse_incident_equip.Row_created_date, hse_incident_equip.Row_effective_date, hse_incident_equip.Row_expiry_date, hse_incident_equip.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateHseIncidentEquip(c *fiber.Ctx) error {
	var hse_incident_equip dto.Hse_incident_equip

	setDefaults(&hse_incident_equip)

	if err := json.Unmarshal(c.Body(), &hse_incident_equip); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	hse_incident_equip.Incident_id = id

    if hse_incident_equip.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update hse_incident_equip set equip_obs_no = :1, equip_role_obs_no = :2, active_ind = :3, crew_role = :4, effective_date = :5, equipment_group = :6, equipment_id = :7, equipment_type = :8, expiry_date = :9, period_obs_no = :10, ppdm_guid = :11, remark = :12, source = :13, uwi = :14, row_changed_by = :15, row_changed_date = :16, row_created_by = :17, row_effective_date = :18, row_expiry_date = :19, row_quality = :20 where incident_id = :22")
	if err != nil {
		return err
	}

	hse_incident_equip.Row_changed_date = formatDate(hse_incident_equip.Row_changed_date)
	hse_incident_equip.Effective_date = formatDateString(hse_incident_equip.Effective_date)
	hse_incident_equip.Expiry_date = formatDateString(hse_incident_equip.Expiry_date)
	hse_incident_equip.Row_effective_date = formatDateString(hse_incident_equip.Row_effective_date)
	hse_incident_equip.Row_expiry_date = formatDateString(hse_incident_equip.Row_expiry_date)






	rows, err := stmt.Exec(hse_incident_equip.Equip_obs_no, hse_incident_equip.Equip_role_obs_no, hse_incident_equip.Active_ind, hse_incident_equip.Crew_role, hse_incident_equip.Effective_date, hse_incident_equip.Equipment_group, hse_incident_equip.Equipment_id, hse_incident_equip.Equipment_type, hse_incident_equip.Expiry_date, hse_incident_equip.Period_obs_no, hse_incident_equip.Ppdm_guid, hse_incident_equip.Remark, hse_incident_equip.Source, hse_incident_equip.Uwi, hse_incident_equip.Row_changed_by, hse_incident_equip.Row_changed_date, hse_incident_equip.Row_created_by, hse_incident_equip.Row_effective_date, hse_incident_equip.Row_expiry_date, hse_incident_equip.Row_quality, hse_incident_equip.Incident_id)
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

func PatchHseIncidentEquip(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update hse_incident_equip set "
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

func DeleteHseIncidentEquip(c *fiber.Ctx) error {
	id := c.Params("id")
	var hse_incident_equip dto.Hse_incident_equip
	hse_incident_equip.Incident_id = id

	stmt, err := db.Prepare("delete from hse_incident_equip where incident_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(hse_incident_equip.Incident_id)
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


