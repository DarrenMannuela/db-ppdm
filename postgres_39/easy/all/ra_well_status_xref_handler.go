package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetRaWellStatusXref(c *fiber.Ctx) error {
	rows, err := db.Query("select * from ra_well_status_xref")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Ra_well_status_xref

	for rows.Next() {
		var ra_well_status_xref dto.Ra_well_status_xref
		if err := rows.Scan(&ra_well_status_xref.Status_xref_id, &ra_well_status_xref.Status_xref_obs_no, &ra_well_status_xref.Alias_id, &ra_well_status_xref.Abbreviation, &ra_well_status_xref.Active_ind, &ra_well_status_xref.Alias_long_name, &ra_well_status_xref.Alias_reason_type, &ra_well_status_xref.Alias_short_name, &ra_well_status_xref.Alias_type, &ra_well_status_xref.Amended_date, &ra_well_status_xref.Created_date, &ra_well_status_xref.Effective_date, &ra_well_status_xref.Expiry_date, &ra_well_status_xref.Status_xref_id, &ra_well_status_xref.Original_ind, &ra_well_status_xref.Owner_ba_id, &ra_well_status_xref.Ppdm_guid, &ra_well_status_xref.Preferred_ind, &ra_well_status_xref.Reason_desc, &ra_well_status_xref.Remark, &ra_well_status_xref.Source, &ra_well_status_xref.Source_document, &ra_well_status_xref.Struckoff_date, &ra_well_status_xref.Sw_application_id, &ra_well_status_xref.Use_rule, &ra_well_status_xref.Row_changed_by, &ra_well_status_xref.Row_changed_date, &ra_well_status_xref.Row_created_by, &ra_well_status_xref.Row_created_date, &ra_well_status_xref.Row_effective_date, &ra_well_status_xref.Row_expiry_date, &ra_well_status_xref.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append ra_well_status_xref to result
		result = append(result, ra_well_status_xref)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetRaWellStatusXref(c *fiber.Ctx) error {
	var ra_well_status_xref dto.Ra_well_status_xref

	setDefaults(&ra_well_status_xref)

	if err := json.Unmarshal(c.Body(), &ra_well_status_xref); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into ra_well_status_xref values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32)")
	if err != nil {
		return err
	}
	ra_well_status_xref.Row_created_date = formatDate(ra_well_status_xref.Row_created_date)
	ra_well_status_xref.Row_changed_date = formatDate(ra_well_status_xref.Row_changed_date)
	ra_well_status_xref.Amended_date = formatDateString(ra_well_status_xref.Amended_date)
	ra_well_status_xref.Created_date = formatDateString(ra_well_status_xref.Created_date)
	ra_well_status_xref.Effective_date = formatDateString(ra_well_status_xref.Effective_date)
	ra_well_status_xref.Expiry_date = formatDateString(ra_well_status_xref.Expiry_date)
	ra_well_status_xref.Struckoff_date = formatDateString(ra_well_status_xref.Struckoff_date)
	ra_well_status_xref.Row_effective_date = formatDateString(ra_well_status_xref.Row_effective_date)
	ra_well_status_xref.Row_expiry_date = formatDateString(ra_well_status_xref.Row_expiry_date)






	rows, err := stmt.Exec(ra_well_status_xref.Status_xref_id, ra_well_status_xref.Status_xref_obs_no, ra_well_status_xref.Alias_id, ra_well_status_xref.Abbreviation, ra_well_status_xref.Active_ind, ra_well_status_xref.Alias_long_name, ra_well_status_xref.Alias_reason_type, ra_well_status_xref.Alias_short_name, ra_well_status_xref.Alias_type, ra_well_status_xref.Amended_date, ra_well_status_xref.Created_date, ra_well_status_xref.Effective_date, ra_well_status_xref.Expiry_date, ra_well_status_xref.Status_xref_id, ra_well_status_xref.Original_ind, ra_well_status_xref.Owner_ba_id, ra_well_status_xref.Ppdm_guid, ra_well_status_xref.Preferred_ind, ra_well_status_xref.Reason_desc, ra_well_status_xref.Remark, ra_well_status_xref.Source, ra_well_status_xref.Source_document, ra_well_status_xref.Struckoff_date, ra_well_status_xref.Sw_application_id, ra_well_status_xref.Use_rule, ra_well_status_xref.Row_changed_by, ra_well_status_xref.Row_changed_date, ra_well_status_xref.Row_created_by, ra_well_status_xref.Row_created_date, ra_well_status_xref.Row_effective_date, ra_well_status_xref.Row_expiry_date, ra_well_status_xref.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateRaWellStatusXref(c *fiber.Ctx) error {
	var ra_well_status_xref dto.Ra_well_status_xref

	setDefaults(&ra_well_status_xref)

	if err := json.Unmarshal(c.Body(), &ra_well_status_xref); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	ra_well_status_xref.Status_xref_id = id

    if ra_well_status_xref.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update ra_well_status_xref set status_xref_obs_no = :1, alias_id = :2, abbreviation = :3, active_ind = :4, alias_long_name = :5, alias_reason_type = :6, alias_short_name = :7, alias_type = :8, amended_date = :9, created_date = :10, effective_date = :11, expiry_date = :12, status_xref_id = :13, original_ind = :14, owner_ba_id = :15, ppdm_guid = :16, preferred_ind = :17, reason_desc = :18, remark = :19, source = :20, source_document = :21, struckoff_date = :22, sw_application_id = :23, use_rule = :24, row_changed_by = :25, row_changed_date = :26, row_created_by = :27, row_effective_date = :28, row_expiry_date = :29, row_quality = :30 where status_xref_id = :32")
	if err != nil {
		return err
	}

	ra_well_status_xref.Row_changed_date = formatDate(ra_well_status_xref.Row_changed_date)
	ra_well_status_xref.Amended_date = formatDateString(ra_well_status_xref.Amended_date)
	ra_well_status_xref.Created_date = formatDateString(ra_well_status_xref.Created_date)
	ra_well_status_xref.Effective_date = formatDateString(ra_well_status_xref.Effective_date)
	ra_well_status_xref.Expiry_date = formatDateString(ra_well_status_xref.Expiry_date)
	ra_well_status_xref.Struckoff_date = formatDateString(ra_well_status_xref.Struckoff_date)
	ra_well_status_xref.Row_effective_date = formatDateString(ra_well_status_xref.Row_effective_date)
	ra_well_status_xref.Row_expiry_date = formatDateString(ra_well_status_xref.Row_expiry_date)






	rows, err := stmt.Exec(ra_well_status_xref.Status_xref_obs_no, ra_well_status_xref.Alias_id, ra_well_status_xref.Abbreviation, ra_well_status_xref.Active_ind, ra_well_status_xref.Alias_long_name, ra_well_status_xref.Alias_reason_type, ra_well_status_xref.Alias_short_name, ra_well_status_xref.Alias_type, ra_well_status_xref.Amended_date, ra_well_status_xref.Created_date, ra_well_status_xref.Effective_date, ra_well_status_xref.Expiry_date, ra_well_status_xref.Status_xref_id, ra_well_status_xref.Original_ind, ra_well_status_xref.Owner_ba_id, ra_well_status_xref.Ppdm_guid, ra_well_status_xref.Preferred_ind, ra_well_status_xref.Reason_desc, ra_well_status_xref.Remark, ra_well_status_xref.Source, ra_well_status_xref.Source_document, ra_well_status_xref.Struckoff_date, ra_well_status_xref.Sw_application_id, ra_well_status_xref.Use_rule, ra_well_status_xref.Row_changed_by, ra_well_status_xref.Row_changed_date, ra_well_status_xref.Row_created_by, ra_well_status_xref.Row_effective_date, ra_well_status_xref.Row_expiry_date, ra_well_status_xref.Row_quality, ra_well_status_xref.Status_xref_id)
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

func PatchRaWellStatusXref(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update ra_well_status_xref set "
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
	query += " where status_xref_id = :id"

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

func DeleteRaWellStatusXref(c *fiber.Ctx) error {
	id := c.Params("id")
	var ra_well_status_xref dto.Ra_well_status_xref
	ra_well_status_xref.Status_xref_id = id

	stmt, err := db.Prepare("delete from ra_well_status_xref where status_xref_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(ra_well_status_xref.Status_xref_id)
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


