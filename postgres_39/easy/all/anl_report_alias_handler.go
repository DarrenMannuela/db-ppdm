package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetAnlReportAlias(c *fiber.Ctx) error {
	rows, err := db.Query("select * from anl_report_alias")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Anl_report_alias

	for rows.Next() {
		var anl_report_alias dto.Anl_report_alias
		if err := rows.Scan(&anl_report_alias.Analysis_id, &anl_report_alias.Source, &anl_report_alias.Alias_id, &anl_report_alias.Abbreviation, &anl_report_alias.Active_ind, &anl_report_alias.Alias_long_name, &anl_report_alias.Alias_reason_type, &anl_report_alias.Alias_short_name, &anl_report_alias.Alias_type, &anl_report_alias.Amended_date, &anl_report_alias.Created_date, &anl_report_alias.Effective_date, &anl_report_alias.Expiry_date, &anl_report_alias.Analysis_id, &anl_report_alias.Original_ind, &anl_report_alias.Owner_ba_id, &anl_report_alias.Ppdm_guid, &anl_report_alias.Preferred_ind, &anl_report_alias.Reason_desc, &anl_report_alias.Remark, &anl_report_alias.Source_document_id, &anl_report_alias.Struckoff_date, &anl_report_alias.Sw_application_id, &anl_report_alias.Use_rule, &anl_report_alias.Row_changed_by, &anl_report_alias.Row_changed_date, &anl_report_alias.Row_created_by, &anl_report_alias.Row_created_date, &anl_report_alias.Row_effective_date, &anl_report_alias.Row_expiry_date, &anl_report_alias.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append anl_report_alias to result
		result = append(result, anl_report_alias)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetAnlReportAlias(c *fiber.Ctx) error {
	var anl_report_alias dto.Anl_report_alias

	setDefaults(&anl_report_alias)

	if err := json.Unmarshal(c.Body(), &anl_report_alias); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into anl_report_alias values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31)")
	if err != nil {
		return err
	}
	anl_report_alias.Row_created_date = formatDate(anl_report_alias.Row_created_date)
	anl_report_alias.Row_changed_date = formatDate(anl_report_alias.Row_changed_date)
	anl_report_alias.Amended_date = formatDateString(anl_report_alias.Amended_date)
	anl_report_alias.Created_date = formatDateString(anl_report_alias.Created_date)
	anl_report_alias.Effective_date = formatDateString(anl_report_alias.Effective_date)
	anl_report_alias.Expiry_date = formatDateString(anl_report_alias.Expiry_date)
	anl_report_alias.Struckoff_date = formatDateString(anl_report_alias.Struckoff_date)
	anl_report_alias.Row_effective_date = formatDateString(anl_report_alias.Row_effective_date)
	anl_report_alias.Row_expiry_date = formatDateString(anl_report_alias.Row_expiry_date)






	rows, err := stmt.Exec(anl_report_alias.Analysis_id, anl_report_alias.Source, anl_report_alias.Alias_id, anl_report_alias.Abbreviation, anl_report_alias.Active_ind, anl_report_alias.Alias_long_name, anl_report_alias.Alias_reason_type, anl_report_alias.Alias_short_name, anl_report_alias.Alias_type, anl_report_alias.Amended_date, anl_report_alias.Created_date, anl_report_alias.Effective_date, anl_report_alias.Expiry_date, anl_report_alias.Analysis_id, anl_report_alias.Original_ind, anl_report_alias.Owner_ba_id, anl_report_alias.Ppdm_guid, anl_report_alias.Preferred_ind, anl_report_alias.Reason_desc, anl_report_alias.Remark, anl_report_alias.Source_document_id, anl_report_alias.Struckoff_date, anl_report_alias.Sw_application_id, anl_report_alias.Use_rule, anl_report_alias.Row_changed_by, anl_report_alias.Row_changed_date, anl_report_alias.Row_created_by, anl_report_alias.Row_created_date, anl_report_alias.Row_effective_date, anl_report_alias.Row_expiry_date, anl_report_alias.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateAnlReportAlias(c *fiber.Ctx) error {
	var anl_report_alias dto.Anl_report_alias

	setDefaults(&anl_report_alias)

	if err := json.Unmarshal(c.Body(), &anl_report_alias); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	anl_report_alias.Analysis_id = id

    if anl_report_alias.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update anl_report_alias set source = :1, alias_id = :2, abbreviation = :3, active_ind = :4, alias_long_name = :5, alias_reason_type = :6, alias_short_name = :7, alias_type = :8, amended_date = :9, created_date = :10, effective_date = :11, expiry_date = :12, analysis_id = :13, original_ind = :14, owner_ba_id = :15, ppdm_guid = :16, preferred_ind = :17, reason_desc = :18, remark = :19, source_document_id = :20, struckoff_date = :21, sw_application_id = :22, use_rule = :23, row_changed_by = :24, row_changed_date = :25, row_created_by = :26, row_effective_date = :27, row_expiry_date = :28, row_quality = :29 where analysis_id = :31")
	if err != nil {
		return err
	}

	anl_report_alias.Row_changed_date = formatDate(anl_report_alias.Row_changed_date)
	anl_report_alias.Amended_date = formatDateString(anl_report_alias.Amended_date)
	anl_report_alias.Created_date = formatDateString(anl_report_alias.Created_date)
	anl_report_alias.Effective_date = formatDateString(anl_report_alias.Effective_date)
	anl_report_alias.Expiry_date = formatDateString(anl_report_alias.Expiry_date)
	anl_report_alias.Struckoff_date = formatDateString(anl_report_alias.Struckoff_date)
	anl_report_alias.Row_effective_date = formatDateString(anl_report_alias.Row_effective_date)
	anl_report_alias.Row_expiry_date = formatDateString(anl_report_alias.Row_expiry_date)






	rows, err := stmt.Exec(anl_report_alias.Source, anl_report_alias.Alias_id, anl_report_alias.Abbreviation, anl_report_alias.Active_ind, anl_report_alias.Alias_long_name, anl_report_alias.Alias_reason_type, anl_report_alias.Alias_short_name, anl_report_alias.Alias_type, anl_report_alias.Amended_date, anl_report_alias.Created_date, anl_report_alias.Effective_date, anl_report_alias.Expiry_date, anl_report_alias.Analysis_id, anl_report_alias.Original_ind, anl_report_alias.Owner_ba_id, anl_report_alias.Ppdm_guid, anl_report_alias.Preferred_ind, anl_report_alias.Reason_desc, anl_report_alias.Remark, anl_report_alias.Source_document_id, anl_report_alias.Struckoff_date, anl_report_alias.Sw_application_id, anl_report_alias.Use_rule, anl_report_alias.Row_changed_by, anl_report_alias.Row_changed_date, anl_report_alias.Row_created_by, anl_report_alias.Row_effective_date, anl_report_alias.Row_expiry_date, anl_report_alias.Row_quality, anl_report_alias.Analysis_id)
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

func PatchAnlReportAlias(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update anl_report_alias set "
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
	query += " where analysis_id = :id"

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

func DeleteAnlReportAlias(c *fiber.Ctx) error {
	id := c.Params("id")
	var anl_report_alias dto.Anl_report_alias
	anl_report_alias.Analysis_id = id

	stmt, err := db.Prepare("delete from anl_report_alias where analysis_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(anl_report_alias.Analysis_id)
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


