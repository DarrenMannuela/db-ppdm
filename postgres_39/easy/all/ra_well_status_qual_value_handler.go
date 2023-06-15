package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetRaWellStatusQualValue(c *fiber.Ctx) error {
	rows, err := db.Query("select * from ra_well_status_qual_value")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Ra_well_status_qual_value

	for rows.Next() {
		var ra_well_status_qual_value dto.Ra_well_status_qual_value
		if err := rows.Scan(&ra_well_status_qual_value.Status_type, &ra_well_status_qual_value.Status, &ra_well_status_qual_value.Status_qualifier, &ra_well_status_qual_value.Status_qualifier_value, &ra_well_status_qual_value.Alias_id, &ra_well_status_qual_value.Abbreviation, &ra_well_status_qual_value.Active_ind, &ra_well_status_qual_value.Alias_long_name, &ra_well_status_qual_value.Alias_reason_type, &ra_well_status_qual_value.Alias_short_name, &ra_well_status_qual_value.Alias_type, &ra_well_status_qual_value.Amended_date, &ra_well_status_qual_value.Created_date, &ra_well_status_qual_value.Effective_date, &ra_well_status_qual_value.Expiry_date, &ra_well_status_qual_value.Status_type, &ra_well_status_qual_value.Original_ind, &ra_well_status_qual_value.Owner_ba_id, &ra_well_status_qual_value.Ppdm_guid, &ra_well_status_qual_value.Preferred_ind, &ra_well_status_qual_value.Reason_desc, &ra_well_status_qual_value.Remark, &ra_well_status_qual_value.Source, &ra_well_status_qual_value.Source_document, &ra_well_status_qual_value.Struckoff_date, &ra_well_status_qual_value.Sw_application_id, &ra_well_status_qual_value.Use_rule, &ra_well_status_qual_value.Row_changed_by, &ra_well_status_qual_value.Row_changed_date, &ra_well_status_qual_value.Row_created_by, &ra_well_status_qual_value.Row_created_date, &ra_well_status_qual_value.Row_effective_date, &ra_well_status_qual_value.Row_expiry_date, &ra_well_status_qual_value.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append ra_well_status_qual_value to result
		result = append(result, ra_well_status_qual_value)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetRaWellStatusQualValue(c *fiber.Ctx) error {
	var ra_well_status_qual_value dto.Ra_well_status_qual_value

	setDefaults(&ra_well_status_qual_value)

	if err := json.Unmarshal(c.Body(), &ra_well_status_qual_value); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into ra_well_status_qual_value values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32, :33, :34)")
	if err != nil {
		return err
	}
	ra_well_status_qual_value.Row_created_date = formatDate(ra_well_status_qual_value.Row_created_date)
	ra_well_status_qual_value.Row_changed_date = formatDate(ra_well_status_qual_value.Row_changed_date)
	ra_well_status_qual_value.Amended_date = formatDateString(ra_well_status_qual_value.Amended_date)
	ra_well_status_qual_value.Created_date = formatDateString(ra_well_status_qual_value.Created_date)
	ra_well_status_qual_value.Effective_date = formatDateString(ra_well_status_qual_value.Effective_date)
	ra_well_status_qual_value.Expiry_date = formatDateString(ra_well_status_qual_value.Expiry_date)
	ra_well_status_qual_value.Struckoff_date = formatDateString(ra_well_status_qual_value.Struckoff_date)
	ra_well_status_qual_value.Row_effective_date = formatDateString(ra_well_status_qual_value.Row_effective_date)
	ra_well_status_qual_value.Row_expiry_date = formatDateString(ra_well_status_qual_value.Row_expiry_date)






	rows, err := stmt.Exec(ra_well_status_qual_value.Status_type, ra_well_status_qual_value.Status, ra_well_status_qual_value.Status_qualifier, ra_well_status_qual_value.Status_qualifier_value, ra_well_status_qual_value.Alias_id, ra_well_status_qual_value.Abbreviation, ra_well_status_qual_value.Active_ind, ra_well_status_qual_value.Alias_long_name, ra_well_status_qual_value.Alias_reason_type, ra_well_status_qual_value.Alias_short_name, ra_well_status_qual_value.Alias_type, ra_well_status_qual_value.Amended_date, ra_well_status_qual_value.Created_date, ra_well_status_qual_value.Effective_date, ra_well_status_qual_value.Expiry_date, ra_well_status_qual_value.Status_type, ra_well_status_qual_value.Original_ind, ra_well_status_qual_value.Owner_ba_id, ra_well_status_qual_value.Ppdm_guid, ra_well_status_qual_value.Preferred_ind, ra_well_status_qual_value.Reason_desc, ra_well_status_qual_value.Remark, ra_well_status_qual_value.Source, ra_well_status_qual_value.Source_document, ra_well_status_qual_value.Struckoff_date, ra_well_status_qual_value.Sw_application_id, ra_well_status_qual_value.Use_rule, ra_well_status_qual_value.Row_changed_by, ra_well_status_qual_value.Row_changed_date, ra_well_status_qual_value.Row_created_by, ra_well_status_qual_value.Row_created_date, ra_well_status_qual_value.Row_effective_date, ra_well_status_qual_value.Row_expiry_date, ra_well_status_qual_value.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateRaWellStatusQualValue(c *fiber.Ctx) error {
	var ra_well_status_qual_value dto.Ra_well_status_qual_value

	setDefaults(&ra_well_status_qual_value)

	if err := json.Unmarshal(c.Body(), &ra_well_status_qual_value); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	ra_well_status_qual_value.Status_type = id

    if ra_well_status_qual_value.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update ra_well_status_qual_value set status = :1, status_qualifier = :2, status_qualifier_value = :3, alias_id = :4, abbreviation = :5, active_ind = :6, alias_long_name = :7, alias_reason_type = :8, alias_short_name = :9, alias_type = :10, amended_date = :11, created_date = :12, effective_date = :13, expiry_date = :14, status_type = :15, original_ind = :16, owner_ba_id = :17, ppdm_guid = :18, preferred_ind = :19, reason_desc = :20, remark = :21, source = :22, source_document = :23, struckoff_date = :24, sw_application_id = :25, use_rule = :26, row_changed_by = :27, row_changed_date = :28, row_created_by = :29, row_effective_date = :30, row_expiry_date = :31, row_quality = :32 where status_type = :34")
	if err != nil {
		return err
	}

	ra_well_status_qual_value.Row_changed_date = formatDate(ra_well_status_qual_value.Row_changed_date)
	ra_well_status_qual_value.Amended_date = formatDateString(ra_well_status_qual_value.Amended_date)
	ra_well_status_qual_value.Created_date = formatDateString(ra_well_status_qual_value.Created_date)
	ra_well_status_qual_value.Effective_date = formatDateString(ra_well_status_qual_value.Effective_date)
	ra_well_status_qual_value.Expiry_date = formatDateString(ra_well_status_qual_value.Expiry_date)
	ra_well_status_qual_value.Struckoff_date = formatDateString(ra_well_status_qual_value.Struckoff_date)
	ra_well_status_qual_value.Row_effective_date = formatDateString(ra_well_status_qual_value.Row_effective_date)
	ra_well_status_qual_value.Row_expiry_date = formatDateString(ra_well_status_qual_value.Row_expiry_date)






	rows, err := stmt.Exec(ra_well_status_qual_value.Status, ra_well_status_qual_value.Status_qualifier, ra_well_status_qual_value.Status_qualifier_value, ra_well_status_qual_value.Alias_id, ra_well_status_qual_value.Abbreviation, ra_well_status_qual_value.Active_ind, ra_well_status_qual_value.Alias_long_name, ra_well_status_qual_value.Alias_reason_type, ra_well_status_qual_value.Alias_short_name, ra_well_status_qual_value.Alias_type, ra_well_status_qual_value.Amended_date, ra_well_status_qual_value.Created_date, ra_well_status_qual_value.Effective_date, ra_well_status_qual_value.Expiry_date, ra_well_status_qual_value.Status_type, ra_well_status_qual_value.Original_ind, ra_well_status_qual_value.Owner_ba_id, ra_well_status_qual_value.Ppdm_guid, ra_well_status_qual_value.Preferred_ind, ra_well_status_qual_value.Reason_desc, ra_well_status_qual_value.Remark, ra_well_status_qual_value.Source, ra_well_status_qual_value.Source_document, ra_well_status_qual_value.Struckoff_date, ra_well_status_qual_value.Sw_application_id, ra_well_status_qual_value.Use_rule, ra_well_status_qual_value.Row_changed_by, ra_well_status_qual_value.Row_changed_date, ra_well_status_qual_value.Row_created_by, ra_well_status_qual_value.Row_effective_date, ra_well_status_qual_value.Row_expiry_date, ra_well_status_qual_value.Row_quality, ra_well_status_qual_value.Status_type)
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

func PatchRaWellStatusQualValue(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update ra_well_status_qual_value set "
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
	query += " where status_type = :id"

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

func DeleteRaWellStatusQualValue(c *fiber.Ctx) error {
	id := c.Params("id")
	var ra_well_status_qual_value dto.Ra_well_status_qual_value
	ra_well_status_qual_value.Status_type = id

	stmt, err := db.Prepare("delete from ra_well_status_qual_value where status_type = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(ra_well_status_qual_value.Status_type)
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


