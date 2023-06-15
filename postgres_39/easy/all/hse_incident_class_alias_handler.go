package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetHseIncidentClassAlias(c *fiber.Ctx) error {
	rows, err := db.Query("select * from hse_incident_class_alias")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Hse_incident_class_alias

	for rows.Next() {
		var hse_incident_class_alias dto.Hse_incident_class_alias
		if err := rows.Scan(&hse_incident_class_alias.Incident_class_id, &hse_incident_class_alias.Alias_obs_no, &hse_incident_class_alias.Abbreviation, &hse_incident_class_alias.Active_ind, &hse_incident_class_alias.Alias_long_name, &hse_incident_class_alias.Alias_reason_type, &hse_incident_class_alias.Alias_short_name, &hse_incident_class_alias.Alias_type, &hse_incident_class_alias.Amended_date, &hse_incident_class_alias.Created_date, &hse_incident_class_alias.Effective_date, &hse_incident_class_alias.Expiry_date, &hse_incident_class_alias.Incident_class_id, &hse_incident_class_alias.Original_ind, &hse_incident_class_alias.Owner_ba_id, &hse_incident_class_alias.Ppdm_guid, &hse_incident_class_alias.Preferred_ind, &hse_incident_class_alias.Reason_desc, &hse_incident_class_alias.Remark, &hse_incident_class_alias.Source, &hse_incident_class_alias.Source_document_id, &hse_incident_class_alias.Struckoff_date, &hse_incident_class_alias.Sw_application_id, &hse_incident_class_alias.Use_rule, &hse_incident_class_alias.Row_changed_by, &hse_incident_class_alias.Row_changed_date, &hse_incident_class_alias.Row_created_by, &hse_incident_class_alias.Row_created_date, &hse_incident_class_alias.Row_effective_date, &hse_incident_class_alias.Row_expiry_date, &hse_incident_class_alias.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append hse_incident_class_alias to result
		result = append(result, hse_incident_class_alias)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetHseIncidentClassAlias(c *fiber.Ctx) error {
	var hse_incident_class_alias dto.Hse_incident_class_alias

	setDefaults(&hse_incident_class_alias)

	if err := json.Unmarshal(c.Body(), &hse_incident_class_alias); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into hse_incident_class_alias values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31)")
	if err != nil {
		return err
	}
	hse_incident_class_alias.Row_created_date = formatDate(hse_incident_class_alias.Row_created_date)
	hse_incident_class_alias.Row_changed_date = formatDate(hse_incident_class_alias.Row_changed_date)
	hse_incident_class_alias.Amended_date = formatDateString(hse_incident_class_alias.Amended_date)
	hse_incident_class_alias.Created_date = formatDateString(hse_incident_class_alias.Created_date)
	hse_incident_class_alias.Effective_date = formatDateString(hse_incident_class_alias.Effective_date)
	hse_incident_class_alias.Expiry_date = formatDateString(hse_incident_class_alias.Expiry_date)
	hse_incident_class_alias.Struckoff_date = formatDateString(hse_incident_class_alias.Struckoff_date)
	hse_incident_class_alias.Row_effective_date = formatDateString(hse_incident_class_alias.Row_effective_date)
	hse_incident_class_alias.Row_expiry_date = formatDateString(hse_incident_class_alias.Row_expiry_date)






	rows, err := stmt.Exec(hse_incident_class_alias.Incident_class_id, hse_incident_class_alias.Alias_obs_no, hse_incident_class_alias.Abbreviation, hse_incident_class_alias.Active_ind, hse_incident_class_alias.Alias_long_name, hse_incident_class_alias.Alias_reason_type, hse_incident_class_alias.Alias_short_name, hse_incident_class_alias.Alias_type, hse_incident_class_alias.Amended_date, hse_incident_class_alias.Created_date, hse_incident_class_alias.Effective_date, hse_incident_class_alias.Expiry_date, hse_incident_class_alias.Incident_class_id, hse_incident_class_alias.Original_ind, hse_incident_class_alias.Owner_ba_id, hse_incident_class_alias.Ppdm_guid, hse_incident_class_alias.Preferred_ind, hse_incident_class_alias.Reason_desc, hse_incident_class_alias.Remark, hse_incident_class_alias.Source, hse_incident_class_alias.Source_document_id, hse_incident_class_alias.Struckoff_date, hse_incident_class_alias.Sw_application_id, hse_incident_class_alias.Use_rule, hse_incident_class_alias.Row_changed_by, hse_incident_class_alias.Row_changed_date, hse_incident_class_alias.Row_created_by, hse_incident_class_alias.Row_created_date, hse_incident_class_alias.Row_effective_date, hse_incident_class_alias.Row_expiry_date, hse_incident_class_alias.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateHseIncidentClassAlias(c *fiber.Ctx) error {
	var hse_incident_class_alias dto.Hse_incident_class_alias

	setDefaults(&hse_incident_class_alias)

	if err := json.Unmarshal(c.Body(), &hse_incident_class_alias); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	hse_incident_class_alias.Incident_class_id = id

    if hse_incident_class_alias.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update hse_incident_class_alias set alias_obs_no = :1, abbreviation = :2, active_ind = :3, alias_long_name = :4, alias_reason_type = :5, alias_short_name = :6, alias_type = :7, amended_date = :8, created_date = :9, effective_date = :10, expiry_date = :11, incident_class_id = :12, original_ind = :13, owner_ba_id = :14, ppdm_guid = :15, preferred_ind = :16, reason_desc = :17, remark = :18, source = :19, source_document_id = :20, struckoff_date = :21, sw_application_id = :22, use_rule = :23, row_changed_by = :24, row_changed_date = :25, row_created_by = :26, row_effective_date = :27, row_expiry_date = :28, row_quality = :29 where incident_class_id = :31")
	if err != nil {
		return err
	}

	hse_incident_class_alias.Row_changed_date = formatDate(hse_incident_class_alias.Row_changed_date)
	hse_incident_class_alias.Amended_date = formatDateString(hse_incident_class_alias.Amended_date)
	hse_incident_class_alias.Created_date = formatDateString(hse_incident_class_alias.Created_date)
	hse_incident_class_alias.Effective_date = formatDateString(hse_incident_class_alias.Effective_date)
	hse_incident_class_alias.Expiry_date = formatDateString(hse_incident_class_alias.Expiry_date)
	hse_incident_class_alias.Struckoff_date = formatDateString(hse_incident_class_alias.Struckoff_date)
	hse_incident_class_alias.Row_effective_date = formatDateString(hse_incident_class_alias.Row_effective_date)
	hse_incident_class_alias.Row_expiry_date = formatDateString(hse_incident_class_alias.Row_expiry_date)






	rows, err := stmt.Exec(hse_incident_class_alias.Alias_obs_no, hse_incident_class_alias.Abbreviation, hse_incident_class_alias.Active_ind, hse_incident_class_alias.Alias_long_name, hse_incident_class_alias.Alias_reason_type, hse_incident_class_alias.Alias_short_name, hse_incident_class_alias.Alias_type, hse_incident_class_alias.Amended_date, hse_incident_class_alias.Created_date, hse_incident_class_alias.Effective_date, hse_incident_class_alias.Expiry_date, hse_incident_class_alias.Incident_class_id, hse_incident_class_alias.Original_ind, hse_incident_class_alias.Owner_ba_id, hse_incident_class_alias.Ppdm_guid, hse_incident_class_alias.Preferred_ind, hse_incident_class_alias.Reason_desc, hse_incident_class_alias.Remark, hse_incident_class_alias.Source, hse_incident_class_alias.Source_document_id, hse_incident_class_alias.Struckoff_date, hse_incident_class_alias.Sw_application_id, hse_incident_class_alias.Use_rule, hse_incident_class_alias.Row_changed_by, hse_incident_class_alias.Row_changed_date, hse_incident_class_alias.Row_created_by, hse_incident_class_alias.Row_effective_date, hse_incident_class_alias.Row_expiry_date, hse_incident_class_alias.Row_quality, hse_incident_class_alias.Incident_class_id)
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

func PatchHseIncidentClassAlias(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update hse_incident_class_alias set "
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
		} else if key == "amended_date" || key == "created_date" || key == "effective_date" || key == "expiry_date" || key == "struckoff_date" || key == "row_effective_date" || key == "row_expiry_date"      {
			if date, ok := value.(string); ok && len(date) > 0 {
				formattedDate := formatDateString(&date)
				value = formattedDate
			}
		}
		values = append(values, value)
	}
	query += " where incident_class_id = :id"

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

func DeleteHseIncidentClassAlias(c *fiber.Ctx) error {
	id := c.Params("id")
	var hse_incident_class_alias dto.Hse_incident_class_alias
	hse_incident_class_alias.Incident_class_id = id

	stmt, err := db.Prepare("delete from hse_incident_class_alias where incident_class_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(hse_incident_class_alias.Incident_class_id)
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


