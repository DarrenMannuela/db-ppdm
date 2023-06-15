package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetHseIncidentTypeAlias(c *fiber.Ctx) error {
	rows, err := db.Query("select * from hse_incident_type_alias")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Hse_incident_type_alias

	for rows.Next() {
		var hse_incident_type_alias dto.Hse_incident_type_alias
		if err := rows.Scan(&hse_incident_type_alias.Incident_set_id, &hse_incident_type_alias.Incident_type_id, &hse_incident_type_alias.Incident_alias_id, &hse_incident_type_alias.Abbreviation, &hse_incident_type_alias.Active_ind, &hse_incident_type_alias.Alias_long_name, &hse_incident_type_alias.Alias_reason_type, &hse_incident_type_alias.Alias_short_name, &hse_incident_type_alias.Alias_type, &hse_incident_type_alias.Amended_date, &hse_incident_type_alias.Created_date, &hse_incident_type_alias.Effective_date, &hse_incident_type_alias.Expiry_date, &hse_incident_type_alias.Incident_set_id, &hse_incident_type_alias.Original_ind, &hse_incident_type_alias.Owner_ba_id, &hse_incident_type_alias.Ppdm_guid, &hse_incident_type_alias.Preferred_ind, &hse_incident_type_alias.Reason_desc, &hse_incident_type_alias.Remark, &hse_incident_type_alias.Source, &hse_incident_type_alias.Source_document_id, &hse_incident_type_alias.Struckoff_date, &hse_incident_type_alias.Sw_application_id, &hse_incident_type_alias.Use_rule, &hse_incident_type_alias.Row_changed_by, &hse_incident_type_alias.Row_changed_date, &hse_incident_type_alias.Row_created_by, &hse_incident_type_alias.Row_created_date, &hse_incident_type_alias.Row_effective_date, &hse_incident_type_alias.Row_expiry_date, &hse_incident_type_alias.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append hse_incident_type_alias to result
		result = append(result, hse_incident_type_alias)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetHseIncidentTypeAlias(c *fiber.Ctx) error {
	var hse_incident_type_alias dto.Hse_incident_type_alias

	setDefaults(&hse_incident_type_alias)

	if err := json.Unmarshal(c.Body(), &hse_incident_type_alias); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into hse_incident_type_alias values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32)")
	if err != nil {
		return err
	}
	hse_incident_type_alias.Row_created_date = formatDate(hse_incident_type_alias.Row_created_date)
	hse_incident_type_alias.Row_changed_date = formatDate(hse_incident_type_alias.Row_changed_date)
	hse_incident_type_alias.Amended_date = formatDateString(hse_incident_type_alias.Amended_date)
	hse_incident_type_alias.Created_date = formatDateString(hse_incident_type_alias.Created_date)
	hse_incident_type_alias.Effective_date = formatDateString(hse_incident_type_alias.Effective_date)
	hse_incident_type_alias.Expiry_date = formatDateString(hse_incident_type_alias.Expiry_date)
	hse_incident_type_alias.Struckoff_date = formatDateString(hse_incident_type_alias.Struckoff_date)
	hse_incident_type_alias.Row_effective_date = formatDateString(hse_incident_type_alias.Row_effective_date)
	hse_incident_type_alias.Row_expiry_date = formatDateString(hse_incident_type_alias.Row_expiry_date)






	rows, err := stmt.Exec(hse_incident_type_alias.Incident_set_id, hse_incident_type_alias.Incident_type_id, hse_incident_type_alias.Incident_alias_id, hse_incident_type_alias.Abbreviation, hse_incident_type_alias.Active_ind, hse_incident_type_alias.Alias_long_name, hse_incident_type_alias.Alias_reason_type, hse_incident_type_alias.Alias_short_name, hse_incident_type_alias.Alias_type, hse_incident_type_alias.Amended_date, hse_incident_type_alias.Created_date, hse_incident_type_alias.Effective_date, hse_incident_type_alias.Expiry_date, hse_incident_type_alias.Incident_set_id, hse_incident_type_alias.Original_ind, hse_incident_type_alias.Owner_ba_id, hse_incident_type_alias.Ppdm_guid, hse_incident_type_alias.Preferred_ind, hse_incident_type_alias.Reason_desc, hse_incident_type_alias.Remark, hse_incident_type_alias.Source, hse_incident_type_alias.Source_document_id, hse_incident_type_alias.Struckoff_date, hse_incident_type_alias.Sw_application_id, hse_incident_type_alias.Use_rule, hse_incident_type_alias.Row_changed_by, hse_incident_type_alias.Row_changed_date, hse_incident_type_alias.Row_created_by, hse_incident_type_alias.Row_created_date, hse_incident_type_alias.Row_effective_date, hse_incident_type_alias.Row_expiry_date, hse_incident_type_alias.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateHseIncidentTypeAlias(c *fiber.Ctx) error {
	var hse_incident_type_alias dto.Hse_incident_type_alias

	setDefaults(&hse_incident_type_alias)

	if err := json.Unmarshal(c.Body(), &hse_incident_type_alias); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	hse_incident_type_alias.Incident_set_id = id

    if hse_incident_type_alias.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update hse_incident_type_alias set incident_type_id = :1, incident_alias_id = :2, abbreviation = :3, active_ind = :4, alias_long_name = :5, alias_reason_type = :6, alias_short_name = :7, alias_type = :8, amended_date = :9, created_date = :10, effective_date = :11, expiry_date = :12, incident_set_id = :13, original_ind = :14, owner_ba_id = :15, ppdm_guid = :16, preferred_ind = :17, reason_desc = :18, remark = :19, source = :20, source_document_id = :21, struckoff_date = :22, sw_application_id = :23, use_rule = :24, row_changed_by = :25, row_changed_date = :26, row_created_by = :27, row_effective_date = :28, row_expiry_date = :29, row_quality = :30 where incident_set_id = :32")
	if err != nil {
		return err
	}

	hse_incident_type_alias.Row_changed_date = formatDate(hse_incident_type_alias.Row_changed_date)
	hse_incident_type_alias.Amended_date = formatDateString(hse_incident_type_alias.Amended_date)
	hse_incident_type_alias.Created_date = formatDateString(hse_incident_type_alias.Created_date)
	hse_incident_type_alias.Effective_date = formatDateString(hse_incident_type_alias.Effective_date)
	hse_incident_type_alias.Expiry_date = formatDateString(hse_incident_type_alias.Expiry_date)
	hse_incident_type_alias.Struckoff_date = formatDateString(hse_incident_type_alias.Struckoff_date)
	hse_incident_type_alias.Row_effective_date = formatDateString(hse_incident_type_alias.Row_effective_date)
	hse_incident_type_alias.Row_expiry_date = formatDateString(hse_incident_type_alias.Row_expiry_date)






	rows, err := stmt.Exec(hse_incident_type_alias.Incident_type_id, hse_incident_type_alias.Incident_alias_id, hse_incident_type_alias.Abbreviation, hse_incident_type_alias.Active_ind, hse_incident_type_alias.Alias_long_name, hse_incident_type_alias.Alias_reason_type, hse_incident_type_alias.Alias_short_name, hse_incident_type_alias.Alias_type, hse_incident_type_alias.Amended_date, hse_incident_type_alias.Created_date, hse_incident_type_alias.Effective_date, hse_incident_type_alias.Expiry_date, hse_incident_type_alias.Incident_set_id, hse_incident_type_alias.Original_ind, hse_incident_type_alias.Owner_ba_id, hse_incident_type_alias.Ppdm_guid, hse_incident_type_alias.Preferred_ind, hse_incident_type_alias.Reason_desc, hse_incident_type_alias.Remark, hse_incident_type_alias.Source, hse_incident_type_alias.Source_document_id, hse_incident_type_alias.Struckoff_date, hse_incident_type_alias.Sw_application_id, hse_incident_type_alias.Use_rule, hse_incident_type_alias.Row_changed_by, hse_incident_type_alias.Row_changed_date, hse_incident_type_alias.Row_created_by, hse_incident_type_alias.Row_effective_date, hse_incident_type_alias.Row_expiry_date, hse_incident_type_alias.Row_quality, hse_incident_type_alias.Incident_set_id)
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

func PatchHseIncidentTypeAlias(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update hse_incident_type_alias set "
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

func DeleteHseIncidentTypeAlias(c *fiber.Ctx) error {
	id := c.Params("id")
	var hse_incident_type_alias dto.Hse_incident_type_alias
	hse_incident_type_alias.Incident_set_id = id

	stmt, err := db.Prepare("delete from hse_incident_type_alias where incident_set_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(hse_incident_type_alias.Incident_set_id)
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


